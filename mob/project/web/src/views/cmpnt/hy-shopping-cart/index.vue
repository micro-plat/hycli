<template>
    <div class="hy-shopping-cart">
        <van-row class="head">
            <van-col span="21">
                <span class="title">购物车</span>
                <span class="addr"><van-icon name="location-o" />{{ user.addr }}</span>
            </van-col>
            <van-col span="3" class="edit">编辑</van-col>
        </van-row>
        <van-row v-for="item in list" :key="item" class="cart-item">
            <van-col span="9" class="check-img">
                <van-checkbox v-model="item.checked"></van-checkbox>
                <span><img :src="item.img" @click="item.checked = !item.checked" /></span>
            </van-col>
            <van-col span="15" class="content">
                <div class="title">{{ item.label }}</div>
                <div class="sku" v-if="item.sku.length > 0"><span>{{ item.sku[0].name }}
                        <van-icon name="arrow-down" v-if="item.sku.length > 1" /></span>
                </div>
                <div class="tag" v-if="(item.tags || []).length > 0">
                    <van-tag plain type="primary" v-for="tag in item.tags" :key="tag">{{ tag.name }}</van-tag>
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

        <van-submit-bar :price="3050" button-text="去结算" @submit="onSubmit">
            <van-checkbox v-model="checked">全选</van-checkbox>
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
            user: { addr: "成都市天府新区协和下街588号" },
            list: [
                { id: 1, count: 1, img: "http://img13.360buyimg.com/n1/s800x800_jfs/t1/196489/29/29167/80251/63295a12Edbcac09a/1b18907d13c59abd.jpg", price: { symbol: "￥", integer: "98", decimal: "99", market: "168.59" }, label: "洗护礼包", sku: [], url: "/product/item", tags: [{ name: "自营" }] },
                { id: 1, count: 1, img: "http://img13.360buyimg.com/n1/s800x800_jfs/t1/123594/18/21645/283330/61f29f7dEa15539d1/dafd831cc28e7c37.png", price: { symbol: "￥", integer: "298", decimal: "99", market: "356.59" }, label: "海拔3000手撕牦牛肉198g中通包邮", sku: [{ name: "牛肉198g", id: 100 }], url: "/product/item" },
                { id: 1, count: 1, img: "https://cdn.fs.qxbpp.cn/fastdfs/file/v2/view/scm/scm20221104154225418ruaqsf069qzk", price: { symbol: "￥", integer: "28", decimal: "99", market: "98.59" }, label: "山真美鹿茸菌80g", sku: [{ name: "茸菌80g", id: 100 }], url: "/product/item" },
                { id: 1, count: 1, img: "https://img13.360buyimg.com/n1/s800x800_jfs/t1/77698/40/23782/91824/6400b067F9309866d/3aa5618a746c041f.jpg", price: { symbol: "￥", integer: "478", decimal: "19", market: "868.59" }, label: "米油礼包", sku: [{ name: "粮油", id: 100 }, { name: "大米", id: 100 }], url: "/product/item" },
            ]
        }
    },
    methods: {
        onSubmit() {
            this.$router.push("/order/cnfrm")
        }
    },

}
</script>
<style scoped>
.hy-shopping-cart {
    margin: 8px;

}

.hy-shopping-cart .head {
    padding: 8px;
    border-radius: 8px;
    line-height: 36px;
    vertical-align: middle;


}

.hy-shopping-cart .head .title {
    font-size: 1.2rem;
    font-weight: 800;
    margin: 0px 6px;
}

.hy-shopping-cart .head .addr {
    background-color: #e9e9ee;
    border-radius: 8px;
    padding: 0px 6px;
}

.hy-shopping-cart .head .edit {
    font-weight: 500;
    text-align: right;
    vertical-align: bottom;
}

.cart-item {
    border-radius: 8px;
    background-color: white;
    margin: 8px;


}

.cart-item .check-img {
    display: flex;
    text-align: center;
    vertical-align: middle;
    padding-left: 10px;
    padding-right: 12px;
    line-height: 1;
}


.cart-item img {
    width: 80px;
    height: 80px;
    border-radius: 8px;
    margin-top: 6px;
    margin-left: 10px;
    object-fit: cover;

}

.hy-shopping-cart .content {
    padding: 10px;
}

.hy-shopping-cart .content .title {
    font-size: 1rem;
    font-weight: 500;
}

.hy-shopping-cart .content .sku span {
    background-color: #f1f1f1;
    padding: 1px 10px;
    border-radius: 8px;
    font-size: 0.7rem;
}

.hy-shopping-cart .content .sku .van-icon {
    font-size: 0.1rem;
    margin-left: 2px;
}



.hy-shopping-cart .price {
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

>>>.van-tag--plain {
    margin-bottom: 2px;
}
</style>
