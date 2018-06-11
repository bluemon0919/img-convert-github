// Package imgconv provides image format conversion function.
// Supported formats are PNG, JPEG.
// Specify the path and format of the input image and the format you want to convert.
package imgconv

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

type ImgConvert interface {
	ConvertTo() error
	decodeTo() (image.Image, error)
	encodeTo(img image.Image) error
}

// ImgConvert is a parameter required for image format ImgConvert.
// Set input / output format, input image path, etc.
type ImgConvInfo struct {
	InFormat  string // InFormat is input image format.
	OutFormat string // OutFormat is output image format.
	Path      string // TargetDir is input image path
	extPath   string
	//	Jquality  int    // Jquality is the quality when converting JPEG.
}

type ImgJpgConvert struct {
	*ImgConvInfo
	Jquality int // Jquality is the quality when converting JPEG.
}

type ImgPngConvert struct {
	*ImgConvInfo
}

// NewImgConvert generates an image change object from the input parameters.
func NewImgConvert(inFormat string, outFormat string, path string, jquality int) ImgConvert {
	p := &ImgConvInfo{inFormat, outFormat, path, ""}
	if "jpg" == outFormat {
		return &ImgJpgConvert{p, jquality}
	}
	if "png" == outFormat {
		return &ImgPngConvert{p}
	}
	return nil
}

// ConvertTo converts the image format according to the given parameters.
func (ic *ImgConvInfo) ConvertTo() error {
	return errors.New("there is no implementation in the specified type")
}

// ConvertTo converts the image format according to the given parameters.
func (ic *ImgJpgConvert) ConvertTo() error {
	img, err := ic.decodeTo()
	if err != nil {
		return err
	}
	ic.extPath = ic.Path[:strings.LastIndex(ic.Path, path.Ext(ic.Path))]
	err = ic.encodeTo(img)
	if err != nil {
		return err
	}
	return nil
}

// ConvertTo converts the image format according to the given parameters.
func (ic *ImgPngConvert) ConvertTo() error {
	img, err := ic.decodeTo()
	if err != nil {
		return err
	}
	ic.extPath = ic.Path[:strings.LastIndex(ic.Path, path.Ext(ic.Path))]
	err = ic.encodeTo(img)
	if err != nil {
		return err
	}
	return nil
}

// decodeTo decodes image data to image.Image format.
func (ic *ImgConvInfo) decodeTo() (image.Image, error) {
	var img image.Image
	inputFile, err := os.Open(ic.Path)
	defer inputFile.Close()
	if nil != err {
		return img, errors.Wrap(err, "file open error")
	}

	img, ic.InFormat, err = image.Decode(inputFile)
	if err != nil {
		return img, errors.Wrap(err, "decode error")
	}
	return img, nil
}

// encodeTo generates an image file according to specified parameters.
func (ic *ImgConvInfo) encodeTo(img image.Image) error {
	return errors.New("there is no implementation in the specified type")
}

// encodeTo generates an image file according to specified parameters.
func (ic *ImgJpgConvert) encodeTo(img image.Image) error {
	outputFile, err := os.Create(ic.extPath + ".jpg")
	defer outputFile.Close()
	if err != nil {
		return errors.Wrap(err, "create error")
	}
	opts := &jpeg.Options{Quality: ic.Jquality}
	if err := jpeg.Encode(outputFile, img, opts); err != nil {
		return errors.Wrap(err, "encode error")
	}
	return nil
}

// encodeTo generates an image file according to specified parameters.
func (ic *ImgPngConvert) encodeTo(img image.Image) error {
	outputFile, err := os.Create(ic.extPath + ".png")
	defer outputFile.Close()
	if err != nil {
		return errors.Wrap(err, "create error")
	}
	if err := png.Encode(outputFile, img); err != nil {
		return errors.Wrap(err, "encode error")
	}
	return nil
}
