import { process } from "./process.js"
import { page } from "./page.js"


function build(penv,router,theia) {
    let main = {}
    main.process=new process()
    main.page=new page(router,theia)
    return main
}

export default {
    load: function (penv,router,theia) {
        return build(penv,router,theia)
    },
}