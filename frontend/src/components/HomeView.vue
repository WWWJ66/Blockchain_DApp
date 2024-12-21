<template>
  <div class="home-view">
    <el-container>
      <el-aside width="200px" style="background-color: rgba(141,248,252,0.3); height: 100vh" type="primary">
        <el-menu default-active="2" class="el-menu-vertical" style="background-color: rgba(0,0,0,0);">
          <el-menu-item index="1" class="menu"
                        style="color: #270055; font-size: 15px;"
                        @click="$router.push('/queries')">溯源查询
          </el-menu-item>
          <el-menu-item index="2" class="menu"
                        style="color: #270055; font-size: 15px;"
                        @click="$router.push('/input')">溯源信息录入
          </el-menu-item>
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
.home-view, el-container, el-aside {
  height: 100vh;
}

menu {
  text-align: center;
  color: white;
}

.content {
  background-color: rgba(141, 248, 252, 0.1);
}
</style>
