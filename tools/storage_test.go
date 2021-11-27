package tools

import (
	"bytes"
	"compress/gzip"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ulikunitz/xz"
)

type StorageTestSuite struct {
	suite.Suite
}

func TestStorageTestSuite(t *testing.T) {
	suite.Run(t, new(StorageTestSuite))
}

func (s *StorageTestSuite) TestCompressedReaderGZIP() {
	name := "bat-v0.18.3-x86_64-unknown-linux-gnu.tar.gz"

	buf := &bytes.Buffer{}
	zw := gzip.NewWriter(buf)
	_, err := zw.Write([]byte(name))
	s.Nil(err)
	err = zw.Close()
	s.Nil(err)

	comp, err := compressedReader(buf, name)
	s.Nil(err)
	_, ok := comp.(*gzip.Reader)
	s.Equal(true, ok)
}

func (s *StorageTestSuite) TestCompressedReaderUnknown() {
	_, err := compressedReader(&bytes.Buffer{}, "file.tar.bz2")
	s.Error(err)
}

func (s *StorageTestSuite) TestCompressedReaderXZ() {
	name := "upx-3.96-amd64_linux.tar.xz"

	buf := &bytes.Buffer{}
	zw, err := xz.NewWriter(buf)
	s.Nil(err)
	_, err = zw.Write([]byte(name))
	s.Nil(err)
	err = zw.Close()
	s.Nil(err)

	comp, err := compressedReader(buf, name)
	s.Nil(err)
	_, ok := comp.(*xz.Reader)
	s.Equal(true, ok)
}
