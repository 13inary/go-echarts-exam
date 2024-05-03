package modules

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var line3DColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

type Line3d struct {
}

func NewLine3d() *Line3d {
	return &Line3d{}
}

func (l *Line3d) Line3dBase(title string, datas []opts.Chart3DData) *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "2400px",
			Height: "1200px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),
		charts.WithGrid3DOpts(opts.Grid3D{
			ViewControl: &opts.ViewControl{
				AutoRotate: false,
			},
		}),
		charts.WithLegendOpts(opts.Legend{Show: false}),
	)
	line3d.AddSeries(title, datas).SetSeriesOptions()
	return line3d
}
