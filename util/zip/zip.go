// Package install provide features to download and install profiles
package zip

import (
	"archive/zip"
    "os"
    "io"
    "path/filepath"
)

// Unpack the file to a destine folder, deleteAfterUnzip can be used to exclude the original file after the process
// Adapted from http://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang
func Unzip(src, dest string, deleteAfterUnzip bool) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer func() {
        if err := r.Close(); err != nil {
            panic(err)
        }
    }()

    os.MkdirAll(dest, 0755)

    // Closure to address file descriptors issue with all the deferred .Close() methods
    extractAndWriteFile := func(f *zip.File) error {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer func() {
            if err := rc.Close(); err != nil {
                panic(err)
            }
        }()

        path := filepath.Join(dest, f.Name)

        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer func() {
                if err := f.Close(); err != nil {
                    panic(err)
                }
            }()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
        return nil
    }

    for _, f := range r.File {
        err := extractAndWriteFile(f)
        if err != nil {
            return err
        }
    }
    
    if deleteAfterUnzip {
        err := os.Remove(src)
        
        if err != nil {
            return err
        }
    }

    return nil
}
