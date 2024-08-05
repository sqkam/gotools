package gotools

import (
	"archive/tar"
	"github.com/pierrec/lz4"
	"io"
	"os"
	"path/filepath"
)

func TarToLz4(src string, buf io.Writer) error {
	w := lz4.NewWriter(buf)
	defer w.Close()
	return TarTo(src, w)
}
func TarTo(src string, buf io.Writer) error {
	tw := tar.NewWriter(buf)
	defer tw.Close()

	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})

}
