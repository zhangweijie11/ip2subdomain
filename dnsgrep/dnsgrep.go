package dnsgrep

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"ip2subdomain/crawler"
	"strings"
)

func GetSubdomain(ip string) []string {
	var subdomains []string
	collyScraper := crawler.NewCollyScraper()
	collyScraper.Collector.OnHTML("#table > tbody > tr", func(element *colly.HTMLElement) {
		ch := element.DOM.Children()
		subdomain := strings.TrimSpace(ch.Eq(0).Text())
		if len(subdomain) > 3 {
			subdomains = append(subdomains, subdomain)
		}
	})

	paramURL := fmt.Sprintf("https://www.dnsgrep.cn/ip/%s", ip)
	collyScraper.Collector.Visit(paramURL)
	return subdomains
}
