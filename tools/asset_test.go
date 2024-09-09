package tools

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AssetTestSuite struct {
	suite.Suite
}

func TestAssetTestSuite(t *testing.T) {
	suite.Run(t, new(AssetTestSuite))
}

func (s *AssetTestSuite) TestDestination() {
	table := []struct {
		name string
		dest string
	}{
		{
			name: Bat,
			dest: "/usr/local/bin/bat",
		},
		{
			name: Bottom,
			dest: "/usr/local/bin/btm",
		},
		{
			name: Cloudflared,
			dest: "/usr/local/bin/cloudflared",
		},
		{
			name: Flyctl,
			dest: "/usr/local/bin/flyctl",
		},
		{
			name: K9s,
			dest: "/usr/local/bin/k9s",
		},
		{
			name: Ripgrep,
			dest: "/usr/local/bin/rg",
		},
		{
			name: Shellcheck,
			dest: "/usr/local/bin/shellcheck",
		},
		{
			name: UPX,
			dest: "/usr/local/bin/upx",
		},
		{
			name: Xh,
			dest: "/usr/local/bin/xh",
		},
		{
			name: Yj,
			dest: "/usr/local/bin/yj",
		},
	}

	for _, tt := range table {
		tool := &Tool{
			User: &UserMock{},
		}
		tool.Name = tt.name
		tool.SetAsset()
		s.Equal(tt.dest, tool.Asset.Destination)
	}
}

func (s *AssetTestSuite) TestIsBinary() {
	table := []struct {
		name    string
		arch    string
		os      string
		version string
		binary  bool
	}{
		{
			name:    Bat,
			arch:    "amd64",
			os:      "linux",
			version: "v0.18.3",
			binary:  false,
		},
		{
			name:    Bottom,
			arch:    "amd64",
			os:      "linux",
			version: "0.6.4",
			binary:  false,
		},
		{
			name:    Cloudflared,
			arch:    "amd64",
			os:      "darwin",
			version: "2021.11.0",
			binary:  false,
		},
		{
			name:    Cloudflared,
			arch:    "amd64",
			os:      "linux",
			version: "2021.11.0",
			binary:  true,
		},
		{
			name:    Flyctl,
			arch:    "amd64",
			os:      "linux",
			version: "v0.0.450",
			binary:  false,
		},
		{
			name:    K9s,
			arch:    "amd64",
			os:      "linux",
			version: "v0.25.4",
			binary:  false,
		},
		{
			name:    Ripgrep,
			arch:    "amd64",
			os:      "linux",
			version: "14.1.0",
			binary:  false,
		},
		{
			name:    Shellcheck,
			arch:    "amd64",
			os:      "linux",
			version: "v0.10.0",
			binary:  false,
		},
		{
			name:    UPX,
			arch:    "amd64",
			os:      "linux",
			version: "v3.96",
			binary:  false,
		},
		{
			name:    Xh,
			arch:    "amd64",
			os:      "linux",
			version: "v0.14.0",
			binary:  false,
		},
		{
			name:    Yj,
			arch:    "amd64",
			os:      "linux",
			version: "v5.0.0",
			binary:  true,
		},
	}

	for _, tt := range table {
		tool := &Tool{
			User: &UserMock{},
		}
		tool.Name = tt.name
		tool.Version = tt.version
		tool.SetRuntime(tt.arch, tt.os)
		tool.SetAsset()
		s.Equal(tt.binary, tool.Asset.IsBinary, "", tool.Name, tt.arch, tt.os)
	}
}

