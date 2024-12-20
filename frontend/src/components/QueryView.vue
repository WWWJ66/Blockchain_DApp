<template>
  <div class="container">
    <h1>溯源信息查询</h1>
    <div class="search-bar">
      <input v-model="searchQuery" type="text" placeholder="请输入溯源码进行查询">
      <button @click="performSearch">搜索</button>
    </div>
    <div class="results">
      <p v-if="!searchResults.length">暂无查询结果</p>
      <ul v-else>
        <li v-for="result in searchResults" :key="result.id">
          <h3>{{ result.name }}</h3>
          <p>{{ result.info }}</p>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

// 定义搜索查询和搜索结果的响应式变量
const searchQuery = ref('');
const searchResults = ref([]);

// 模拟搜索功能
const performSearch = () => {
  // 这里应该调用后端API获取实际的搜索结果
  // 为了示例，我们手动设置一些搜索结果
  searchResults.value = [
    { id: 1, name: '产品A', info: '详细信息A' },
    { id: 2, name: '产品B', info: '详细信息B' },
    { id: 3, name: '产品C', info: '详细信息C' }
  ];
  // 实际开发中，你可能需要使用fetch或axios等库来发送请求
};
</script>

<style scoped>
.container {
  width: 100%;
  padding: 20px;
  box-sizing: border-box;
}

h1 {
  text-align: center;
}

.search-bar {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.search-bar input {
  padding: 10px;
  margin-right: 10px;
  width: 300px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.search-bar button {
  padding: 10px 20px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.search-bar button:hover {
  background-color: #0056b3;
}

.results {
  text-align: center;
}

.results ul {
  list-style-type: none;
  padding: 0;
}

.results li {
  margin: 5px 0;
}
</style>
