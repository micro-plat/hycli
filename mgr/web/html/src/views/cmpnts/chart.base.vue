<template>
    <div :id="unq" :style="{ margin: margin, height: height, width: width }"></div>
</template>
<script>
import * as echarts from 'echarts';
export default {
    name: "ChartBase",//line-bar-pie
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
            default: "line",//line,pie,bar
        },
        stack: {
            type: String,
            default: "", //Total
        },
        rType: {
            type: String, //rose
            default: "false",
        },
        showLabel: {
            type: String,
            default: "true"
        },
        theme: {
            type: String,
            default: "vintage",//dark
        }
    },
    data() {
        return {
            lowerType: this.type.toLowerCase(),
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
            formatter: {
                "pie": '{b} : {c} ({d}%)',
                "line_bar": '{b} : {c} ',
            },
            option: {
                title: {
                    text: this.title,
                    left: 'center'
                },
                tooltip: {
                    trigger: 'item',
                },
                legend: {
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
                    }
                ]
            }
        }
    },
    methods: {
        showChart() {
            // 基于准备好的dom，初始化echarts实例
            var chartDom = document.getElementById(this.unq);
            var myChart = echarts.init(chartDom, this.theme);
            this.option && myChart.setOption(this.option);
        },
        show(form) {
            if (!this.url) {
                return
            }
            let that = this
            this.$theia.http.get(this.url, form).then(res => {
                if (that.resetData(res || {}, that)) {
                    that.showChart()
                }
            });
        },
        resetData(r, that) {
            if (that.IsPIE()) {
                return that.resetPIE(r, that)
            }
            return that.resetLineBar(r, that)
        },
        resetPIE(item, that) {
            let url = that.url
            that.option.tooltip.formatter = that.formatter.pie
            that.option.legend.data = []
            if (item.legend && Array.isArray(item.legend)) {
                that.option.legend.data = item.legend
            }

            //将x坐标兼容为legend
            that.option.xAxis.data = []
            if (item.xAxis && Array.isArray(item.xAxis)) {
                that.option.legend.data = item.xAxis
            }
            that.option.yAxis.data = []
            if (item.yAxis && Array.isArray(item.yAxis)) {
                that.option.yAxis.data = item.yAxis
            }
            if (item.data && Array.isArray(item.data)) {
                that.option.yAxis.data = item.data
            }
            //检查输入参数
            if (that.option.legend.data.length == 0) {
                let msg = `请求(${url})返回数据中未包含lengend或xAxis`
                that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                return false
            }
            if (that.option.yAxis.data.length == 0) {
                let msg = `请求(${url})返回数据中未包含data或yAxis`
                that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                return false
            }

            //[[320,332,301,334,390,330,320]] [[{"name":"2023-3","value":100}]]
            //[320,332,301,334,390,330,320],[{"name":"2023-3","value":100}]
            //[{"name":"2023-3","value":100}]
            //处理饼图的名称问题https://echarts.apache.org/examples/zh/editor.html?c=pie-simple
            let currentYAxis = Array.isArray(that.option.yAxis.data[0]) ? that.option.yAxis.data[0] : that.option.yAxis.data
            if (currentYAxis.length == 0) {
                let msg = `请求(${url})返回数据(data或yAxis)数据为空`
                that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                return
            }
            //将 [320,332,301,334,390,330,320] 转为：[{"name":"2023-3","value":100}]
            if (!(currentYAxis[0]||{}).hasOwnProperty("name")) {
                currentYAxis = that.fullYAxis(that.option.legend.data, currentYAxis)
            }

            //[{"name":"2023-3","value":100}]
            let s = {
                type: that.lowerType,
                stack: that.stack,
                barWidth: "8%",//柱子宽度
                barGap: "0%",
                label: { show: that.showLabel == "true" },
                labelLine: { show: that.showLabel == "true" },
                roseType: that.roseType[that.rType] || "",
                itemStyle: that.styles[that.rType] || {},
                data: currentYAxis
            }
            if (that.radius[that.rType]) { //处理空心
                s.radius = that.radius[that.rType]
            }
            that.option.series = []
            that.option.series.push(s)
            return true
        },
        resetLineBar(item, that) {
            let url = that.url
            that.option.legend.data = []
            that.option.tooltip.formatter = that.formatter.line_bar
            if (item.legend && Array.isArray(item.legend)) {
                that.option.legend.data = item.legend
            }
            if (!item.xAxis || !Array.isArray(item.xAxis) || item.xAxis.length == 0) {
                let msg = `请求(${url})返回数据中未包含xAxis或不是数组`
                that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                return false
            }
            if (!item.yAxis || !Array.isArray(item.yAxis) || item.yAxis.length == 0) {
                let msg = `请求(${url})返回数据中未包含yAxis或不是数组`
                that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                return false

            }
            that.option.series = []
            that.option.xAxis.data = item.xAxis

            //yAxis 为单个元素
            if (!Array.isArray(item.yAxis[0] || {})) {
                that.option.series.push({
                    type: that.lowerType,
                    stack: that.stack,
                    barWidth: "10%",//柱子宽度
                    barGap: "0%",
                    labelLine: { show: that.showLabel == "true" },
                    data: item.yAxi
                })
                return true

            }
            //yAxis 为多元素
            item.yAxis.forEach((y, i) => {
                let name = that.option.legend.data.length > i ? that.option.legend.data[i] : ""
                that.option.series.push({
                    name: name,
                    type: that.lowerType,
                    stack: that.stack,
                    barWidth: "10%",//柱子宽度
                    barGap: "0%",
                    labelLine: { show: true },
                    data: y
                })
            })
            return true
        },
        fullYAxis(xAxis, yAxis) { //转为[{"name":"2023-3","value":100}]格式
            if (xAxis.length == 0 || xAxis.length != yAxis.length) {
                return yAxis
            }
            let lst = []
            yAxis.forEach((y, i) => {
                lst.push({
                    name: xAxis[i],
                    value: y,
                })
            })
            return lst
        },
        IsPIE() {
            return this.lowerType == "pie"
        }
    },
}
</script>