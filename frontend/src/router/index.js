import HomeView from "../components/HomeView.vue";
import InputView from "../components/InputView.vue";
import QueryView from '../components/QueryView.vue'
import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            name: 'home',
            path: '/',
            component: HomeView,
            children: [
                {
                    name: 'queries',
                    meta: {
                        title: '溯源查询',
                    },
                    path: 'queries',
                    component: QueryView,
                },
                {
                    name: 'input',
                    meta: {
                        title: '溯源信息录入',
                    },
                    path: 'input',
                    component: InputView,
                },
            ],
        },
        // {
        //     name: 'login',
        //     path: '/login',
        //     component: LoginView,
        // },
    ],
})

export default router