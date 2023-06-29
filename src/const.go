package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Server struct {
	Addr     string `json:"addr,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Tag      string `json:"tag,omitempty"`
	Jumper   string `json:"jumper,omitempty"`

	Regex bool `json:"regex,omitempty"` // 是否使用正则匹配，如果是正则匹配的话，则针对该节点禁用前缀查询
}

func (s *Server) Show(name string) {
	fmt.Printf("\033[0;32m%-10s\033[0m\t\033[0;31m%v\033[0m: %s@%s:%d\n", s.Tag, name, s.User, s.Addr, s.Port)
}

func (s *Server) command() *exec.Cmd {
	var args []string

	if s.Password == "" {
		args = append(args, "ssh")
	} else {
		args = append(args, "sshpass", "-p", s.Password, "ssh")
	}

	if s.Jumper != "" {
		args = append(args, "-J", s.Jumper)
	}

	args = append(args, s.User+"@"+s.Addr, "-p", strconv.Itoa(s.Port))
	//fmt.Println(args)
	return exec.Command(args[0], args[1:]...)
}

func (s *Server) Run() {
	cmd := s.command()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

type ServerList struct {
	list  map[string]*Server
	names []string
}

func (sl *ServerList) Load(path string) error {
	jsFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsFile.Close()

	byteValue, _ := ioutil.ReadAll(jsFile)
	json.Unmarshal(byteValue, &sl.list)

	sl.sortName()
	return nil
}

func (sl *ServerList) sortName() {
	for idx := range sl.list {
		sl.names = append(sl.names, idx)
		if sl.list[idx].Port == 0 {
			sl.list[idx].Port = 22
		}
		if sl.list[idx].User == "" {
			sl.list[idx].User = "root"
		}
	}

	sort.Strings(sl.names)
}

func (sl *ServerList) Show() {
	fmt.Println("\033[0;32mMy ConfigFile: ", cfgFile, "\033[0m")
	fmt.Println("----------------------------------------------------")
	for _, idx := range sl.names {
		server := sl.list[idx]
		server.Show(idx)
	}
}

func (sl *ServerList) Find(target string) *Server {
	for _, idx := range sl.names {
		if sl.list[idx].Regex {
			fmt.Println("idx:", idx)
			if m := regexp.MustCompile(idx); m.MatchString(target) {
				sl.list[idx].Addr = target
				return sl.list[idx]
			}
		} else if strings.HasPrefix(idx, target) {
			return sl.list[idx]
		}
	}
	return nil
}
