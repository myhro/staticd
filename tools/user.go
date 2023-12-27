package tools

import "os/user"

type User interface {
	HomeDir() string
	IsRoot() bool
}

type UserOS struct{}

func (u *UserOS) HomeDir() string {
	whoami, err := user.Current()
	if err != nil {
		return "/tmp/"
	}

	return whoami.HomeDir
}

func (u *UserOS) IsRoot() bool {
	whoami, err := user.Current()
	if err != nil {
		return false
	}

	return whoami.Uid == "0"
}
