const routes = [
    {
      path: "/",
      component: () => import("../views/menu/index.vue"),
      children: [
        {-{- range $i,$c:=. }-}
        {
          path: '{-{$c.Name.MainPath}-}{-{flterJoinColumnNames $c "rp" "/:"}-}',
          component: () => import("../views/{-{.Name.Prefix}-}/{-{$c.Name.Main}-}/{-{$c.Name}-}.list.vue"),
        },
       {-{- end}-}
      ]
    },
  ]
export default routes