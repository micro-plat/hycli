const routes = [
    {
      path: "/",
      component: () => import("../views/menu/index.vue"),
      children: [
        {-{- range $i,$c:=. }-}
        {
          path: '{-{$c.Name.MainPath}-}{-{ $c.JoinNames "rp" "/:" "?"}-}',
          component: () => import("../views/{-{.Name.Prefix}-}/{-{$c.Name.Main}-}/{-{$c.Name}-}.list.vue"),
        },
       {-{- end}-}
      ]
    },
  ]
export default routes