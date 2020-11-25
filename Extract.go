package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

func fname(s string) string {
	Pat := strings.Split(s, "/")
	Name := strings.Split(Pat[len(Pat)-1], ".")
	return Name[0]
}

func openPPT(FName string) error {

	fz, err := zip.OpenReader(FName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fz.Close()

	err = os.Mkdir(fname(FName), 0)

	for _, file := range fz.File {

		if m, _ := regexp.Match(`ppt/media/media[1234567890]*.m4a`, []byte(file.Name)); m {

			fmt.Println("  - ", fname(file.Name))
			ft, _ := file.Open()
			fd, _ := os.Create(fname(FName) + "/" + fname(file.Name) + ".m4a")
			io.Copy(fd, ft)

			ft.Close()
			fd.Close()
		}
	}

	return nil
}

func main() {

	if len(os.Args) > 1 {

		for i := 1; i < len(os.Args); i++ {
			openPPT(os.Args[1])
		}

	} else {

		dir, _ := os.Getwd()
		FL, _ := ioutil.ReadDir(dir)

		for _, E := range FL {
			if ext := path.Ext(E.Name()); ext == ".ppt" || ext == ".pptx" || ext == ".ppsx" || ext == ".pptm" || ext == ".potx" || ext == ".ppsm" || ext == ".pps" {
				fmt.Println(E.Name())
				openPPT(E.Name())
			}
		}
	}
}
