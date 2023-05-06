<template>
    <div>
        <el-tag v-for="tag in tags" :key="tag" class="mx-1" closable :disable-transitions="false" @close="handleClose(tag)">
            {{ tag }}
        </el-tag>
        <el-input v-if="inputVisible" ref="input" v-model="inputValue" class="ml-1 w-20" 
            @keyup.enter="handleInputConfirm" @blur="handleInputConfirm" />
        <el-button v-else class="button-new-tag ml-1"  @click="showInput">
            + 添加
        </el-button>
    </div>
</template>
<script>
export default {
    props: {
        tags: {
            type: Array,
            default: []
        }
    },
    data() {
        return {
            inputValue: "",
            inputVisible: false,
        }
    },
    methods: {
        showInput() {
            this.inputVisible = true
            this.$nextTick(() => {
                this.$refs.input.focus()
            })
        },

        handleClose(tag) {
            this.tags.splice(this.tags.indexOf(tag), 1)
        },
        handleInputConfirm() {
            let that = this
            if (this.inputValue) {
                // this.$theia.http.xget(this.inputValue, {}).then(res => {
                //     var matches = data.match(/<title>(.*?)<\/title>/);
                //     if ((matches || []).length > 0) {
                //         that.tags.push(matches[0])
                //         return
                //     }
                //     that.tags.push(that.inputValue)
                // }).catch(err => {
                //     that.tags.push(that.inputValue)
                // });
                that.tags.push(that.inputValue)
            }
            this.inputVisible = false
            this.inputValue = ''


        }
    },

}
</script>
<style scoped></style>