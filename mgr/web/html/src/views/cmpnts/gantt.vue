<template>
    <div class="gantt">
        <div v-if="mode == 'header'" class="gantt-header">
            <div class="gantt-title">甘特图
                <span class="gantt-tools" @click="select(x % conf.progressColor.length)" v-for="(v, x) in labels"
                    :key="v"><span class="gantt-rdo"
                        :style="{ opacity: (selectIdx == -1 || selectIdx == x % conf.progressColor.length) ? 1 : 0.5, backgroundColor: this.conf.progressColor[x % conf.progressColor.length].color }"></span>{{
                            v }}</span>
            </div>
            <div class="gantt-container">
                <div class="gantt-line gantt-hd" v-for="v in headerNo" :key="v" :style="{ left: v.left }" v-html="v.fmt">
                </div>
            </div>
        </div>
        <div v-if="mode == 'line'" class="gantt-container">
            <div class="gantt-cell gantt-float gantt-process" :title="x.title" v-for="(x) in dataItem" :key="x"
                style="bottom:4px" :style="{ left: x.left, width: x.width }">
                <div class="gantt-bar" v-if="showBar(x)"
                    :style="{ backgroundColor: conf.progressColor[x.idx % conf.progressColor.length].color }">
                </div>
            </div>
            <div class="gantt-line" v-for="v in headerNo" :key="v" :style="{ left: v.left }"></div>
            <div class="gantt-line" :style="{ left: headerNo.length * 20 + 'px' }"></div>
            <div class="gantt-line cr" :style="todayStyle"></div>

        </div>
    </div>
</template>
<script>
export default {
    props: {
        mode: {
            type: String,
            default: "line"
        },
        header: {
            type: Boolean,
        },
        data: {
            types: Array,
            default: []
        },
        labels: {
            types: Array,
            default: []
        },
        dates: {
            types: Array,
            default: []
        },
        maxNo: {
            types: Number,
            default: 20,
        },
        row: {
            types: Object,
            default: {}
        },
        ganttIdx: {
            types: Number,
            default: -1,
        }
    },

    data() {
        return {
            conf: {
                progressColor: this.$theia.env.conf.progress || []
            },
            headerNo: [],
            dataItem: [],
            todayStyle: {},
            selectIdx: -1,
        }
    },
    watch: {
        'data'() {
            this.loadHeaderNo()
            this.loadLine()
        }
    },
    mounted() {

        this.loadHeaderNo()
        this.loadLine()
    },
    computed: {
        showBar: function () {
            let that = this
            return (x) => {
                let value = that.ganttIdx == -1 || that.ganttIdx == (x.idx % that.conf.progressColor.length)
                return value
            }
        }
    },
    methods: {
        select(idx) {
            this.selectIdx = this.selectIdx == idx ? -1 : idx;
            this.$emit('selectChange', this.selectIdx)
        },
        loadHeaderNo() {
            let that = this
            let range = this.$theia.xdate.formatRange(this.getDateRange(), "ddW", false)
            range.forEach((v, i) => {
                v.left = i * 20 + 'px'
            })
            this.headerNo = range
        },
        getTodayStyle(ranges) {
            let offset = this.getOffset(ranges, new Date())
            if (offset < 0) {
                return { display: "none" }
            }
            return { left: (offset) * 20 + 10 + 'px' }
        },
        getDateRange() {
            let headRange = this.getDateMinMax(this.data, this.dates)
            if (this.$theia.xdate.diffDay(headRange.min, new Date()) < this.maxNo) {
                headRange.max = this.$theia.xdate.addDay(headRange.min, this.maxNo)
            }
            let ranges = this.$theia.xdate.fillRange(headRange.min, headRange.max)
            return ranges.slice(0, this.maxNo)
        },
        loadLine() {
            if (this.mode != "line") {
                return
            }
            //构建甘特图头部日期
            let headRange = this.getDateMinMax(this.data, this.dates)
            if (this.$theia.xdate.diffDay(headRange.min, new Date()) < this.maxNo) {
                headRange.max = this.$theia.xdate.addDay(headRange.min, this.maxNo)
            }
            let ranges = this.$theia.xdate.fillRange(headRange.min, headRange.max)
            this.todayStyle = this.getTodayStyle(ranges)

            //构建任务进度数据,日期范围，当前数据源，需要处理的日期，日期标签名
            let ganntItems = this.buildDataRange(ranges, this.row, this.dates, this.labels)

            this.dataItem = []
            let that = this
            ganntItems.forEach((v, i) => {
                if (v.end > that.maxNo) {
                    return
                }
                let current = Object.assign({}, v)
                current.left = (v.end) * 20 + 'px'
                current.width = (v.xlen * 20) + 'px'
                current.title = that.$theia.xdate.format(v.value, "yyyy/MM/dd 周W ") + v.label
                that.dataItem.push(current)
            })
        },


        buildDataRange(ranges, row, dates, labels) {
            let ganntIstems = []
            let that = this
            //获取数据源中的每一个日期值，从日期范围中查找日期偏移值
            dates.forEach((v, i) => {
                //处理单日期，最终描绘为一个点(长度为1)
                if (row[v]) {
                    let current = new Date(row[v])
                    let offset = that.getOffset(ranges, current)
                    if (offset >= 0) {
                        let idx = { end: offset }
                        idx.label = labels.length > i ? labels[i] : "";
                        idx.idx = i
                        idx.value = current
                        idx.xlen = 1
                        ganntIstems.push(idx)
                    }
                    return
                }

                //处理起始日期，最终描绘为一条线(长度>=1)
                let start = row["start_" + v]
                let finish = row["finish_" + v]

                if (start && finish) {
                    let left = that.getOffset(ranges, new Date(start), true)
                    let right = that.getOffset(ranges, new Date(finish), false)

                    left = left < 0 ? 0 : left
                    right = right < 0 ? left : right
                    if (right >= left && left >= 0) {
                        let idx = { end: left }
                        idx.label = labels.length > i ? labels[i] : "";
                        idx.idx = i
                        idx.value = new Date(start)
                        idx.finish = new Date(finish)
                        idx.xlen = right - left + 1
                        ganntIstems.push(idx)
                    }
                }
            })
            return ganntIstems
        },
        getDateMinMax(rows, dates) {
            let max = new Date()
            let min = new Date()
            rows.forEach(rw => {
                let row = rw.__raw || {}
                dates.forEach(d => {
                    let lst = [d, "start_" + d, "finish_" + d]
                    lst.forEach(m => {
                        if (row[m]) {
                            let current = new Date(row[m])
                            min = current < min ? current : min
                            max = current > max ? current : max
                        }
                    })

                })
            })
            if (this.$theia.xdate.diffDay(max, new Date()) < 0) {
                min = this.$theia.xdate.diffDay(min, new Date()) > 7 ? this.$theia.xdate.addDay(new Date(), -7) : min
            }

            return { min: min, max: max }
        },
        //指定日期数组，判断开始结束日期在数组中的位置
        getOffset(dateRange, current, fmin = true) {
            let offset = -1
            if (
                dateRange.length == 0 ||
                !this.$theia.xdate.valid(current)
            ) {
                return offset;
            }
            let min = dateRange.length
            let max = 0
            dateRange.forEach((v, i) => {
                let vidx = this.$theia.xdate.diffDay(this.$theia.xdate.toDay(v),
                    this.$theia.xdate.toDay(current))
                min = vidx < min ? vidx : min
                max = vidx > max ? vidx : max

                if (vidx == 0) {
                    offset = i;
                }
            });
            if (offset == -1) {
                return fmin ? min : (max >= dateRange.length ? dateRange.length - 1 : max)
            }
            return offset;
        }

    },
}
</script>
<style>
:root {

    --gant-cell-hd-width: 20px;
    --gant-cell-hd-height: 22px;
    --gant-cell-hd-lineheight: 22px;
    --gant-cell-width: 20px;
    --gant-cell-height: 6px;
    --gant-cell-lineheight: 6px;
}

