<template>
    <div>
        <searchbar></searchbar>
        <banner></banner>
        <king></king>
        <van-row>
            <cmdprod :span="span"></cmdprod>
            <cmdimg :span="span"></cmdimg>
        </van-row>
        <prodlst></prodlst>

    </div>
</template>
<script>
import searchbar from "@/views/cmpnt/hy-searchbar"
import cmdprod from "@/views/cmpnt/hy-cmd-prod"
import banner from "@/views/cmpnt/hy-banner"
import king from "@/views/cmpnt/hy-king"
import cmdimg from "@/views/cmpnt/hy-cmd-img"
import prodlst from "@/views/cmpnt/hy-prodlst"
export default {
    components: {
        searchbar,
        banner,
        king,
        cmdimg,
        cmdprod,
        prodlst,
    },
    data() {
        return {
            span:12,

        }
    },

}
</script>
