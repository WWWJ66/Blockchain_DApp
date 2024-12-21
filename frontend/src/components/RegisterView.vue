<template>
    <div>
        <h1>水产品养殖溯源系统</h1>
        <div class="container" style="width: 33.33%; margin: 0 auto; text-align: center;">
            <form @submit.prevent="handleSubmit" id="registration-form" style="text-align: left; max-width: 300px; margin: 0 auto;">
                <label for="phone">手机号:</label>
                <input type="text" id="phone" v-model="phone" required style="width: 100%;">

                <label for="password">设置密码:</label>
                <input type="password" id="password" v-model="password" required style="width: 100%;">

                <label for="confirm-password">确认密码:</label>
                <input type="password" id="confirm-password" v-model="confirmPassword" required style="width: 100%;">

                <p v-if="errorMessage" style="color: red;">{{ errorMessage }}</p>

                <label for="role">选择角色:</label>
                <select id="role" v-model="role" required style="width: 100%;">
                    <option value="水产品养殖户">水产品养殖户</option>
                    <option value="工厂">工厂</option>
                    <option value="运输司机">运输司机</option>
                    <option value="商店">商店</option>
                    <option value="消费者">消费者</option>
                </select>

                <button type="submit" style="width: 100%; max-width: 300px;">注册</button>
            </form>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            phone: '',
            password: '',
            confirmPassword: '',
            role: '',
            errorMessage: ''
        };
    },
    methods: {
        handleSubmit() {
            if (this.password !== this.confirmPassword) {
                this.errorMessage = '密码不匹配，请重新输入。';
                return;
            }

            // const payload = {
            //     phone: this.phone,
            //     password: this.password,
            //     role: this.role
            // };

            alert('注册成功！即将跳转到登陆界面...');
            setTimeout(() => {
                this.$router.push({name: 'Login'});
            }, 1500);
        }
    }
};
</script>

<style scoped>
body {
    background-image: url('../assets/registration4.png');
    background-repeat: no-repeat;
    background-size: 100%;
    background-attachment: fixed;
    font-family: Arial, sans-serif;
    color: #000;
}
.container {
    background-color: rgba(255, 255, 255, 0.8);
    width: 100%;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
h1 {
    text-align: center;
    color: rgb(46, 53, 53);
}
label {
    display: block;
    margin-top: 10px;
}
input[type="text"], input[type="password"], select {
    width: 100%;
    padding: 8px;
    margin-top: 5px;
    border: 1px solid #ccc;
    border-radius: 5px;
}
button {
    width: 100%;
    padding: 10px;
    background-color: rgb(141, 248, 252);
    border: none;
    border-radius: 5px;
    color: #fff;
    font-size: 16px;
    cursor: pointer;
    margin-top: 20px;
}
button:hover {
    background-color: rgb(100, 200, 200);
}
</style>