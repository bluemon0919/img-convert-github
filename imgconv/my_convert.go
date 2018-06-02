// Package imgconv provides image format conversion function.
// Supported formats are PNG, JPEG.
// Specify the path and format of the input image and the format you want to convert.
package imgconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

// Format is defines the image format.
type Format int

// ImgConvert is a parameter required for image format ImgConvert.
// Set input / output format, input image path, etc.
// by tennten ImgConvertは動詞っぽい（本来はConvertImgだが。)型名や変数は名詞を使う（e.g. Converter)
type ImgConvert struct {
	InFormat  Format // InFormat is input image format.
	OutFormat Format // OutFormat is output image format.
	Path      string // TargetDir is input image path
	Jquality  int    // Jquality is the quality when converting JPEG.
}

// by tenntenn 整数じゃなくて文字列のままで良かったんでは？
// by tenntenn Format型の定義の真下に書いたほうがわかりやすい
const (
	// NON is the initial value and is prohibited.
	NON Format = 0
	//JPG is JPEG format.
	JPG Format = 1
	// PNG is PNG format.
	PNG Format = 2
)

// ConvertTo converts the image format according to the given parameters.
// by tenntenn レシーバがポインタじゃない理由はあるのか？ポインタにしないとコピーが発生する
// by tenntenn see https://play.golang.org/p/uWzUm-uF-ZL
func (ic ImgConvert) ConvertTo() error {
	var err error
	var img image.Image
	if img, err = ic.decodeTo(); err != nil {
		return err
	}
	if err = ic.encodeTo(img); err != nil {
		return err
	}
	return nil
}

// decodeTo decodes image data to image.Image format.
func (ic ImgConvert) decodeTo() (image.Image, error) {
	var img image.Image
	var inputFile *os.File
	var err error
	if inputFile, err = os.Open(ic.Path); err != nil {
		/* by tenntenn
		元のエラーを隠蔽しない。
		やるならerrors.Wrapを使う。
		see: https://deeeet.com/writing/2016/04/25/go-pkg-errors/
		*/
		return img, fmt.Errorf("open error")
	}
	defer inputFile.Close()

	/* by tenntenn
	image.Decodeを使えばいい
	see https://golang.org/pkg/image/#Decode
	ただし、以下のようにpngとjpgパッケージをインポートしておく必要がある

	import (
		_ "image/png"
		_ "image/jpg"
	)
	*/
	if JPG == ic.InFormat {
		if img, err = jpeg.Decode(inputFile); err != nil {
			// by tenntenn 元のエラーを隠蔽しない
			return img, fmt.Errorf("jpeg.decode error")
		}
	} else {
		if img, err = png.Decode(inputFile); err != nil {
			// by tenntenn 元のエラーを隠蔽しない
			return img, fmt.Errorf("png.decode error")
		}
	}

	return img, nil
}

// encodeTo generates an image file according to specified parameters.
func (ic ImgConvert) encodeTo(img image.Image) error {
	var outputFile *os.File
	var err error
	withoutExt := ic.Path[:strings.LastIndex(ic.Path, path.Ext(ic.Path))]

	// by tennten JPGとPNGの分岐が多すぎる。インタフェースなりを使って賢く分岐する
	// by tenntenn ifのでかすぎる。4行以上かかない気持ちでいる。ネストは基本しない。
	if JPG == ic.OutFormat {
		if outputFile, err = os.Create(withoutExt + ".jpg"); err != nil {
			return fmt.Errorf("create error")
		}
		// by tenntenn デバッグコードを消す
		println("path" + path.Ext(ic.Path))
		defer outputFile.Close()

		opts := &jpeg.Options{Quality: ic.Jquality}
		if err := jpeg.Encode(outputFile, img, opts); err != nil {
			return fmt.Errorf("jpeg.encode error")
		}
	} else {
		if outputFile, err = os.Create(withoutExt + ".png"); err != nil {
			return fmt.Errorf("create error")
		}
		// by tenntenn デバッグコードを消す
		println("path" + path.Ext(ic.Path))
		defer outputFile.Close()

		if err := png.Encode(outputFile, img); err != nil {
			return fmt.Errorf("png.encode error")
		}
	}

	return nil
}
