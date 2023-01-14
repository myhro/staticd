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
		tool *Tool
		dest string
	}{
		{
			tool: &Tool{
				Name: Bat,
			},
			dest: "/usr/local/bin/bat",
		},
		{
			tool: &Tool{
				Name: Bottom,
			},
			dest: "/usr/local/bin/btm",
		},
		{
			tool: &Tool{
				Name: Cloudflared,
			},
			dest: "/usr/local/bin/cloudflared",
		},
		{
			tool: &Tool{
				Name: DockerCompose,
			},
			dest: "/usr/libexec/docker/cli-plugins/docker-compose",
		},
		{
			tool: &Tool{
				Name: Flyctl,
			},
			dest: "/usr/local/bin/flyctl",
		},
		{
			tool: &Tool{
				Name: K9s,
			},
			dest: "/usr/local/bin/k9s",
		},
		{
			tool: &Tool{
				Name: UPX,
			},
			dest: "/usr/local/bin/upx",
		},
		{
			tool: &Tool{
				Name: Xh,
			},
			dest: "/usr/local/bin/xh",
		},
		{
			tool: &Tool{
				Name: Yj,
			},
			dest: "/usr/local/bin/yj",
		},
	}

	for _, tt := range table {
		tt.tool.SetAsset()
		s.Equal(tt.dest, tt.tool.Asset.Destination)
	}
}

func (s *AssetTestSuite) TestIsBinary() {
	table := []struct {
		tool   *Tool
		arch   string
		os     string
		binary bool
	}{
		{
			tool: &Tool{
				Name:    Bat,
				Version: "v0.18.3",
			},
			arch:   "amd64",
			os:     "linux",
			binary: false,
		},
		{
			tool: &Tool{
				Name:    Bottom,
				Version: "0.6.4",
			},
			arch:   "amd64",
			os:     "linux",
			binary: false,
		},
		{
			tool: &Tool{
				Name:    Cloudflared,
				Version: "2021.11.0",
			},
			arch:   "amd64",
			os:     "darwin",
			binary: false,
		},
		{
			tool: &Tool{
				Name:    Cloudflared,
				Version: "2021.11.0",
			},
			arch:   "amd64",
			os:     "linux",
			binary: true,
		},
		{
			tool: &Tool{
				Name:    DockerCompose,
				Version: "v2.1.1",
			},
			arch:   "amd64",
			os:     "linux",
			binary: true,
		},
		{
			tool: &Tool{
				Name:    Flyctl,
				Version: "v0.0.450",
			},
			arch:   "amd64",
			os:     "linux",
			binary: false,
		},
		{
			tool: &Tool{
				Name:    K9s,
				Version: "v0.25.4",
			},
			arch:   "amd64",
			os:     "linux",
			binary: false,
		},
		{
			tool: &Tool{
				Name:    UPX,
				Version: "v3.96",
			},
			arch:   "amd64",
			os:     "linux",
			binary: false,
		},
		{
			tool: &Tool{
				Name:    Xh,
				Version: "v0.14.0",
			},
			arch:   "amd64",
			os:     "linux",
			binary: false,
		},
		{
			tool: &Tool{
				Name:    Yj,
				Version: "v5.0.0",
			},
			arch:   "amd64",
			os:     "linux",
			binary: true,
		},
	}

	for _, tt := range table {
		tt.tool.SetRuntime(tt.arch, tt.os)
		tt.tool.SetAsset()
		s.Equal(tt.binary, tt.tool.Asset.IsBinary, "", tt.tool.Name, tt.arch, tt.os)
	}
}

