package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type urlSet struct {
	URLList []url `xml:"url"`
}

type url struct {
	Loc      string `xml:"loc"`
	Lastmod  string `xml:"lastmod"`
	Priority string `xml:"priority"`
}

func main() {
	sitemapPath := os.Args[1]
	formspreeAccount := os.Args[2]
	fmt.Println("Path: " + sitemapPath)

	xmlFile, err := os.Open(sitemapPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	var q urlSet
	err = xml.NewDecoder(xmlFile).Decode(&q)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	for _, episode := range q.URLList {
		sendTestEmail(episode.Loc, formspreeAccount)
	}
}

func sendTestEmail(urlPath string, formspreeAccount string) {
	fmt.Println("Send test email for path: " + urlPath)

	body := strings.NewReader(`name=MrTest&email=mr%40test.com&phone=555666777&inquiry=Hi&send=Send+message`)
	req, err := http.NewRequest("POST", "https://formspree.io/"+formspreeAccount, body)
	if err != nil {
		fmt.Println("Error building request:", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", urlPath)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()
}
