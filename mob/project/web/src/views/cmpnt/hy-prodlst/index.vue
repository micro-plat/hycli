<template>
    <div class="hy-prodlst">
        <van-list :loading="conf.loading" :error="conf.error" finished-text="没有更多了" error-text="请求失败，点击重新加载" @load="onLoad">
            <van-row gutter="20">
                <van-col span="24" ref="head" class="head">
                    <van-sticky :container="head">
                        <van-row>
                            <van-col class="item" span="6" v-for="cat in commend.catgory" :key="cat" @click="query(cat)">
                                <van-col class="title" span="24">{{ cat.title }}</van-col>
                                <van-col class="name" span="24">{{ cat.name }}</van-col>
                            </van-col>
                        </van-row>
                    </van-sticky>
                </van-col>
                <van-col class="prod-item" @click="goto(prod.url)" :span="12" v-for="prod in commend.prods" :key="prod">
                    <van-row>
                        <van-col span="24" class="img">
                            <img :src="prod.img" /></van-col>
                        <van-col span="24" class="label">{{ prod.label }}</van-col>
                        <van-col span="24" class="price">
                            <span class="symbol" v-if="prod.price.symbol">{{ prod.price.symbol }}</span>
                            <span class="integer">{{ prod.price.integer }}</span>
                            <span class="decimal" v-if="prod.price.decimal">.{{ prod.price.decimal }}</span>
                            <span class="market" v-if="prod.price.market">{{ prod.price.symbol }}{{ prod.price.market
                            }}</span>
                        </van-col>
                    </van-row>
                </van-col>
            </van-row>
        </van-list>
        <div style="height:100px"></div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            conf: {
                loading: false,
                error: null,
            },
            commend: {
                catgory: [
                    { title: "优选推荐", name: "猜你喜欢", value: "100" },
                    { title: "美味狂欢", name: "解馋零食", value: "200" },
                    { title: "美妆护肤", name: "打牌钜惠", value: "300" },
                    { title: "厨具家居", name: "品质生活", value: "400" },

                ],
                more: "/product",
                prods: [
                    { img: "http://img13.360buyimg.com/n1/s800x800_jfs/t1/196489/29/29167/80251/63295a12Edbcac09a/1b18907d13c59abd.jpg", price: { symbol: "￥", integer: "98", decimal: "99", market: "168.59" }, label: "洗护礼包" ,url:"/product/item"},
                    { img: "http://img13.360buyimg.com/n1/s800x800_jfs/t1/123594/18/21645/283330/61f29f7dEa15539d1/dafd831cc28e7c37.png", price: { symbol: "￥", integer: "298", decimal: "99", market: "356.59" }, label: "海拔3000手撕牦牛肉198g中通包邮" ,url:"/product/item"},
                    { img: "https://cdn.fs.qxbpp.cn/fastdfs/file/v2/view/scm/scm20221104154225418ruaqsf069qzk", price: { symbol: "￥", integer: "28", decimal: "99", market: "98.59" }, label: "山真美鹿茸菌80g" ,url:"/product/item"},
                    { img: "https://img13.360buyimg.com/n1/s800x800_jfs/t1/77698/40/23782/91824/6400b067F9309866d/3aa5618a746c041f.jpg", price: { symbol: "￥", integer: "478", decimal: "19", market: "868.59" }, label: "米油礼包",url:"/product/item" },
                ]
            }

        }
    },
    methods: {
        onLoad() {

        },
        query(catgory) {

        },
        goto(url){
            this.$router.push(url)
        }
    },
}
</script>
<style scoped>
.hy-prodlst {
    margin: 0px 8px;
    border-radius: 8px;

}

.hy-prodlst .head {
    line-height: 1.4rem;
    padding: 4px 4px;
    color: var(--van-grid-item-text-color);
    margin-bottom: 10px;
    font-size: 1.1rem;
    text-align: center;
    background-color: white;
}

.hy-prodlst .head .title,
.name {
    padding: 0px;
    margin: 0px;
}

.hy-prodlst .head .name {
    font-size: 0.4rem;
}

.hy-prodlst .head .item:hover,
.head .item:active {
    color: red;
}

.hy-prodlst .prod-item {
    padding: 0px 8px;
    margin: 8px 0px;
    background-color: white;
}

.hy-prodlst img {
    width: 100%;
    border-radius: 8px;
}

.hy-prodlst .label,
.price {
    text-align: left;
    padding: 4px;
}

.hy-prodlst .price {
    text-align: left;
    font-size: 1.2rem;
    color: red;
}


.price .symbol,
.decimal,
.market {
    font-size: 1rem;
}

.price .market {
    margin-left: 6px;
    color: var(--van-gray-6);
    text-decoration: line-through;
}

.van-sticky {
    margin-left: 0px;
    margin-right: 0px;
    background-color: white;
}</style>