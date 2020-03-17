package main

import (
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var (
	config = flag.String("conf", "/opt/yt2watch/yt2watch.yml", "config file")
	c      conf
)

type conf struct {
	API struct {
		URL       string `yaml:"url"`
		Key       string `yaml:"key"`
		Recursive bool   `yaml:"recursive"`
	} `yaml:"api"`
	Channels []struct {
		Name string `yaml:"name"`
		ID   string `yaml:"id"`
	} `yaml:"channels"`
}

type Results struct {
	NextPageToken string `json:"nextPageToken"`
	Items         []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
	} `json:"items"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile(*config)
	if err != nil {
		log.Fatalf("ioutil.ReadFile [%v]", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("yaml.Unmarshal [%v]", err)
	}
	return c
}

// getIDs returns a
func getIDs(channelID, token string, recursive bool) (next string, ids []string) {
	url := fmt.Sprintf("%s?key=%s&channelId=%s&part=snippet,id&order=date&maxResults=50", c.API.URL, c.API.Key, channelID)
	if token != "" {
		url = fmt.Sprintf("%s&pageToken=%s", url, token)
	}
	log.Debugf("url [%s]", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("http.NewRequest [%v]", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do [%v]", err)
	}

	var results Results
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		log.Fatalf("json.NewDecoder [%v]", err)
	}

	if results.Error.Message != "" {
		log.Fatalf("%v [%v]", results.Error.Message, results.Error.Code)
	}

	var videos []string
	if results.NextPageToken != "" && recursive {
		next, videos = getIDs(channelID, results.NextPageToken, recursive)
		for _, videoID := range videos {
			ids = append(ids, videoID)
		}
	}

	for _, item := range results.Items {
		ids = append(ids, item.ID.VideoID)
	}

	return
}

func main() {
	flag.Parse()
	fmt.Println(*config)
	if _, err := os.Stat(*config); err != nil {
		log.Fatalf("missing config [%v]", config)
	}
	c.getConf()
	log.SetLevel(log.DebugLevel)

	rand.Seed(time.Now().Unix())
	channel := c.Channels[rand.Intn(len(c.Channels))]

	_, videos := getIDs(channel.ID, "", c.API.Recursive)
	log.Debugf("channel [%v], videos [%v]", channel.Name, len(videos))
	id := videos[rand.Intn(len(videos))]

	err := exec.Command("xdg-open", fmt.Sprintf("https://www.youtube.com/watch?v=%s", id)).Start()
	if err != nil {
		log.Fatalf("exec.Command [%v]", err)
	}
}