.gantt {
    width: 100%;
}

.gantt-header .gantt-title {
    width: 100%;
    height: 60px;
}

.gantt-container {
    display: flex;
}

.gantt-tools {
    font-weight: normal;
    font-size: 0.1rem;
}

.gantt-rdo {
    border-radius: 2px;
    width: 6px;
    height: 6px;
    display: inline-block;
    margin: auto 4px;
}

.gantt-cell {
    box-sizing: border-box;
    height: var(--gant-cell-hd-height);
    line-height: var(--gant-cell-hd-lineheight);
    text-align: center;
    vertical-align: middle;
    border: 1px solid #eee;
    font-weight: 600;
    border-right-width: 0;

}

.gantt-hd {
    width: var(--gant-cell-hd-width);
    border-width: 0px;
    border-width: 0 !important;
    top: 34px !important;
    cursor: pointer;
    line-height: 0.8rem;
    font-weight: 600;
    margin-top: 4px;
}

.gantt-cell:last-child {
    border-right-width: 1px;
    border-top-right-radius: 2px;
    border-bottom-right-radius: 2px;
}

.gantt-cell:first-child {
    border-top-left-radius: 2px;
    border-bottom-left-radius: 2px;
}

.gantt-float {
    position: relative;
}

.gantt-cell.gantt-float.gantt-process {
    border-width: 0;
    font-weight: normal;
    font-size: 0.1rem;
    position: absolute;
    height: 100%;

}

.gantt-cell.gantt-float.gantt-process .gantt-label {
    height: 18px;
    font-size: 0.1rem;
    float: left;
    margin-top: 2px;
    width: 100%;
    text-align: right;
    min-width: 50px;
}

.gantt-line {
    box-sizing: border-box;
    position: absolute;
    top: 0;
    z-index: -1;
    height: 102%;
    border-left: 1px solid #e9ecef;
}

.gantt-line.cr {
    border-left: 1px solid red;
}

.gantt-cell.gantt-float.gantt-process .gantt-bar {
    box-sizing: border-box;
    height: var(--gant-cell-height);
    line-height: var(--gant-cell-lineheight);
    width: 100%;
    border-radius: 2px;
    margin-top: 24px;
}
</style>