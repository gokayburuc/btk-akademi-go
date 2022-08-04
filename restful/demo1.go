package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Photos struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/photos")
	defer response.Body.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Baglanti Saglandi")
	}

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(bodyBytes)

	// JSON dataset collected
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var photo Photos
	json.Unmarshal(bodyBytes, &photo)
	fmt.Println(photo)

}

func Demo2() {
	photo := Photos{Title: "Yeni Fotoğraf", AlbumID: 1, ID: 1004}
	jsonPhoto, _ := json.Marshal(photo)
	response, err := http.Post("https://jsonplaceholder.typicode.com/photos", "application/json;charset=utf-8", bytes.NewBuffer(jsonPhoto))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

}
