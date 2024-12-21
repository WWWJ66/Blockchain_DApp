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
    <div class="no-result-info" v-if="JSON.stringify(details) === '{}'">暂无查询结果</div>
    <div class="result-info" v-else>
      <div class="info" v-for="(detail, infoType) in details" :key="infoType">
        <h3>{{ infoType }}</h3>
        <div v-for="(value, key) in detail" :key="key">
          {{ key }}: {{ value }}
        </div>
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
const details = ref({});

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
  details.value = {
    "basicInfo": {
      "name": "黄鱼",
      "fisherName": "渔民A",
      "fishedTime": "2024.1.1",
      "fishedArea": "黄海",
      "tradeId": "",
      "tradeTime": ''
    },
    "factoryInfo": {
      "name": "黄鱼酥",
      "factoryName": "加工厂A",
      "processTime": "2024.1.2",
      "factoryAddress": "A市B区",
      "tradeId": "",
      "tradeTime": ''
    },
    "transportInfo": {
      "name": "黄鱼酥",
      "driverName": "司机A",
      "transportTime": "2024.1.3",
      "destination": "B市A区",
      "tradeId": "",
      "tradeTime": ''
    },
    "shopInfo": {
      "name": "黄鱼酥",
      "shopName": "商店A",
      "storageTime": "2024.1.4",
      "sellTime": "2024.1.4",
      "tradeId": "",
      "tradeTime": ''
    }

  };
  console.log(details.value);
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

  input {
    padding: 10px;
    margin-right: 10px;
    width: 300px;
    border: 1px solid #ccc;
  }

  button {
    padding: 10px 20px;
    background-color: #5eaaff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
}

.search-bar button:hover {
  background-color: #1d89fd;
}

.results-table {
  border-radius: 4px;
  width: 400px;
  margin: 10px auto;
  border-collapse: collapse;
  text-align: center;

  th {
    background-color: #d2ffff;
    padding: 10px;
    text-align: center;
  }

  td {
    padding: 7px;
    background-color: rgba(204, 226, 255, 0.7);
    border-bottom: 1px solid #ddd;
    justify-content: center;

    button {
      padding: 8px 15px;
      background-color: #72b4ff;
      color: white;
      border: none;
      border-radius: 4px;
      cursor: pointer;
    }
  }
}

.results-table td button:hover {
  background-color: #1e77d8;
}

.no-result-info {
  margin-top: 40px;
  justify-content: center;
  text-align: center;
  display: flex;
}

.info {
  justify-content: left;
  text-align: left;
  h3 {
    margin-left:20%;
  }
  div {
    margin-left: 30%;
  }
}

</style>
