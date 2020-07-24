package tiktok

type VideoProps struct {
	Url []string `json:"urls"`
}

type Data struct {
	Video VideoProps `json:"video"`
	Image []string `json:"covers"`
	Text string `json:"text"`
}
type ItemInfos struct {
	ItemInfos Data `json:"itemInfos"`
}

type VideoData struct {
	VideoData ItemInfos `json:"videoData"`
}

type PageProps struct {
	PageProps VideoData `json:"pageProps"`
}

type Props struct {
	Props PageProps `json:"props"`
}
