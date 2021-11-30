/**
 * to anywhere
 **/
package main

import (
	"os"
)

const cfgFile = "/usr/local/etc/to.json"

func main() {
	var serverList ServerList

	if err := serverList.Load(cfgFile); err != nil {
		panic(err)
	}

	if len(os.Args) != 2 {
		serverList.Show()
		return
	}

	target := os.Args[1]
	if server := serverList.Find(target); server != nil {
		server.Run()
		return
	}

	serverList.Show()
}
