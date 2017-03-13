package main

import (
	"github.com/stackcats/TRSpeexGo/handler"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"os"
	"runtime"
)

/*
#include <stdio.h>
#include <sys/time.h>
#include <sys/resource.h>

int rlimit_init() {
    printf("setting rlimit\n");

    struct rlimit limit;

    if (getrlimit(RLIMIT_NOFILE, &limit) == -1) {
        printf("getrlimit error\n");
        return 1;
    }

    limit.rlim_cur = limit.rlim_max = 50000;

    if (setrlimit(RLIMIT_NOFILE, &limit) == -1) {
        printf("setrlimit error\n");
        return 1;
    }

    printf("set limit ok\n");
    return 0;
}
*/
import "C"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	C.rlimit_init()
	sep := "/"
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		sep = "\\"
	}
	dir, _ := os.Getwd()
	os.Mkdir(dir+sep+"uploads", os.ModePerm)

	app := iris.New()
	app.Adapt(httprouter.New())
	app.Post("spx-to-wav", iris.LimitRequestBodySize(10<<20), handler.SpxToWav)
	app.Listen(":8888")
}
