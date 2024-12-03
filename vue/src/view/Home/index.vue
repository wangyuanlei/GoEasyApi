<script setup lang="ts">
import { ref,onMounted,watch } from 'vue';
import router from '@/router';
import { useRoute } from 'vue-router';
import { ElDropdown, ElDropdownMenu, ElDropdownItem, ElAvatar,ElMessage,ElMessageBox } from 'element-plus';
import HomeApi from "@/api/home";
// import SvgIcon from '@/components/svg/svg.vue';
const activeOption = ref<string | null>(null);
//修改密码逻辑
const showDialog = ref(false);
const formRef = ref();

const form = ref({
  oldPass: '',
  newPass: ''
});
//重置表单
const resetForm = () => {
  form.value = {
    oldPass: '',
    newPass: '',
  };
  if (formRef.value) {
    formRef.value.clearValidate();
  }
};
//添加表单检验逻辑
const rules = {
  oldPass: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ],
  newPass: [
    { required: true, message: '请输入确认密码', trigger: 'blur' }
  ],
};
const hasToken = <string>localStorage.getItem('accessToken');
const submitForm = () => {
  formRef.value?.validate((valid: boolean) => {
    if (valid) {
                  //新增保存   
        HomeApi.changePsd(hasToken, form.value).then(res => {
          ElMessage.success('密码修改成功');
          showDialog.value = false;
          resetForm();
        })
        }else{
          ElMessage.error('请检查输入信息');
          resetForm();
        }
  });
};
  const handleNavClick = (option: string) => {
  activeOption.value = option;
  switch (option) {
    case 'option1':
      router.push('/Home/DataSource');
      break;
    case 'option2':
      router.push('/Home/SetApi');
      break;
    case 'option3':
      router.push('/Home/SysTemInfo');
      break;
    case 'option4':
        router.push('/Home/UserList');
      break;
  }
};
const logout = () => {
  ElMessageBox.confirm('确定要退出登录吗？', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    localStorage.removeItem('accessToken');
    router.push('/');
  }); 
};
// 在组件挂载时选中第一个选项
onMounted(() => {
  // activeOption.value = 'option1';
});
// 监听路由变化，更新 activeOption
const route = useRoute();
watch(route, (newRoute) => {
  
  if (newRoute.path === '/Home/DataSource') {
    activeOption.value = 'option1';
  }else if(newRoute.path === '/Home/SetApi'){
    activeOption.value = 'option2';
  }else if (newRoute.path === '/Home/SysTemInfo') {
    activeOption.value = 'option3';
  } else if (newRoute.path === '/Home/UserList') {
    activeOption.value = 'option4';
  }
}, { immediate: true });
</script>

<template>
  <div class="home">
    <header class="header">
      <div class="logo">
        <!-- <img src="../../../public/images/login.jpeg" alt="Logo" /> -->
        <svg-icon name="index-top" style="margin-right: 8px;width:35px;height: 35px;"/>
      </div>
      <nav class="nav">
        <ul class="nav-list">
          <li 
            :class="{ active: activeOption === 'option1' }" 
            class="nav-item"
            @click="handleNavClick('option1')">
              <svg-icon name="data-source" style="margin-right: 8px;"/>
              数据源
          </li>
          <li 
          :class="{ active: activeOption === 'option2' }" 
          class="nav-item" @click="handleNavClick('option2')">
            <svg-icon name="application--data" style="margin-right: 8px;"/>
          应用程序接口
        </li>
          <li 
            :class="{ active: activeOption === 'option3' }" 
            class="nav-item" @click="handleNavClick('option3')">
            <svg-icon name="system-settings" style="margin-right: 8px;"/>
            系统设置
        </li>
          <li 
          :class="{ active: activeOption === 'option4' }" 
          class="nav-item"
           @click="handleNavClick('option4')">
           <svg-icon name="user-management" style="margin-right: 8px;" width="20px" height="20px"/>
           用户管理</li>
        </ul>
      </nav>
      <div class="user-info">
        <!-- <el-button type="primary" :icon="EditPen" circle @click="showDialog = true"></el-button> -->
        <el-dropdown>
          <!-- <el-avatar src="/path/to/your/avatar.png" size="small"></el-avatar> -->
          <svg-icon name="setting" style="width: 21px;height: 21px;margin-right: 15px;"></svg-icon>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="showDialog = true">
                <svg-icon name="editPsd" style="margin-right: 8px;"/>
                修改密码
              </el-dropdown-item>
              <el-dropdown-item @click="logout">
                <svg-icon name="log-out" style="margin-right: 8px;"/>
                退出登录
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </header>
    <main class="content">
      <router-view></router-view>
    </main>
    <el-dialog title="修改密码" v-model="showDialog" width="30%" @close="resetForm">
      <el-form label-width="80px" :rules="rules" :model="form" ref="formRef">
        <el-form-item label="老密码" prop="oldPass" required>
          <el-input v-model="form.oldPass" type="password" placeholder="请输入老密码"></el-input>
        </el-form-item>
        <el-form-item label="新密码" prop="newPass" required>
          <el-input v-model="form.newPass" type="password" placeholder="请输入新密码"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showDialog = false">取 消</el-button>
          <el-button type="primary" @click="submitForm">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<style lang="scss" scoped>
.home {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background-color: #333;
  color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.logo img {
  height: 30px;
}

.nav {
  flex: 1;
  display: flex;
  // justify-content: center;
  height: 100%;
}

.nav-list {
  list-style: none;
  display: flex;
  gap: 20px;
  margin: 0;
}
.nav-item{
  display: flex;
  align-items: center;
  padding:20px 8px;
  &:hover {
    background-color: #555;
    color: #fff;
    cursor: pointer;
  }

  &.active {
    background-color: #444;
    color: #fff;
  }
}
.nav-item a {
  color: #fff;
  text-decoration: none;
  font-size: 16px;
  padding: 0 8px;
}

.user-info {
  display: flex;
  align-items: center;
}

.content {
  flex: 1;
  padding: 20px;
  background-color: #f0f2f5;
}
.el-avatar{
  width: 32px;
  height: 32px;
  margin-left: 10px;
}
::v-deep(.el-dropdown .el-tooltip__trigger) {
  outline: none; /* 移除默认的聚焦边框 */
  border: none; /* 移除默认的边框 */
  background-color: transparent; /* 移除默认的背景色 */
}
</style>