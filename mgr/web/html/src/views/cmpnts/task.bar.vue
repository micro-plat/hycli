<template>
    <div v-if="list.length > 0 " :style="{ margin: margin, height: height, width: width }">
        <div class="title"> {{ title }} <span>
                <Back v-if="redirect == 1" @click="add(-1)" />
                <Right v-if="redirect == 0 && pc > pi" @click="add(1)" />
            </span></div>
        <el-row>

            <el-col :span="24 / rows" v-for="(item, i) in list" :key="item" class="panel"
                v-show="i >= (pi - 1) * count && i < pi * count">
                <el-alert :description="item.desc || item.memo" :title="item.title || item.name" :type="types[i % 4]"
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
            default: "8px 0px 16px 0px"
        },
        count: {
            type: String,
            default: 10
        }
    },
    data() {

        return {
            pi: 1,
            pc: 1,
            redirect: 0,
            types: ["success", "warning", "error", "info"],
            list: [],
        }
    },
    methods: {
        add(i) {
            this.pi = this.pi + i
            if (this.pi == this.pc) {
                this.redirect = 1
            }
            if (this.pi == 1 && this.pc > 1) {
                this.redirect = 0
            }
        },
        show(queryForm = {}) {
            if (!this.url) {
                return
            }
            //构建统计查询
            //数据查询
            let that = this
            this.$theia.http.post(this.url, queryForm).then(res => {
                that.list = res
                that.pc = Math.ceil(that.list.length / that.count)
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

.title span {
    float: right;
    font-size: 0.8rem;
}

.title span * {
    cursor: pointer;
    width: 1.2em;
    height: 1.2em;
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