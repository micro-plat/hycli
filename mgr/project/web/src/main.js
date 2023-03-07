import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/router'

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
  
app.use(theia, { env: process.env })
app.use(ElementPlus, { locale })
app.use(router)
app.mount('#app')
