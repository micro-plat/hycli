import { createApp } from "vue";
import App from "@/App.vue";
import router from "@/router";
import "./assets/main.css";

import "font-awesome/css/font-awesome.min.css";
// import 'bootstrap/dist/css/bootstrap.min.css'
// import theia from "theia-mob"

import {
  Cell,
  CellGroup,
  Image as VanImage,
  Col,
  Row,
  Tag,
  Grid,
  GridItem,
  Icon,
  Tabbar,
  TabbarItem,
  Lazyload,
  Field,
  Swipe,
  SwipeItem,
  Search,
  Sticky,
  ActionBar,
  ActionBarIcon,
  ActionBarButton,
  Sidebar,
  SidebarItem,
  Card,
  Checkbox,
  CheckboxGroup,
  Stepper,
  SubmitBar,
  ContactCard,
  Tab,
  Tabs,
  Button,
  Divider,
} from "vant";

var app = createApp(App);
// app.use(theia, { env: process.env })

app.use(router);
app.mount("#app");
app.use(Cell);
app.use(CellGroup);
app.use(VanImage);
app.use(Col);
app.use(Row);
app.use(Tag);
app.use(Grid);
app.use(GridItem);
app.use(Icon);
app.use(Tabbar);
app.use(TabbarItem);
app.use(Lazyload, { lazyComponent: true });
app.use(Field);
app.use(Swipe);
app.use(SwipeItem);
app.use(Search);
app.use(Sticky);
app.use(ActionBar);
app.use(ActionBarIcon);
app.use(ActionBarButton);
app.use(Sidebar);
app.use(SidebarItem);
app.use(Card);
app.use(Checkbox);
app.use(CheckboxGroup);
app.use(Stepper);
app.use(SubmitBar);
app.use(ContactCard);
app.use(Tab);
app.use(Tabs);
app.use(Button);
app.use(Divider );