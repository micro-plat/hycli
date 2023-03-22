<template>
    <div :id="unq" class="chart-line" :style="{ height: height, width: width }"></div>
</template>
<script>
import * as echarts from 'echarts';
export default {
    name: "ChartLine",
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
    },
    data() {
        return {
            option: {
                title: {
                    text: this.title,
                },
                tooltip: {
                    trigger: 'axis'
                },
                legend: {
                    data: ['Email', 'Union Ads', 'Video Ads', 'Direct']
                },
                grid: {
                    left: '3%',
                    right: '4%',
                    bottom: '3%',
                    containLabel: true
                },
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
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
    mounted() {
        this.showChart()
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
            if (!item.xAxis || !Array.isArray(item.xAxis)) {
                let msg = `请求(${url})返回数据中未包含xAxis`
                that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                return false
            }
            if (!item.yAxis || !Array.isArray(item.yAxis)) {
                let msg = `请求(${url})返回数据中未包含yAxis`
                that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                return false
            }
            that.option.xAxis.data = item.xAxis
            that.option.series = []
            let success = true
            item.yAxis.forEach((i, y) => {
                if (!Array.isArray(y)) {
                    let msg = `请求(${url})返回数据中yAxis应是二维数组`
                    that.$notify.error({ title: '失败', message: msg, duration: 5000 })
                    success = false
                    return false
                }
                that.option.series.push({
                    name: this.option.legend.data[i],
                    type: 'line',
                    stack: 'Total',
                    data: y
                })
            });
            return success
        },
    },
}
</script>
<style>
.chart-line {
    margin-top: 48px;
}
</style>