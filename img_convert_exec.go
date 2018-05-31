package main

import (
	"flag"
	"fmt"
	"img-convert-github/imgconv"
	"io/ioutil"
	"os"
	"path"
)

var targetDir = flag.String("dir", "", "フォーマット変換を行うディレクトリを指定する。")
var in = flag.String("in", "jpg", "入力画像フォーマットを指定する。対応フォーマット:jpg, png")
var out = flag.String("out", "png", "出力画像フォーマットを指定する。対応フォーマット:jpg, png")
var jquality *int = flag.Int("jquality", 100, "変換クオリティを指定する。（0-100）")

func main() {
	var ic imgconv.ImgConvert
	var err error
	//convertToPngFromDirectry(`C:\Users\bluem\go\img`)
	flag.Parse()

	var dirList []string
	dirList = getFileName(*targetDir)
	for _, path := range dirList {
		if ic, err = setting(path); err == nil {
			err = ic.ConvertTo()
		}
	}

	if err != nil {
		fmt.Println("[ERROR]", err)
	}
}

// setting receives and sets parameters necessary for image conversion from the command line.
func setting(path string) (imgconv.ImgConvert, error) {
	var ic imgconv.ImgConvert
	ic.InFormat = imgconv.NON
	ic.OutFormat = imgconv.NON
	ic.Path = path
	ic.Jquality = *jquality
	if "jpg" == *in {
		ic.InFormat = imgconv.JPG
	} else if "png" == *in {
		ic.InFormat = imgconv.PNG
	} else {
		return ic, fmt.Errorf("The parameter `in` is incorrect.")
	}
	if "jpg" == *out {
		ic.OutFormat = imgconv.JPG
	} else if "png" == *out {
		ic.OutFormat = imgconv.PNG
	} else {
		return ic, fmt.Errorf("The parameter `out` is incorrect.")
	}
	if ic.InFormat == ic.OutFormat {
		return ic, fmt.Errorf("Cannot convert format.")
	}
	return ic, nil
}

// getFileName get the list of file names under the specified directory.
func getFileName(dir string) []string {
	var paths []string
	var files []os.FileInfo
	var err error
	if files, err = ioutil.ReadDir(dir); err != nil {
		println("Error", err)
	}
	for _, file := range files {
		if file.IsDir() {
			getFileName(dir + `\` + file.Name())
		} else if path.Ext(file.Name()) == ".jpg" {
			paths = append(paths, dir+`\`+file.Name())
		} else {
		}
	}
	return paths
}
