<template>
    <div class="hy-king">
        <van-grid square  :border="false" :column-num="5">
            <van-grid-item v-for="item in menus" :key="item" :icon="item.icon" :text="item.name" :to="item.path" />
          </van-grid>
    </div>
   
</template>

<script>
export default {
    data(){
        return{
            menus: [
                {
                    name: "美食",
                    icon: "paid",
                    path: "/card/list",
                },
                {
                    name: "酒店",
                    icon: "star-o",
                    path: "/fav/lst",
                },
                {
                    name: "休闲娱乐",
                    icon: "balance-list-o",
                    path: "/invoice/list",
                },
                {
                    name: "打车",
                    icon: "phone-o",
                    path: "/refund/list",
                },
                {
                    name: "电影演出",
                    icon: "phone-o",
                    path: "/refund/list",
                }, {
                    name: "饮品小吃",
                    icon: "paid",
                    path: "/card/list",
                },
                {
                    name: "火车机票",
                    icon: "star-o",
                    path: "/fav/lst",
                },
                {
                    name: "生活服务",
                    icon: "balance-list-o",
                    path: "/invoice/list",
                },
                {
                    name: "购车",
                    icon: "phone-o",
                    path: "/refund/list",
                },
                {
                    name: "特价团购",
                    icon: "phone-o",
                    path: "/refund/list",
                }
            ],
        }
    }
}
</script>

<style scoped>
.hy-king{
    margin: 8px;
}
</style>