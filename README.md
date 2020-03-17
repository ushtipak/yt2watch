# yt2watch

Load up a random YouTube video from one of the preferred channels


## Config

YT API key and belowed channels should be saved in YAML format (by default `/opt/yt2watch/yt2watch.yml`, overriden with `--conf`)

For example:

```yaml
api:
    url: "https://www.googleapis.com/youtube/v3/search"
    key: "acquire one from https://console.developers.google.com/apis/ :)"
    recursive: false
channels:
  - name: "Vsauce"
    id: "UC6nSFpj9HTCZ5t-N3Rm3-HA"
  - name: "SmarterEveryDay"
    id: "UC6107grRI4m0o2-emgoDnAA"
  - name: "Isaac Arthur"
    id: "UCZFipeZtQM5CKUjx6grh54g"
  - name: "5-Minute Crafts"
    id: "UC295-Dw_tDNtZXFeAPAW6Aw"
  - name: " HowStuffWorks"
    id: "UCa35qyNpnlZ_u8n9qoAZbMQ"
  - name: "Mental Floss"
    id: "UCpZ5qUqpW4hW4zdfuBxMSJA"
  - name: "minutephysics"
    id: "UCUHW94eEFW7hkUMVaZz4eDg"
  - name: "SciShow"
    id: "UCZYTClx2T1of7BRZ86-8fow"
  - name: "Matthew Santoro"
    id: "UCXhSCMRRPyxSoyLSPFxK7VA"
  - name: "AsapSCIENCE"
    id: "UCC552Sd-3nyi_tk2BudLUzA"
  - name: "Veritasium"
    id: "UCHnyfMqiRRG1u-2MsSQLbXA"
  - name: "CGP Grey"
    id: "UC2C_jShtL725hvbm1arSV9w"
  - name: "Science Channel"
    id: "UCvJiYiBUbw4tmpRSZT2r1Hw"
  - name: "Kurzgesagt – In a Nutshell"
    id: "UCsXVk37bltHxD1rDPwtNM8Q"
```

