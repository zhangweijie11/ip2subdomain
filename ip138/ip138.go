package ip138

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"ip2subdomain/crawler"
)

func GetSubdomain(ip string) []string {
	var subdomains []string
	collyScraper := crawler.NewCollyScraper()
	collyScraper.Collector.OnHTML("#list li	a", func(e *colly.HTMLElement) {
		// 提取元素的 data 属性值
		subdomain := e.Text
		subdomains = append(subdomains, subdomain)
	})

	paramURL := fmt.Sprintf("https://site.ip138.com/%s/", ip)
	collyScraper.Collector.Visit(paramURL)
	return subdomains
}
