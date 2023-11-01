package app

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type Server struct {
	Addr     string `json:"addr,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Tag      string `json:"tag,omitempty"`
	Jumper   string `json:"jumper,omitempty"`
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
