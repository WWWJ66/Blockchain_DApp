<template>
  <div class="container">
    <div class="sidebar">
      <ul>
        <li><router-link to="/queries">溯源查询</router-link></li>
        <li><router-link to="/input">溯源信息录入</router-link></li>
      </ul>
    </div>
    <div class="content">
      <!-- 主要内容区域 -->
      <router-view />
    </div>
  </div>
</template>

<script setup>
import { RouterLink, RouterView } from 'vue-router';
</script>

<script>
export default {
  name: 'HomeView'
}
</script>

<style scoped>
.container {
  display: flex;
  height: 100vh;
}

.sidebar {
  width: 200px;
  background-color: #f4f4f4;
  padding: 20px;
}

.sidebar ul {
  list-style-type: none;
  padding: 0;
}

.sidebar ul li {
  margin-bottom: 10px;
}

.sidebar ul li a {
  text-decoration: none;
  color: #333;
}

.sidebar ul li a:hover {
  color: #007bff;
}

.content {
  flex: 1;
  padding: 20px;
}
</style>