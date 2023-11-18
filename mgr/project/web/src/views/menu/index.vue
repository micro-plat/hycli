<template>
  <te-frame :system="system" :menus="menus" :dropdowns="dropdowns" @menuItemChanged="menuItemChanged"></te-frame>
</template>
<script>
export default {
  name: "App",
  data() {
    return {
      system: {},
      menus: [],
      dropdowns: [],
    };
  },
  created() {
    this.getUserInfo();
    this.system = this.$theia.env.getSystemInfo();
    document.title = this.system.name;
    this.$theia.env.setTheme(this.system.theme);
    this.menus = this.$theia.env.getMenus();


  },

  methods: {
    menuItemChanged(mi) {
      if (!mi.name) {
        return
      }
      document.title = mi.name + " - " + this.system.name;
    },
    getUserInfo() {
      let data = this.$theia.user.get()
      this.dropdowns.push({ icon: "fa fa-user-o", name: data.name });
      this.dropdowns.push({
        icon: "fa fa-sign-out",
        name: "退出",
        url: "/login",
      });
    }
  },
};
</script>