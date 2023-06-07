package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world")
}

func start() {
	r := gin.Default()
	// r.GET(relativePath string, handlers ...gin.HandlerFunc)
	r.Run()
}

func main() {
	r := gin.Default()
	r.POST("/", func(ctx *gin.Context) {
		v := gin.H{}
		ctx.Bind(&v)
		values := ctx.Request.PostForm
		fmt.Printf("formdata: %+v\n", values)
		fmt.Printf("binding data: %+v\n", v)
		ctx.JSON(http.StatusOK, values)
	})
	r.GET("/", func(ctx *gin.Context) {
		// ctx.Request.Header["token"]
		ctx.String(http.StatusOK, "Hello World!!")
	})

	r.GET("/excel", func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
		ctx.Writer.Header().Add("Content-disposition", "attachment;filename="+"test.xls")
		ctx.Writer.Header().Add("Content-Transfer-Encoding", "binary")
		xlFile := xlsx.NewFile()
		xlFile.Write(ctx.Writer)
		ctx.Status(http.StatusOK)
	})

	r.GET("/apis/forward/test", func(ctx *gin.Context) {
		a := ctx.Request.Header.Get("Login")
		b := ctx.Request.Header.Get("Login.name")
		c := ctx.Request.Header.Get("timeoffset")
		fmt.Println(a, b, c)
		fmt.Printf("header: %+v\n", ctx.Request.Header)
		// header: map[Accept-Encoding:[gzip] Authorization:[685c7745-833d-4270-9533-e112be650bfc] Connection:[close] Level:[user] Login:[alan12345] Login.name:[alan12345] User-Agent:[Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36]]
		fmt.Printf("body: %+v\n", ctx.Request.Body)
		ctx.String(http.StatusOK, "Hello world!!")
	})
	///v4/my/ricky/test

	r.GET("/v4/my/ricky/test", func(ctx *gin.Context) {
		fmt.Println("123")
		ctx.String(http.StatusOK, "Hello world!!")
	})

	r.GET("/ping", Hello)

	srv := &http.Server{
		Addr:    ":9999",
		Handler: r,
	}
	go func() {
		srv.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	<-ctx.Done()
}
