<template>
    <div class="hy-serachbar">
        <van-row>
            <van-col span="24">
        <van-search v-model="content" show-action  placeholder="请输入搜索关键词" @search="onSearch">
            <template #action>
                <div @click="onSearch">搜索</div>
            </template>
        </van-search>
        </van-col>
        </van-row>
    </div>
</template>
<script>
export default {
    name: "hy-searchbar",
    data() {
        return {
            content: ""
        }
    },
    methods: {
        onSearch(){
            this.$emit("onsearch",this.content)
        }
    }
}
</script>
<style scoped>

</style>