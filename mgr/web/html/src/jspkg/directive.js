export function directive() {
  directive.prototype.all = {};
  directive.prototype.all.loadMore = this.loadMore;
  console.log("directive:");
}

directive.prototype.loadMore = {
  beforeMount(el, binding) {
    console.log("loadMore:", el);
    const selectWrap = el.querySelector(".el-table__body-wrapper");
    selectWrap.addEventListener("scroll", function () {
      console.log("scroll:", el);
        const scrollDistance =
            this.scrollHeight - this.scrollTop - this.clientHeight;
        if (scrollDistance <= 0.5) {
          binding.value()//执行在使用时绑定的函数，在这里即addData方法
        }
    });
  },
};
