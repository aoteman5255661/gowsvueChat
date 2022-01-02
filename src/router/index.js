import Vue from "vue";
import VueRouter from 'vue-router'
import Chat from "../components/Chat"
// import HelloWorld from "../components/HelloWorld";

Vue.use(VueRouter)

export default new VueRouter({
    routes: [
        {path: '/', redirect: '/chat'},
        {path: '/chat', component: Chat},
    ]
})