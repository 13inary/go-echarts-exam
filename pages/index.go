package pages

import (
	"go-echarts-exam/config"
	"go-echarts-exam/modules"
	"math"
	"math/rand"
	"path/filepath"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}

func (i *Index) Build() error {
	//datas := generatePieItems()
	charts := []components.Charter{
		//modules.NewPie().PieBase("in out", datas),
		//modules.NewBar().BarBase("bar", weeks, generateBarItems()),
		modules.NewLine3d().Line3dBase("queue", genLine3dData()),
	}
	return newPage(filepath.Join(config.HtmlDir, "index.html"), charts...).run()
}

var (
	itemCntPie = 2
	seasons    = []string{"in", "out"}
)

func generatePieItems() []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < itemCntPie; i++ {
		items = append(items, opts.PieData{Name: seasons[i], Value: rand.Intn(100)})
	}
	return items
}

var (
	itemCnt = 7
	weeks   = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
)

func generateBarItems() map[string][]opts.BarData {
	datas := make(map[string][]opts.BarData, 1)
	for i := 0; i < 1; i++ {
		items := make([]opts.BarData, 0)
		for i := 0; i < itemCnt; i++ {
			items = append(items, opts.BarData{Value: rand.Intn(300)})
		}
		datas[strconv.Itoa(i)] = items
	}
	return datas
}

func genLine3dData() []opts.Chart3DData {
	data := make([][3]float64, 0)
	for i := 0; i < 25000; i++ {
		t := float64(i) / 1000
		data = append(data,
			[3]float64{(1 + 0.25*math.Cos(75*t)) * math.Cos(t), (1 + 0.25*math.Cos(75*t)) * math.Sin(t), t + 2.0*math.Sin(75.0*t)},
		)
	}

	ret := make([]opts.Chart3DData, 0, len(data))
	//for _, d := range data {
	ret = append(ret, opts.Chart3DData{Name: "1", Value: []interface{}{1, 0, 0}})
	ret = append(ret, opts.Chart3DData{Name: "1", Value: []interface{}{1, 1, 1}})
	ret = append(ret, opts.Chart3DData{Name: "1", Value: []interface{}{1, 2, 2}})
	ret = append(ret, opts.Chart3DData{Name: "1", Value: []interface{}{1, 3, 3}})

	ret = append(ret, opts.Chart3DData{Name: "2", Value: []interface{}{2, 0, 0}})
	ret = append(ret, opts.Chart3DData{Name: "2", Value: []interface{}{2, 1, 1}})
	//}
	return ret
}

func genLine3dData0() []opts.Chart3DData {
	data := make([][3]float64, 0)
	for i := 0; i < 25000; i++ {
		t := float64(i) / 1000
		data = append(data,
			[3]float64{(1 + 0.25*math.Cos(75*t)) * math.Cos(t), (1 + 0.25*math.Cos(75*t)) * math.Sin(t), t + 2.0*math.Sin(75.0*t)},
		)
	}

	ret := make([]opts.Chart3DData, 0, len(data))
	for _, d := range data {
		ret = append(ret, opts.Chart3DData{Value: []interface{}{d[0], d[1], d[2]}})
	}
	return ret
}
