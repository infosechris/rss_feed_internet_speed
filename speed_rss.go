package main

import (
	"fmt"
	"time"
  "github.com/adolfosilva/speedtest"
)

const rssTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <title>Verizon Fios Internet Speed Alerts</title>
    <description>This RSS feed reports on slow internet speeds for Verizon Fios.</description>
    <link>http://example.com/verizon-fios-internet-speed-alerts</link>
    <lastBuildDate>%s</lastBuildDate>
    <item>
      <title>Slow Internet Speed Alert</title>
      <description>Verizon Fios internet speed is currently slow. Download: %f Mbps Upload: %f Mbps</description>
      <link>http://example.com/verizon-fios-internet-speed-alerts</link>
      <pubDate>%s</pubDate>
    </item>
  </channel>
</rss>`

func InternetSpeedCheck() (float64, float64) {
	client := speedtest.NewClient()
	downloadSpeed, err := client.Download()
	if err != nil {
		log.Fatal(err)
	}
	uploadSpeed, err := client.Upload()
	if err != nil {
		log.Fatal(err)
	}
	return downloadSpeed, uploadSpeed
}

func main() {
	downloadSpeed, uploadSpeed := InternetSpeedCheck()
	if downloadSpeed < 600.0 || uploadSpeed < 800.0 {
		// Internet speeds are slow
		now := time.Now().Format(time.RFC1123)
		rssXml := fmt.Sprintf(rssTemplate, now, downloadSpeed, uploadSpeed, now)
		fmt.Println(rssXml)
	}
}
