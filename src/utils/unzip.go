package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type zipper struct {
	source      string
	destination string
}

func NewZipper(source, destination string) *zipper {
	return &zipper{
		source:      source,
		destination: destination,
	}
}

func (zipper *zipper) UnzipSource() error {
	// Source: https://gosamples.dev/unzip-file/

	logger.Infof("Extraction of %s started.", zipper.source)

	// 1. Open the zip file
	reader, err := zip.OpenReader(zipper.source)
	if err != nil {
		return err
	}
	defer func(reader *zip.ReadCloser) {
		err := reader.Close()
		if err != nil {
			logger.Fatalf("An error occurred while closing zip file reader: %s", err)
		}
	}(reader)

	// 2. Get the absolute destination path
	zipper.destination, err = filepath.Abs(zipper.destination)
	if err != nil {
		return err
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		if f.Name == "chromedriver" {
			err := unzipFile(f, zipper.destination)
			if err != nil {
				return err
			}
		}
	}

	logger.Infof("Extraction of %s finished.", zipper.source)

	return nil
}

func unzipFile(file *zip.File, destination string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	filePath, err := sanitizeArchivePath(destination, file.Name)
	if err != nil {
		logger.Fatal(err)
	}

	// 5. Create directory tree
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			logger.Fatalf("An error occurred while creating directory tree: %s", err)
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		logger.Fatalf("An error occurred while creating directory tree: %s", err)
		return err
	}
	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		logger.Fatalf("An error occurred while creating a destination file: %s", err)
		return err
	}
	defer func(destinationFile *os.File) {
		err := destinationFile.Close()
		if err != nil {
			logger.Fatalf("An error occurred while closing destination file reader: %s", err)
		}
	}(destinationFile)

	logger.Infof("Extracting Chromedriver binary to: %s.", filePath)

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer func(zippedFile io.ReadCloser) {
		err := zippedFile.Close()
		if err != nil {
			logger.Fatalf("An error occurred while closing zip file reader: %s", err)
		}
	}(zippedFile)

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}

	return nil
}

func sanitizeArchivePath(destination, file string) (path string, err error) {
	path = filepath.Join(destination, file)
	if strings.HasPrefix(path, filepath.Clean(destination)) {
		return path, nil
	}

	return "", fmt.Errorf("%s: %s", "content filepath is tainted", file)
}
