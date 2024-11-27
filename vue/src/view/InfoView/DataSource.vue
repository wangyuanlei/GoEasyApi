<template>
  <div class="dataSource-box">
    <el-card v-loading="loading">
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
            </el-form>
            <div class="btn-box">
                <el-button size="large" @click="defaultClick">设为默认值</el-button>
                <el-button size="large" type="primary" @click="saveClick">保存</el-button>
            </div>
        </el-card>
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
}
//定义类型
interface defineSetting {
    DatabaseName?: string | null,
    Description?: string | null,
    OrmType?: string | null,
    Dns?: string | null,
};
//默认参数
const formSetting = ref<defineSetting>({
    DatabaseName: '',
    Description: '',
    OrmType:'',
    Dns:'',
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

}
// 监听 OrmType 的变化
watch(() => formSetting.value.OrmType, (newVal) => {
  if (newVal === 'mysql') {
    formSetting.value.Dns = 'username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local';
  } else if (newVal === 'postgresql') {
    formSetting.value.Dns = 'host=localhost user=your_username password=your_password dbname=your_db port=5432 sslmode=disable';
  } else if (newVal === 'sqlserver') {
    formSetting.value.Dns = 'sqlserver://username:password@localhost:1433?database=dbname';
  } else if (newVal === 'sqlite') {
    formSetting.value.Dns = './db.sql';
  }
});
//保存
const saveClick =()=>{
  dataSource.saveData(hasToken,formSetting.value).then((res:any) => {

  ElMessage.success('保存成功');
  getData()
}).finally(() => {
        loading.value = false; // 开始加载
    });
}
onMounted(()=>{
  getData()
});

const getData = ()=>{
  loading.value = true; // 开始加载
  dataSource.getDatas(hasToken).then((res:any) => {
    
  Object.assign(alldata, res.data);
  // console.log('alldataalldata',alldata);
  formSetting.value={
    ...res.data
  }
  
}).finally(() => {
        loading.value = false; // 开始加载
        // ElMessage.error('获取配置信息失败');
    });
}
</script>
<style scoped lang="scss">
.dataSource-box{
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #ebeef5;
}
.el-card {
    width: 50%;
    margin-bottom: 30px;
}

.card-header {
    text-align: center;
    font-size: 18px;
    margin-bottom: 12px;
}
.btn-box{
  display: flex;
  justify-content: end;
}
.el-radio__label {
  text-transform: none;
} 
</style>