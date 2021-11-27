package tools

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ulikunitz/xz"
)

func compressedReader(file *os.File) (io.Reader, error) {
	var comp io.Reader

	var err error

	switch {
	case strings.HasSuffix(file.Name(), ".gz"):
		comp, err = gzip.NewReader(file)
		if err != nil {
			return nil, fmt.Errorf("gzip.NewReader: %w", err)
		}
	case strings.HasSuffix(file.Name(), ".xz"):
		comp, err = xz.NewReader(file)
		if err != nil {
			return nil, fmt.Errorf("xz.NewReader: %w", err)
		}
	default:
		return nil, fmt.Errorf("unknown archive type for %v", file.Name())
	}

	return comp, nil
}

func saveBinary(src io.Reader, dest string) error {
	err := saveFile(src, dest)
	if err != nil {
		return fmt.Errorf("saveFile: %w", err)
	}

	err = os.Chmod(dest, 0755)
	if err != nil {
		return fmt.Errorf("os.Chmod: %w", err)
	}

	return nil
}

func saveFile(src io.Reader, dest string) error {
	file, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, src)
	if !errors.Is(err, io.EOF) && err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}

	return nil
}
