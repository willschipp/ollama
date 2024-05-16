package registry

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var catalog string = "/_catalog"
var tags string = "/tags/list"

type Images struct {
	Repositories []string
}

type Image struct {
	Name string
	Tags []string
}

func getImages(url string) []string {
	response, err := http.Get(url + catalog)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Images
	json.Unmarshal(body, &data)

	return data.Repositories
}

func getTags(url string, name string) []string {
	full_url := url + name + tags
	log.Println(full_url)
	response, err := http.Get(full_url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(body)

	var data Image
	json.Unmarshal(body, &data)

	log.Println(data)

	// return data.Repositories
	return data.Tags

}

func GetImages() []string {
	url := "http://localhost:5000/v2/"
	response := getImages(url)
	return response
}

func GetTags(name string) []string {
	url := "http://localhost:5000/v2/"
	response := getTags(url, name)
	return response
}
