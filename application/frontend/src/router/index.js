import HomeView from "../components/HomeView.vue";
import InputView from "../components/InputView.vue";
import QueryView from '../components/QueryView.vue';
import Register from '../components/RegisterView.vue';
import Login from '../components/LoginView.vue';
import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: '/',
            name: 'Login',
            component: Login
        },
        {
            path: '/register',
            name: 'Register',
            component: Register
        },
        {
            name: 'home',
            path: '/home',
            component: HomeView,
            children: [
                {
                    name: 'queries',
                    meta: {
                        title: '溯源查询',
                    },
                    path: '/queries',
                    component: QueryView,
                },
                {
                    name: 'input',
                    meta: {
                        title: '溯源信息录入',
                    },
                    path: '/input',
                    component: InputView,
                },
                {
                    name: 'blockchain',
                    meta: {
                        title: '区块链网络',
                    },
                    path: 'http://192.168.133.131:8080',
                    component: InputView,
                }
            ],
        },
    ],
})

export default router