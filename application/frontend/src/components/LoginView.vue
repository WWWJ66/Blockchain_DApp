<template>
  <div>
    <h1>水产品养殖溯源系统</h1>
    <div class="container" style="width: 33.33%; margin: 0 auto; text-align: center;">
      <form @submit.prevent="handleSubmit" style="text-align: left; max-width: 300px; margin: 0 auto;">
        <label for="username">用户名:</label>
        <input type="text" v-model="username" id="username" required style="width: 100%;">

        <label for="password">密码:</label>
        <input type="password" v-model="password" id="password" required style="width: 100%;">

        <button type="submit" style="width: 100%; max-width: 300px;">登录</button>
        <div style="text-align: center; margin-top: 20px;">
          <button @click="goToRegistration"
                  style="width: auto; max-width: none; background-color: transparent; border: none; color: rgb(43, 43, 57); text-decoration: underline; font-size: 70%; cursor: pointer; font-family: 'Microsoft YaHei', sans-serif; font-weight: bold;">
            没有账户？立即注册
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import {useUserStore} from '@/store/userStore';
import {ref} from "vue";
import { useRouter } from 'vue-router';
import axios from "axios";

const userStore = useUserStore();
const username = ref('');
const password = ref('');
const router = useRouter();

const handleSubmit = async () => {
  try {
    const formData = new FormData();
    formData.append("username", username.value.trim());
    formData.append("password", password.value);

    const res = await axios.post("http://192.168.133.131:9090/login", formData)
    if (res.data.code === 200) {
      alert('登录成功');
      userStore.setUsername(username.value);
      userStore.setToken(res.data.jwt);
      setUserType();
      router.push({ name: 'home' })
    } else {
      alert('用户名或密码错误，请重新输入');
    }
  } catch(error) {
    console.error("登录请求失败：", error);
    alert('登录失败，请稍后再试');
  }
};

const setUserType = async () => {
  try {
    const res = await axios.post(
      "http://192.168.133.131:9090/getInfo", 
      {},
      {
        headers: {
          'Authorization': userStore.token
        }
      }
    );
    if (res.data.code === 200) {
      userStore.setUserType(res.data.userType);
    } else {
      console.log(res.data);
    }
  } catch(error) {
    console.error("获取用户类型请求失败：", error);
  }
}

const goToRegistration = () => {
  router.push({name: 'Register'});
};
</script>

<style scoped>
body {
  background-image: url('../assets/login3.png');
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