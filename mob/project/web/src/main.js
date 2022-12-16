import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/router'

import "font-awesome/css/font-awesome.min.css";
import 'bootstrap/dist/css/bootstrap.min.css'
import theia from "theia-mob"
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

var app = createApp(App)
app.use(theia, { env: process.env })

app.use(router)
app.mount('#app')

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
