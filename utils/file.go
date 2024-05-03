package utils

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/components"
)

func WriteHtml(page *components.Page, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	return page.Render(io.MultiWriter(f))
}
