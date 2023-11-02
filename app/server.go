package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
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
}

func (s *Server) String(name string) string {
	return fmt.Sprintf("%-10s\t%v: %s@%s:%d", s.Tag, name, s.User, s.Addr, s.Port)
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

func (sl *ServerList) Load(path string) {
	jsFile, err := os.Open(path)
	Must(err)
	defer func() { _ = jsFile.Close() }()
	byteValue, err := io.ReadAll(jsFile)
	Must(err)
	Must(json.Unmarshal(byteValue, &sl.list))
	sl.sortName()
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
	//fmt.Println("\033[0;32mMy ConfigFile: ", cfgFile, "\033[0m")
	fmt.Println("----------------------------------------------------")
	for _, idx := range sl.names {
		server := sl.list[idx]
		server.String(idx)
	}
}

func (sl *ServerList) Find(target string) *Server {
	for _, idx := range sl.names {
		if strings.HasPrefix(idx, target) {
			return sl.list[idx]
		}
	}
	return nil
}
