<template>
    <div>
        <router-view />
        <van-tabbar v-model="active" :active-color="conf.active_color">
            <van-tabbar-item :name="item.name"  v-for="(item, index) in menus" :key="index" :to="item.path" :icon="active!=item.name ? item.icon : item.aicon">
                <span>{{ item.name }}</span>
            </van-tabbar-item>
        </van-tabbar>
    </div>
</template>
<script>

export default {
    name: "hy-tabbar",
    data() {
        return {
            conf: {
                active_color: "#333"
            },
            active:"首页",
            menus: [
                {
                    name: "首页",
                    icon: "wap-home-o",
                    aicon: "wap-home",
                    path: "/home",
                },
                {
                    name: "分类",
                    icon: "qr-invalid",
                    aicon: "qr-invalid",
                    path: "/catgory",
                },
                {
                    name: "购物车",
                    icon: "shopping-cart-o",
                    aicon: "shopping-cart",
                    path: "/shopping",
                },
                {
                    name: "个人",
                    icon: "manager-o",
                    aicon: "manager",
                    path: "/user",
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