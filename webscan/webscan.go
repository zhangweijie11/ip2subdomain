package webscan

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"ip2subdomain/crawler"
)

type webscanResponse []struct {
	Domain string `json:"domain"`
	Title  string `json:"title"`
}

func GetSubdomain(ip string) []string {
	var subdomains []string
	var webResponse webscanResponse
	collyScraper := crawler.NewCollyScraper()
	collyScraper.Collector.OnResponse(func(response *colly.Response) {
		err := json.Unmarshal(response.Body, &webResponse)
		if err == nil {
			for _, web := range webResponse {
				subdomains = append(subdomains, web.Domain)
			}
		}
		response.Body = nil
	})

	collyScraper.Collector.OnError(func(response *colly.Response, err error) {
		err = json.Unmarshal(response.Body, &webResponse)
		if err == nil {
			for _, web := range webResponse {
				subdomains = append(subdomains, web.Domain)
			}
		}
		response.Body = nil
	})

	paramURL := fmt.Sprintf("http://api.webscan.cc/?action=query&ip=%s", ip)
	collyScraper.Collector.Visit(paramURL)
	return subdomains
}
