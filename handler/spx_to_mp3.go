package handler

import (
	"fmt"
	"github.com/stackcats/TRSpeexGo/util"
	"gopkg.in/kataras/iris.v6"
	"io"
	"os"
	"os/exec"
)

// SpxToMp3 spx convert to mp3
func SpxToMp3(ctx *iris.Context) {
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

	wavfile := fpath + ".wav"

	mp3file := fpath + ".mp3"

	if err := exec.Command("lame", "-S", "-V", "5", wavfile, mp3file).Run(); err != nil {
		panic(err)
	}

	os.Remove(wavfile)

	defer os.Remove(mp3file)

	ctx.SendFile(mp3file, fname+".mp3")
}
