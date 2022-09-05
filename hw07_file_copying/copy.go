package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if len(fromPath) == 0 || len(toPath) == 0 {
		return ErrUnsupportedFile
	}

	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := fileStat.Size()
	if fileSize < offset {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 {
		limit = fileSize
	}

	if fileSize < offset+limit {
		limit = fileSize - offset
	}

	if offset != 0 {
		file.Seek(offset, io.SeekStart)
	}

	copyFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer copyFile.Close()

	bar := pb.Start64(limit)
	barReader := bar.NewProxyReader(file)
	defer bar.Finish()

	_, err = io.CopyN(copyFile, barReader, limit)
	if err != nil {
		return err
	}

	return nil
}
