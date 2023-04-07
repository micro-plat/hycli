const routes = [
    {
      path: "/",
      component: () => import("../views/menu/index.vue"),
      children: [
        {
          path: "",
          component: () => import("../views/home/index.vue"),
        },
        {
          path: "user",
          component: () => import("../views/user/index.vue"),
        },
        {
          path: "home",
          component: () => import("../views/home/index.vue"),
        },
        {
          path: "shopping",
          component: () => import("../views/shopping/cat.vue"),
        },
        {
          path: "catgory",
          component: () => import("../views/catgory/index.vue"),
        },
      ],
    },
    {
      path: "/product/item",
      component: () => import("../views/prod/item.vue"),
    },
    {
      path: "/order/cnfrm",
      component: () => import("../views/order/cnfrm.vue"),
    },
    {
      path: "/order/success",
      component: () => import("../views/order/success.vue"),
    }, {
      path: "/order/list",
      component: () => import("../views/order/list.vue"),
    },{
      path: "/order/detail",
      component: () => import("../views/order/detail.vue"),
    },
  ];
  export default routes;
  