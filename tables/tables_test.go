package tables_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/myhro/staticd/tables"
	"github.com/myhro/staticd/tools"
)

type TablesTestSuite struct {
	suite.Suite
}

func TestTablesTestSuite(t *testing.T) {
	suite.Run(t, new(TablesTestSuite))
}

func (s *TablesTestSuite) TestToolsNames() {
	table := []struct {
		local  string
		remote string
	}{
		{
			local:  tables.Bat,
			remote: tools.Bat,
		},
		{
			local:  tables.Bottom,
			remote: tools.Bottom,
		},
		{
			local:  tables.Cloudflared,
			remote: tools.Cloudflared,
		},
		{
			local:  tables.Flyctl,
			remote: tools.Flyctl,
		},
		{
			local:  tables.K9s,
			remote: tools.K9s,
		},
		{
			local:  tables.Kubectx,
			remote: tools.Kubectx,
		},
		{
			local:  tables.UPX,
			remote: tools.UPX,
		},
		{
			local:  tables.Xh,
			remote: tools.Xh,
		},
	}

	for _, tt := range table {
		s.Equal(tt.remote, tt.local)
	}
}
