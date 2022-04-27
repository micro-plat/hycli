import { createRouter, createWebHistory } from "vue-router";
import routes from "./routes"

//登录
routes.push({
  path: "/login",
  component: () => import("../views/login/index.vue"),
})

//默认
routes.push({
  path: "",
  component: () => import("../views/login/index.vue"),
})
const router = createRouter({
  history: createWebHistory(),
  routes
});
export default router;