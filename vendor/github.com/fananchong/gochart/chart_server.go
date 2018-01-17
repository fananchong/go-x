package gochart

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

type ChartServer struct {
	charts map[string]IChartInner
}

func (this *ChartServer) AddChart(chartname string, chart IChartInner, savedata bool) {
	if this.charts == nil {
		this.charts = make(map[string]IChartInner)
	}
	chart.Init()
	this.charts[chartname] = chart
	if savedata {
		chart.GoSaveData(chartname)
	}
}

func (this *ChartServer) ListenAndServe(addr string) error {
	http.HandleFunc("/", this.handler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/js/", this.js)
	return http.ListenAndServe(addr, nil)
}

func (this *ChartServer) handler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	chartname := values.Get("query")
	if chartname == "" {
		xlog.Errorln("usage: http://your_ip:8000?query=cpu")
		return
	}
	if _, ok := this.charts[chartname]; ok {
		this.queryChart(chartname, w, r)
	} else if ok, path := this.isExistFile(chartname); ok {
		this.queryChartFile(chartname, path, w, r)
	} else {
		xlog.Errorln("no find the chart, chartname =", chartname)
		return
	}
}

func (this *ChartServer) queryChart(chartname string, w http.ResponseWriter, r *http.Request) {
	now := time.Now().Unix()
	chart := this.charts[chartname]
	datas := chart.Update(now)
	chart.SaveData(datas)
	outdatas := chart.AddData(datas, now)
	json := simplejson.New()
	json.Set("DataArray", outdatas)
	b, _ := json.Get("DataArray").Encode()
	chart.Build(string(b))
	if t, err := template.New("foo").Parse(chart.Template()); err != nil {
		w.Write([]byte(err.Error()))
	} else {
		if err = t.ExecuteTemplate(w, "T", chart.Data()); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

func (this *ChartServer) queryChartFile(chartname, path string, w http.ResponseWriter, r *http.Request) {
	s := strings.Split(chartname, "_")
	if len(s) < 2 || len(s[1]) < 5 {
		xlog.Errorln("chart file name error! file =", chartname)
		return
	}

	var chart IChartFile
	ct, _ := strconv.Atoi(string(s[1][4]))
	switch ChartClassType(ct) {
	case CCT_TIME:
		chart = &ChartTime{}
	default:
		xlog.Errorln("chart file type error! file =", chartname)
		return
	}

	ok, outdatas := chart.Load(path)
	if !ok {
		xlog.Errorln("load chart file fail! file =", chartname)
		return
	}
	chart.Init()
	json := simplejson.New()
	json.Set("DataArray", outdatas)
	b, _ := json.Get("DataArray").Encode()
	chart.Build(string(b))
	if t, err := template.New("foo").Parse(chart.TemplateScrollBars()); err != nil {
		w.Write([]byte(err.Error()))
	} else {
		data := chart.Data()
		data["RefreshTime"] = "99999"
		if err = t.ExecuteTemplate(w, "T", data); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

func (this *ChartServer) isExistFile(chartname string) (bool, string) {
	wd, err1 := os.Getwd()
	if err1 != nil {
		xlog.Errorln(err1)
		return false, ""
	}
	filename := wd + "/" + chartname
	_, err2 := os.Stat(filename)
	return err2 == nil || os.IsExist(err2), filename
}

func (this *ChartServer) js(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		xlog.Errorln(err)
		return
	}
	http.FileServer(http.Dir(wd)).ServeHTTP(w, r)
}
