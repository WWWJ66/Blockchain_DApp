<template>
  <div class="container">
    <h1>溯源信息查询</h1>
    <div class="search-bar">
      <input v-model="searchQuery" type="text" placeholder="请输入溯源码进行查询">
      <button @click="viewProduct(searchQuery)">搜索</button>
    </div>
    <table class="results-table">
      <thead>
      <tr>
        <th>溯源码</th>
        <th>操作</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="product in products" :key="product.id">
        <td>{{ product }}</td>
        <td>
          <button @click="addToMy(product)">添加至我的</button>
          <button @click="viewProduct(product)">查看</button>
        </td>
      </tr>
      </tbody>
    </table>
    <div class="result-info">
      <p v-if="!details.length">暂无查询结果</p>
      <div v-else>
        <table class="details-table" v-for="detail in details" :key="detail.id">
          <thead>
          <tr>
            <th>产品名称</th>
            <th>详细信息</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="info in detail" :key="info.id">
            <td>{{ info.name }}</td>
            <td>{{ info.info }}</td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue';
// import axios from "axios";

const username = "111111";

// 定义搜索查询和搜索结果的响应式变量
const searchQuery = ref('');
const products = ref([]);
const details = ref([]);

const getProductsList = async () => {
  console.log(`查询 ${username} 的产品列表`);
  // const productlistResponse = await axios.post("localhost:9090/getProductList", {username: username});
  // for (const item of productlistResponse.data) {
  //   products.value.push(item["tracecode"]);
  // }
  products.value = ["1", "2", "3"];
}

const addToMy = async (traceCode) => {
  // await axios.post("localhost:9090/addProduct", {username: username, tracecode: traceCode});
  console.log(`添加产品: ${traceCode}`);
};

const viewProduct = async (traceCode) => {
  // const res = await axios.post("localhost:9090/getProductInfo", {tracecode: traceCode});
  // searchResults.value = res.data;
  console.log(`展示产品详情: ${traceCode}`);
  details.value = [[
    {id: 1, name: '产品A', info: '详细信息A'},
    {id: 2, name: '产品B', info: '详细信息B'},
    {id: 3, name: '产品C', info: '详细信息C'}
  ],[
    {id: 4, name: '产品1', info: '详细信息A'},
    {id: 5, name: '产品2', info: '详细信息B'},
    {id: 6, name: '产品3', info: '详细信息C'}
  ]];
};

onMounted(async () => {
  await getProductsList();
});
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
  background-color: #5eaaff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.search-bar button:hover {
  background-color: #1d89fd;
}

.result-info {
  margin-top: 40px;
  justify-content: center;
  text-align: center;
  display: flex;
}

.results-table, .details-table {
  width: 400px;
  margin: 0 auto;
  border-collapse: collapse;
  text-align: center;
}

.results-table th, .details-table th {
  background-color: #d2ffff;
  padding: 10px;
  text-align: center;
}

.results-table td, .details-table td {
  padding: 7px;
  border-bottom: 1px solid #ddd;
  justify-content: center;
}

.results-table td button{
  padding: 8px 15px;
  background-color: #72b4ff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.results-table td button:hover, .details-table td button:hover {
  background-color: #1e77d8;
}
</style>
