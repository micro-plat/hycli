<template>
    <van-col :span="span">
        <div class="hy-cmdprod">
            <van-row>
                <van-col span="20" class="title">{{ commend.title }}
                    <van-tag v-if="commend.tag" type="danger">
                        {{ commend.tag }}
                    </van-tag> </van-col>
                <van-col span="4" class="more" v-if="commend.more">
                    <van-icon name="arrow"><a href="commend.more"></a></van-icon>
                </van-col>
                <van-col :span="24/num" v-for="prod in commend.prods" :key="prod">
                    <van-row>
                        <van-col span="24" class="img">
                            <img :src="prod.img" /></van-col>
                        <van-col span="24" class="price" v-if="prod.price">{{ prod.price }}</van-col>
                        <van-col span="24" class="label" v-if="prod.label">{{ prod.label }}</van-col>
                    </van-row>
                </van-col>
            </van-row>
        </div>
    </van-col>
</template>
<script>
export default {
    props: {
        span: {
            type: String,
            default: 12,
        },
        num: {
            type: String,
            default: 2,
        }
    },
    data() {
        return {
            commend: {
                title: "精品推荐",
                tag: "惊喜不断",
                more: "/product",
                prods: [
                    { img: "http://img13.360buyimg.com/n1/s800x800_jfs/t1/196489/29/29167/80251/63295a12Edbcac09a/1b18907d13c59abd.jpg", price: "￥98.99", label: "" },
                    { img: "http://img13.360buyimg.com/n1/s800x800_jfs/t1/123594/18/21645/283330/61f29f7dEa15539d1/dafd831cc28e7c37.png", label: "速溶咖啡" },
                ],
            }

        }
    },
}
</script>
<style scoped>
.hy-cmdprod {
    margin: 8px;
    border-radius: 8px;
    background-color: white;

}

.hy-cmdprod .title,
.more {
    line-height: 2rem;
    padding: 4px 8px;
    ;

}

.hy-cmdprod .more {
    text-align: right;
}

.hy-cmdprod .title {
    color: var(--van-grid-item-text-color);
    font-size: var(--van-grid-item-title-font-size);
}

.hy-cmdprod .title .van-tag {
    vertical-align: top;
}

.hy-cmdprod img {
    width: 100%;
}

.hy-cmdprod .label,
.price {
    font-size: 0.8rem;
    text-align: center;
    margin: 4px;

}

.hy-cmdprod .price {
    color: red;
}
</style>