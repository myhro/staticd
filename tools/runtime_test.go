package tools

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RuntimeTestSuite struct {
	suite.Suite
}

func TestRuntimeTestSuite(t *testing.T) {
	suite.Run(t, new(RuntimeTestSuite))
}

func (s *RuntimeTestSuite) TestBatRuntime() {
	table := []struct {
		arch    string
		os      string
		archOut string
		osOut   string
	}{
		{
			arch:    "amd64",
			os:      "linux",
			archOut: "x86_64",
			osOut:   "unknown-linux-gnu",
		},
		{
			arch:    "amd64",
			os:      "darwin",
			archOut: "x86_64",
			osOut:   "apple-darwin",
		},
	}

	for _, tt := range table {
		t := &Tool{
			Name: Bat,
		}
		err := t.SetRuntime(tt.arch, tt.os)
		s.Nil(err)
		s.Equal(tt.archOut, t.Arch)
		s.Equal(tt.osOut, t.OS)
	}
}

func (s *RuntimeTestSuite) TestBottomRuntime() {
	table := []struct {
		arch    string
		os      string
		archOut string
		osOut   string
	}{
		{
			arch:    "amd64",
			os:      "linux",
			archOut: "x86_64",
			osOut:   "unknown-linux-gnu",
		},
		{
			arch:    "amd64",
			os:      "darwin",
			archOut: "x86_64",
			osOut:   "apple-darwin",
		},
	}

	for _, tt := range table {
		t := &Tool{
			Name: Bottom,
		}
		err := t.SetRuntime(tt.arch, tt.os)
		s.Nil(err)
		s.Equal(tt.archOut, t.Arch)
		s.Equal(tt.osOut, t.OS)
	}
}

func (s *RuntimeTestSuite) TestCloudflaredRuntime() {
	table := []struct {
		arch    string
		os      string
		archOut string
		osOut   string
	}{
		{
			arch:    "amd64",
			os:      "linux",
			archOut: "amd64",
			osOut:   "linux",
		},
		{
			arch:    "amd64",
			os:      "darwin",
			archOut: "amd64",
			osOut:   "darwin",
		},
	}

	for _, tt := range table {
		t := &Tool{
			Name: Cloudflared,
		}
		err := t.SetRuntime(tt.arch, tt.os)
		s.Nil(err)
		s.Equal(tt.archOut, t.Arch)
		s.Equal(tt.osOut, t.OS)
	}
}

func (s *RuntimeTestSuite) TestInvalidRuntime() {
	t := &Tool{
		Name: Bat,
	}
	err := t.SetRuntime("riscv", "darwin")
	s.Error(err)
}

func (s *RuntimeTestSuite) TestFlyctlRuntime() {
	table := []struct {
		arch    string
		os      string
		archOut string
		osOut   string
	}{
		{
			arch:    "amd64",
			os:      "linux",
			archOut: "x86_64",
			osOut:   "Linux",
		},
		{
			arch:    "amd64",
			os:      "darwin",
			archOut: "x86_64",
			osOut:   "macOS",
		},
		{
			arch:    "arm64",
			os:      "linux",
			archOut: "arm64",
			osOut:   "Linux",
		},
		{
			arch:    "arm64",
			os:      "darwin",
			archOut: "arm64",
			osOut:   "macOS",
		},
	}

	for _, tt := range table {
		t := &Tool{
			Name: Flyctl,
		}
		err := t.SetRuntime(tt.arch, tt.os)
		s.Nil(err)
		s.Equal(tt.archOut, t.Arch)
		s.Equal(tt.osOut, t.OS)
	}
}

func (s *RuntimeTestSuite) TestK9sRuntime() {
	table := []struct {
		arch    string
		os      string
		archOut string
		osOut   string
	}{
		{
			arch:    "amd64",
			os:      "linux",
			archOut: "amd64",
			osOut:   "Linux",
		},
		{
			arch:    "amd64",
			os:      "darwin",
			archOut: "amd64",
			osOut:   "Darwin",
		},
	}

	for _, tt := range table {
		t := &Tool{
			Name: K9s,
		}
		err := t.SetRuntime(tt.arch, tt.os)
		s.Nil(err)
		s.Equal(tt.archOut, t.Arch)
		s.Equal(tt.osOut, t.OS)
	}
}

func (s *RuntimeTestSuite) TestRipgrepRuntime() {
	table := []struct {
		arch    string
		os      string
		archOut string
		osOut   string
	}{
		{
			arch:    "amd64",
			os:      "linux",
			archOut: "x86_64",
			osOut:   "unknown-linux-musl",
		},
		{
			arch:    "amd64",
			os:      "darwin",
			archOut: "x86_64",
			osOut:   "apple-darwin",
		},
		{
			arch:    "arm64",
			os:      "linux",
			archOut: "aarch64",
			osOut:   "unknown-linux-gnu",
		},
		{
			arch:    "arm64",
			os:      "darwin",
			archOut: "aarch64",
			osOut:   "apple-darwin",
		},
	}

	for _, tt := range table {
		t := &Tool{
			Name: Ripgrep,
		}
		err := t.SetRuntime(tt.arch, tt.os)
		s.Nil(err)
		s.Equal(tt.archOut, t.Arch)
		s.Equal(tt.osOut, t.OS)
	}
}

func (s *RuntimeTestSuite) TestShellcheckRuntime() {
	table := []struct {
		arch    string
		os      string
		archOut string
		osOut   string
	}{
		{
			arch:    "amd64",
			os:      "linux",
			archOut: "x86_64",
			osOut:   "linux",
		},
		{
			arch:    "amd64",
			os:      "darwin",
			archOut: "x86_64",
			osOut:   "darwin",
		},
		{
			arch:    "arm64",
			os:      "linux",
			archOut: "aarch64",
			osOut:   "linux",
		},
		{
			arch:    "arm64",
			os:      "darwin",
			archOut: "aarch64",
			osOut:   "darwin",
		},
	}

	for _, tt := range table {
		t := &Tool{
			Name: Shellcheck,
		}
		err := t.SetRuntime(tt.arch, tt.os)
		s.Nil(err)
		s.Equal(tt.archOut, t.Arch)
		s.Equal(tt.osOut, t.OS)
	}
}

func (s *RuntimeTestSuite) TestXhRuntime() {
	table := []struct {
		arch    string
		os      string
		archOut string
		osOut   string
	}{
		{
			arch:    "amd64",
			os:      "linux",
			archOut: "x86_64",
			osOut:   "unknown-linux-musl",
		},
		{
			arch:    "amd64",
			os:      "darwin",
			archOut: "x86_64",
			osOut:   "apple-darwin",
		},
	}

	for _, tt := range table {
		t := &Tool{
			Name: Xh,
		}
		err := t.SetRuntime(tt.arch, tt.os)
		s.Nil(err)
		s.Equal(tt.archOut, t.Arch)
		s.Equal(tt.osOut, t.OS)
	}
}
