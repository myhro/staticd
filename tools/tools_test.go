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
