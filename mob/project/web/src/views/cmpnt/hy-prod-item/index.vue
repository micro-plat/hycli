<template>
    <div class="hy-proditem">
        <!-- 图片轮播 -->
        <van-swipe class="my-swipe" :autoplay="3000" indicator-color="white" :loop="true" touchable>
            <van-swipe-item v-for="(img, i) in prod.imgs" :key="i">
                <img :src="img.url" />
            </van-swipe-item>
            <template #indicator="{ active, total }">
                <div class="custom-indicator">{{ active + 1 }}/{{ total }}</div>
            </template>
        </van-swipe>
        <!-- 价格信息 -->
        <van-row class=" price-info">
            <van-col span="24" class="price">
                到手价
                <span class="symbol" v-if="prod.price.symbol">{{ prod.price.symbol }}</span>
                <span class="integer">{{ prod.price.integer }}</span>
                <span class="decimal" v-if="prod.price.decimal">.{{ prod.price.decimal }}</span>
                <span class="market" v-if="prod.price.market">{{ prod.price.symbol }}{{ prod.price.market
                }}</span>
            </van-col>
            <van-col span="24" class="label">
                {{ prod.name }}
            </van-col>
        </van-row>
        <!-- 商品信息 -->
        <van-cell-group class="prod-info">
            <van-cell title="已选" :value="prod.checkSKU.name" is-link />
            <van-cell title="送至" :value="user.addr" is-link />
            <van-cell title="运费" :value="prod.carriageFree" />
        </van-cell-group>

        <!-- 详情 -->
        <van-sidebar class="prod-detail">
            <van-sidebar-item title="详情">
                <div v-html="prod.details"></div>
            </van-sidebar-item>
        </van-sidebar>

        <van-action-bar>
            <van-action-bar-icon icon="chat-o" text="客服" @click="onClickIcon" />
            <van-action-bar-icon icon="cart-o" text="购物车" @click="onClickIcon" />
            <van-action-bar-icon icon="shop-o" text="店铺" @click="onClickIcon" />
            <van-action-bar-button type="warning" text="加入购物车" />
            <van-action-bar-button type="danger" text="立即购买" @click="onBuyNow" />
        </van-action-bar>
    </div>
