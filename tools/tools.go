package tools

import (
	"archive/tar"
	"errors"
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
	UPX           = "upx"
	Xh            = "xh"
	Yj            = "yj"
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

	if t.Asset.IsBinary {
		err := saveBinary(resp.Body, t.Asset.Destination)
		if err != nil {
			return fmt.Errorf("saveBinary: %w", err)
		}
	} else {
		err := saveFile(resp.Body, t.Asset.Name)
		if err != nil {
			return fmt.Errorf("saveFile: %w", err)
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

	compFile, err := compressedReader(file, t.Asset.Name)
	if err != nil {
		return fmt.Errorf("compressedReader: %w", err)
	}

	tr := tar.NewReader(compFile)

	for {
		hdr, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return fmt.Errorf("tr.Next: %w", err)
		}

		if hdr.Name == t.Asset.WithinArchive {
			err := saveBinary(tr, t.Asset.Destination)
			if err != nil {
				return fmt.Errorf("saveBinary: %w", err)
			}

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
	defer resp.Body.Close()

	loc, err := resp.Location()
	if err != nil {
		return fmt.Errorf("resp.Location: %w", err)
	}

	list := strings.Split(loc.Path, "/")
	if len(list) == 0 {
		return errors.New("strings.Split: empty slice")
	}

	t.Version = list[len(list)-1]

	return nil
}

func (t *Tool) SetAsset() error {
	t.Asset.Destination = path.Join("/usr/local/bin/", t.Name)
	t.Asset.WithinArchive = t.Name

	switch t.Name {
	case Bat:
		t.AssetBat()
	case Bottom:
		t.AssetBottom()
	case Cloudflared:
		t.AssetCloudflared()
	case DockerCompose:
		t.AssetDockerCompose()
	case K9s:
		t.AssetK9s()
	case UPX:
		t.AssetUPX()
	case Xh:
		t.AssetXh()
	case Yj:
		t.AssetYj()
	}

	if t.Asset.Name == "" {
		return errors.New("empty asset name")
	}

	return nil
}

func (t *Tool) SetRuntime(arch string, os string) error {
	var archOk, osOk bool
	t.Arch, archOk = tables.Arch[t.Name][os][arch]
	t.OS, osOk = tables.OS[t.Name][os][arch]

	if !archOk || !osOk {
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
