<template>
    <div>
        <van-cell-group inset clickable>
            <van-cell title="常用功能" />
        </van-cell-group>
        <van-grid :border="false">
            <van-grid-item v-for="item in menus" :key="item" :icon="item.icon" :text="item.name" :to="item.path" />
        </van-grid>
    </div>
</template>
<script>

export default {
    name: "hy-user-king",
    data() {
        return {
            menus: [
                {
                    name: "我的卡包",
                    icon: "paid",
                    path: "/card/list",
                },
                {
                    name: "我的收藏",
                    icon: "star-o",
                    path: "/fav/lst",
                },
                {
                    name: "发票",
                    icon: "balance-list-o",
                    path: "/invoice/list",
                },
                {
                    name: "售后",
                    icon: "phone-o",
                    path: "/refund/list",
                }
            ],

        }
    },
    mounted() {
        this.query()
    },
    methods: {
        query() {
        },

    }
}
</script>