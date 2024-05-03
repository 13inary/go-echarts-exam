package modules

import (
	"go-echarts-exam/config"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Pie struct {
}

func NewPie() *Pie {
	return &Pie{}
}

func (p *Pie) PieBase(title string, datas []opts.PieData) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		//charts.WithInitializationOpts(opts.Initialization{
		//	Theme: types.ThemeInfographic,
		//}),
		charts.WithTitleOpts(opts.Title{
			Title: title,
		}),
		charts.WithColorsOpts(opts.Colors{config.Green, config.DarkRed}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Trigger:   "item",
			Formatter: "{a} <br/>{b} : {c} ({d}%)",
		}),
		charts.WithLegendOpts(opts.Legend{Show: false}),
	)
	pie.AddSeries(title, datas).SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:      true,
			FontSize:  18,
			Formatter: "{b}: {c}",
		}),
		charts.WithPieChartOpts(opts.PieChart{
			//Radius: []string{"40%", "75%"},
			RoseType: "radius",
		}),
	)

	return pie
}
