<template>
    <div>
        <van-cell center class="user-panel">
            <template #title>
                <van-row type="flex" align="center">
                    <van-col span="2" align="center"></van-col>
                    <van-col span="3" align="center">
                        <van-image v-if="user.header" :src="user.header" round fit="cover" width="5rem" height="5rem">
                        </van-image>
                        <van-icon v-if="!user.header" name="user-circle-o" />
                    </van-col>
                    <van-col span="3"> </van-col>
                    <van-col span="12">
                        <div class="h6" v-if="user.name" v-text="user.name"></div>
                        <div class="h6" v-if="!user.name">匿名用户</div>
                        <div class="text-muted"></div>
                        <div>
                            <van-tag type="warning">普通会员</van-tag>
                        </div>
                    </van-col>
                </van-row>
            </template>
        </van-cell>
        <div>
            <van-cell-group inset clickable>
                <van-cell title="我的订单" value="全部订单" is-link size="large" to="/order/list" />
            </van-cell-group>
            <van-grid :border="false">
                <van-grid-item icon="balance-pay" inset text="待付款" to="/order/list/needpay" />
                <van-grid-item icon="idcard" text="待发货" to="/order/list/needelivery" />
                <van-grid-item icon="logistics" text="待收货" to="/order/list/needrecv" />
                <van-grid-item icon="comment-circle-o" text="待确认" to="/order/list/needconfirm" />
            </van-grid>
        </div>
        <div>
            <van-cell-group inset clickable>
                <van-cell title="常用功能" />
            </van-cell-group>
            <van-grid :border="false">
                <van-grid-item icon="paid" text="我的卡包" to="/card/list" />
                <van-grid-item icon="star" text="我的收藏" icon-color="red" to="/fav/lst" />
                <van-grid-item icon="balance-list-o" inset text="发票" to="/invoice/list" />
                <van-grid-item icon="phone-o" text="售后" to="/refund/list" />
            </van-grid>
        </div>
        <van-cell-group inset>
            <van-cell title="" size="large" />
        </van-cell-group>
        <van-cell title="我的地址" icon="location-o" is-link size="large" to="/user/address" />
        <van-cell title="常见问题" icon="question-o" is-link size="large" to="/service/question" />
        <van-cell title="联系客服" icon="phone-circle-o" is-link size="large" to="/service/contact" />
    </div>
</template>
<script>
import { Cell, Image as VanImage, Col, Row, Tag, Grid, GridItem, Icon } from 'vant';
export default {
    components: {
        [Cell.name]: Cell,
        [VanImage.name]: VanImage,
        [Row.name]: Row,
        [Col.name]: Col,
        [Tag.name]: Tag,
        [Grid.name]: Grid,
        [GridItem.name]: GridItem,
        [Icon.name]: Icon
    },
    data() {
        return {
            active: 0,
            user:{},
            // user: {header:"/head/zz.png"},
        }
    }
}
</script>

<style>
body {
    background-color: white;
}

.user-panel {
    height: 140px;
    margin: 0;
    padding: 0;
    color: white;
    background-image: linear-gradient(125deg, #0fd3b0, #04e382, #04e382, #0fd3b0);
    background-size: 500%;
}

.user-panel img {
    border-radius: 40px;
    border: 2px solid white;
}

.user-panel i {
    font-size: 3.9rem;
}

.order {
    height: 200px;
}
</style>