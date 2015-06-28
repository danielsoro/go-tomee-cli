// Package responsable for tracking the percentual value of IO operations
// Based on http://stackoverflow.com/questions/22421375/how-to-print-the-bytes-while-the-file-is-being-downloaded-golang
package progress

import (
    "fmt"
    "io"
)

// ProgressLoader wraps an existing io.Reader.
//
// It simply forwards the Read() call, while displaying
// the results from individual calls to it.
type ProgressLoader struct {
    io.Reader
    total    int64 // Total # of bytes transferred
    progress float64
    Length   int64 // Expected length
    Filepath string
}

// Read 'overrides' the underlying io.Reader's Read method.
// This is the one that will be called by io.Copy(). We simply
// use it to keep track of byte counts and then forward the call.
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
