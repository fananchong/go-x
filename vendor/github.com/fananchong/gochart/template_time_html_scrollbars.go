package gochart

// see the resource of : http://www.freecdn.cn/libs/highcharts/

// spline,line,column,area,bar
var TemplateTimeHtml_ScrollBars = `{{define "T"}}
<!DOCTYPE HTML>
<html>
    <head>
	    <meta http-equiv="refresh" content='{{.RefreshTime}}'>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title>Gochart - {{.Title}}</title>
        <style>
            .center{
                position: fixed;
                top: 50%;
                left: 50%;
                width:100%;
                height: 50%;
                -webkit-transform: translateX(-50%) translateY(-50%);
        }
        </style>
        <link rel="stylesheet" href="/js/highcharts.css">
        <script type="text/javascript" src="/js/jquery-1.8.3.min.js"></script>
        <script type="text/javascript">
        $(function () {
            Highcharts.stockChart('container', {
                chart: {
                    // type: 'spline'
                    type: '{{.ChartType}}',
                    animation:false
                                        // 默认是折线图(line)。type取值如下：
                                        // line:折线图
                                        // spline:曲线图，线条圆滑一些
                                        // column:竖向柱状图
                                        // area:面积图
                                        // bar:横向柱状图
                                        // 更多选项请参考：http://api.highcharts.com/highcharts#plotOptions
                },
                title: {
                    // text: 'Monthly Average Temperature',
                    text: '{{.Title}}',
                },
                subtitle: {
                    // text: 'Source: WorldClimate.com',
                    text: '{{.SubTitle}}',
                },
                xAxis: {
                    type: 'datetime',
                    tickInterval: {{.TickInterval}},
                    labels: {
                        format: '{value: %H:%M:%S}',
                        step: {{.TickLabelStep}},
                        staggerLines: 1
                    }
                },
                yAxis: {
                    title: {
                        // text: 'Temperature (°C)'
                        text: '{{.YAxisText}}'
                    },
                    plotLines: [{{.PlotLinesY}}],
                    max: {{.YMax}},
                    min: 0
                },
                tooltip: {
                    shared: true,
                    // valueSuffix: '°C'
                    valueSuffix: '{{.ValueSuffix}}'
                },
                legend: {
                    layout: 'vertical',
                    align: 'right',
                    verticalAlign: 'middle',
                    borderWidth: 0
                },
                plotOptions: {
                    series: {
                        animation: false,
                        marker: {
	                        radius: 1
                        }
                    }
                },
                credits : {
                    enabled: false
                },
                series: {{.DataArray}}
                /*
                [
	                // sample data is bellow
	                {
	                    name: 'Tokyo',
	                    data: [7.0, 6.9, 9.5, 14.5, 18.2, 21.5, 25.2, 26.5, 23.3, 18.3, 13.9, 9.6]
	                    pointInterval: 306000,
                        pointStart: Date.UTC(2014, 6, 10,0,0,0),
                        pointEnd: Date.UTC(2014,6,10,23,59,59),
	                }, 
	                {
	                    name: 'New York',
	                    data: [-0.2, 0.8, 5.7, 11.3, 17.0, 22.0, 24.8, 24.1, 20.1, 14.1, 8.6, 2.5]
	                    pointInterval: 306000,
                        pointStart: Date.UTC(2014, 6, 10,0,0,0),
                        pointEnd: Date.UTC(2014,6,10,23,59,59),
	                }, 
	                {
	                    name: 'Berlin',
	                    data: [-0.9, 0.6, 3.5, 8.4, 13.5, 17.0, 18.6, 17.9, 14.3, 9.0, 3.9, 1.0]
	                    pointInterval: 306000,
                        pointStart: Date.UTC(2014, 6, 10,0,0,0),
                        pointEnd: Date.UTC(2014,6,10,23,59,59),
	                }, 
	                {
	                    name: 'London',
	                    data: [3.9, 4.2, 5.7, 8.5, 11.9, 15.2, 17.0, 16.6, 14.2, 10.3, 6.6, 4.8]
	                    pointInterval: 306000,
                        pointStart: Date.UTC(2014, 6, 10,0,0,0),
                        pointEnd: Date.UTC(2014,6,10,23,59,59),
	                }
                ]
                */
            });
        });    
        </script>
    </head>
    <body>
    <script type="text/javascript" src="/js/highstock.js"></script>
    <script type="text/javascript" src="/js/exporting.js"></script>
    <div id="container" style="height: 500px"></div>
    
</body>
</html>
{{end}}
`
