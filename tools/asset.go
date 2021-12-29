package tools

import (
	"fmt"
	"path"
	"strings"
)

func (t *Tool) AssetBat() {
	baseName := fmt.Sprintf("bat-%v-%v-%v", t.Version, t.Arch, t.OS)
	t.Asset.Name = baseName + ".tar.gz"
	t.Asset.WithinArchive = path.Join(baseName, t.Name)
}

func (t *Tool) AssetBottom() {
	t.Asset.Name = fmt.Sprintf("bottom_%v-%v.tar.gz", t.Arch, t.OS)
}

func (t *Tool) AssetCloudflared() {
	t.Asset.Name = fmt.Sprintf("cloudflared-%v-%v", t.OS, t.Arch)
	t.Asset.IsBinary = true

	if t.OS == "darwin" {
		t.Asset.Name += ".tgz"
		t.Asset.IsBinary = false
	}
}

func (t *Tool) AssetDockerCompose() {
	t.Asset.Name = fmt.Sprintf("docker-compose-%v-%v", t.OS, t.Arch)
	t.Asset.Destination = path.Join("/usr/libexec/docker/cli-plugins/", t.Name)
	t.Asset.IsBinary = true
}

func (t *Tool) AssetK9s() {
	t.Asset.Name = fmt.Sprintf("k9s_%v_%v.tar.gz", t.OS, t.Arch)
}

func (t *Tool) AssetUPX() {
	version := strings.TrimPrefix(t.Version, "v")
	baseName := fmt.Sprintf("upx-%v-%v_%v", version, t.Arch, t.OS)
	t.Asset.Name = baseName + ".tar.xz"
	t.Asset.WithinArchive = path.Join(baseName, t.Name)
}

func (t *Tool) AssetXh() {
	baseName := fmt.Sprintf("xh-%v-%v-%v", t.Version, t.Arch, t.OS)
	t.Asset.Name = baseName + ".tar.gz"
	t.Asset.WithinArchive = path.Join(baseName, t.Name)
}
