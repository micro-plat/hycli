<template>
    <div class="hy-banner">
        <van-row>
            <van-col span="24">
                <van-swipe class="my-swipe" :autoplay="3000" indicator-color="white" :loop="true"  touchable>
                <van-swipe-item v-for="(banner,i) in banners" :key="i">
                   <img :src="banner.img" @click="goto(banner.url)"/>
                </van-swipe-item>
            </van-swipe>
        </van-col>
          </van-row>
        
    </div>
</template>
<script>

export default {
    data() {
        return {
            banners:
                [
                    
                    { img: "https://file.mall2.18jiayou.com/core/ee9559c3-fc5c-4b1f-9cbe-2695a52d73e1.jpg" },
                    { img: "https://file.mall2.18jiayou.com/core/5070816b-de3d-4732-bc61-0a1d0d9ddf68.png" },
                ],
        }
    },
    mounted() {
    },
}
</script>
<style>
.my-swipe img {
    width: 100%;
}
.hy-banner{
    margin: 8px;
}
</style>
