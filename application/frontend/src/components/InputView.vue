<template>
  <div class="uplink-container">
    <!-- 用户信息 -->
    <div style="color: #909399; margin-bottom: 30px">
      当前用户：{{ name }}; 用户角色: {{ userType }}
    </div>
    <!-- 表单部分 -->
    <div>
      <el-form
        ref="form"
        :model="tracedata"
        label-width="120px"
        size="mini"
      >
        <!-- 通用溯源码 -->
        <el-form-item
          v-if="userType !== '水产品养殖者' && userType !== '消费者'"
          label="溯源码:"
          style="width: 300px"
        >
          <el-input v-model="tracedata.tracecode" />
        </el-form-item>

        <!-- 捕捞者表单 -->
        <template v-if="userType === '水产品养殖者'">
          <el-form-item label="水产品名称:" style="width: 300px">
            <el-input v-model="tracedata.Fisher_input.Fi_fishName" />
          </el-form-item>
          <el-form-item label="养殖者名称:" style="width: 300px">
            <el-input v-model="tracedata.Fisher_input.Fi_fisherName" />
          </el-form-item>
          <el-form-item label="捕捞时间:" style="width: 300px">
            <el-input v-model="tracedata.Fisher_input.Fi_fishedTime" />
          </el-form-item>
          <el-form-item label="捕捞海域:" style="width: 300px">
            <el-input v-model="tracedata.Fisher_input.Fi_fishedArea" />
          </el-form-item>
        </template>

        <!-- 工厂表单 -->
        <template v-if="userType === '工厂'">
          <el-form-item label="商品名称:" style="width: 300px">
            <el-input v-model="tracedata.Factory_input.Fac_productName" />
          </el-form-item>
          <el-form-item label="工厂名称:" style="width: 300px">
            <el-input v-model="tracedata.Factory_input.Fac_factoryName" />
          </el-form-item>
          <el-form-item label="加工时间:" style="width: 300px">
            <el-input v-model="tracedata.Factory_input.Fac_processTime" />
          </el-form-item>
          <el-form-item label="工厂厂址:" style="width: 300px">
            <el-input v-model="tracedata.Factory_input.Fac_factoryAdress" />
          </el-form-item>
        </template>

        <!-- 物流公司表单 -->
        <template v-if="userType === '物流公司'">
          <el-form-item label="公司名称:" style="width: 300px">
            <el-input v-model="tracedata.Driver_input.Dr_name" />
          </el-form-item>
          <el-form-item label="联系电话:" style="width: 300px">
            <el-input v-model="tracedata.Driver_input.Dr_phone" />
          </el-form-item>
          <el-form-item label="运输时间:" style="width: 300px">
            <el-input v-model="tracedata.Driver_input.Dr_transportTime" />
          </el-form-item>
          <el-form-item label="运输单号:" style="width: 300px">
            <el-input v-model="tracedata.Driver_input.Dr_id" />
          </el-form-item>
        </template>

        <!-- 商店表单 -->
        <template v-if="userType === '商店'">
          <el-form-item label="商店名称:" style="width: 300px">
            <el-input v-model="tracedata.Shop_input.Sh_shopName" />
          </el-form-item>
          <el-form-item label="商店地址:" style="width: 300px">
            <el-input v-model="tracedata.Shop_input.Sh_shopAddress" />
          </el-form-item>
          <el-form-item label="存入时间:" style="width: 300px">
            <el-input v-model="tracedata.Shop_input.Sh_storageTime" />
          </el-form-item>
          <el-form-item label="销售时间:" style="width: 300px">
            <el-input v-model="tracedata.Shop_input.Sh_sellTime" />
          </el-form-item>
        </template>
      </el-form>

      <!-- 底部操作 -->
      <div class="dialog-footer">
        <el-button
          v-if="userType !== '消费者'"
          type="primary"
          plain
          style="margin-left: 220px"
          @click="submittracedata"
        >
          提 交
        </el-button>
        <div v-else style="color: gray">消费者没有权限录入！请使用溯源功能！</div>
      </div>
    </div>
  </div>
</template>


<script setup>
import { ref } from "vue";
import axios from "axios";

import {useUserStore} from "@/store/userStore";
const name = useUserStore().username;
const userType = useUserStore().userType;
const token = useUserStore().token;
const tracedata = ref({
  tracecode: "",
  Fisher_input: {
    Fi_fishName: "",
    Fi_fisherName: "",
    Fi_fishedTime: "",
    Fi_fishedArea: "",
  },
  Factory_input: {
    Fac_productName: "",
    Fac_factoryName: "",
    Fac_processTime: "",
    Fac_factoryAddress: "",
  },
  Driver_input: {
    Dr_name: "",
    Dr_phone: "",
    Dr_transportTime: "",
    Dr_id: "",
  },
  Shop_input: {
    Sh_shopName: "",
    Sh_shopAddress: "",
    Sh_storageTime: "",
    Sh_sellTime: ""
  },
});

const submittracedata = () => {
  const formData = new FormData();
  formData.append("username", name);
  formData.append("tracecode", tracedata.value.tracecode);

  function uplink(data) {
    return axios.post("http://192.168.133.131:9090/uplink", data, {
      headers: {
          'Authorization': token
      },
    });
  }

  switch (userType) {
    case "水产品养殖者":
      formData.append("arg1", tracedata.value.Fisher_input.Fi_fishName);
      formData.append("arg2", tracedata.value.Fisher_input.Fi_fisherName);
      formData.append("arg3", tracedata.value.Fisher_input.Fi_fishedTime);
      formData.append("arg4", tracedata.value.Fisher_input.Fi_fishedArea);
      break;
    case "工厂":
      formData.append("arg1", tracedata.value.Factory_input.Fac_productName);
      formData.append("arg2", tracedata.value.Factory_input.Fac_factoryName);
      formData.append("arg3", tracedata.value.Factory_input.Fac_processTime);
      formData.append("arg4", tracedata.value.Factory_input.Fac_factoryAddress);
      break;
    case "物流公司":
      formData.append("arg1", tracedata.value.Driver_input.Dr_name);
      formData.append("arg2", tracedata.value.Driver_input.Dr_phone);
      formData.append("arg3", tracedata.value.Driver_input.Dr_transportTime);
      formData.append("arg4", tracedata.value.Driver_input.Dr_id);
      break;
    case "商店":
      formData.append("arg1", tracedata.value.Shop_input.Sh_shopName);
      formData.append("arg2", tracedata.value.Shop_input.Sh_shopAddress);
      formData.append("arg3", tracedata.value.Shop_input.Sh_storageTime);
      formData.append("arg4", tracedata.value.Shop_input.Sh_sellTime);
      break;
  }

  uplink(formData)
    .then((res) => {
      if (res.data.code === 200) {
        alert("上链成功，交易ID：" + res.data.txid + "\n溯源码：" + res.data.tracecode);
      } else {
        alert("上链失败");
        console.log(res.data);
      }
    })
    .catch((err) => {
      console.error(err);
    });
};
</script>

<style lang="scss" scoped>
.uplink {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
