package fetcher

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"

	_type "go-spider/type"
)

func Fetch(body io.ReadCloser) (resources _type.Resources) {
	defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println("Fetch err: ", err)
		return
	}
	resources = _type.Resources{}
	resources.Requests = make([]_type.Request, 0)
	resources.Datas = make([]_type.Data, 0)

	doc.Find("[data-role='ershoufang'] [href]").Each(func(i int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			resources.Requests = append(resources.Requests, _type.Request{
				Url:     "https://cd.lianjia.com" + href,
				Fetcher: Fetch,
			})
		}
	})

	doc.Find(".house-lst-page-box  [href]").Each(func(i int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			resources.Requests = append(resources.Requests, _type.Request{
				Url:     "https://cd.lianjia.com" + href,
				Fetcher: Fetch,
			})
		}
	})

	doc.Find(".listContent   li > a ").Each(func(i int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			resources.Requests = append(resources.Requests, _type.Request{
				Url:     href,
				Fetcher: FetchInfo,
			})
		}
	})

	doc.Find(".listContent   li ").Each(func(i int, selection *goquery.Selection) {
		resources.Datas = append(resources.Datas, _type.Data{
			Type:    "name",
			Content: selection.Find(".info > .title > a").Text()+"---"+selection.Find(".xiaoquListItemRight > .xiaoquListItemPrice > .totalPrice > span").Text(),
		})
	})

	return
}

func FetchInfo(body io.ReadCloser) (resources _type.Resources) {
	defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println("Fetch err: ", err)
		return
	}
	resources = _type.Resources{}
	resources.Requests = make([]_type.Request, 0)
	resources.Datas = make([]_type.Data, 0)

	resources.Datas = append(resources.Datas, _type.Data{
		Type:    "info",
		Content: doc.Find(".xiaoquInfo").Text(),
	})

	return resources
}
