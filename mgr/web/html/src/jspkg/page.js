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

page.prototype.latestDays=function(){
  return [
      {
        text: '今天',
        value: [new Date(),new Date()],
      },
      {
        text: '昨天',
        value: () => {
          const date = new Date()
          date.setTime(date.getTime() - 3600 * 1000 * 24)
          return [date,new Date()]
        },
      },
      {
        text: '前天',
        value: () => {
          const date = new Date()
          date.setTime(date.getTime() - 3600 * 1000 * 48)
          return [date,new Date()]
        },
      },
      {
        text: '本周',
        value: () => {
          const now = new Date()
          return [new Date(now.getFullYear(), now.getMonth(), now.getDate() - now.getDay()),new Date()]
        },
      },
      {
        text: '本月',
        value: () => {
          const now = new Date()
          return [new Date(now.getFullYear(), now.getMonth(), 1),new Date()]
        },
      },
    ]
}

page.prototype.columnfilter=function(value,row,column){
    const property = column['property']
    return row[property] === value
}