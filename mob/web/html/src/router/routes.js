
const routes = [
    {
      path: "/",
      component: () => import("../views/menu/index.vue"),
      children: [
        {{range $i,$c:=. }}
        {
          path: "{{$c.Name.MainPath}}",
          component: () => import("../views/{{$c.Name.Main}}/{{$c.Name}}.list.vue"),
        },
       {{- end}}
      ]
    },
  ]
export default routes