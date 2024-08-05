package gotools

import (
	"io"
	"os"
	"testing"
)

func Test_TarToLz4(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		path   string
		target string
	}{
		{
			path:   "./dir",
			target: "dir.tar.lz4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			pr, pw, _ := os.Pipe()

			errCh := make(chan error)
			go func() {
				err = TarToLz4(tt.path, pw)
				pw.Close()
				errCh <- err
			}()

			file, err := os.OpenFile(tt.target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				return
			}
			defer file.Close()
			_, err = io.Copy(file, pr)
			if err != nil {
				//return err
			}
			subErr := <-errCh
			if subErr != nil {
				//return subErr
			}

		})
	}
}
func Test_TarTo(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		path   string
		target string
	}{
		{
			path:   "./dir",
			target: "dir.tar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			pr, pw, _ := os.Pipe()

			errCh := make(chan error)
			go func() {
				err = TarTo(tt.path, pw)
				pw.Close()
				errCh <- err
			}()

			file, err := os.OpenFile(tt.target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				return
			}
			defer file.Close()
			_, err = io.Copy(file, pr)
			if err != nil {
				//return err
			}
			subErr := <-errCh
			if subErr != nil {
				//return subErr
			}

		})
	}
}
