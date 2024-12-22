import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => ({
        username: '', // 全局 username
        userType:'',
        token:'',
    }),
    actions: {
        setUsername(name) {
            this.username = name; // 更新 username
        },
        setUserType(type) {
            this.userType = type;
        },
        setToken(token) {
            this.token = token;
        }
    },
});
