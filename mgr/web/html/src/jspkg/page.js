export function page(router, theia) {
  page.prototype.$router = router;
  page.prototype.$theia = theia;
}

page.prototype.colorful = function (r, name) {
  if (this.$theia.env.conf.colorful[name]) {
    return this.$theia.env.conf.colorful[name][r] || "";
  }
  return this.$theia.env.conf.colorful[r] || "";
};
page.prototype.tagColor = function (r, name) {
  if (this.$theia.env.conf.colorTag[name]) {
    return this.$theia.env.conf.colorTag[name][r] || "";
  }
  return this.$theia.env.conf.colorTag[r] || "";
};
page.prototype.goto = function (url, param) {
  if (!url) {
    return;
  }
  if (!url.indexOf("://")) {
    this.$router.push({ path: url, query: param });
    return;
  }
  let p = this.$theia.url.encode(param);
  window.location = `${url}?${p}`;
};
