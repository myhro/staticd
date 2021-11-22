package tools

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/myhro/staticd/tables"
)

const (
	Bat           = "bat"
	Bottom        = "btm"
	Cloudflared   = "cloudflared"
	DockerCompose = "docker-compose"
	K9s           = "k9s"
	Xh            = "xh"
)

type Asset struct {
	Destination   string
	IsBinary      bool
	Name          string
	WithinArchive string
}

type Tool struct {
	Arch    string
	Asset   Asset
	Name    string
	OS      string
	URL     string
	Version string
}

func (t *Tool) Download() error {
	resp, err := http.Get(fmt.Sprintf("%v/download/%v/%v", t.URL, t.Version, t.Asset.Name))
	if err != nil {
		return fmt.Errorf("http.Get: %w", err)
	}
	defer resp.Body.Close()

	filename := t.Asset.Name
	if t.Asset.IsBinary {
		filename = t.Asset.Destination
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer file.Close()

	io.Copy(file, resp.Body)
	if t.Asset.IsBinary {
		err := file.Chmod(0755)
		if err != nil {
			return fmt.Errorf("file.Chmod: %w", err)
		}
	}

	return nil
}

func (t *Tool) Extract() error {
	file, err := os.Open(t.Asset.Name)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}
	defer os.Remove(t.Asset.Name)
	defer file.Close()

	gzfile, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("gzip.NewReader: %w", err)
	}

	binary, err := os.Create(t.Asset.Destination)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer binary.Close()

	err = binary.Chmod(0755)
	if err != nil {
		return fmt.Errorf("binary.Chmod: %w", err)
	}

	tr := tar.NewReader(gzfile)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("tr.Next: %w", err)
		}

		if hdr.Name == t.Asset.WithinArchive {
			io.Copy(binary, tr)
			break
		}
	}

	return nil
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

func (t *Tool) SetAsset() error {
	t.Asset.Destination = path.Join("/usr/local/bin/", t.Name)
	t.Asset.WithinArchive = t.Name

	switch t.Name {
	case Bat:
		baseName := fmt.Sprintf("bat-%v-%v-%v", t.Version, t.Arch, t.OS)
		t.Asset.Name = baseName + ".tar.gz"
		t.Asset.WithinArchive = path.Join(baseName, t.Name)
	case Bottom:
		t.Asset.Name = fmt.Sprintf("bottom_%v-%v.tar.gz", t.Arch, t.OS)
	case Cloudflared:
		t.Asset.Name = fmt.Sprintf("cloudflared-%v-%v", t.OS, t.Arch)
		t.Asset.IsBinary = true
		if t.OS == "darwin" {
			t.Asset.Name += ".tgz"
			t.Asset.IsBinary = false
		}
	case DockerCompose:
		t.Asset.Name = fmt.Sprintf("docker-compose-%v-%v", t.OS, t.Arch)
		t.Asset.Destination = path.Join("/usr/libexec/docker/cli-plugins/", t.Name)
		t.Asset.IsBinary = true
	case K9s:
		t.Asset.Name = fmt.Sprintf("k9s_%v_%v.tar.gz", t.OS, t.Arch)
	case Xh:
		baseName := fmt.Sprintf("xh-%v-%v-%v", t.Version, t.Arch, t.OS)
		t.Asset.Name = baseName + ".tar.gz"
		t.Asset.WithinArchive = path.Join(baseName, t.Name)
	}

	if t.Asset.Name == "" {
		return fmt.Errorf("empty asset name")
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
