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
		{
			arch:    "arm",
			os:      "linux",
			archOut: "arm",
			osOut:   "unknown-linux-gnueabihf",
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
		{
			arch:    "arm",
			os:      "linux",
			archOut: "armv7",
			osOut:   "unknown-linux-gnueabihf",
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
		{
			arch:    "arm",
			os:      "linux",
			archOut: "arm",
			osOut:   "linux",
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

func (s *RuntimeTestSuite) TestDockerComposeRuntime() {
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
			arch:    "arm",
			os:      "linux",
			archOut: "armv7",
			osOut:   "linux",
		},
	}

	for _, tt := range table {
		t := &Tool{
			Name: DockerCompose,
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
	err := t.SetRuntime("arm", "darwin")
	s.Error(err)
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
			archOut: "x86_64",
			osOut:   "Linux",
		},
		{
			arch:    "amd64",
			os:      "darwin",
			archOut: "x86_64",
			osOut:   "Darwin",
		},
		{
			arch:    "arm",
			os:      "linux",
			archOut: "arm",
			osOut:   "Linux",
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
		{
			arch:    "arm",
			os:      "linux",
			archOut: "arm",
			osOut:   "unknown-linux-gnueabihf",
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