func (s *AssetTestSuite) TestName() {
	table := []struct {
		name     string
		arch     string
		os       string
		version  string
		filename string
	}{
		{
			name:     Bat,
			arch:     "amd64",
			os:       "linux",
			version:  "v0.18.3",
			filename: "bat-v0.18.3-x86_64-unknown-linux-gnu.tar.gz",
		},
		{
			name:     Bottom,
			arch:     "amd64",
			os:       "linux",
			version:  "0.6.4",
			filename: "bottom_x86_64-unknown-linux-gnu.tar.gz",
		},
		{
			name:     Cloudflared,
			arch:     "amd64",
			os:       "linux",
			version:  "2021.11.0",
			filename: "cloudflared-linux-amd64",
		},
		{
			name:     Flyctl,
			arch:     "amd64",
			os:       "linux",
			version:  "v0.0.450",
			filename: "flyctl_0.0.450_Linux_x86_64.tar.gz",
		},
		{
			name:     K9s,
			arch:     "amd64",
			os:       "linux",
			version:  "v0.30.4",
			filename: "k9s_Linux_amd64.tar.gz",
		},
		{
			name:     Kubectx,
			arch:     "arm64",
			os:       "darwin",
			version:  "v0.9.5",
			filename: "kubectx_v0.9.5_darwin_arm64.tar.gz",
		},
		{
			name:     Ripgrep,
			arch:     "amd64",
			os:       "linux",
			version:  "14.1.0",
			filename: "ripgrep-14.1.0-x86_64-unknown-linux-musl.tar.gz",
		},
		{
			name:     Shellcheck,
			arch:     "amd64",
			os:       "linux",
			version:  "v0.10.0",
			filename: "shellcheck-v0.10.0.linux.x86_64.tar.xz",
		},
		{
			name:     UPX,
			arch:     "amd64",
			os:       "linux",
			version:  "v3.96",
			filename: "upx-3.96-amd64_linux.tar.xz",
		},
		{
			name:     Xh,
			arch:     "amd64",
			os:       "linux",
			version:  "v0.14.0",
			filename: "xh-v0.14.0-x86_64-unknown-linux-musl.tar.gz",
		},
		{
			name:     Yj,
			version:  "v5.0.0",
			arch:     "amd64",
			os:       "linux",
			filename: "yj-linux-amd64",
		},
	}

	for _, tt := range table {
		tool := &Tool{
			User: &UserMock{},
		}
		tool.Name = tt.name
		tool.Version = tt.version
		tool.SetRuntime(tt.arch, tt.os)
		tool.SetAsset()
		s.Equal(tt.filename, tool.Asset.Name)
	}
}

func (s *AssetTestSuite) TestWithinArchive() {
	table := []struct {
		name          string
		arch          string
		os            string
		version       string
		withinArchive string
	}{
		{
			name:          Bat,
			arch:          "amd64",
			os:            "linux",
			version:       "v0.18.3",
			withinArchive: "bat-v0.18.3-x86_64-unknown-linux-gnu/bat",
		},
		{
			name:          Bottom,
			arch:          "amd64",
			os:            "linux",
			version:       "0.6.4",
			withinArchive: "btm",
		},
		{
			name:          Cloudflared,
			arch:          "amd64",
			os:            "linux",
			version:       "2021.11.0",
			withinArchive: "cloudflared",
		},
		{
			name:          Flyctl,
			arch:          "amd64",
			os:            "linux",
			version:       "v0.0.450",
			withinArchive: "flyctl",
		},
		{
			name:          K9s,
			arch:          "amd64",
			os:            "linux",
			version:       "v0.25.4",
			withinArchive: "k9s",
		},
		{
			name:          Ripgrep,
			arch:          "amd64",
			os:            "linux",
			version:       "14.1.0",
			withinArchive: "ripgrep-14.1.0-x86_64-unknown-linux-musl/rg",
		},
		{
			name:          Shellcheck,
			arch:          "amd64",
			os:            "linux",
			version:       "v0.10.0",
			withinArchive: "shellcheck-v0.10.0/shellcheck",
		},
		{
			name:          UPX,
			arch:          "amd64",
			os:            "linux",
			version:       "v3.96",
			withinArchive: "upx-3.96-amd64_linux/upx",
		},
		{
			name:          Xh,
			arch:          "amd64",
			os:            "linux",
			version:       "v0.14.0",
			withinArchive: "xh-v0.14.0-x86_64-unknown-linux-musl/xh",
		},
	}

	for _, tt := range table {
		tool := &Tool{
			User: &UserMock{},
		}
		tool.Name = tt.name
		tool.Version = tt.version
		tool.SetRuntime(tt.arch, tt.os)
		tool.SetAsset()
		s.Equal(tt.withinArchive, tool.Asset.WithinArchive)
	}
}
