<template>
  <div class="container">
    <h1>溯源信息查询</h1>
    <div style="color: #000; margin-bottom: 30px">
      当前用户：{{ username }}; 用户角色: {{ userType }}
    </div>
    <div class="search-bar">
      <input v-model="searchQuery" type="text" placeholder="请输入溯源码进行查询">
      <button @click="doSearch(searchQuery)">搜索</button>
    </div>
    <table class="results-table" v-if="products.length">
      <thead>
      <tr>
        <th>溯源码</th>
        <th>操作</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="product in products" :key="product.tracecode">
        <td>{{ product.tracecode }}</td>
        <td>
          <button @click="checkProduct(product.tracecode)">查看</button>
        </td>
      </tr>
      </tbody>
    </table>
    <div v-if="Object.keys(details).length">
      <div v-if="searchingCode !== ''">
        <h4>搜索{{ searchingCode }}的结果：</h4>
        <button @click="addToMy(searchingCode)">添加至我的</button>
      </div>
      <div class="result-info">
        <div class="info" v-for="(detail, infoType) in details" :key="infoType">
          <h3>{{ infoType }}</h3>
          <div v-for="(value, key) in detail" :key="key">
            {{ key }}: {{ value }}
          </div>
        </div>
      </div>
    </div>
    <div class="no-result-info" v-else>暂无查询结果</div>
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue';
import axios from "axios";
import {useUserStore} from "@/store/userStore";
const username = useUserStore().username; // 获取当前用户名
const userType = useUserStore().type;

// 定义搜索查询和搜索结果的响应式变量
const searchQuery = ref('');
const products = ref([]);
const details = ref({});
const searchingCode = ref('');

const doSearch = async (code) => {
  searchingCode.value = code;
  await getDetails(searchingCode.value);
}

const getProductsList = async () => {
  console.log(`查询 ${username} 的产品列表`);
  const productlistResponse = await axios.post("localhost:9090/getProductList", {username: username});
  products.value = productlistResponse.data;
  // products.value = [{"tracecode": "1",}, {"tracecode": "2",}, {"tracecode": "3",}];
}

const addToMy = async (traceCode) => {
  await axios.post("localhost:9090/addProduct", {username: username, tracecode: traceCode});
  console.log(`添加产品: ${traceCode}`);
  await getProductsList();
};

// 定义属性名与中文名的映射关系
const propertyNameToChinese = {
  "basicInfo": "基本信息",
  "factoryInfo": "工厂信息",
  "transportInfo": "运输信息",
  "shopInfo": "商店信息",
  "name": "产品名",
  "fisherName": "渔民名",
  "fishedTime": "捕捞时间",
  "fishedArea": "捕捞区域",
  "tradeId": "交易ID",
  "tradeTime": "交易时间",
  "factoryName": "加工厂名",
  "processTime": "加工时间",
  "factoryAddress": "加工厂地址",
  "driverName": "司机名",
  "transportTime": "运输时间",
  "destination": "目的地",
  "shopName": "商店名",
  "storageTime": "入库时间",
  "sellTime": "销售时间"
};

// 定义一个函数来转换属性名
function convertPropertyNamesToChinese(obj) {
  // 遍历对象的每个键值对
  for (let key in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, key)) {
      // 如果值是一个对象，则递归调用函数
      if (typeof obj[key] === 'object' && obj[key] !== null) {
        obj[propertyNameToChinese[key]] = obj[key];
        delete obj[key];
        convertPropertyNamesToChinese(obj[propertyNameToChinese[key]]);
      } else if (propertyNameToChinese[key] !== undefined) {
        // 如果键在映射关系中，则替换键名
        obj[propertyNameToChinese[key]] = obj[key];
        delete obj[key];
      }
    }
  }
}

const checkProduct = async (product) => {
  searchingCode.value = '';
  await getDetails(product);
}

const getDetails = async (traceCode) => {
  const res = await axios.post("localhost:9090/getProductInfo", {tracecode: traceCode});
  details.value = res.data;
  console.log(`展示产品详情: ${traceCode}`);
  // details.value = {
  //   "basicInfo": {
  //     "name": "黄鱼",
  //     "fisherName": "渔民A",
  //     "fishedTime": "2024.1.1",
  //     "fishedArea": "黄海",
  //     "tradeId": "",
  //     "tradeTime": ''
  //   },
  //   "factoryInfo": {
  //     "name": "黄鱼酥",
  //     "factoryName": "加工厂A",
  //     "processTime": "2024.1.2",
  //     "factoryAddress": "A市B区",
  //     "tradeId": "",
  //     "tradeTime": ''
  //   },
  //   "transportInfo": {
  //     "name": "黄鱼酥",
  //     "driverName": "司机A",
  //     "transportTime": "2024.1.3",
  //     "destination": "B市A区",
  //     "tradeId": "",
  //     "tradeTime": ''
  //   },
  //   "shopInfo": {
  //     "name": "黄鱼酥",
  //     "shopName": "商店A",
  //     "storageTime": "2024.1.4",
  //     "sellTime": "2024.1.4",
  //     "tradeId": "",
  //     "tradeTime": ''
  //   }
  //
  // };
  convertPropertyNamesToChinese(details.value);  // 使用函数转换属性名
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
  background-color: rgba(249, 249, 249, 0.3);
}

h1 {
  text-align: center;
  color: #333;
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
  border: 2px solid #ccc;
  border-radius: 4px;
  transition: border-color 0.3s;
}

.search-bar input:focus {
  border-color: #5eaaff;
  outline: none;
}

.search-bar button {
  padding: 10px 20px;
  background-color: #5eaaff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.search-bar button:hover {
  background-color: #1d89fd;
}

.results-table {
  width: 100%;
  max-width: 600px;
  margin: 0 auto;
  border-collapse: collapse;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.results-table th {
  background-color: #d2ffff;
  padding: 10px;
}

.results-table td {
  padding: 10px;
  background-color: #ffffff;
  border-bottom: 1px solid #ddd;
}

.results-table td button {
  margin: 0 5px;
  padding: 8px 15px;
  background-color: #72b4ff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.results-table td button:hover {
  background-color: #1e77d8;
}

.no-result-info {
  margin-top: 40px;
  text-align: center;
  color: #999;
}

.result-info {
  margin-top: 20px;
  display: flex; /* 添加此行 */
  flex-wrap: wrap; /* 添加此行 */
  gap: 10px; /* 添加此行 */
}

.info {
  background-color: #ffffff;
  padding: 20px;
  margin-right: 0;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  border-right: 10px solid #ddd; /* 修改此行 */
  flex: 1 1 20%; /* 添加此行 */
}

.info:last-child {
  border-right: none; /* 添加此行 */
}

h3 {
  color: #5eaaff;
}
</style>
