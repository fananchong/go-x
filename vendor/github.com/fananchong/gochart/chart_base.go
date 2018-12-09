package gochart

import (
	"fmt"
	"io/ioutil"
	"runtime/debug"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

type IChart interface {
	Update(now int64) map[string][]interface{}
}

type IChartInner interface {
	IChart
	IChartSave
	ICharNormal
}

type ICharNormal interface {
	Init()
	Template() string
	Build(dataArray string)
	Data() map[string]string
	AddData(map[string][]interface{}, int64) []interface{}
}

type IChartSave interface {
	GoSaveData(filename string)
	IsEnableSaveData() bool
	SaveData(datas map[string][]interface{})
}

type IChartFile interface {
	ICharNormal
	Load(filename string) (bool, []interface{})
	TemplateScrollBars() string
}

type ChartClassType int

const (
	CCT_UNKNOW ChartClassType = iota
	CCT_TIME
)

type ChartBase struct {
	ChartType    string
	Title        string
	SubTitle     string
	YAxisText    string
	XAxisNumbers string
	ValueSuffix  string
	YMax         string
	RefreshTime  int
	SampleNum    int

	chartArgs map[string]string
	m         sync.RWMutex

	//chart data
	chartData         map[string][]interface{}
	chartDataSamleNum int64

	// save data
	filename     string
	saveData     map[string][]interface{}
	chanSaveData chan map[string][]interface{}
	beginTime    int64

	chartClassType ChartClassType
}

func (this *ChartBase) InitBase() {
	this.chartArgs = make(map[string]string)
	this.chartArgs["ChartType"] = this.ChartType
	this.chartArgs["Title"] = this.Title
	this.chartArgs["SubTitle"] = this.SubTitle
	this.chartArgs["YAxisText"] = this.YAxisText
	this.chartArgs["XAxisNumbers"] = this.XAxisNumbers
	this.chartArgs["ValueSuffix"] = this.ValueSuffix
	this.chartArgs["YMax"] = this.YMax
	if this.RefreshTime == 0 {
		this.RefreshTime = 60
	}
	this.chartArgs["RefreshTime"] = strconv.Itoa(this.RefreshTime)

	this.chartData = make(map[string][]interface{})
}

func (this *ChartBase) Build(dataArray string) {
	this.m.Lock()
	this.chartArgs["DataArray"] = dataArray
	this.m.Unlock()
}

func (this *ChartBase) Data() map[string]string {
	return this.chartArgs
}

func (this *ChartBase) GoSaveData(filename string) {
	this.filename = fmt.Sprintf("%s_type%d.chart", filename, int(this.chartClassType))
	this.chanSaveData = make(chan map[string][]interface{}, 1)
	this.saveData = make(map[string][]interface{})

	go func() {
		defer func() {
			if err := recover(); err != nil {
				xlog.Errorln("[异常] ", err, "\n", string(debug.Stack()))
			}
		}()

		newDataFlag := false
		tick := time.NewTicker(30 * time.Second)
		for {
			select {
			case datas := <-this.chanSaveData:
				if len(datas) > 0 {

					if this.beginTime == 0 {
						this.beginTime = time.Now().Unix()
					}

					newDataFlag = true
					for k, v := range datas {
						if _, ok := this.saveData[k]; !ok {
							this.saveData[k] = make([]interface{}, 0)
						}
						for _, tempv := range v {
							this.saveData[k] = append(this.saveData[k], tempv)
						}
					}
				}
			case <-tick.C:
				if newDataFlag {
					newDataFlag = false

					root := simplejson.New()
					this.m.RLock()
					for key, val := range this.chartArgs {
						root.Set(key, val)
					}
					this.m.RUnlock()

					root.Set("beginTime", this.beginTime)

					temlen := atomic.LoadInt64(&this.chartDataSamleNum)
					root.Set("SampleNum", temlen)

					outdatas := make([]interface{}, 0)
					for k, v := range this.saveData {
						json := simplejson.New()
						json.Set("name", k)
						json.Set("data", v)
						outdatas = append(outdatas, json)
					}
					json := simplejson.New()
					json.Set("DataArray", outdatas)
					b, _ := json.Get("DataArray").Encode()
					root.Set("DataArray", string(b))

					s, _ := root.MarshalJSON()
					ioutil.WriteFile(this.filename, []byte(s), 0666)
				}
			}
		}
	}()
}

func (this *ChartBase) IsEnableSaveData() bool {
	return this.filename != ""
}

func (this *ChartBase) SaveData(datas map[string][]interface{}) {
	if this.chanSaveData != nil {
		this.chanSaveData <- datas
	}
}

func (this *ChartTime) LoadBase(filename string) (bool, *simplejson.Json) {
	data, err1 := ioutil.ReadFile(filename)
	if err1 != nil {
		xlog.Errorln(err1)
		return false, nil
	}

	json, err2 := simplejson.NewJson(data)
	if err2 != nil {
		xlog.Errorln(err2)
		return false, nil
	}

	this.ChartType, _ = json.Get("ChartType").String()
	this.Title, _ = json.Get("Title").String()
	this.SubTitle, _ = json.Get("SubTitle").String()
	this.YAxisText, _ = json.Get("YAxisText").String()
	this.XAxisNumbers, _ = json.Get("XAxisNumbers").String()
	this.ValueSuffix, _ = json.Get("ValueSuffix").String()
	this.YMax, _ = json.Get("YMax").String()
	tmpv, _ := json.Get("RefreshTime").String()
	this.RefreshTime, _ = strconv.Atoi(tmpv)
	this.SampleNum, _ = json.Get("SampleNum").Int()
	this.beginTime, _ = json.Get("beginTime").Int64()

	return true, json
}
