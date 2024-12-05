<template>
    <div class="overall-plugin-table">
      <div class="overall-plugin-top">
        <div class="search-left">
        </div>
        <div class="search-right">
          <el-button type="primary" @click="addInterface" size="large">+ 添加新接口</el-button>
        </div>
      </div>
      <div class="overall-plugin-bottom">
        <el-table :data="tableData" style="width: 100%; height: 100%;" border stripe :header-cell-style="{background: '#f5f7fa',color: '#909399'}" empty-text="没有数据">
          <el-table-column type="index" label="序号" width="60" align="left" header-align="center"></el-table-column>
          <el-table-column prop="Path" label="接口路径" width="100" header-align="center"></el-table-column>
          <el-table-column prop="Method" label="请求方式" width="100" header-align="center" align="center"></el-table-column>
          <el-table-column prop="Description" label="接口描述" width="100" header-align="center"></el-table-column>
          <el-table-column prop="CacheEnabled" label="接口缓存启用" width="120" align="center" header-align="center"></el-table-column>
          <el-table-column prop="CacheTime" label="缓存时间" show-overflow-tooltip width="230" align="center" header-align="center"></el-table-column>
          <el-table-column prop="RateLimitEnabled" label="接口限流启用" width="120" header-align="center"></el-table-column>
          <el-table-column prop="RateLimitCount" label="接口限流次数" width="120" header-align="center"></el-table-column>
          <el-table-column prop="RateLimitTime" label="接口限流时间" width="120" header-align="center"></el-table-column>
          <el-table-column prop="TokenValidationEnabled" label="Token验证启用" width="150" header-align="center"></el-table-column>
          <el-table-column prop="ReturnType" label="接口返回类型" width="150" header-align="center"></el-table-column>
          <el-table-column prop="ReturnValMode" label="接口返回模式" width="150" header-align="center"></el-table-column>
          <el-table-column label="操作" header-align="center" align="center" min-width="120" fixed="right">
            <template #default="scope">
            <el-button-group>
              <el-tooltip class="box-item" effect="dark" content="编辑" placement="top">
                <el-button type="primary" :icon="Edit" size="small" @click="EditRow(scope.row)" />
              </el-tooltip>
              <el-tooltip class="box-item" effect="dark" content="删除" placement="top">
                <el-button type="danger" :icon="Delete" size="small" @click="handleDelete(scope.row)" />
              </el-tooltip>
            </el-button-group>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </template>
<script lang="ts" setup>
import { ref , onMounted} from 'vue'
import SetApi from '@/api/setApi';
import { ElMessageBox,ElMessage } from 'element-plus'
import { Delete,Edit} from '@element-plus/icons-vue';
const loading = ref(false);
const hasToken = <string>localStorage.getItem('accessToken');

  import { defineEmits } from 'vue';

  const emit = defineEmits(['switchToActuator']);

const tableData = ref([]);
//加载列表接口
const getTableData = () => {
  loading.value = true; // 开始加载
  SetApi.getUesrList(hasToken).then(res => {
    tableData.value = res.data;
  })
};
onMounted(()=>{
    getTableData();
});

const handleDelete = (row:any) => {
  ElMessageBox.confirm('此操作将永久删除该接口, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    SetApi.deleteInterface(hasToken,row.Id).then(() => {
      ElMessage.success('删除成功');
      getTableData();
    })
  })
}
const EditRow = (row:any) => {
  emit('switchToActuator', row ,'actuator');
};
const addInterface = () => {
  emit('switchToActuator', null,'actuator');
};  
</script>
<style scoped>
.overall-plugin-table {
  width: 100%;
  height: 100%;
  padding: 20px;
  box-sizing: border-box;
}

.overall-plugin-top {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
}

.el-input,
.el-select {
  margin-right: 10px;
}

.overall-plugin-bottom {
  height: calc(100% - 90px);
  /* overflow: auto; */
}


.el-table .el-button {
  margin: 0;
}
.state-span{
  display: flex;
  align-items: center;
}
.state-span .el-icon , .state-span i{
  margin-right: 5px;
}
</style>