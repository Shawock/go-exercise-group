package matchers

import "encoding/xml"

type (
	item struct {
		XMLName xml.Name `xml:"item"`
		PubDate string   `xml:"pubDate"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	channel struct {
		XMLName xml.Name `xml:"channel"`
		Title   string   `xml:"title"`
		Image   image    `xml:"image"`
		Item    item     `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)
