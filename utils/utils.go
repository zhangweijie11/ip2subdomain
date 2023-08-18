package utils

import "log"

var Logger *log.Logger

const (
	SourceDnsgrep  = "dnsgrep"
	SourceIp138    = "ip138"
	SourceRapiddns = "rapiddns"
	SourceWebscan  = "webscan"
)

// RemoveDuplicates 去重
func RemoveDuplicates(slice []string) []string {
	var result []string

	for _, item := range slice {
		found := false
		for _, val := range result {
			if val == item {
				found = true
				break
			}
		}
		if !found {
			result = append(result, item)
		}
	}

	return result
}
