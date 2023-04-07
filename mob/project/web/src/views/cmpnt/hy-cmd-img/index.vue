<template>
    <van-col :span="span">
    <div class="hy-cmdimg">
        <van-swipe :autoplay="3000" indicator-color="white" :loop="true" touchable>
            <van-swipe-item v-for="(banner, i) in banners" :key="i">
              <a :href="banner.url?banner.url:'#'"> <img :src="banner.img" /></a>
            </van-swipe-item>
        </van-swipe>
    </div>
    </van-col>
</template>
<script>
export default {
    props: {
        span: {
            type: String,
            default: 12,
        },
    },
    data() {
        return {
            banners:
                [
                    // { img: "https://file.mall2.18jiayou.com/core/a19bf161-bf7f-4fa9-ba2a-7b3a9c7d71ee.png" },
                    // {img:"http://img13.360buyimg.com/n1/s800x800_jfs/t1/196489/29/29167/80251/63295a12Edbcac09a/1b18907d13c59abd.jpg"}
                {img:"https://cdn.file.mall.18jiayou.com/ces-file/bannerImage/2023020617475734149c96dc844bb58e07fdd09a84fd2a.png"}
                ],
        }
    },
}
</script>
<style scoped>
.hy-cmdimg {
    margin: 8px;
}

.hy-cmdimg img {
    width: 100%;
    border-radius: 8px;
}
</style>