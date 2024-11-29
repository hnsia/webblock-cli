package utils

import "fmt"

func AddWww(sites []string) (updatedSites []string) {
	updatedSites = sites
	for _, site := range sites {
		prefixedSite := fmt.Sprintf("www.%s", site)
		updatedSites = append(updatedSites, prefixedSite)
	}
	return
}
