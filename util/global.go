package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func FullObject(prefix, filepath string) string {
	if len(prefix) == 0 {
		return path.Base(filepath)
	}
	return fmt.Sprintf("%s/%s", strings.Trim(prefix, "/"), path.Base(filepath))
}

func IsFileExist(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsDirectory(filepath string) bool {
	if IsFileExist(filepath) {
		if fi, _ := os.Stat(filepath); fi.IsDir() {
			return true
		}
		return false
	}
	return false
}

func WalkDirectory(dirpath string) (dirlist, filelist []string, err error) {

	if !IsDirectory(dirpath) {
		return nil, nil, errors.New(fmt.Sprintf("%s: not exist or is not a direcitory)", dirpath))
	}

	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, nil, err
	}

	for _, fi := range files {
		// fmt.Printf("k-> %v ; v -> %s\n", k, fi.Name())
		if fi.IsDir() {
			dirlist = append(dirlist, fmt.Sprintf("%s/%s", dirpath, fi.Name()))
		} else {
			filelist = append(filelist, fmt.Sprintf("%s/%s", dirpath, fi.Name()))
		}
	}
	return dirlist, filelist, nil
}

// GetMd5 returns md5hash
func GetMd5(filepath string) (s string) {
	if IsDirectory(filepath) {
		return ""
	}

	f, _ := os.Open(filepath)
	h := md5.New()
	io.Copy(h, f)

	return strings.ToLower(fmt.Sprintf("%x", h.Sum(nil)))

}
