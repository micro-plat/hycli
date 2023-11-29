import { process } from "./process.js";
import { page } from "./page.js";
import { directive } from "./directive.js";

function build(penv, router, theia) {
  let pkg = {};
  pkg.process = new process();
  pkg.page = new page(router, theia);
  return pkg;
}

export default {
  load: function (penv, router, theia) {
    return build(penv, router, theia);
  },
  directive: function (app) {
    let dir = new directive();
    Object.keys(dir.all).forEach((key) => {
      app.directive(key, dir[key]);
    });
  },
};
