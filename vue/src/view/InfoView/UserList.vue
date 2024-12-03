<template>
    <div class="search-user-table">
      <div class="search-user-top">
        <div class="search-left">
          <el-input v-model="input" style="width: 240px" ref="inputRef" size="large" placeholder="用户姓名 / 登陆账号" clearable @blur="handleSearch" @clear="handleClear" @keyup.enter="handleSearch"/>
          <el-button type="primary" @click="getTableData" :icon="Search" circle />
        </div>
        <div class="search-right">
          <el-button type="primary" @click="dialogFormVisible = true" size="large">+ 添加新用户</el-button>
        </div>
      </div>
      <div class="search-user-bottom">
        <el-table :data="tableData" style="width: 100%; height: 98%;" border stripe empty-text="没有数据">
          <el-table-column type="index" label="序号" width="60" align="left" header-align="center"></el-table-column>
          <el-table-column prop="Name" label="用户姓名" header-align="center"></el-table-column>
          <el-table-column prop="IsValid" label="状态" header-align="center" align="center">
            <template #default="scope">
              <el-tag v-if="scope.row.IsValid == 2" type="success">正常</el-tag>
              <el-tag v-else type="danger">禁用</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="Account" label="登录账号" header-align="center"></el-table-column>
          <el-table-column prop="DeptId" label="部门" show-overflow-tooltip width="230" header-align="center"></el-table-column>
          <!-- <el-table-column prop="Salt" label="盐值" header-align="center"></el-table-column> -->
          <el-table-column prop="RegisterTime" label="注册时间" show-overflow-tooltip width="230" header-align="center"></el-table-column>
          <el-table-column label="操作" header-align="center" align="center">
            <template #default="scope">
            <el-button-group>
              <el-tooltip class="box-item" effect="dark" content="编辑" placement="top">
                <el-button type="primary" :icon="Edit" size="small" @click="EditRow(scope.row)" />
              </el-tooltip>
              <el-tooltip class="box-item" effect="dark" content="修改密码" placement="top">
                <el-button type="warning" :icon="Setting" size="small" style="border-radius: 0;"
                  @click="changePassword(scope.row)" />
              </el-tooltip>
              <!-- <span v-if="scope.row.state * 1 === 20"> -->
              <!-- <el-tooltip v-if="scope.row.invalid * 1 === 1" class="box-item" effect="dark" content="启用账号"
                placement="top">
                <el-button type="success" size="small" style="border-radius: 0;" :icon="CircleCheckFilled"
                  @click="InvalidData(scope.row)" />
              </el-tooltip>
              <el-tooltip v-else class="box-item" effect="dark" content="停用账号" placement="top">
                <el-button type="info"  style="border-radius: 0;" size="small"
                class="iconfont icon-jinyong"
                  @click="InvalidData(scope.row)" />
              </el-tooltip> -->
              <el-tooltip class="box-item" effect="dark" content="删除" placement="top">
                <el-button type="danger" :icon="Delete" size="small" @click="handleDelete(scope.row)" />
              </el-tooltip>
            </el-button-group>
          </template>
          </el-table-column>
        </el-table>

      </div>
      <el-config-provider :locale="zhCn">
        <el-pagination background style="text-align: center;margin-top: 30px;height: 20px;" v-model:current-page="currentPage"
          v-model:page-size="pageSize" :page-sizes="[15, 20, 30, 50]" layout="total, sizes, prev, pager, next, jumper"
          :total="total" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
      </el-config-provider>
    </div>
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle" width="500" @close="resetForm">
    <el-form :model="form" :rules="rules" ref="formRef">
      <el-form-item label="姓名" prop="name" :label-width="formLabelWidth" required>
        <el-input v-model="form.name" autocomplete="off" />
      </el-form-item>
      <el-form-item  v-if="!isEdit" label="账号名" prop="account" :label-width="formLabelWidth" required>
        <el-input v-model="form.account" autocomplete="off" />
      </el-form-item>
      <el-form-item v-if="!isEdit" label="密码:" prop="password" :label-width="formLabelWidth" required>
          <el-input v-model="form.password" autocomplete="off" show-password></el-input>
      </el-form-item>
      <el-form-item label="部门" prop="deptId" :label-width="formLabelWidth" required>
        <el-input v-model="form.deptId" autocomplete="off" />
      </el-form-item>
      <el-form-item label="状态" prop="isValid" :label-width="formLabelWidth" required>
        <el-radio-group v-model="form.isValid">
          <el-radio :label="2">开启</el-radio>
          <el-radio :label="1">禁用</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSave">
          确定
        </el-button>
      </div>
    </template>
  </el-dialog>
  </template>
<script lang="ts" setup>
import { ref, watch ,onMounted} from 'vue'
import UserManage from "@/api/user";
import { ElMessageBox,ElMessage } from 'element-plus'
import { Delete,Edit,Setting,Search} from '@element-plus/icons-vue';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
// 模糊搜索框
const input = ref('')
const inputRef =ref(null);
const formLabelWidth = '100px';

// 处理模糊搜索
const handleSearch = () => {
    filterData.value.search = input.value;
    getTableData();
};
//清空输入框
const handleClear = () => {
    filterData.value.search = '';
  //   if (inputRef.value) {
  //   inputRef.value.blur(); // 让输入框失去焦点
  // }
    getTableData();
};
//弹窗是否打开
const dialogFormVisible = ref(false);


