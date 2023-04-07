<template>
    <div class="hy-order-detail">
        <van-row>
            <van-col span="12" class="status"><van-icon name="passed" /></van-col>
            <van-col span="12" class="label">完成</van-col>
        </van-row>
        <van-row v-for="item in list" :key="item" class="prod-item">
            <van-col span="24" class="spp"> {{item.spp}}<van-icon name="arrow" /></van-col>
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
                <van-button type="default" size="mini">加购物车</van-button>
            </van-col>
 
        </van-row>

        <van-cell-group inset>
            <van-cell title="订单编号" :value="order.no" />
            <van-cell title="下单时间" :value="order.order_time" />
            <van-cell title="支付方式" :value="order.pay_way" />
            <van-cell title="支付时间" :value="order.pay_time" />
          </van-cell-group>

          <van-cell-group inset class="logistics">
            <van-cell title="配送方式" :value="order.logistics" />
            <van-cell title="收货人" :value="contact.name" />
            <van-cell title="收货地址" :value="contact.addr" />
          </van-cell-group>
    </div>
</template>

<script>
export default {
    data() {
        return {
            list: [
                { id: "R03040687230009923",spp:"京东慧采直供", num: 1, img: "http://img13.360buyimg.com/n1/s800x800_jfs/t1/196489/29/29167/80251/63295a12Edbcac09a/1b18907d13c59abd.jpg", price: { symbol: "￥", integer: "98", decimal: "99", market: "168.59" }, label: "家庭疫情洗护礼包", sku: [{ name: "家庭礼包" }], url: "/product/item", tags: [{ name: "自营" }], status_label: "已完成" },
            ],
            order:{no:"R209348909885",order_time:"2023/4/7 9:16","pay_way":"在线支付",pay_time:"2023/4/7 9:18",logistics:"中通快递"},
            contact: { tag: "家", short: "四川省成都市", addr: "天府新区协和下街588号", tel: "158****0877", name: "杨磊" },
        }
    },
}
</script>

<style scoped>
.status {
    font-size: 2rem;
    text-align: right;
    padding: 4px;
    color: var(--van-success-color);
}

.label {
    font-size: 1.4rem;
    text-align: left;
    padding: 4px;
    padding-top: 10px;
    display: block;
    color: var(--van-success-color);
}
.prod-item {
    background-color: white;
    margin: 8px;
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
.logistics{
    margin-top: 16px;
}
</style>