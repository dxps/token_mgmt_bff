import { createRouter, createWebHistory } from "vue-router"
import WebsocketsTest from "@/pages/WebsocketsTest.vue"
import HomePage from "@/pages/HomePage.vue"

const routes = [
  {
    path: "/",
    name: "HomePage",
    component: HomePage,
  },
  {
    path: "/wstest",
    name: "WebSockets Test",
    component: WebsocketsTest,
  },
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
