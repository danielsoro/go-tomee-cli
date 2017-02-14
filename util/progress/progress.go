package progress

import (
	"fmt"
	"io"
)

type ProgressLoader struct {
	io.Reader
	total    int64 // Total # of bytes transferred
	progress float64
	Length   int64 // Expected length
	Filepath string
}

func (pt *ProgressLoader) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	if n > 0 {
		pt.total += int64(n)

		totalMb := pt.total / 1048576
		lengthMb := pt.Length / 1048576
		percentage := float64(pt.total) / float64(pt.Length) * float64(100)
		i := int(percentage)

		fmt.Printf("\rDownloading %s [%d/%d Mb] %d%%", pt.Filepath, totalMb, lengthMb, i)
		pt.progress = percentage
	}

	return n, err
}
