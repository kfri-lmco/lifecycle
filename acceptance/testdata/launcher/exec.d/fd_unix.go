//go:build unix

package main

import "os"

func outputFile() (*os.File, error) {
	return os.NewFile(3, "outputFile"), nil
}
