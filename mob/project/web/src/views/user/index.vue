<template>
    <div class="user">
        <userheader></userheader>
        <div class="content">
            <orderbar></orderbar>
            <userking></userking>
            <userlstbar></userlstbar>
        </div>
       
        <!-- <cmpnt v-for="c in cmpnts" :type="c" :key="c"></cmpnt> -->

    </div>
</template>
<script>
import userheader from "@/views/cmpnt/hy-user-head"
import orderbar from "@/views/cmpnt/hy-user-order"
import userking from "@/views/cmpnt/hy-user-king"
import userlstbar from "@/views/cmpnt/hy-user-lstbar"
export default {
    components: {
        userheader,
        orderbar,
        userking,
        userlstbar,
    },
    data() {
        return {
            cmpnts: ["userhead", "orderbar", "userking", "userlstbar"]
        }
    },

}
</script>
<style  scoped>
.user{
    position: relative;
}
.content{
    position: absolute;
    z-index: 98;
    top:120px;
    left:10px;
    width:95%;
}
/deep/ .van-cell-group--inset{
    margin:0;
}
</style>
