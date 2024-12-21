import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => ({
        username: '', // 全局 username
        identity:'',
        loggedIn: 0, // 初始状态为 0
        messages:[
            {"role": "system", "content": "请告诉我你遇到的问题，我会尽力为你提供帮助。"},
          ],
        audiourllist:[
        {"src":"null"},
        ],
    }),
    actions: {
        setUsername(name) {
            this.username = name; // 更新 username
        },
        logIn() {
            this.loggedIn = 1; // 设置为登录状态
        },
        logOut() {
            this.loggedIn = 0; // 设置为未登录状态
        },
    },
});
