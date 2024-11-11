<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router';
import { ElForm, ElFormItem, ElInput, ElButton, ElMessage } from 'element-plus';
import AuthAPI from "@/api/login";

const router = useRouter();
const formRef = ref<InstanceType<typeof ElForm>>();
const accountInputRef = ref<InstanceType<typeof ElInput>>();
const passwordInputRef = ref<InstanceType<typeof ElInput>>();

const form = ref({
  account: '',
  password: ''
});

const title = ref('用户');

interface LoginData {
  // 用户名
  account: string;
  // 密码
  pwd: string;
}

const handleLogin = () => {
  if (!formRef.value) return;

  formRef.value.validate((valid) => {
    if (valid) {
      // 表单验证通过，进行登录操作
      const loginData = ref<LoginData>({
        account: form.value.account,
        pwd: form.value.password
      });
      AuthAPI.getUesrLogin(loginData.value).then(res => {
        localStorage.setItem('accessToken', `${res.data}`);
        ElMessage.success("登录成功");
        router.replace({ path: "/Home" });
      }).catch(err => {
        console.log('err=', err);
        ElMessage.error('请输入正确的账号或密码');
      });
    } else {
      // 表单验证失败
      if (!form.value.account && !form.value.password) {
        ElMessage.error('请输入账号和密码');
      } else if (!form.value.account) {
        ElMessage.error('请输入账号');
      } else if (!form.value.password) {
        ElMessage.error('请输入密码');
      }
    }
  });
};

const handleAccountEnter = () => {
  if (passwordInputRef.value) {
    passwordInputRef.value.focus();
  }
};

const handlePasswordEnter = () => {
  handleLogin();
};
</script>

<template>
  <div>
    <div class="login">
      <div class="container">
        <div class="content">
          <h4>{{ title }}登录</h4>
          <div class="account-area">
            <el-form ref="formRef" :model="form" label-width="auto" :hide-required-asterisk="true">
              <el-form-item label="账号:" prop="account" :rules="[{ required: true, message: '请输入账号', trigger: 'blur' }]">
                <el-input
                  v-model="form.account"
                  placeholder="请输入账号"
                  ref="accountInputRef"
                  @keyup.enter="handleAccountEnter"
                ></el-input>
              </el-form-item>
              <el-form-item label="密码:" prop="password" :rules="[{ required: true, message: '请输入密码', trigger: 'blur' }]">
                <el-input
                  v-model="form.password"
                  placeholder="请输入密码"
                  show-password
                  ref="passwordInputRef"
                  @keyup.enter="handlePasswordEnter"
                ></el-input>
              </el-form-item>
            </el-form>
            <div class="operation-area">
              <el-button class="submit-btn" type="primary" @click="handleLogin">登录</el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style lang="scss" scoped>
%bg {
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
}


.login {
  // min-height: 675px;
  // height: 100vh;
  // width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: url(../../../public/images/login.jpeg) ;
  // @extend %bg;
  background-size: 100% 100%;
  background-position: center;
  background-attachment: fixed; /* 背景图片固定 */
  height: 100vh; /* 至少为视口高度 */
  overflow: hidden;
}

.container {
  width: 350px;
}

.content {
  color: #fff;
  width: 300px;
  background-color: #fff;
  border: 1px solid #eee;
  padding: 20px;
  border-radius: 10px;
  box-shadow: -15px 15px 15px rgba(6, 17, 47, 0.7);
  background: linear-gradient(230deg,
      rgba(53, 57, 74, 0) 0%,
      rgb(0, 0, 0) 100%);
  h4{
    // color: #000;
    text-align: center;
  }
}

.account-area {
  margin-top: 50px;
}

:deep(.el-form-item__label) {
  color: #fff;
}

.qr-code-area {
  width: 200px;
  height: 200px;
  margin: 40px auto 20px auto;

  img {
    width: 100%;
    height: 100%;
  }
}

.forget-area {
  float: right;
  font-size: 13px;
  line-height: 16px;
  color: #267EF0;
}

.submit-btn {
  margin: 50px 0 30px 0;
  width: 100%;
}

.change-login-type {
  display: flex;
  border-top: 1px solid #eee;
  justify-content: space-between;
  padding-top: 10px;

  .login-type {
    font-size: 12px;
    line-height: 16px;
    text-align: center;
    color: #fff;
    width: 25%;
    position: relative;

    &:hover {
      cursor: pointer;
      color: #267EF0;
    }
  }

  .active {
    color: #267EF0;
  }

}
</style>

