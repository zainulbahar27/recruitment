package main

import (
	"fmt"
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

func flagfile(banding string, bandingkan []string) bool {
	for _, i := range bandingkan {
		if i == banding {
			return false
		}
	}
	return true
}

func hasil(data []string, x []string, y []string, z bool) []string {
	for _, i := range x {
		if flagfile(i, y) {
			if z {
				data = append(data, i+" Deleted ")
			} else {
				data = append(data, i+" New ")
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
	status = hasil(status, source, target, false)
	status = hasil(status, target, source, true)
	for _, i := range status {
		fmt.Println(i)
	}
}
