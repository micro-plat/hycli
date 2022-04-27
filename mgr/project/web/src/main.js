import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/router'

import  "font-awesome/css/font-awesome.min.css";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import locale from 'element-plus/lib/locale/lang/zh-cn'

import theia from "theia-ui"

var app=createApp(App)
app.use(theia)
app.use(ElementPlus, { locale })
app.use(router)
app.mount('#app')
