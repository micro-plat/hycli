<template>
    <div>

        <van-search v-model="value" placeholder="请输入搜索关键词" />
        <van-tabs :active="active">
            <van-tab title="全部"></van-tab>
            <van-tab title="待付款"></van-tab>
            <van-tab title="待发货/使用"></van-tab>
            <van-tab title="已完成"></van-tab>
        </van-tabs>
        <van-row v-for="item in list" :key="item" class="prod-item">
            <van-col span="18" class="spp"> {{item.spp}}<van-icon name="arrow" /></van-col>
            <van-col span="6" class="status">
                <span>{{ item.status_label }}</span>
            </van-col>
            <van-col span="6">
                <span><img :src="item.img" /></span>
            </van-col>
            <van-col span="18" class="content" @click="gotoDetail">
                <div class="title">{{ item.label }}</div>
                <div class="sku" v-if="item.sku.length > 0"><span>{{ item.sku[0].name }}</span>
                </div>
                <div class="price">
                    <van-row>
                        <van-col span="12">
                            <span class="symbol" v-if="item.price.symbol">{{ item.price.symbol }}</span>
                            <span class="integer">{{ item.price.integer }}</span>
                            <span class="decimal" v-if="item.price.decimal">.{{ item.price.decimal }}</span>
                        </van-col>
                        <van-col span="12" class="num" >
                            <span>X {{ item.num }}</span>
                        </van-col>
                    </van-row>
                </div>
            </van-col>
            <van-col span="24" class="toolbar">
                <van-button type="default" size="mini">删除订单</van-button>
                <van-button type="default" size="mini">退换/售后</van-button>
                <van-button type="default" size="mini">再来一单</van-button>
            </van-col>
 
        </van-row>
    </div>
</template>

<script>
export default {
    data() {
        return {
            active: 0,
            list: [
                { id: "R03040687230009923",spp:"京东慧采直供", num: 1, img: "http://img13.360buyimg.com/n1/s800x800_jfs/t1/196489/29/29167/80251/63295a12Edbcac09a/1b18907d13c59abd.jpg", price: { symbol: "￥", integer: "98", decimal: "99", market: "168.59" }, label: "家庭疫情洗护礼包", sku: [{ name: "家庭礼包" }], url: "/product/item", tags: [{ name: "自营" }], status_label: "已完成" },
            ]
        }
    },
    methods: {
        gotoDetail(){
            this.$router.push("/order/detail")
        }
    },
}
</script>

<style scoped>
.prod-item {
    margin: 8px;
    background-color: white;
    border-radius: 8px;

}

.prod-item img {
    width: 80px;
    height: 80px;
    border-radius: 8px;
    margin-top: 6px;
    margin-left: 10px;
    object-fit: cover;

}

.content {
    padding: 10px;
}

.content .title {
    font-size: 1rem;
    font-weight: 500;
}

.content .sku span {
    background-color: #f1f1f1;
    padding: 1px 10px;
    border-radius: 8px;
    font-size: 0.7rem;
}

.content .sku .van-icon {
    font-size: 0.1rem;
    margin-left: 2px;
}



.price {
    text-align: left;
    font-size: 1.2rem;
    font-weight: 500;
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
.prod-item .spp{
    font-size: 1rem;
    padding: 8px;
    font-weight: 600;
}
.prod-item .status {
    text-align: right;
    font-size: 0.96rem;
    padding: 8px;
    color: red;
}
.num{
    color: #c1c1c1;
    font-size: 0.9rem;
    font-weight: 400;
    text-align: right;
}
.toolbar{
    text-align: right;
    padding: 8px;

}
</style>