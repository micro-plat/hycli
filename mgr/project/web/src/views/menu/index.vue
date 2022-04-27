<template>
  <te-frame-plus :system="system" :menus="menus"></te-frame-plus>
</template>
<script>
export default {
  name: "App",
  data() {
    return {
      system: {},
      menus: [],
    };
  },
  mounted() {},
  created() {
    this.system = this.$theia.env.getSystemInfo();
    document.title = this.system.name;
    this.$theia.env.setTheme(this.system.theme);
    this.menus = this.$theia.env.getMenus();
  },

  methods: {},
};
</script>