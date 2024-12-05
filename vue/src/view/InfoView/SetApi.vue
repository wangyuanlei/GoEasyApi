<template>
    <div class="set-api">
        <!-- <div class="set-api-left"> -->
            <!-- <el-tabs :tab-position="'left'" class="demo-tabs" v-model="activeTab" @tab-change="handleTabChange">
                <el-tab-pane label="接口列表" name="apiList"></el-tab-pane>
                <el-tab-pane label="执行器" name="actuator" ></el-tab-pane>
            </el-tabs> -->
        <!-- </div> -->
        <div class="set-api-right">
            <!-- <div class="set-api-info-header">
                创建API
            </div> -->
            <div class="set-api-info-content">
                <component :is="currentComponent" @switchToActuator="handleSwitchToActuator" :data="actuatorData"></component>
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import Actuator from '@/components/Actuator.vue';
import ApiInterfacelist from '@/components/ApiInterface.vue';


    
const currentComponent = ref<any>(ApiInterfacelist);
// const activeTab = ref('apiList');
const actuatorData = ref<any>(null);
const handleTabChange = (tabName: string) => {
  switch (tabName) {
    case 'apiList':
      currentComponent.value = ApiInterfacelist;
      actuatorData.value=null
      break;
    case 'actuator':
      currentComponent.value = Actuator;
      break;
  }
};
const handleSwitchToActuator = (data: any,type:any) => {
    // console.log('type',data);
    
//   activeTab.value = type;
  handleTabChange(type);
  
  const actuatorId =data?data.Id:null;
  // 传递数据给 Actuator 组件
  if(!data){
    // console.log('data',data);
    actuatorData.value=null
  }
  actuatorData.value = actuatorId;

};
onMounted(() => {
  // 初始化时加载基本信息组件
//   handleTabChange(activeTab.value);
});

</script>
<style scoped lang="scss">
.set-api{
    width: 100%;
    height: calc(100%);
    display: flex;
    background-color: #ebeef5;
}
.set-api-left{
    width: 150px;
    height: 100%;
}
::v-deep(.el-tabs){
    margin-top: 56px;
}   
.set-api-right{
    width: 100%;
    height: 100%;
}
.set-api-info-header{
    text-align: center;
    font-weight: bold;
    font-size: 20px;
    margin-bottom: 30px;
}
::v-deep(.el-tabs__content){
    display: none;
}
.set-api-info-content{
    width: 100%;
    height: 100%;
}
</style>