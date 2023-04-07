<template>
    <div class="hy-order-cnfrm">
        <van-cell is-link center class="contact">
            <!-- 使用 title 插槽来自定义标题 -->
            <template #title>
                <div><van-tag type="danger" v-if="contact.tag">{{ contact.tag }}</van-tag> {{ contact.short }}</div>
                <div class="addr">{{ contact.addr }}</div>
                <div>{{ contact.name }}{{ contact.tel }}</div>
            </template>
        </van-cell>

        <van-row v-for="item in list" :key="item" class="prod-item">
            <van-col span="6">
                <span><img :src="item.img" /></span>
            </van-col>
            <van-col span="15" class="content">
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
                        <van-col span="12" class="stepper">
                            <van-stepper v-model="item.count" step="1" button-size="22" disable-input />
                        </van-col>
                    </van-row>
                </div>
            </van-col>

        </van-row>
        <van-cell title="商品金额" value="￥98.99" />
        <van-cell title="运费" value="￥0" />
        <van-cell title="优惠券" value="￥0" is-link to="index" />
        <van-row class="total">
            <van-col span="24"> 合计: ￥98.99</van-col>
        </van-row>

        <van-row class="pay-way">
            <van-cell title="支付方式" value="在线支付" is-link to="index" />
        </van-row>

        

        <van-submit-bar :price="3050" button-text="提交订单" @submit="onSubmit">
            <template #tip>
                你的收货地址不支持配送, <span @click="onClickLink">修改地址</span>
            </template>
        </van-submit-bar>
    </div>
</template>
<script>
export default {
    data() {
        return {
            contact: { tag: "家", short: "四川省成都市", addr: "天府新区协和下街588号", tel: "158****0877", name: "杨磊" },
            list: [
                { id: 1, count: 1, img: "http://img13.360buyimg.com/n1/s800x800_jfs/t1/196489/29/29167/80251/63295a12Edbcac09a/1b18907d13c59abd.jpg", price: { symbol: "￥", integer: "98", decimal: "99", market: "168.59" }, label: "家庭疫情洗护礼包", sku: [{ name: "家庭礼包" }], url: "/product/item", tags: [{ name: "自营" }] },
            ]
        }
    },
    methods: {
        onSubmit(){
            this.$router.push("/order/success")
        }
    },

}
</script>
<style scoped>
.hy-order-cnfrm .addr {
    font-size: 1.1rem;
    font-weight: 600;
}

.contact:before {
    position: absolute;
    right: 0;
    bottom: 0;
    left: 0;
    height: 3px;
    background: repeating-linear-gradient(-60deg, #1989fa, #1989fa 20%, transparent 0, transparent 25%, #ff6c6c 0, #ff6c6c 45%, transparent 0, transparent 50%);
    background-size: 90px;
    content: "";
}

.prod-item {
    margin: 8px 0px;
    background-color: white;
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

.price .van-stepper {
    border-radius: 40px;
}

.price .stepper {
    text-align: right;
}

/deep/ .prod-item .van-cell__value {
    color: red;
}
/deep/.van-submit-bar__text{
    text-align: left !important;
}
/deep/.van-submit-bar__text>span:first-child{
    display: none !important;
}
/deep/ .van-submit-bar__price .van-submit-bar__price-integer{
    font-size: 1.6rem;
}

.total{
    text-align: right;
    background-color: white;
    height: 32px;
    font-size: 1.1rem;
    padding: 16px;
}

.pay-way{
    margin: 8px 0px;
}
</style>