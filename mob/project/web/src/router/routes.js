
const routes = [{
    path: "/",
    component: () => import("../views/menu/index.vue"),
    children: [{
        path: "user",
        component: () => import("../views/user/index.vue"),
    }]
}]
export default routes