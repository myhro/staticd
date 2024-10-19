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

func compressedReader(reader io.Reader, name string) (io.Reader, error) {
	switch {
	case strings.HasSuffix(name, ".gz") || strings.HasSuffix(name, ".tgz"):
		comp, err := gzip.NewReader(reader)
		if err != nil {
			return nil, fmt.Errorf("gzip.NewReader: %w", err)
		}

		return comp, nil
	case strings.HasSuffix(name, ".xz"):
		comp, err := xz.NewReader(reader)
		if err != nil {
			return nil, fmt.Errorf("xz.NewReader: %w", err)
		}

		return comp, nil
	}

	return nil, fmt.Errorf("unknown archive type for %v", name)
}

func saveBinary(src io.Reader, dest string) error {
	err := saveFile(src, dest)
	if err != nil {
		return fmt.Errorf("saveFile: %w", err)
	}

	//nolint:mnd
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