</template>
<script>
export default {
    data() {
        return {
            active: 0,
            prod: {
                name: "京东京造软管枕黑科技食品级PE软管填充高度可调缓可机洗国标A类 60*40cm",
                skus: [],
                details: "<style>.ssd-module-wrap{position:relative;margin:0 auto;width:750px;text-align:left;background-color:#fff}.ssd-module-wrap .ssd-module,.ssd-module-wrap .ssd-module-heading{width:750px;position:relative;overflow:hidden}.ssd-module-wrap .ssd-floor-activity{max-height:380px;overflow:hidden}.ssd-module-wrap .ssd-floor-shopPrior{max-height:900px;overflow:hidden}.ssd-module-wrap .ssd-floor-authorize{max-height:300px;overflow:hidden}.ssd-module-wrap .ssd-floor-reminder{max-height:4000px;overflow:hidden}.ssd-module-wrap .ssd-module{background-repeat:no-repeat;background-position:left top;background-size:100% 100%}.ssd-module-wrap .ssd-module-heading{background-repeat:no-repeat;background-position:left center;background-size:100% 100%}.ssd-module-wrap .ssd-module-heading .ssd-module-heading-layout{display:inline-block}.ssd-module-wrap .ssd-module-heading .ssd-widget-heading-ch{float:left;display:inline-block;margin:0 6px 0 15px;height:100%}.ssd-module-wrap .ssd-module-heading .ssd-widget-heading-en{float:left;display:inline-block;margin:0 15px 0 6px;height:100%}.ssd-module-wrap .ssd-widget-circle,.ssd-module-wrap .ssd-widget-line,.ssd-module-wrap .ssd-widget-pic,.ssd-module-wrap .ssd-widget-rectangle,.ssd-module-wrap .ssd-widget-table,.ssd-module-wrap .ssd-widget-text,.ssd-module-wrap .ssd-widget-triangle{position:absolute;overflow:hidden}.ssd-module-wrap .ssd-widget-rectangle{box-sizing:border-box;-moz-box-sizing:border-box;-webkit-box-sizing:border-box}.ssd-module-wrap .ssd-widget-triangle svg{width:100%;height:100%}.ssd-module-wrap .ssd-widget-table table{width:100%;height:100%}.ssd-module-wrap .ssd-widget-table td{position:relative;white-space:pre-line;word-break:break-all}.ssd-module-wrap .ssd-widget-pic img{display:block;width:100%;height:100%}.ssd-module-wrap .ssd-widget-text{line-height:1.5;word-break:break-all}.ssd-module-wrap .ssd-widget-text span{display:block;overflow:hidden;width:100%;height:100%;padding:0;margin:0;word-break:break-all;word-wrap:break-word;white-space:normal}.ssd-module-wrap .ssd-widget-link{position:absolute;left:0;top:0;width:100%;height:100%;background:0 0;z-index:100}.ssd-module-wrap .ssd-cell-text{position:absolute;top:0;left:0;right:0;width:100%;height:100%;overflow:auto}.ssd-module-wrap .ssd-show{display:block}.ssd-module-wrap .ssd-hide{display:none}.ssd-module-wrap .M16686466234081{width:750px; background-color:#e9e9e9; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/121643/2/33811/60480/637586cfEb0561c67/b3e464e7c3ab99a9.jpg); height:1202px}\n.ssd-module-wrap .M16686466234452{width:750px; background-color:#d7d7d7; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/178736/35/30761/26849/637586cdE8198b5f9/713d4cb59b765bd4.jpg); height:473px}\n.ssd-module-wrap .M16686466234713{width:750px; background-color:#f2f2f2; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/88124/6/33438/115286/637586cfE728f4395/591b6a43ea0bee36.jpg); height:1166px}\n.ssd-module-wrap .M16686466234914{width:750px; background-color:#e9e9e9; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/46877/18/22094/120005/637586ceEbac4ef66/a43a874e4814a2fc.jpg); height:1130px}\n.ssd-module-wrap .M16686466235105{width:750px; background-color:#e9e9e9; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/196287/9/29462/109229/637586cfEf15cfef1/e1a9f2450634d44b.jpg); height:1433px}\n.ssd-module-wrap .M16686466235317{width:750px; background-color:#b3b3b3; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/91121/39/27944/18383/637586cdE720c3b51/0facc75544c2ca09.jpg); height:519px}\n.ssd-module-wrap .M16686466235538{width:750px; background-color:#f2f2f2; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/72740/20/23599/51738/637586ceEeb81daee/72f55a533745dfd1.jpg); height:1104px}\n.ssd-module-wrap .M16686466235759{width:750px; background-color:#e9e9e9; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/194303/18/30150/58944/637586ceEfdc58069/93b1afe109b34514.jpg); height:1058px}\n.ssd-module-wrap .M166864662359710{width:750px; background-color:#b3b3b3; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/163166/31/32394/147267/637586d0E9ad56729/d2219b6c9ac60dc0.jpg); height:1813px}\n.ssd-module-wrap .M166864662362411{width:750px; background-color:#cbcbcb; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/182973/16/30562/83318/637586cfE52c9d571/28be0ab4dcf1ded5.jpg); height:1004px}\n.ssd-module-wrap .M166864662364712{width:750px; background-color:#b3b3b3; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/124620/22/32768/82883/637586cfE165d0193/34f371935b4c0055.jpg); height:1312px}\n.ssd-module-wrap .M166864662368313{width:750px; background-color:#d7d7d7; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/133672/23/26387/36567/637586ceE417d62b8/a3665a5a8a855fa4.jpg); height:995px}\n.ssd-module-wrap .M166864662371114{width:750px; background-color:#cbcbcb; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/216065/11/23546/83606/637586cfEf777311b/73a1532defaa004e.jpg); height:1247px}\n.ssd-module-wrap .M166864662375215{width:750px; background-color:#e9e9e9; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/165386/39/29271/106843/637586cfE90f76efc/4b91f84c9bc3faf7.jpg); height:1451px}\n.ssd-module-wrap .M166864662378416{width:750px; background-color:#f2f2f2; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/140483/8/31308/37886/637586ceE42f8b759/2f29088b8d7d614d.jpg); height:916px}\n.ssd-module-wrap .M166864662381117{width:750px; background-color:#e9e9e9; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/141928/35/31353/48895/637586ceEf3aada14/bc6a14fb4ae1913b.jpg); height:1279px}\n.ssd-module-wrap .M166864662383818{width:750px; background-color:#cbcbcb; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/111434/26/33483/50403/637586cdE18ed4191/787ae3c13f4b59c9.jpg); height:617px}\n.ssd-module-wrap .M166864662388819{width:750px; background-color:#b3b3b3; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/55690/37/21751/190963/637586d0E8aa3383c/8ff3448e1d7af11f.jpg); height:2070px}\n.ssd-module-wrap .M166864662395320{width:750px; background-color:#e9e9e9; background-size:100% 100%; background-image:url(//img30.360buyimg.com/sku/jfs/t1/80828/10/22795/235343/637586d0Edb144357/b00a8800a32b5787.jpg); height:1951px}\n</style><div cssurl='//sku-market-gw.jd.com/css/pc/100026707574.css?t=1678872031473'></div><div id='zbViewModulesH'  value='22740'></div><input id='zbViewModulesHeight' type='hidden' value='22740'/><div skudesign=\"100010\"></div><div class=\"ssd-module-wrap\" >\n            <div class=\"ssd-module M16686466234081 animate-M16686466234081\" data-id=\"M16686466234081\">\n        \n</div>\n<div class=\"ssd-module M16686466234452 animate-M16686466234452\" data-id=\"M16686466234452\">\n        \n</div>\n<div class=\"ssd-module M16686466234713 animate-M16686466234713\" data-id=\"M16686466234713\">\n        \n</div>\n<div class=\"ssd-module M16686466234914 animate-M16686466234914\" data-id=\"M16686466234914\">\n        \n</div>\n<div class=\"ssd-module M16686466235105 animate-M16686466235105\" data-id=\"M16686466235105\">\n        \n</div>\n<div class=\"ssd-module M16686466235317 animate-M16686466235317\" data-id=\"M16686466235317\">\n        \n</div>\n<div class=\"ssd-module M16686466235538 animate-M16686466235538\" data-id=\"M16686466235538\">\n        \n</div>\n<div class=\"ssd-module M16686466235759 animate-M16686466235759\" data-id=\"M16686466235759\">\n        \n</div>\n<div class=\"ssd-module M166864662359710 animate-M166864662359710\" data-id=\"M166864662359710\">\n        \n</div>\n<div class=\"ssd-module M166864662362411 animate-M166864662362411\" data-id=\"M166864662362411\">\n        \n</div>\n<div class=\"ssd-module M166864662364712 animate-M166864662364712\" data-id=\"M166864662364712\">\n        \n</div>\n<div class=\"ssd-module M166864662368313 animate-M166864662368313\" data-id=\"M166864662368313\">\n        \n</div>\n<div class=\"ssd-module M166864662371114 animate-M166864662371114\" data-id=\"M166864662371114\">\n        \n</div>\n<div class=\"ssd-module M166864662375215 animate-M166864662375215\" data-id=\"M166864662375215\">\n        \n</div>\n<div class=\"ssd-module M166864662378416 animate-M166864662378416\" data-id=\"M166864662378416\">\n        \n</div>\n<div class=\"ssd-module M166864662381117 animate-M166864662381117\" data-id=\"M166864662381117\">\n        \n</div>\n<div class=\"ssd-module M166864662383818 animate-M166864662383818\" data-id=\"M166864662383818\">\n        \n</div>\n<div class=\"ssd-module M166864662388819 animate-M166864662388819\" data-id=\"M166864662388819\">\n        \n</div>\n<div class=\"ssd-module M166864662395320 animate-M166864662395320\" data-id=\"M166864662395320\">\n        \n</div>\n\n    </div>\n<!-- 2022-11-17 08:57:13 --> ",
                checkSKU: { name: "可机洗国标A类 60*40cm" },
                carriageFree: "包邮",
                price: { symbol: "￥", integer: "98", decimal: "99", market: "168.59" },
                imgs: [{ "url": "https://img13.360buyimg.com/n1/s800x800_jfs/t1/169352/8/35771/77819/64118d57F5f51fff8/450dcb4957f5ae02.jpg", "type": 1, "sort": 1, "isDefault": 0 }, { "url": "https://img13.360buyimg.com/n1/s800x800_jfs/t1/216946/40/18841/114458/627dc6cfE6e661b9a/4826dc61a73ded61.jpg", "type": 1, "sort": 1, "isDefault": 1 }, { "url": "https://img13.360buyimg.com/n1/s800x800_jfs/t1/64010/19/25109/421737/62c7f639E1bad9094/1af6d6f1f8aea7ac.jpg", "type": 1, "sort": 1, "isDefault": 1 }, { "url": "https://img13.360buyimg.com/n1/s800x800_jfs/t1/36765/31/16438/263010/62c7f631Ed4e22143/0610776ba2b05900.jpg", "type": 1, "sort": 1, "isDefault": 1 }, { "url": "https://img13.360buyimg.com/n1/s800x800_jfs/t1/151046/35/25655/116081/62c7f713E4413c5d1/a65eda7f003a32ff.jpg", "type": 1, "sort": 1, "isDefault": 1 }, { "url": "https://img13.360buyimg.com/n1/s800x800_jfs/t1/147218/23/28676/206558/62c7f62dEaba6e20a/ec4d7645650f01fb.jpg", "type": 1, "sort": 1, "isDefault": 1 }, { "url": "https://img13.360buyimg.com/n1/s800x800_jfs/t1/8962/35/18790/353464/62c7f638Ee3ac2f75/d2e1809b9c852120.jpg", "type": 1, "sort": 1, "isDefault": 1 }, { "url": "https://img13.360buyimg.com/n1/s800x800_jfs/t1/148509/19/28858/405758/62c7f638E5c8ebcd3/2be7d18c75e223b9.jpg", "type": 1, "sort": 1, "isDefault": 1 }]
            },
            user: {
                addr: "成都市天府新区协和下街588号"
            }

        }
    },
    methods: {
        onBuyNow(){
            this.$router.push("/shopping")
        }
    },
}
</script>
<style scoped>
.my-swipe img {
    width: 100%;
}

.price-info {
    margin: 8px;
    background: white;
    border-radius: 8px;
    padding: 8px;
}

.price-info .label {
    margin-top: 4px;
    font-weight: 550;
    font-size: 1.1rem;
}

.price-info .price {
    text-align: left;
    font-size: 0.86rem;
}

.price .integer {
    color: red;
    font-weight: 500;
    font-size: 2rem;
}

.price .decimal,
.price .symbol {
    color: red;
}

.price .market {
    margin-left: 10px;
    color: var(--van-gray-6);
    text-decoration: line-through;
}

.price .symbol,
.decimal,
.market {
    font-size: 1rem;
}

.price .market {
    margin-left: 16px;
}

.custom-indicator {
    position: absolute;
    right: 6px;
    bottom: 10px;
    padding: 2px 10px;
    font-size: 1.1rem;
    background: rgba(0, 0, 0, 0.1);
    border-radius: 10px;
}

.prod-info {
    margin: 8px;
}

.prod-detail {
    margin: 8px;
    width: calc(100% - 16px);
    border-radius: 8px;
}
</style>