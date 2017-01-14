package main

import (
	"encoding/xml"
	"fmt"
	"os"
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
	argsWithProg := os.Args
	fmt.Println("Path: " + argsWithProg[1])

	xmlFile, err := os.Open(argsWithProg[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	var q urlSet
	errorer := xml.NewDecoder(xmlFile).Decode(&q)
	if errorer != nil {
		fmt.Println("Error opening file:", errorer)
		return
	}

	for _, episode := range q.URLList {
		fmt.Printf(episode.Loc + "\n")
	}
}
