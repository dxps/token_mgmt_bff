import { createApp } from 'vue'
import VueSSE from "vue-sse"
import router from "@/router"
import App from './App.vue'

import './assets/styles/tailwind.css';

const app = createApp(App)

app.use(VueSSE)
app.use(router)

app.mount('#app')
