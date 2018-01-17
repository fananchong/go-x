package gochart

import (
	"fmt"
	"strconv"
	"sync/atomic"

	simplejson "github.com/bitly/go-simplejson"
)

type ChartTime struct {
	ChartBase
	TickInterval  string
	TickLabelStep string
	PlotLinesY    string
	TickUnit      int
}

func (this *ChartTime) Init() {
	this.ChartBase.InitBase()
	this.chartClassType = CCT_TIME
	if this.TickLabelStep == "" {
		this.TickLabelStep = "60"
	}
	this.chartArgs["TickLabelStep"] = this.TickLabelStep
	this.chartArgs["PlotLinesY"] = this.PlotLinesY
	if this.TickUnit == 0 {
		this.TickUnit = 1000
	}
	this.chartArgs["TickInterval"] = strconv.Itoa(this.RefreshTime * this.TickUnit)
}

func (this *ChartTime) Template() string {
	return TemplateTimeHtml
}

func (this *ChartTime) TemplateScrollBars() string {
	return TemplateTimeHtml_ScrollBars
}

func (this *ChartTime) AddData(newDatas map[string][]interface{}, UTCTime int64) []interface{} {
	endtime := 1000 * int(8*60*60+UTCTime)
	begintime := endtime - this.TickUnit*this.SampleNum*this.RefreshTime
	datas := make([]interface{}, 0)
	for k, v := range newDatas {
		if _, ok := this.chartData[k]; !ok {
			this.chartData[k] = make([]interface{}, this.SampleNum)
		}
		for _, tempv := range v {
			this.chartData[k] = append(this.chartData[k], tempv)
		}
		for len(this.chartData[k]) > this.SampleNum {
			this.chartData[k] = this.chartData[k][1:]
		}
		var json *simplejson.Json
		json = simplejson.New()
		json.Set("name", k)
		json.Set("data", this.chartData[k])
		json.Set("pointInterval", this.RefreshTime*this.TickUnit)
		json.Set("pointStart", begintime)
		json.Set("pointEnd", endtime)
		datas = append(datas, json)

		temlen := int64(len(this.chartData[k]))
		atomic.StoreInt64(&this.chartDataSamleNum, temlen)
	}
	return datas
}

func (this *ChartTime) Load(filename string) (bool, []interface{}) {
	ok, root := this.LoadBase(filename)
	if !ok {
		return false, nil
	}
	this.TickInterval, _ = root.Get("TickInterval").String()
	this.TickLabelStep, _ = root.Get("TickLabelStep").String()
	this.PlotLinesY, _ = root.Get("PlotLinesY").String()
	tmpv, _ := strconv.Atoi(this.TickInterval)
	this.TickUnit = tmpv / this.RefreshTime

	outdatas, _ := root.Get("DataArray").String()
	outdatas = fmt.Sprintf("{\"DataArray\":%s}", outdatas)
	json, _ := simplejson.NewJson([]byte(outdatas))
	arrays, _ := json.Get("DataArray").Array()

	datas := make([]interface{}, 0)
	for _, val := range arrays {
		temp := val.(map[string]interface{})
		temparray := temp["data"].([]interface{})
		begintime := int64(0)
		endtime := int64(0)
		templen := len(temparray)
		if templen < this.SampleNum {
			temparray2 := make([]interface{}, this.SampleNum-templen)
			for _, val := range temparray {
				temparray2 = append(temparray2, val)
			}
			temparray = temparray2
			begintime = 1000*(8*60*60+this.beginTime) - int64(this.TickUnit*(this.SampleNum-templen)*this.RefreshTime)
			endtime = begintime + int64(this.TickUnit*this.SampleNum*this.RefreshTime)
		} else {
			begintime = 1000 * (8*60*60 + this.beginTime)
			endtime = begintime + int64(this.TickUnit*templen*this.RefreshTime)
		}
		json := simplejson.New()
		json.Set("name", temp["name"])
		json.Set("data", temparray)
		json.Set("pointInterval", this.RefreshTime*this.TickUnit)
		json.Set("pointStart", begintime)
		json.Set("pointEnd", endtime)
		datas = append(datas, json)
	}
	return true, datas
}
