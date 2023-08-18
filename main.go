package main

import (
	"fmt"
	"github.com/forease/gotld"
	"ip2subdomain/dnsgrep"
	"ip2subdomain/ip138"
	"ip2subdomain/rapiddns"
	"ip2subdomain/utils"
	"ip2subdomain/webscan"
)

type SubdomainInfo struct {
	Subdomain string
	Domain    string
}

func GetSubdomain(ip string) []SubdomainInfo {
	var allSubdomain []string
	var allSubdomainInfo []SubdomainInfo
	subdomainSources := []string{utils.SourceDnsgrep, utils.SourceIp138, utils.SourceRapiddns, utils.SourceWebscan}
	for _, source := range subdomainSources {
		if source == utils.SourceDnsgrep {
			dnsGrewpSubdomains := dnsgrep.GetSubdomain(ip)
			for _, subdomain := range dnsGrewpSubdomains {
				allSubdomain = append(allSubdomain, subdomain)
			}
		}
		if source == utils.SourceIp138 {
			ip138Subdomains := ip138.GetSubdomain(ip)
			for _, subdomain := range ip138Subdomains {
				allSubdomain = append(allSubdomain, subdomain)
			}
		}
		if source == utils.SourceRapiddns {
			rapiddnsSubdomains := rapiddns.GetSubdomain(ip)
			for _, subdomain := range rapiddnsSubdomains {
				allSubdomain = append(allSubdomain, subdomain)
			}
		}
		if source == utils.SourceWebscan {
			webscanSubdomains := webscan.GetSubdomain(ip)
			for _, subdomain := range webscanSubdomains {
				allSubdomain = append(allSubdomain, subdomain)
			}
		}
	}
	// 去重
	newAllSubdomain := utils.RemoveDuplicates(allSubdomain)
	for _, subdomain := range newAllSubdomain {
		_, domain, err := gotld.GetTld(subdomain)
		if err == nil {
			subdomainInfo := SubdomainInfo{
				Subdomain: subdomain,
				Domain:    domain,
			}
			allSubdomainInfo = append(allSubdomainInfo, subdomainInfo)
		}
	}

	return allSubdomainInfo
}

func main() {
	subdomainInfo := GetSubdomain("1.1.1.1")
	fmt.Println("------------>", subdomainInfo)
}
