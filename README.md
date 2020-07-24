## Go Tiktok API

### Example of Usage
```cassandraql
package main

import (
	"fmt"
	"github.com/guaychou/go-tiktok-api"
	"log"
)

func main() {
	t := tiktok.NewTiktok()
	data,err:=t.GetVideo("https://vt.tiktok.com/UwXGbG/")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(data.Text,data.VideoURL,data.ImageURL)
}
```