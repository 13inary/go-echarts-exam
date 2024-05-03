package modules

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Bar struct {
}

func NewBar() *Bar {
	return &Bar{}
}

func (p *Bar) BarBase(title string, x []string, datas map[string][]opts.BarData) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		//charts.WithInitializationOpts(opts.Initialization{
		//	Theme: types.ThemeInfographic,
		//}),
		charts.WithTitleOpts(opts.Title{
			Title: title,
			//Link:  "http://" + config.Socket,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "bar",
			//AxisLabel: &opts.AxisLabel{
			//	Show:      true,
			//	Formatter: "{value} x-unit",
			//},
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "one",
			//AxisLabel: &opts.AxisLabel{
			//	Show:      true,
			//	Formatter: "{value} y-unit",
			//},
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
		//charts.WithColorsOpts(opts.Colors{"blue", "pink"}),
		charts.WithLegendOpts(opts.Legend{Show: false}),
		//charts.WithToolboxOpts(opts.Toolbox{
		//	Right: "20%",
		//	Feature: &opts.ToolBoxFeature{
		//		SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
		//			Type:  "jpg",
		//			Title: "Anything you want",
		//		},
		//		DataView: &opts.ToolBoxFeatureDataView{
		//			Title: "DataView",
		//			// set the language
		//			// Chinese version: ["数据视图", "关闭", "刷新"]
		//			Lang: []string{"data view", "turn off", "refresh"},
		//		},
		//	}},
		//),
	)
	bar.SetXAxis(x)
	//bar.SetSeriesOptions(
	//	charts.WithBarChartOpts(opts.BarChart{
	//		BarGap: "150%",
	//	}),
	//)
	for k, v := range datas {
		bar.AddSeries(k, v)
		//bar.AddSeries(k, v).SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{
		//	Stack: "stackA",
		//}))
	}
	//bar.XYReversal()
	return bar
}
