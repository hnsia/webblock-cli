package block

import (
	"fmt"

	"github.com/txn2/txeh"
)

func CleanBlocks(hosts *txeh.Hosts, blacklistConfiguredSites []string) (cleanedTargets []string) {
	// stateSites := state.ListSites()
	stateSites := blacklistConfiguredSites

	for _, stateSite := range stateSites {
		// exists := isInList(stateSite.Url, blacklistConfiguredSites)
		// if !exists {
		// hosts.RemoveHost(stateSite.Url)
		hosts.RemoveHost(stateSite)
		err := hosts.Save()
		if err != nil {
			fmt.Printf("Something went wrong trying to unblock the site. %s\n", err)
		}
		// state.Remove(stateSite.Url)
		cleanedTargets = append(cleanedTargets, stateSite)
		// }
	}
	return
}
