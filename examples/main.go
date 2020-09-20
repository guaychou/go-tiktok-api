package main

import (
	"fmt"
	"github.com/guaychou/go-tiktok-api"
	"log"
)

func main() {
t := tiktok.NewTiktok()

// if you want just to get the video props (e.g: video link, image preview, and caption ) use GetVideoProperties Function
data, err := t.GetVideoProperties("https://vt.tiktok.com/UwXGbG/")
if err != nil {
	log.Fatal(err)
}
fmt.Println(&data)
fmt.Println(data.Text, data.VideoURL, data.ImageURL)

// if you want to download the video use Download function
	t.Download("https://vt.tiktok.com/UKnSvB/")
}