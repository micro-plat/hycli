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
</div>
</template>
<script>
export default {
    name:"hy-user-head",
    data() {
        return {
            active: 0,
            user:{},
        }
    }
}
</script>
<style>

.user-panel {
    height: 160px;
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