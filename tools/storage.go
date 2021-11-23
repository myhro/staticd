package tools

import (
	"errors"
	"fmt"
	"io"
	"os"
)

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
