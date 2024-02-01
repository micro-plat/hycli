import register from "./register.js";

const sizeMap = { small: 0, medium: 1, large: 2, s: 0, m: 1, l: 2 };
const components = [register];

// 定义 install 方法，接收 Vue 作为参数。如果使用 use 注册插件，则所有的组件都将被注册
const install = function (app, p) {
  if (install.installed) return;
  // 遍历注册全局组件
  components.map((component) => app.component(component.name, component));

  //注册js组件
  app.config.globalProperties.$js = register.load(
    p.env || {},
    app.config.globalProperties.$router,
    app.config.globalProperties.$theia
  );

  //注册组件
  register.directive(app);
};

// 判断是否是直接引入文件
if (typeof window !== "undefined" && window.Vue) {
  install(window.Vue);
}

export default {
  // 导出的对象必须具有 install，才能被 Vue.use() 方法安装
  install,
  screenSize,
  // 以下是具体的组件列表
  load: register.load,
};

function screenSize(app, elenentPlus, locale) {
  let size = "small";
  let clientWidth = document.documentElement.clientWidth;
  if (clientWidth <= 1280) {
    size = "small";
  } else if (clientWidth <= 1920) {
    size = "medium";
  } else {
    size = "large";
  }
  app.config.globalProperties.$size = {
    name: size,
    idx: sizeMap[size],
    eq: function (name = "s") {
      return sizeMap[name.toLowerCase()] == sizeMap[size];
    },
    gt: function (name = "s") {
      return sizeMap[size] > sizeMap[name.toLowerCase()];
    },
    lt: function (name = "s") {
      return sizeMap[size] < sizeMap[name.toLowerCase()];
    },
  };
  if (elenentPlus) {
    app.use(elenentPlus, { size: size, locale });
    return;
  }
  app.config.globalProperties.$ELEMENT = { size: size };
}
