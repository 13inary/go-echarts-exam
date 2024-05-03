package main

import (
	"go-echarts-exam/config"
	"go-echarts-exam/logs"
	"go-echarts-exam/pages"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := os.MkdirAll(config.HtmlDir, 0755); err != nil {
		panic(err)
	}

	allPages := []pages.IPage{
		pages.NewIndex(),
	}

	for _, v := range allPages {
		if err := v.Build(); err != nil {
			panic(err)
		}
	}

	fs := http.FileServer(http.Dir(config.HtmlDir))
	log.Println("running server at http://" + config.Socket)
	log.Fatal(http.ListenAndServe(config.Socket, logs.LogRequest(fs)))
}