const loading = ref(false);
const hasToken = <string>localStorage.getItem('accessToken');

const tableData = ref([]);
//分页代码
const currentPage = ref(1);
const pageSize = ref(15);
const total = ref(0);
const handleSizeChange = (val: number) => {
  pageSize.value = val;
};
const handleCurrentChange = (val: number) => {
  currentPage.value = val;
};
const filterData = ref<any>({
  page: currentPage.value.toString(),
  page_size: pageSize.value.toString(),
  search: input.value,
});
// 监听分页和搜索输入的变化
watch([currentPage, pageSize], () => {
  filterData.value.page = currentPage.value.toString();
  filterData.value.page_size = pageSize.value.toString();
  // filterData.value.search = input.value;
  getTableData();
});
//加载列表接口
const getTableData = () => {
  loading.value = true; // 开始加载
  UserManage.getUesrList(hasToken, filterData.value).then(res => {
    // console.log('res', res.data.list);
    tableData.value = res.data.list;
    total.value = res.data.total;
  })
  .finally(() => {
    loading.value = false; // 加载完成
  });
};
onMounted(()=>{
    getTableData();
})  
// getTableData();
//注册用户
const dialogTitle = ref('注册用户');

//重置表单内容
const formRef = ref();
//弹窗打开默认参数 
const form = ref<any>({
  name: '',
  account: '',
  password: '',
  deptId: '',
  isValid:2
});
//重置表单
const resetForm = () => {
  form.value = {
    name: '',
    password: '',
    account: '',
    deptId: '',
    isValid:2

  };
  if (formRef.value) {
    formRef.value.clearValidate();
  }
  dialogTitle.value = '用户注册';
  isEdit.value = false;
};
//添加表单检验逻辑
const rules = {
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' }
  ],
  deptId: [
    { required: true, message: '请输入部门id', trigger: 'blur' }
  ],
  account: [
    { required: true, message: '请输入登陆账号', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ],
};
const changePassword =(row: any)=>{
  ElMessageBox.prompt('请输入新密码', '修改密码', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /^.{6,20}$/,
    inputErrorMessage: '密码长度6-20位',
  }).then(({ value }) => {
    UserManage.updtPsd(hasToken, row.UserId.toString() ,value).then(res => {
      ElMessage.success('密码修改成功');
    })
  })
}
//弹框保存逻辑
const handleSave = () => {
  formRef.value.validate((valid: boolean) => {
    if (valid) {
      if (isEdit.value && currentUserId.value) {
        console.log('form.value', form.value);
        
        UserManage.updateUser(hasToken, currentUserId.value, form.value).then(res => {
          ElMessage.success('修改成功');
          dialogFormVisible.value = false;
          getTableData();
        }).catch(error => {
          console.error(error);
          ElMessage.error('修改失败');
        });
      }else{
        //新增保存   
        UserManage.createUser(hasToken, form.value).then(res => {
          // console.log('res', res);
          dialogFormVisible.value = false;
          ElMessage.success('用户添加成功');
          resetForm();
          getTableData()
        })
      }
    } else {
      ElMessage.error('请填写完整的用户信息');
    }
  });

};
//编辑按钮点击逻辑
const isEdit = ref(false); // 用于区分是新增还是编辑
const currentUserId = ref<string | null>(null); // 当前编辑的用户ID
const EditRow = (row: any) => {
  dialogFormVisible.value = true;
  isEdit.value = true;
  dialogTitle.value = '编辑用户';
  currentUserId.value = row.UserId.toString();
  const str = row.UserId.toString();
  UserManage.getUser(hasToken, str).then(res => {
    // console.log('res', res);
    //数据回填
    form.value = {
      name: res.data.Name,
      account: res.data.Account,
      deptId: res.data.DeptId,
      password: res.data.Password,
      isValid: res.data.IsValid
    };
  }).catch(error => {
    console.error(error);
    ElMessage.error('获取用户信息失败');
  });
};
//删除按钮点击逻辑
const handleDelete = (row: any) => {
  ElMessageBox.confirm('此操作将永久删除该用户, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
      ElMessage.success(row.UserId.toString());
    // UserManage.deleteUser(hasToken, row.UserId.toString()).then(res => {
    //   ElMessage.success('删除成功');
    // }).catch(() => {
    //   ElMessage.info('已取消删除');
    // });
  })
}
</script>
<style scoped>
.search-user-table {
  width: 100%;
  height: 100%;
  padding: 20px;
  box-sizing: border-box;
  background-color: #ebeef5;
}
.search-user-table /deep/ .is-leaf {
    background-color: var(--v-table-bg-color);
    color: var(--v-table-color);
    font-weight: bold;
}

.search-user-top {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
}

.search-right .el-button {
  height: 40px;
}

.el-input{
  margin-right: 10px;
}

.search-user-bottom {
  height: calc(100% - 90px);
  /* overflow: auto; */
}

.el-pagination {
  float: right;
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
.el-table img{
  width: 18px;
  height: 18px;
  margin-right: 5px;
  vertical-align: sub;
}
</style>