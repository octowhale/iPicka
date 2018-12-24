package util

import (
	"fmt"
	"testing"
)

func Test_Global(t *testing.T) {

	dirpath := "/data/tmp"

	WalkDirectory(dirpath)

	fmt.Println("file md5:", GetMd5("/data/tmp/naruto.jpg"))
}
