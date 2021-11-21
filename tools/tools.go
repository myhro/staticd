package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/myhro/staticd/tables"
)

const (
	Bat           = "bat"
	Bottom        = "bottom"
	Cloudflared   = "cloudflared"
	DockerCompose = "docker-compose"
	K9s           = "k9s"
	Xh            = "xh"
)

type Tool struct {
	Arch    string
	Archive string
	Folder  string
	Name    string
	OS      string
	URL     string
	Version string
}

func (t *Tool) Download() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%v/download/%v/%v", t.URL, t.Version, t.Archive))
	if err != nil {
		return "", fmt.Errorf("http.Get: %w", err)
	}
	defer resp.Body.Close()

	filename := t.Archive
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("os.Create: %w", err)
	}
	defer file.Close()
	io.Copy(file, resp.Body)

	return filename, nil
}

func (t *Tool) GetVersion() error {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Head(t.URL + "/latest")
	if err != nil {
		return fmt.Errorf("client.Head: %w", err)
	}

	loc, err := resp.Location()
	if err != nil {
		return fmt.Errorf("resp.Location: %w", err)
	}

	list := strings.Split(loc.Path, "/")
	if len(list) == 0 {
		return fmt.Errorf("strings.Split: empty slice")
	}
	t.Version = list[len(list)-1]

	return nil
}

func (t *Tool) SetArchive() error {
	switch t.Name {
	case Bat:
		t.Folder = fmt.Sprintf("bat-%v-%v-%v", t.Version, t.Arch, t.OS)
		t.Archive = t.Folder + ".tar.gz"
	case Bottom:
		t.Archive = fmt.Sprintf("bottom_%v-%v.tar.gz", t.Arch, t.OS)
	case Cloudflared:
		t.Archive = fmt.Sprintf("cloudflared-%v-%v", t.OS, t.Arch)
		if t.OS == "darwin" {
			t.Archive += ".tgz"
		}
	case DockerCompose:
		t.Archive = fmt.Sprintf("docker-compose-%v-%v", t.OS, t.Arch)
	case K9s:
		t.Archive = fmt.Sprintf("k9s_%v_%v.tar.gz", t.OS, t.Arch)
	case Xh:
		t.Folder = fmt.Sprintf("xh-%v-%v-%v", t.Version, t.Arch, t.OS)
		t.Archive = t.Folder + ".tar.gz"
	}

	if t.Archive == "" {
		return fmt.Errorf("empty archive name")
	}

	return nil
}

func (t *Tool) SetRuntime(arch string, os string) error {
	t.Arch = tables.Arch[t.Name][os][arch]
	t.OS = tables.OS[t.Name][os][arch]

	if t.Arch == "" || t.OS == "" {
		return fmt.Errorf("no candidate for %v on %v/%v", t.Name, os, arch)
	}

	return nil
}

func (t *Tool) SetURL() error {
	t.URL = tables.URL[t.Name]
	if t.URL == "" {
		return fmt.Errorf("no url defined for: %v", t.Name)
	}
	return nil
}
