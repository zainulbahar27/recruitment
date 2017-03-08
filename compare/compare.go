package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func caripath(cari string) []string {
	isifolder := []string{}
	_ = filepath.Walk(cari, func(path string, f os.FileInfo, err error) error {
		path = strings.Replace(path, cari, "", -1)
		if path != cari {
			isifolder = append(isifolder, path)
		}
		return nil
	})
	return isifolder
}

func minimalize(Path string) (result string, err error) {
	file, err := os.Open(Path)
	if err != nil {
		return
	}
	defer file.Close()
	x := md5.New()
	_, err = io.Copy(x, file)
	if err != nil {
		return
	}
	result = hex.EncodeToString(x.Sum(nil))
	return
}

func flagfile(banding string, bandingkan []string) bool {
	for _, i := range bandingkan {
		if i == banding {
			return false
		}
	}
	return true
}

func hasil(data []string, x []string, y []string, z int8) []string {
	for _, i := range x {
		if flagfile(i, y) {
			if z == 0 {
				data = append(data, i+" Deleted ")
			} else if z == 1 {
				data = append(data, i+" New ")
			}
		} else if z > 1 {
			a, _ := minimalize("source/" + i)
			b, _ := minimalize("target/" + i)
			if a != b {
				data = append(data, i+" Modified ")
			}
		}
	}
	return data
}

func main() {
	var status []string
	source := caripath("source")
	target := caripath("target")
	// fmt.Println(source)
	// fmt.Println(target)
	status = hasil(status, source, target, 1)
	status = hasil(status, target, source, 0)
	status = hasil(status, source, target, 2)
	for _, i := range status {
		fmt.Println(i)
	}
}
