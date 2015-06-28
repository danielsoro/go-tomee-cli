// Package install provide features to download and install profiles
package install

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	progress "github.com/danielsoro/tomee-cli/util/progress"
	zip "github.com/danielsoro/tomee-cli/util/zip"
)

// Install an specif version of a profile in an path
func Install(tomeePath, dist string, version string) error {
	archiveURL, error := fetchArchiveURL(dist, version)

	if error == nil {
		filename, error := downloadArchive(tomeePath, archiveURL)

		if error == nil {
			unzipArchive(tomeePath, filename)

			if error == nil {
				filepath.Walk(strings.TrimSuffix(filename, ".zip"), grantPermition)
			} else {
			}
		} else {
			return error
		}
	} else {
		return error
	}
	return nil
}

// Performs a GET request and return the content in []byte
func getFromURL(url string) ([]byte, error) {
	response, error := http.Get(url)
	var b []byte
	if error == nil {
		b, error = ioutil.ReadAll(response.Body)
		response.Body.Close()
		return b, nil
	} else {
		return nil, error
	}
}

// Find the correct URL mirror for the archive
func fetchArchiveURL(dist string, version string) (string, error) {
	var archiveURL string
	projectPathURL := fmt.Sprintf("/tomee/tomee-%s/apache-tomee-%s-%s.zip", version, version, dist)
	profileMirrorURL := fmt.Sprintf("http://www.apache.org/dyn/closer.cgi/%s", projectPathURL)

	htmlBody, error := getFromURL(profileMirrorURL)

	if error == nil {
		archiveURL = findArchiveURLFromHTMLBody(string(htmlBody), projectPathURL)
		return archiveURL, nil
	} else {
		return "", error
	}
}

// Parses the HTML body seeking the mirror link
func findArchiveURLFromHTMLBody(htmlBody string, projectPathURL string) string {
	// Retrieve the first occurence for the mirror link
	archiveURLRegex := "(https?://)?([0-9a-z.-]+)\\.([a-z.]{2,6})([\\/\\w.-]*)" + projectPathURL
	re := regexp.MustCompile(archiveURLRegex)
	archiveURL := re.FindString(htmlBody)
	return archiveURL
}

// Download from the specified URL to the path informed in the path flag
func downloadArchive(tomeePath string, archiveURL string) (string, error) {
	archiveURLSlice := strings.Split(archiveURL, "/")
	filepath := filepath.Join(tomeePath, archiveURLSlice[len(archiveURLSlice)-1])

	out, err := os.Create(filepath)
	defer out.Close()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		panic(err)
		return "", err
	}

	response, err := http.Get(archiveURL)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		panic(err)
		return "", err
	}

	_, er := io.Copy(out, &progress.ProgressLoader{Reader: response.Body, Length: response.ContentLength, Filepath: archiveURL})
	if er != nil {
		panic(err)
		return "", err
	}

	fmt.Printf("\nDownloaded %s with %d\n", archiveURLSlice[len(archiveURLSlice)-1], response.ContentLength)

	return filepath, nil
}

// Unpack the archive to the specified folder
func unzipArchive(tomeePath string, filename string) {
	zip.Unzip(filename, tomeePath, true)
}

// Scans the bin directory, looking for executable files, and gives execution permission for them.
func grantPermition(path string, f os.FileInfo, err error) error {
	var extension string
	if runtime.GOOS == "windows" {
		extension = ".exe"
	} else {
		extension = ".sh"
	}

	if strings.HasSuffix(path, extension) {
		err := os.Chmod(path, 755)
		if err != nil {
			panic(err)
			return err
		}
	}
	return nil
}
