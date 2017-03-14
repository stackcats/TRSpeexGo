package handler

import (
	"errors"
	"fmt"
	"github.com/stackcats/TRSpeexGo/util"
	"gopkg.in/kataras/iris.v6"
	"os"
	"os/exec"
)

// SpxToQN spx convert to qiniu
func SpxToQN(ctx *iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(iris.StatusOK, iris.Map{
				"code":    0,
				"message": fmt.Sprint(err),
			})
		}
	}()

	url := ctx.FormValue("url")
	if url == "" {
		panic(errors.New("url不存在"))
	}

	fname, err := util.Download(url)
	if err != nil {
		panic(err)
	}

	fpath := "./" + fname

	defer os.Remove(fpath)

	util.Convert(fpath)

	wavfile := fpath + ".wav"

	mp3file := fpath + ".mp3"

	if err := exec.Command("lame", wavfile, mp3file).Run(); err != nil {
		panic(err)
	}

	os.Remove(wavfile)

	ret, err := util.Upload(mp3file)
	if err != nil {
		panic(err)
	}

	os.Remove(mp3file)

	ctx.JSON(iris.StatusOK, iris.Map{
		"code":   "1",
		"result": ret,
	})
}
