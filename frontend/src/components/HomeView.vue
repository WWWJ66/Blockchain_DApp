<template>
  <div class="home-view">
    <el-container>
      <el-aside type="primary">
        <el-menu default-active="2" class="el-menu-vertical">
          <el-menu-item index="1" class="menu" @click="$router.push('/queries')">溯源查询</el-menu-item>
          <el-menu-item index="2" class="menu" @click="$router.push('/input')">溯源信息录入</el-menu-item>
        </el-menu>
      </el-aside>
      <el-container class="content">
        <router-view/>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import {RouterView} from 'vue-router';
import {ElContainer, ElAside} from 'element-plus';
</script>

<style scoped>
.home-view, el-container, el-aside, el-menu {
  height: 100vh;
}

el-aside{
  width: 150px;
}
</style>
