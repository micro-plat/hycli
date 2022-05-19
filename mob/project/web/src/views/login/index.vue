<template>
  <tm-login
    :system="$theia.system"
    :error="error"
    mode="msg"
    ref="login"
    @login="login"
  ></tm-login>
</template>

<script>
export default {
  data() {
    return {
      error: "",
     errs: {
        901: "用户名密码错误",
        902: "用户已锁定，暂时不能登录",
        903: "用户已禁用,暂时无法登录",
        904: "没有权限或角色已锁定,暂时无法登录",
        905: "系统已锁定,暂时无法登录",
        999: "系统繁忙,暂时无法登录",
      },
    };
  },
  created(){
    this.$theia.system =  this.$theia.env.getSystemInfo();
  },

  mounted() {
    this.$theia.http.post("/member/logout", {})
  },
  methods: {
    login(u, p) {
      let that = this;
      this.$theia.http
        .post("/member/login", { uname: u, pwd: this.$theia.crypto.md5(p) })
        .then(() => {
          that.$router.push("/#/file");
        })
        .catch((err) => {
          let msg = "";
          if (
            err.response &&
            err.response.data &&
            err.response.data.remain_count
          ) {
            msg = ",剩余" + err.response.data.remain_count + "次数";
          }
          that.$refs.login.error(
            (that.errs[err.response.status] || "用户名密码错误") + msg
          );
        });
    },
  },
};
</script>