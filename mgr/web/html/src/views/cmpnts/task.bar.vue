<template>
    <div :style="{ margin: margin, height: height, width: width }">
        <div class="title"> {{ title }}</div>
        <el-row>
            <el-col :span="24 / rows" v-for="(item, i) in list" :key="item" class="panel">
                <el-alert v-if="i < count" :description="item.desc || item.memo" :title="item.title || item.name" :type="types[i % 4]"
                    closable="false" />
            </el-col>
        </el-row>
    </div>
</template>
<script>
export default {
    props: {
        title: {
            type: String,
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
        rows: {
            type: String,
            default: "2"
        },
        margin: {
            type: String,
            default: "32px 0px 16px 0px"
        },
        count:{
            type:String,
            default:10
        }
    },
    data() {

        return {
            types: ["success", "warning", "error", "info"],
            list: [{ title: "售后二期延期处理问题修复", desc: "2023/2/15上线" },
            { title: "商城取消订单处理", desc: "2023/4/15上线" },
            { title: "数字商品后端改造", desc: "2023/4/15上线" },
            { title: "淘金客服系统对接", desc: "2023/4/15上线" }],
        }
    },
    methods: {
        show(queryForm = {}) {
            if (!this.url) {
                return
            }
            //构建统计查询
            //数据查询
            let that = this
            this.$theia.http.post(this.url, queryForm).then(res => {
                that.list = res
            });
        }
    },

}
</script>

<style scoped>
.title {
    font-size: 1rem;
    font-weight: 600;
    margin-bottom: 8px;
}

.panel {
    padding: 5px 5px;
}

.panel .el-alert {
    text-align: left;
}

/deep/ .el-alert__close-btn {
    display: none;
}
</style>