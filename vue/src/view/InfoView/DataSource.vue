<template>
  <div class="dataSource-box">
    <el-card v-loading="loading" class="custom-card">
      <!-- <div class="custom-card"> -->
      <div class="card-header">
        <span>数据源</span>
      </div>
      <el-form :model="formSetting" :rules="rules" ref="formRef1">
        <el-form-item label="数据库名称" :label-width="formLabelWidth" prop="DatabaseName">
          <el-input v-model="formSetting.DatabaseName" size="large" clearable />
        </el-form-item>
        <el-form-item label="描述" :label-width="formLabelWidth" prop="Description">
          <el-input v-model="formSetting.Description" size="large" clearable />
        </el-form-item>
        <el-form-item label="数据库类型" prop="OrmType" :label-width="formLabelWidth" required>
          <el-radio-group v-model="formSetting.OrmType">
            <el-radio :label="'mysql'">mysql</el-radio>
            <el-radio :label="'postgresql'">postgresql</el-radio>
            <el-radio :label="'sqlserver'">sqlserver</el-radio>
            <el-radio :label="'sqlite'">sqlite</el-radio>                        
          </el-radio-group>
        </el-form-item>
        <el-form-item label="链接地址" :label-width="formLabelWidth" prop="Dns">
          <el-input v-model="formSetting.Dns" size="large" clearable />
        </el-form-item>
        <el-form-item label="参考值" :label-width="formLabelWidth">
          <span>{{ formSetting.DemoDns }}</span>
        </el-form-item>
      </el-form>
      <div class="btn-box">
        <el-button size="large" @click="defaultClick">设为默认值</el-button>
        <el-button size="large" type="primary" @click="saveClick">保存</el-button>
      </div>
    </el-card>
  <!-- </div> -->
  </div>
</template>
<script lang="ts" setup>
import { ref ,onMounted , watch} from 'vue';
import dataSource from '@/api/dataSource';
const hasToken = <string>localStorage.getItem('accessToken');
import { ElMessage } from 'element-plus'
const loading = ref(false);
const formLabelWidth = '110px';
//默认数据
const alldata ={
    DatabaseName: '',
    Description: '',
    OrmType:'',
    Dns:'',
    DemoDns:''
}
//定义类型
interface defineSetting {
    DatabaseName?: string | null,
    Description?: string | null,
    OrmType?: string | null,
    Dns?: string | null,
    DemoDns?: string | null,
};
//默认参数
const formSetting = ref<defineSetting>({
    DatabaseName: '',
    Description: '',
    OrmType:'',
    Dns:'',
    DemoDns:''
});

const rules = {
  DatabaseName: [
    { required: true, message: '请输入数据库名称', trigger: 'blur' }
  ],
  Description: [
    { required: true, message: '请输入描述', trigger: 'blur' }
  ],
  OrmType: [
    { required: true, message: '请输入数据库类型', trigger: 'blur' }
  ],
  Dns: [
    { required: true, message: '请输入链接地址', trigger: 'blur' }
  ],
};
//设为默认值
const defaultClick =()=>{
  formSetting.value={
    ...alldata
  }
  switch (formSetting.value.OrmType) {
    case 'mysql':
      formSetting.value.DemoDns = 'username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local';
      break;
    case 'postgresql':
      formSetting.value.DemoDns = 'host=localhost user=your_username password=your_password dbname=your_db port=5432 sslmode=disable'
      break;
    case'sqlserver':
      formSetting.value.DemoDns ='sqlserver://username:password@localhost:1433?database=dbname';
      break;
    case'sqlite':
      formSetting.value.DemoDns = './db.sql';
      break;
  }
}
// 监听 OrmType 的变化
watch(() => formSetting.value.OrmType, (newVal) => {
  if (newVal === 'mysql') {
    formSetting.value.DemoDns = 'username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local';
  } else if (newVal === 'postgresql') {
    formSetting.value.DemoDns = 'host=localhost user=your_username password=your_password dbname=your_db port=5432 sslmode=disable';
  } else if (newVal === 'sqlserver') {
    formSetting.value.DemoDns = 'sqlserver://username:password@localhost:1433?database=dbname';
  } else if (newVal === 'sqlite') {
    formSetting.value.DemoDns = './db.sql';
  }
});
//保存
const saveClick =()=>{
  dataSource.saveData(hasToken,formSetting.value).then((res:any) => {

  ElMessage.success('保存成功');
  getData()
})
// .finally(() => {
//         loading.value = false; // 开始加载
//     });
}
onMounted(()=>{
  getData()
});

const getData = ()=>{
  loading.value = true; // 开始加载
  dataSource.getDatas(hasToken).then((res:any) => {
    loading.value = false; ;
  Object.assign(alldata, res.data);
  // console.log('alldataalldata',alldata);
  formSetting.value={
    ...res.data
  }
  
})
// .finally(() => {
//         loading.value = false; // 开始加载
//         // ElMessage.error('获取配置信息失败');
//     });
}
</script>
<style scoped lang="scss">
.dataSource-box {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  // align-items: center;
  background: linear-gradient(135deg, #f5f5f5, #ebeef5); /* 渐变背景 */
  padding: 20px; /* 增加内边距 */
}

.custom-card {
  width: 90%;
  height: 70%;
  // box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* 添加阴影 */
  // border-radius: 12px; /* 圆角 */
  // transition: transform 0.3s ease-in-out; /* 平滑过渡效果 */
}

.card-header {
  text-align: center;
  font-size: 25px;
  margin-bottom: 30px;
  color: #333; /* 深色文字 */
}

.btn-box {
  display: flex;
  justify-content: end;
  margin-top: 20px; /* 增加顶部间距 */
}

.el-radio__label {
  text-transform: none;
}

/* 可以添加一些图标或辅助信息 */
.icon-box {
  position: absolute;
  top: 20px;
  left: 20px;
  font-size: 24px;
  color: #999;
}
</style>