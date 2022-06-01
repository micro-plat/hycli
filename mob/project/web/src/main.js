import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/router'

import "font-awesome/css/font-awesome.min.css";
import 'bootstrap/dist/css/bootstrap.min.css'
import theia from "theia-mob"


var app = createApp(App)
app.use(theia, { env: process.env })

app.use(router)
app.mount('#app')
