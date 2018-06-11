package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"

	"img-convert-github/imgconv"
)

var (
	targetDir = flag.String("dir", "", "フォーマット変換を行うディレクトリを指定する。")
	in        = flag.String("in", "jpg", "入力画像フォーマットを指定する。対応フォーマット:jpg, png")
	out       = flag.String("out", "png", "出力画像フォーマットを指定する。対応フォーマット:jpg, png")
	jquality  = flag.Int("jquality", 100, "変換クオリティを指定する。（0-100）")
)

func main() {
	var err error
	//convertToPngFromDirectry(`C:\Users\bluem\go\img`)
	td := `C:\Users\bluem\go\img`
	flag.Parse()
	targetDir = &td

	var dirList []string
	dirList, err = getFileName(*targetDir)
	if err == nil {
		for _, path := range dirList {
			ic := imgconv.NewImgConvert(*in, *out, path, *jquality)
			if ic == nil {
				break
			}
			err = ic.ConvertTo()
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		fmt.Println("[ERROR]", err)
		fmt.Println("Error occurrece")
	}
}

// getFileName get the list of file names under the specified directory.
func getFileName(dir string) ([]string, error) {
	var paths []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.Wrap(err, "ReadDir faild")
	}
	for _, file := range files {
		if file.IsDir() {
			underPaths, err := getFileName(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, errors.Wrap(err, "getFileName faild")
			}
			for _, path := range underPaths {
				paths = append(paths, path)
			}
		} else if ext := filepath.Ext(file.Name()); ext != "" {
			if (*in == "jpg" && ext == ".jpg") || (*in == "png" && ext == ".png") {
				paths = append(paths, filepath.Join(dir, file.Name()))
			}
		}
	}
	return paths, nil
}
