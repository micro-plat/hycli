<template>
  <div>
    <router-view/>
    <van-tabbar v-model="active" :active-color="conf.active_color">
      <van-tabbar-item to="/home">
        <span>首页</span>
        <template #icon="props">
          <i></i>
          <van-icon v-if="!props.active" name="wap-home-o" />
          <van-icon v-if="props.active" name="wap-home" />
        </template>
      </van-tabbar-item>
      <van-tabbar-item to="/cart">
        <span>购物车</span>
        <template #icon="props">
          <van-icon v-if="!props.active" name="shopping-cart-o" />
          <van-icon v-if="props.active" name="shopping-cart" />
        </template>
      </van-tabbar-item>
      <van-tabbar-item to="/user">
        <span>个人</span>
        <template #icon="props">
          <van-icon v-if="!props.active" name="manager-o" />
          <van-icon v-if="props.active" name="manager" />
        </template>
      </van-tabbar-item>
    </van-tabbar>
  </div>

</template>
<script>
import { Tabbar, TabbarItem, Icon } from 'vant';
export default {
  components: {
    [Tabbar.name]: Tabbar,
    [TabbarItem.name]: TabbarItem,
    [Icon.name]: Icon
  },
  data() {
    return {
      conf: {
        active_color: "#333"
      },
      active: 2
    }
  },
  mounted() {
    let system = this.$theia.env.getSystemInfo();
    let theme = this.$theia.env.getTheme(system.theme);
    this.conf.active_color = theme.active_bg_color;
  },
}
</script>

<style>
</style>