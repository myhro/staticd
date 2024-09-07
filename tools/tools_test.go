package tools

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ToolsTestSuite struct {
	suite.Suite
}

func TestToolsTestSuite(t *testing.T) {
	suite.Run(t, new(ToolsTestSuite))
}

func (s *ToolsTestSuite) TestDownload() {
	s.Run("shouldn't download anything on 404 errors", func() {
		ts := httptest.NewServer(http.NotFoundHandler())

		tool := Tool{}
		tool.URL = ts.URL
		err := tool.Download()

		s.Error(err)
		s.Contains(err.Error(), "404 Not Found")
	})
}

func (s *ToolsTestSuite) TestSetURL() {
	s.Run("should set URL correctly for known tools", func() {
		testCases := []struct {
			name string
			url  string
		}{
			{
				name: Bat,
				url:  "https://github.com/sharkdp/bat/releases",
			},
			{
				name: Bottom,
				url:  "https://github.com/ClementTsang/bottom/releases",
			},
			{
				name: Cloudflared,
				url:  "https://github.com/cloudflare/cloudflared/releases",
			},
			{
				name: Flyctl,
				url:  "https://github.com/superfly/flyctl/releases",
			},
			{
				name: K9s,
				url:  "https://github.com/derailed/k9s/releases",
			},
			{
				name: Kubectx,
				url:  "https://github.com/ahmetb/kubectx/releases",
			},
			{
				name: UPX,
				url:  "https://github.com/upx/upx/releases",
			},
			{
				name: Xh,
				url:  "https://github.com/ducaale/xh/releases",
			},
			{
				name: Yj,
				url:  "https://github.com/sclevine/yj/releases",
			},
		}

		for _, tc := range testCases {
			tool := Tool{Name: tc.name}
			err := tool.SetURL()
			s.NoError(err)
			s.Equal(tc.url, tool.URL)
		}
	})

	s.Run("should return error for unknown tool", func() {
		tool := Tool{
			Name: "unknown_tool",
		}
		err := tool.SetURL()
		s.Error(err)
		s.Contains(err.Error(), "no url defined for: unknown_tool")
	})
}
