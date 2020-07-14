package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
)

var (
	baseURL = "https://mcs-notes2.open.ac.uk/files/"
)

func main() {
	doc, err := htmlquery.LoadURL(baseURL)

	if err != nil {
		fmt.Println(err)
	}

	linkNodes := htmlquery.Find(doc, "//a/@href")

	for i, link := range linkNodes {
		if i > 4 { // ignore the first 5 links on the page as these aren't valid files
			fileName := htmlquery.SelectAttr(link, "href")
			if strings.HasSuffix(fileName, ".mp4") {
				continue
			}
			downloadFile(fileName)
		}
	}
}

func downloadFile(link string) {
	url := fmt.Sprintf("%s%s", baseURL, link)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	filePath := fmt.Sprintf("files/%s", link)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println(err)
	}
}
