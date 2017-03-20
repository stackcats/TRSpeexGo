package handler

import (
	"fmt"
	"github.com/cryptix/wav"
	"github.com/stackcats/TRSpeexGo/util"
	"gopkg.in/kataras/iris.v6"
	"io"
	"os"
)

// SpxToWav spx convert to wav
func SpxToWav(ctx *iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(iris.StatusOK, iris.Map{
				"code":    0,
				"message": fmt.Sprint(err),
			})
		}
	}()

	file, info, err := ctx.FormFile("uploadfile")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fname := info.Filename

	fpath := "./uploads/" + fname
	out, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(out, file)

	util.Convert(fpath)

	wavPath := fpath + ".wav"

	fileInfo, _ := os.Stat(wavPath)

	wavFile, _ := os.Open(wavPath)
	wavReader, err := wav.NewReader(wavFile, fileInfo.Size())

	meta := wavReader.GetFile()
	fmt.Println(meta.Duration)
	ctx.SendFile(fpath+".wav", fname+".wav")
}
