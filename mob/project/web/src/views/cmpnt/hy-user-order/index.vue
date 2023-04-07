<template>
    <div>
        <van-cell-group inset clickable>
            <van-cell title="我的订单" value="全部" is-link size="large" to="/order/list" />
        </van-cell-group>
        <van-grid :border="false">
            <van-grid-item v-for="item in menus" :icon="item.icon" :text="item.name" :to="item.path" :badge="item.badge" />
        </van-grid>
    </div>
</template>
<script>

export default {
    name: "hy-user-order",
    data() {
        return {
            menus: [
                {
                    name: "待付款",
                    icon: "balance-pay",
                    path: "/order/list",
                    badge: "3"
                },
                {
                    name: "待发货",
                    icon: "idcard",
                    path: "/order/list",
                },
                {
                    name: "待收货",
                    icon: "logistics",
                    path: "/order/list",
                },
                {
                    name: "待确认",
                    icon: "comment-circle-o",
                    path: "/order/list",
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