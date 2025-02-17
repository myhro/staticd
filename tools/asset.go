package tools

import (
	"fmt"
	"path"
)

func (t *Tool) AssetAsdf() {
	t.Asset.Name = fmt.Sprintf("asdf-%v-%v-%v.tar.gz", t.Version, t.OS, t.Arch)
}

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

func (t *Tool) AssetFlyctl() {
	t.Asset.Name = fmt.Sprintf("flyctl_%v_%v_%v.tar.gz", t.TrimVersion(), t.OS, t.Arch)
}

func (t *Tool) AssetHugo() {
	t.Asset.Name = fmt.Sprintf("hugo_%v_%v-%v.tar.gz", t.TrimVersion(), t.OS, t.Arch)
}

func (t *Tool) AssetK9s() {
	t.Asset.Name = fmt.Sprintf("k9s_%v_%v.tar.gz", t.OS, t.Arch)
}

func (t *Tool) AssetKubectx() {
	t.Asset.Name = fmt.Sprintf("kubectx_%v_%v_%v.tar.gz", t.Version, t.OS, t.Arch)
}

func (t *Tool) AssetRipgrep() {
	baseName := fmt.Sprintf("ripgrep-%v-%v-%v", t.Version, t.Arch, t.OS)
	t.Asset.Name = baseName + ".tar.gz"
	t.Asset.WithinArchive = path.Join(baseName, t.Name)
}

func (t *Tool) AssetShellcheck() {
	baseName := fmt.Sprintf("shellcheck-%v.%v.%v", t.Version, t.OS, t.Arch)
	t.Asset.Name = baseName + ".tar.xz"

	folderName := fmt.Sprintf("shellcheck-%v", t.Version)
	t.Asset.WithinArchive = path.Join(folderName, t.Name)
}

func (t *Tool) AssetUPX() {
	baseName := fmt.Sprintf("upx-%v-%v_%v", t.TrimVersion(), t.Arch, t.OS)
	t.Asset.Name = baseName + ".tar.xz"
	t.Asset.WithinArchive = path.Join(baseName, t.Name)
}

func (t *Tool) AssetUv() {
	baseName := fmt.Sprintf("uv-%v-%v", t.Arch, t.OS)
	t.Asset.Name = baseName + ".tar.gz"
	t.Asset.WithinArchive = path.Join(baseName, t.Name)
}

func (t *Tool) AssetXh() {
	baseName := fmt.Sprintf("xh-%v-%v-%v", t.Version, t.Arch, t.OS)
	t.Asset.Name = baseName + ".tar.gz"
	t.Asset.WithinArchive = path.Join(baseName, t.Name)
}

func (t *Tool) AssetYj() {
	t.Asset.IsBinary = true
	t.Asset.Name = fmt.Sprintf("yj-%v-%v", t.OS, t.Arch)
}
