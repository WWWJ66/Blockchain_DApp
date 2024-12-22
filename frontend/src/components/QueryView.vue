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
    
    <div v-if="detailsVisible" class="overlay" @click="closeDetails">
      <div class="result-info-box" @click.stop>
        <button class="close-btn" @click="closeDetails">关闭</button>
        <div class="result-info">
          <!-- 基本信息 -->
          <div class="info-box">
            <h3>基本信息</h3>
            <p><strong>水产品名称:</strong> {{ details.fisher_input.fi_fishName }}</p>
            <p><strong>水产品养殖者:</strong> {{ details.fisher_input.fi_fisherName }}</p>
            <p><strong>捕捞时间:</strong> {{ details.fisher_input.fi_fishedTime }}</p>
            <p><strong>捕捞区域:</strong> {{ details.fisher_input.fi_fishedArea }}</p>
            <div class="transaction-info">
              <p>交易ID: {{ details.fisher_input.fi_txid }}</p>
              <p>交易时间: {{ details.fisher_input.fi_timestamp }}</p>
            </div>
          </div>
    
          <!-- 加工信息 -->
          <div class="info-box">
            <h3>加工信息</h3>
            <p><strong>产品名称:</strong> {{ details.factory_input.fac_productName || '暂无数据' }}</p>
            <p><strong>工厂名称:</strong> {{ details.factory_input.fac_factoryName || '暂无数据' }}</p>
            <p><strong>加工时间:</strong> {{ details.factory_input.fac_processTime || '暂无数据' }}</p>
            <p><strong>工厂地址:</strong> {{ details.factory_input.fac_factoryAddress || '暂无数据' }}</p>
            <div class="transaction-info">
              <p>交易ID: {{ details.factory_input.fac_txid || '暂无数据' }}</p>
              <p>交易时间: {{ details.factory_input.fac_timestamp || '暂无数据' }}</p>
            </div>
          </div>
    
          <!-- 运输信息 -->
          <div class="info-box">
            <h3>运输信息</h3>
            <p><strong>公司名称:</strong> {{ details.driver_input.dr_name || '暂无数据' }}</p>
            <p><strong>联系电话:</strong> {{ details.driver_input.dr_phone || '暂无数据' }}</p>
            <p><strong>运输时间:</strong> {{ details.driver_input.dr_transportTime || '暂无数据' }}</p>
            <p><strong>运输单号:</strong> {{ details.driver_input.dr_id || '暂无数据' }}</p>
            <div class="transaction-info">
              <p>交易ID: {{ details.driver_input.dr_txid || '暂无数据' }}</p>
              <p>交易时间: {{ details.driver_input.dr_timestamp || '暂无数据' }}</p>
            </div>
          </div>
    
          <!-- 商店信息 -->
          <div class="info-box">
            <h3>商店信息</h3>
            <p><strong>商店名称:</strong> {{ details.shop_input.sh_shopName || '暂无数据' }}</p>
            <p><strong>商店地址:</strong> {{ details.shop_input.sh_shopAddress || '暂无数据' }}</p>
            <p><strong>上架时间:</strong> {{ details.shop_input.sh_storageTime || '暂无数据' }}</p>
            <p><strong>售出时间:</strong> {{ details.shop_input.sh_sellTime || '暂无数据' }}</p>
            <div class="transaction-info">
              <p>交易ID: {{ details.shop_input.sh_txid || '暂无数据' }}</p>
              <p>交易时间: {{ details.shop_input.sh_timestamp || '暂无数据' }}</p>
            </div>
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
const username = useUserStore().username;
const userType = useUserStore().userType;
const token = useUserStore().token;
const detailsVisible = ref(false);

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
  const productlistResponse = 
    await axios.post("http://192.168.133.131:9090/getFishList", 
      {},
      { headers: {
        'Authorization': token
      },}
    );
  products.value = JSON.parse(productlistResponse.data.data);
  console.log(products.value);
  console.log(products.value.length);
}

const checkProduct = async (product) => {
  searchingCode.value = '';
  await getDetails(product);
}

const getDetails = async (traceCode) => {
  var formData = new FormData()
  formData.append('tracecode',traceCode)
  const res = await axios.post("http://192.168.133.131:9090/getFishInfo",formData);
  details.value = JSON.parse(res.data.data);
  console.log(details.value);
  detailsVisible.value = true;  // 显示弹窗
};

const closeDetails = () => {
  detailsVisible.value = false;  // 关闭弹窗
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
  display: flex;
  gap: 50px; /* 框之间的间距 */
  justify-content: center; /* 均匀分布 */
  align-items: flex-start; /* 元素顶部对齐 */
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

.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.result-info-box {
  background: white;
  padding: 20px;
  width: 90%; /* 调整宽度使弹窗更加狭长 */
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  position: relative;
  overflow-y: auto;
  max-height: 80vh; /* 允许滚动 */
} 

.info-box {
  background-color: #f9f4f4;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  flex: 1 1 15%; /* 每个框占据 22% 宽度，自动分配空间 */
  max-width: 15%; /* 保证框不超过最大宽度 */
  min-width: 100px; /* 设置最小宽度，避免在屏幕较小时过小 */
}
.info-box:hover {
  transform: translateY(-5px); /* 悬浮时上移效果 */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
}
.info-box h3 {
  color: #5eaaff;
}


.info-box p {
  margin: 8px 0;
}

.close-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  padding: 8px 16px;
  background-color: #ff4d4d;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.transaction-info {
  display: none;
  background-color: #f0f8ff;
  padding: 10px;
  margin-top: 10px;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
</style>
