<template>
  <div>
    <tabbar ref="tabbar"></tabbar>
  </div>

</template>
<script>
import tabbar from "@/views/cmpnt/hy-tabbar"
export default {
  components:{
    tabbar,
  },
  data() {
    return {
      conf: {
        active_color: "#333"
      },
      active: 0
    }
  },
  mounted() {
  },
}
</script>

<style>
</style>