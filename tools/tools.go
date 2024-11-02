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
)

const (
	Bat         = "bat"
	Bottom      = "btm"
	Cloudflared = "cloudflared"
	Flyctl      = "flyctl"
	K9s         = "k9s"
	Kubectx     = "kubectx"
	Ripgrep     = "rg"
	Shellcheck  = "shellcheck"
	UPX         = "upx"
	Uv          = "uv"
	Xh          = "xh"
	Yj          = "yj"
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
	User    User
	Version string
}

func (t *Tool) BaseDir() string {
	if t.User.IsRoot() {
		return "/usr/local/bin/"
	}

	return path.Join(t.User.HomeDir(), ".local", "bin")
}

func (t *Tool) Download() error {
	resp, err := http.Get(fmt.Sprintf("%v/download/%v/%v", t.URL, t.Version, t.Asset.Name))
	if err != nil {
		return fmt.Errorf("http.Get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("couldn't download %v: %v", t.Asset.Name, resp.Status)
	}

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

	found := false

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

			err = os.Chtimes(t.Asset.Destination, hdr.AccessTime, hdr.ModTime)
			if err != nil {
				return fmt.Errorf("os.Chtimes: %w", err)
			}

			found = true

			break
		}
	}

	if !found {
		return fmt.Errorf("binary not found in archive: %v", t.Asset.WithinArchive)
	}

	return nil
}

func (t *Tool) GetVersion() error {
	if t.Version != "latest" {
		return nil
	}

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

//nolint:cyclop
func (t *Tool) SetAsset() error {
	t.Asset.Destination = path.Join(t.BaseDir(), t.Name)
	t.Asset.WithinArchive = t.Name

	switch t.Name {
	case Bat:
		t.AssetBat()
	case Bottom:
		t.AssetBottom()
	case Cloudflared:
		t.AssetCloudflared()
	case Flyctl:
		t.AssetFlyctl()
	case K9s:
		t.AssetK9s()
	case Kubectx:
		t.AssetKubectx()
	case Ripgrep:
		t.AssetRipgrep()
	case Shellcheck:
		t.AssetShellcheck()
	case UPX:
		t.AssetUPX()
	case Uv:
		t.AssetUv()
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
	t.Arch, archOk = Arch[t.Name][os][arch]
	t.OS, osOk = OS[t.Name][os][arch]

	if !archOk || !osOk {
		if arch == "arm64" && os == "darwin" {
			return t.SetRuntime("amd64", os)
		}

		return fmt.Errorf("no candidate for %v on %v/%v", t.Name, os, arch)
	}

	return nil
}

func (t *Tool) SetURL() error {
	t.URL = URL[t.Name]
	if t.URL == "" {
		return fmt.Errorf("no url defined for: %v", t.Name)
	}

	t.URL += "/releases"

	return nil
}

func (t *Tool) TrimVersion() string {
	return strings.TrimPrefix(t.Version, "v")
}
