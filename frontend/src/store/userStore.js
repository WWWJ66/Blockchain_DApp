import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => ({
        username: '', // 全局 username
        type:'',
    }),
    actions: {
        setUsername(name) {
            this.username = name; // 更新 username
        },
        setType(type) {
            this.type = type;
        }
    },
});
