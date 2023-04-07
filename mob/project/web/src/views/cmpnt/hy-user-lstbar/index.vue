<template>
    <div class="userlstbar">
        <van-cell size="large" v-for="item in menus" :title="item.name" :icon="item.icon"
         is-link  :to="item.path" />
    </div>
</template>
<script>

export default {
    name: "hy-user-lstbar",
    data() {
        return {
            menus: [
                {
                    name: "我的地址",
                    icon: "location-o",
                    path: "/user/address",
                },
                {
                    name: "常见问题",
                    icon: "question-o",
                    path: "/service/question",
                },
                {
                    name: "联系客服",
                    icon: "phone-circle-o",
                    path: "/service/contact",
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
<style>
.userlstbar{
    margin-top: 16px;
}
</style>