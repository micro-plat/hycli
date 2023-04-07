<template>
    <div class="hy-tipbar">
        <div class="head">
            <van-row>
                <van-col span="24" class="title"> <van-icon size="3rem" name="passed" color="var(--van-success-color)" />
                    <span>支付成功</span></van-col>

            </van-row>
        </div>
        <van-row class="nav">
            <van-col span="24">
                <a @click="goto('/home')">返回首页</a>
                <a @click="goto('/order/list')">查看订单</a>
            </van-col>
        </van-row>

    </div>
</template>
<script>
export default {
    methods: {
        goto(url) {
            this.$router.push(url)
        }
    },
}

</script>
<style scoped>
.hy-tipbar .head {
    margin: 8px;
    margin-top: 100px;
    vertical-align: middle;
}

.hy-tipbar .head .title {
    text-align: center;
    vertical-align: middle;
    font-size: 1.2rem;
    padding-top: 0.5rem;
    color: var(--van-success-color);
}

.hy-tipbar .head .title span {
    line-height: 3rem;
    display: block;
}

.hy-tipbar .nav {
    text-align: center;
}

.hy-tipbar .nav a {
    margin: 8px;
}</style>