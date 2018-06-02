package main

import (
	"flag"
	"fmt"
	"img-convert-github/imgconv"
	"io/ioutil"
	"os"
	"path"
)

// by tenntenn var ( ... ) でまとめたほうがわかりやすいかも
var targetDir = flag.String("dir", "", "フォーマット変換を行うディレクトリを指定する。")
var in = flag.String("in", "jpg", "入力画像フォーマットを指定する。対応フォーマット:jpg, png")
var out = flag.String("out", "png", "出力画像フォーマットを指定する。対応フォーマット:jpg, png")
var jquality *int = flag.Int("jquality", 100, "変換クオリティを指定する。（0-100）")

func main() {
	// by tenntenn Cとは違うので戦闘で定義する必要なし、できるだけ使う場所の近くで定義する
	// by tenntenn := が使えるところは使う
	var ic imgconv.ImgConvert
	var err error
	//convertToPngFromDirectry(`C:\Users\bluem\go\img`)
	flag.Parse()

	var dirList []string
	dirList = getFileName(*targetDir)
	for _, path := range dirList {
		if ic, err = setting(path); err == nil {
			// by tenntenn これだと最後のio.ConvertToの呼び出しがエラーじゃなかったらエラーがnilで上書きされる
			err = ic.ConvertTo()
		}
	}

	if err != nil {
		// by tenntenn os.Stderr（標準エラー出力に)
		// by tenntenn 終了コードを1とかにする
		fmt.Println("[ERROR]", err)
	}
}

// setting receives and sets parameters necessary for image conversion from the command line.
func setting(path string) (imgconv.ImgConvert, error) {

	/* by tenntenn リテラルで書く
	ic := imgconv.ImgConvert{
		InFormat:  imgconv.NON,
		Outformat: imgconv.NON,
		Path:      path,
		Jquality: *jquality,
	}
	*/
	var ic imgconv.ImgConvert
	ic.InFormat = imgconv.NON
	ic.OutFormat = imgconv.NON
	ic.Path = path
	ic.Jquality = *jquality

	// by tenntenn この部分は関数にまとめたほうがわかりやすいのでは？
	if "jpg" == *in {
		ic.InFormat = imgconv.JPG
	} else if "png" == *in {
		ic.InFormat = imgconv.PNG
	} else {
		return ic, fmt.Errorf("The parameter `in` is incorrect.")
	}

	// by tenntenn 意味のある単位で空行をいれる
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

	// by tenntenn ここも空行いれたほうがわかりやすい
	return ic, nil
}

// getFileName get the list of file names under the specified directory.
// by tenntenn filepath.Walkを使ったほうが簡単
func getFileName(dir string) []string {
	var paths []string
	var files []os.FileInfo
	var err error
	if files, err = ioutil.ReadDir(dir); err != nil {
		// by tennten デバッグ目的以外でprintlnは使わない
		println("Error", err)
	}

	for _, file := range files {
		if file.IsDir() {
			var underPaths []string
			// by tenntenn ファイルパスの結合はfilepath.Joinを使う(see path/filepathパッケージ)
			underPaths = getFileName(dir + `\` + file.Name())
			for _, path := range underPaths {
				paths = append(paths, path)
			}
		} else if path.Ext(file.Name()) == ".jpg" {
			// by tenntenn filepath.Join
			paths = append(paths, dir+`\`+file.Name())
		} else {
			// by tenntenn 無駄なelse
		}
	}
	return paths
}
