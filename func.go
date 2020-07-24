package tiktok

import (
	"encoding/json"
	"errors"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
)

type Tiktok struct {
	scraper *colly.Collector
}

func (t *Tiktok) GetVideo(copiedLink string) (*TiktokResultData, error) {
	if isTiktokUrl(copiedLink){
		dataInterface:=Props{}
		var err error
		t.scraper.OnHTML("script[id=__NEXT_DATA__]", func(e *colly.HTMLElement) {
			err=json.Unmarshal([]byte(e.Text),&dataInterface)
		})
		err=t.scraper.Visit(copiedLink)
		data :=TiktokResultData{
			ImageURL: dataInterface.Props.PageProps.VideoData.ItemInfos.Image[0],
			VideoURL: dataInterface.Props.PageProps.VideoData.ItemInfos.Video.Url[0],
			Text: dataInterface.Props.PageProps.VideoData.ItemInfos.Text,
		}
		return &data,err
	}
	return nil , errors.New("Invalid tiktok URL")
}

func NewTiktok() *Tiktok {
	c := colly.NewCollector()
	c.UserAgent="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Accept-Encoding","gzip, deflate")
		r.Headers.Set("Accept-Language","en-US,en;q=0.9")
	})
	t:=&Tiktok{
		scraper: c,
	}
	return t
}

func isTiktokUrl(str string) bool {
	if (strings.Contains(str,"tiktok")){
		u, err := url.Parse(str)
		return err == nil && u.Scheme != "" && u.Host != ""
	}
	return false
}