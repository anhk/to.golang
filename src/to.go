/**
 * to anywhere
 **/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

const (
	cfgFile = "/usr/local/etc/to.json"
)

type Server struct {
	Addr string
	Port int
	User string
}

var serverList map[string]*Server

func myinit() error {
	jsFile, err := os.Open(cfgFile)
	if err != nil {
		return err
	}
	defer jsFile.Close()

	byteValue, _ := ioutil.ReadAll(jsFile)
	json.Unmarshal(byteValue, &serverList)
	return nil
}

func getNames() []string {
	var names []string
	for idx := range serverList {
		names = append(names, idx)
	}

	sort.Strings(names)
	return names
}

func show() {
	fmt.Println("\033[0;32mMy ConfigFile: ", cfgFile, "\033[0m")
	names := getNames()

	for _, idx := range names {
		server := serverList[idx]
		fmt.Printf("\033[0;31m%v\033[0m: %s@%s:%d\n", idx, server.User, server.Addr, server.Port)
	}
}

func main() {
	if err := myinit(); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(-1)
	}

	if len(os.Args) != 2 {
		show()
		os.Exit(-1)
		return
	}

	target := os.Args[1]

	names := getNames()
	for _, idx := range names {
		if strings.HasPrefix(idx, target) {
			target = idx
		}
	}

	if server := serverList[target]; server != nil {
		cmd := exec.Command("ssh", server.User+"@"+server.Addr, "-p", strconv.Itoa(server.Port))
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		cmd.Run()
		return
	}
	show()
	os.Exit(-1)
}
