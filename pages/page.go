package pages

import (
	"go-echarts-exam/utils"

	"github.com/go-echarts/go-echarts/v2/components"
)

type IPage interface {
	Build() error
}

type Page struct {
	File   string
	Charts []components.Charter
}

func newPage(file string, charts ...components.Charter) *Page {
	return &Page{
		File:   file,
		Charts: charts,
	}
}

func (p *Page) run() error {
	page := components.NewPage()
	page.AddCharts(p.Charts...)
	page.SetLayout(components.PageFlexLayout)
	return utils.WriteHtml(page, p.File)
}
