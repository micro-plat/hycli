<template>
    <div :id="unq" :style="{ margin: margin, height: height, width: width }"></div>
</template>
<script>
import * as echarts from 'echarts';
export default {
    name: "ChartBase",
    props: {
        title: {
            type: String,
        },
        unq: {
            type: String
        },
        url: {
            type: String,
        },
        height: {
            type: String,
            default: "300px"
        },
        width: {
            type: String,
            default: "100%"
        },
        margin: {
            type: String,
            default: "32px 0px 16px 0px"
        },
        type: {
            type: String,
            default: "line",
        },
        stack: {
            type: String,
            default: "",
        },
        rType: {
            type: String,
            default: "false",
        }
    },
    data() {
        return {
            roseType: {
                "rose": "radius",
            },
            radius: {
                "radius": ['40%', '70%']
            },
            styles: {
                "rose": {
                    borderRadius: 5
                },
                "radius": {
                    borderRadius: 10,
                    borderColor: '#fff',
                    borderWidth: 2,
                },
            },
            option: {
                title: {
                    text: this.title,
                    left: 'center'
                },
                tooltip: {
                    trigger: 'axis'
                },
                legend: {
                    data: ['Email', 'Union Ads', 'Video Ads', 'Direct'],
                    orient: 'vertical',
                    left: 'right',
                },
                grid: {
                    left: '3%',
                    right: '4%',
                    bottom: '3%',
                    containLabel: true
                },
                xAxis: {
                    type: 'category',
                    boundaryGap: ['5%', '5%'],
                    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        name: 'Email',
                        type: 'line',
                        stack: 'Total',
                        data: [120, 132, 101, 134, 90, 230, 210]
                    },
                    {
                        name: 'Union Ads',
                        type: 'line',
                        stack: 'Total',
                        data: [220, 182, 191, 234, 290, 330, 310]
                    },
                    {
                        name: 'Video Ads',
                        type: 'line',
                        stack: 'Total',
                        data: [150, 232, 201, 154, 190, 330, 410]
                    },
                    {
                        name: 'Direct',
                        type: 'line',
                        stack: 'Total',
                        data: [320, 332, 301, 334, 390, 330, 320]
                    },
                    {
                        name: 'Search Engine',
                        type: 'line',
                        stack: 'Total',
                        data: [820, 932, 901, 934, 1290, 1330, 1320]
                    }
                ]
            }
        }
    },
    methods: {
        showChart() {
            // 基于准备好的dom，初始化echarts实例
            var chartDom = document.getElementById(this.unq);
            var myChart = echarts.init(chartDom);
            this.option && myChart.setOption(this.option);
        },
        show(form) {
            if (!this.url) {
                return
            }
            let that = this
            this.$theia.http.get(this.url, form).then(res => {
                if (this.resetData(res || {}, that)) {
                    this.showChart()
                }
            });
        },
        resetData(item, that) {
            let url = that.url
            //处理多个图表
            that.option.legend.data = [this.title]
            if (item.legend && Array.isArray(item.legend)) {
                that.option.legend.data = item.legend
            }
            if ((that.type.toLowerCase() != "pie") && (!item.xAxis || !Array.isArray(item.xAxis))) {
                let msg = `请求(${url})返回数据中未包含xAxis`
                that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                return false
            }
            if (!item.yAxis || !Array.isArray(item.yAxis)) {
                let msg = `请求(${url})返回数据中未包含yAxis`
                that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                return false
            }
            that.option.xAxis.data = item.xAxis || []
            that.option.series = []
            that.option.xAxis.data = item.xAxis || []
            that.option.series = []
            that.setYAxisData(that, item.yAxis)
            return true
        },
        setYAxisData(that, yAxis) {

            //处理饼图的名称问题https://echarts.apache.org/examples/zh/editor.html?c=pie-simple
            //[{"name":"2023-3","value":100},{"name":"2023-4","value":400},{"name":"2023-5","value":600}]}
            if ((yAxis || []).length > 0 && !Array.isArray(yAxis[0])) {
                let s = {
                    type: that.type.toLowerCase(),
                    stack: that.stack,
                    barWidth: "8%",//柱子宽度
                    barGap: "0%",
                    labelLine: { show: yAxis["name"] != null },
                    roseType: that.roseType[that.rType] || "",
                    itemStyle: that.styles[that.rType] || {},
                    data: yAxis
                }
                if (that.radius[that.rType]) {
                    s.radius = that.radius[that.rType]
                }
                that.option.series.push(s)
                return
            }
            //[[320,332,301,334,390,330,320]]}
            yAxis.forEach((y, i) => {
                let s={
                    name: that.option.legend.data[i],
                    type: that.type.toLowerCase(),
                    stack: that.stack,
                    barWidth: "8%",//柱子宽度
                    barGap: "0%",
                    labelLine: { show: y["name"] != null },
                    roseType: that.roseType[that.rType] || "",
                    itemStyle: that.styles[that.rType] || {},
                    data: y
                }
                if (that.radius[that.rType]) {
                    s.radius = that.radius[that.rType]
                }
                that.option.series.push(s)
            });
        },
    },
}
</script>