package main

import (
	"fmt"
	"os"
	"os/user"
)

func printGroupsInfo(gids []string) {
	for _, gid := range gids {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			continue
		}
		fmt.Printf("%-8s - %-10s\n", group.Gid, group.Name)
	}
}

func main() {
	arguments := os.Args
	var username string
	if len(arguments) == 1 {
		usern, err := user.Current()
		username = usern.Name
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		username = arguments[1]
	}
	u, err := user.Lookup(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	gids, err := u.GroupIds()
	if err != nil {
		fmt.Println(err)
		return
	}
	printGroupsInfo(gids)
}