func (s *AssetTestSuite) TestName() {
	table := []struct {
		tool *Tool
		arch string
		os   string
		name string
	}{
		{
			tool: &Tool{
				Name:    Bat,
				Version: "v0.18.3",
			},
			arch: "amd64",
			os:   "linux",
			name: "bat-v0.18.3-x86_64-unknown-linux-gnu.tar.gz",
		},
		{
			tool: &Tool{
				Name:    Bottom,
				Version: "0.6.4",
			},
			arch: "amd64",
			os:   "linux",
			name: "bottom_x86_64-unknown-linux-gnu.tar.gz",
		},
		{
			tool: &Tool{
				Name:    Cloudflared,
				Version: "2021.11.0",
			},
			arch: "amd64",
			os:   "linux",
			name: "cloudflared-linux-amd64",
		},
		{
			tool: &Tool{
				Name:    DockerCompose,
				Version: "v2.1.1",
			},
			arch: "amd64",
			os:   "linux",
			name: "docker-compose-linux-x86_64",
		},
		{
			tool: &Tool{
				Name:    Flyctl,
				Version: "v0.0.450",
			},
			arch: "amd64",
			os:   "linux",
			name: "flyctl_0.0.450_Linux_x86_64.tar.gz",
		},
		{
			tool: &Tool{
				Name:    K9s,
				Version: "v0.25.4",
			},
			arch: "amd64",
			os:   "linux",
			name: "k9s_Linux_x86_64.tar.gz",
		},
		{
			tool: &Tool{
				Name:    UPX,
				Version: "v3.96",
			},
			arch: "amd64",
			os:   "linux",
			name: "upx-3.96-amd64_linux.tar.xz",
		},
		{
			tool: &Tool{
				Name:    Xh,
				Version: "v0.14.0",
			},
			arch: "amd64",
			os:   "linux",
			name: "xh-v0.14.0-x86_64-unknown-linux-musl.tar.gz",
		},
		{
			tool: &Tool{
				Name:    Yj,
				Version: "v5.0.0",
			},
			arch: "amd64",
			os:   "linux",
			name: "yj-linux",
		},
		{
			tool: &Tool{
				Name:    Yj,
				Version: "v5.0.0",
			},
			arch: "arm",
			os:   "linux",
			name: "yj-linux-arm-v7",
		},
	}

	for _, tt := range table {
		tt.tool.SetRuntime(tt.arch, tt.os)
		tt.tool.SetAsset()
		s.Equal(tt.name, tt.tool.Asset.Name)
	}
}

func (s *AssetTestSuite) TestWithinArchive() {
	table := []struct {
		tool          *Tool
		arch          string
		os            string
		withinArchive string
	}{
		{
			tool: &Tool{
				Name:    Bat,
				Version: "v0.18.3",
			},
			arch:          "amd64",
			os:            "linux",
			withinArchive: "bat-v0.18.3-x86_64-unknown-linux-gnu/bat",
		},
		{
			tool: &Tool{
				Name:    Bottom,
				Version: "0.6.4",
			},
			arch:          "amd64",
			os:            "linux",
			withinArchive: "btm",
		},
		{
			tool: &Tool{
				Name:    Cloudflared,
				Version: "2021.11.0",
			},
			arch:          "amd64",
			os:            "linux",
			withinArchive: "cloudflared",
		},
		{
			tool: &Tool{
				Name:    DockerCompose,
				Version: "v2.1.1",
			},
			arch:          "amd64",
			os:            "linux",
			withinArchive: "docker-compose",
		},
		{
			tool: &Tool{
				Name:    Flyctl,
				Version: "v0.0.450",
			},
			arch:          "amd64",
			os:            "linux",
			withinArchive: "flyctl",
		},
		{
			tool: &Tool{
				Name:    K9s,
				Version: "v0.25.4",
			},
			arch:          "amd64",
			os:            "linux",
			withinArchive: "k9s",
		},
		{
			tool: &Tool{
				Name:    UPX,
				Version: "v3.96",
			},
			arch:          "amd64",
			os:            "linux",
			withinArchive: "upx-3.96-amd64_linux/upx",
		},
		{
			tool: &Tool{
				Name:    Xh,
				Version: "v0.14.0",
			},
			arch:          "amd64",
			os:            "linux",
			withinArchive: "xh-v0.14.0-x86_64-unknown-linux-musl/xh",
		},
	}

	for _, tt := range table {
		tt.tool.SetRuntime(tt.arch, tt.os)
		tt.tool.SetAsset()
		s.Equal(tt.withinArchive, tt.tool.Asset.WithinArchive)
	}
}
