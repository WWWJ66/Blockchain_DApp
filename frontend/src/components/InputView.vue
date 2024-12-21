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
          v-if="userType !== '捕捞者' && userType !== '消费者'"
          label="溯源码:"
          style="width: 300px"
        >
          <el-input v-model="tracedata.tracecode" />
        </el-form-item>

        <!-- 捕捞者表单 -->
        <template v-if="userType === '捕捞者'">
          <el-form-item label="农产品名称:" style="width: 300px">
            <el-input v-model="tracedata.Fisherman_input.Fi_fishName" />
          </el-form-item>
          <el-form-item label="产地:" style="width: 300px">
            <el-input v-model="tracedata.Fisherman_input.Fi_origin" />
          </el-form-item>
          <el-form-item label="种植时间:" style="width: 300px">
            <el-input v-model="tracedata.Fisherman_input.Fi_plantTime" />
          </el-form-item>
          <el-form-item label="采摘时间:" style="width: 300px">
            <el-input v-model="tracedata.Fisherman_input.Fi_catchingTime" />
          </el-form-item>
          <el-form-item label="捕捞者名称:" style="width: 300px">
            <el-input v-model="tracedata.Fisherman_input.Fi_FishermanName" />
          </el-form-item>
        </template>

        <!-- 加工厂表单 -->
        <template v-if="userType === '加工厂'">
          <el-form-item label="商品名称:" style="width: 300px">
            <el-input v-model="tracedata.Factory_input.Fac_productName" />
          </el-form-item>
          <el-form-item label="生产批次:" style="width: 300px">
            <el-input v-model="tracedata.Factory_input.Fac_productionbatch" />
          </el-form-item>
          <el-form-item label="生产时间:" style="width: 300px">
            <el-input v-model="tracedata.Factory_input.Fac_productionTime" />
          </el-form-item>
          <el-form-item label="加工厂名称与厂址:" style="width: 300px">
            <el-input v-model="tracedata.Factory_input.Fac_factoryName" />
          </el-form-item>
          <el-form-item label="加工厂电话:" style="width: 300px">
            <el-input v-model="tracedata.Factory_input.Fac_contactNumber" />
          </el-form-item>
        </template>

        <!-- 物流公司表单 -->
        <template v-if="userType === '物流公司'">
          <el-form-item label="公司名称:" style="width: 300px">
            <el-input v-model="tracedata.Company_input.Dr_name" />
          </el-form-item>
          <el-form-item label="公司地址:" style="width: 300px">
            <el-input v-model="tracedata.Company_input.Dr_age" />
          </el-form-item>
          <el-form-item label="联系电话:" style="width: 300px">
            <el-input v-model="tracedata.Company_input.Dr_phone" />
          </el-form-item>
          <el-form-item label="运输车牌号:" style="width: 300px">
            <el-input v-model="tracedata.Company_input.Dr_carNumber" />
          </el-form-item>
          <el-form-item label="运输记录:" style="width: 300px">
            <el-input v-model="tracedata.Company_input.Dr_transport" />
          </el-form-item>
        </template>

        <!-- 商店表单 -->
        <template v-if="userType === '商店'">
          <el-form-item label="存入时间:" style="width: 300px">
            <el-input v-model="tracedata.Shop_input.Sh_storeTime" />
          </el-form-item>
          <el-form-item label="销售时间:" style="width: 300px">
            <el-input v-model="tracedata.Shop_input.Sh_sellTime" />
          </el-form-item>
          <el-form-item label="商店名称:" style="width: 300px">
            <el-input v-model="tracedata.Shop_input.Sh_shopName" />
          </el-form-item>
          <el-form-item label="商店位置:" style="width: 300px">
            <el-input v-model="tracedata.Shop_input.Sh_shopAddress" />
          </el-form-item>
          <el-form-item label="商店电话:" style="width: 300px">
            <el-input v-model="tracedata.Shop_input.Sh_shopPhone" />
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
// import { computed } from "vue";
// import { useStore } from "vuex";
import axios from "axios";

// const store = useStore();
// const name = computed(() => store.getters ? store.getters.name : "default");
// const userType = computed(() => store.getters ? store.getters.userType : "消费者");
const name = "ti"
const userType = "商店"

const tracedata = ref({
  tracecode: "",
  Fisherman_input: {
    Fi_fishName: "",
    Fi_origin: "",
    Fi_plantTime: "",
    Fi_catchingTime: "",
    Fi_FishermanName: "",
  },
  Factory_input: {
    Fac_productName: "",
    Fac_productionbatch: "",
    Fac_productionTime: "",
    Fac_factoryName: "",
    Fac_contactNumber: "",
  },
  Company_input: {
    Dr_name: "",
    Dr_age: "",
    Dr_phone: "",
    Dr_carNumber: "",
    Dr_transport: "",
  },
  Shop_input: {
    Sh_storeTime: "",
    Sh_sellTime: "",
    Sh_shopName: "",
    Sh_shopAddress: "",
    Sh_shopPhone: "",
  },
});

const submittracedata = () => {
  console.log(tracedata.value);
  const loading = window.$loading({
    lock: true,
    text: "数据上链中...",
    spinner: "el-icon-loading",
    background: "rgba(0, 0, 0, 0.7)",
  });

  const formData = new FormData();
  formData.append("username", name);
  formData.append("tracecode", tracedata.value.tracecode);

  function uplink(data) {
    return axios.post("http://localhost:9090/uplink", data, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  }

  switch (userType.value) {
    case "捕捞者":
      formData.append("arg1", tracedata.value.Fisherman_input.Fi_fishName);
      formData.append("arg2", tracedata.value.Fisherman_input.Fi_origin);
      formData.append("arg3", tracedata.value.Fisherman_input.Fi_plantTime);
      formData.append("arg4", tracedata.value.Fisherman_input.Fi_catchingTime);
      formData.append("arg5", tracedata.value.Fisherman_input.Fi_FishermanName);
      break;
    case "加工厂":
      formData.append("arg1", tracedata.value.Factory_input.Fac_productName);
      formData.append(
        "arg2",
        tracedata.value.Factory_input.Fac_productionbatch
      );
      formData.append("arg3", tracedata.value.Factory_input.Fac_productionTime);
      formData.append("arg4", tracedata.value.Factory_input.Fac_factoryName);
      formData.append("arg5", tracedata.value.Factory_input.Fac_contactNumber);
      break;
    case "物流公司":
      formData.append("arg1", tracedata.value.Company_input.Dr_name);
      formData.append("arg2", tracedata.value.Company_input.Dr_age);
      formData.append("arg3", tracedata.value.Company_input.Dr_phone);
      formData.append("arg4", tracedata.value.Company_input.Dr_carNumber);
      formData.append("arg5", tracedata.value.Company_input.Dr_transport);
      break;
    case "商店":
      formData.append("arg1", tracedata.value.Shop_input.Sh_storeTime);
      formData.append("arg2", tracedata.value.Shop_input.Sh_sellTime);
      formData.append("arg3", tracedata.value.Shop_input.Sh_shopName);
      formData.append("arg4", tracedata.value.Shop_input.Sh_shopAddress);
      formData.append("arg5", tracedata.value.Shop_input.Sh_shopPhone);
      break;
  }

  uplink(formData)
    .then((res) => {
      if (res.code === 200) {
        loading.close();
        window.$message({
          message:
            "上链成功，交易ID：" +
            res.txid +
            "\n溯源码：" +
            res.tracecode,
          type: "success",
        });
      } else {
        loading.close();
        window.$message({
          message: "上链失败",
          type: "error",
        });
      }
    })
    .catch((err) => {
      loading.close();
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
