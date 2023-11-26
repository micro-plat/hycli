import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/router'

import jspkg from "./jspkg/index.js"
import  "font-awesome/css/font-awesome.min.css";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import locale from 'element-plus/lib/locale/lang/zh-cn'
import 'bootstrap/dist/css/bootstrap.min.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import theia from "theia-ui"

var app=createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }


app.use(theia, { env: process.env });
if (document.documentElement.clientWidth <= 1280) {
  app.use(ElementPlus, { size: "small", locale });
} else if (document.documentElement.clientWidth <= 1920) {
  app.use(ElementPlus, { size: "medium", locale });
} else {
  app.use(ElementPlus, { size: "large", locale });
}
window.onresize = () => {
  if (document.documentElement.clientWidth <= 1280) {
    app.config.globalProperties.$ELEMENT = { size: 'small' }
  } else if (document.documentElement.clientWidth <= 1920) {
    app.config.globalProperties.$ELEMENT = { size: 'medium' }  
  } else {
    app.config.globalProperties.$ELEMENT = { size: 'large' }
  }
}
app.use(router);
app.use(jspkg,{ env: process.env });
app.mount("#app");