<template>
    <!-- 系统设置 -->
    <div class="w-full h-full bg-gray-100">
        <div class="w-full h-full pt-10 sys-box">
            <section class="center-box rounded m-auto p-5">
                <div class="head-title text-2xl text-center pb-6 font-light">系统设置</div>
                <div class="dashed px-8"></div>
                <div class="info-reply-box" v-loading="loading">
                    <el-scrollbar max-height="505">
                        <div class="info-reply">
                            <el-radio-group v-model="selectedListType">
                                <el-radio :label="0">无</el-radio>
                                <el-radio :label="1">黑名单</el-radio>
                                <el-radio :label="2">白名单</el-radio>                     
                            </el-radio-group>
                        </div>
                        <div style="height:90%">
                            <el-card class="w-full mb-4" v-for="(item, index) in defaultData" :key="index">
                            <div class="card-left">
                                <div class="flex justify-between items-center">
                                    <div class="">IP:{{ item.IP }}</div>
                                    <div class="text-base">描述：{{ item.Description }}</div>
                                </div>
                            </div>
                            <div class="card-right ml-3">
                                <!-- <i class="iconfont icon-xiugai text-2xl text-blue-500 cursor-pointer mr-2" @click="editItem(item, index)"></i> -->
                                <i class="iconfont icon-htmal5icon17 text-2xl text-red-500 cursor-pointer" @click="deleteItem(item, index)"></i>
                            </div>
                        </el-card>
                        </div>

                    </el-scrollbar>
                </div>
                <div class="btn-box">
                    <el-button type="primary" @click="addToList('black')" v-if="selectedListType === 1">添加黑名单</el-button>
                    <el-button type="primary" @click="addToList('white')" v-if="selectedListType === 2">添加白名单</el-button>
                </div>
            </section>
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle" width="500" @close="resetForm">
    <el-form :model="form" :rules="rules" ref="formRef">
      <el-form-item label="ip" prop="ip" :label-width="formLabelWidth" required>
        <el-input v-model="form.ip" autocomplete="off" />
      </el-form-item>
      <el-form-item label="description" prop="description" :label-width="formLabelWidth" required>
        <el-input v-model="form.description" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" v-if="selectedListType === 1" @click="blackSave">
          确定
        </el-button>
        <el-button type="primary" v-else-if="selectedListType === 2" @click="whiteSave">
          确定
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>
<script lang="ts" setup>
import { ref, onMounted, watch } from 'vue'
import { ElMessage,ElMessageBox } from 'element-plus'
import  sysSet  from '@/api/sysSet'
const hasToken = <string>localStorage.getItem('accessToken');

const formLabelWidth = '110px';
const dialogFormVisible = ref(false);
const dialogTitle = ref('');
const loading = ref(false);
interface ListData {
  IP: string;
  Description: string;
}
const defaultData = ref<ListData[]>([]);
const selectedListType = ref<number>(0) // 默认选中黑名单

//加载列表接口
const getListData = (type: number) => {
  loading.value = true; // 开始加载
  // getType.value = type.toString();
  sysSet.SetType(hasToken,type.toString()).then(res => {
    }).catch(err => {
      console.log('err', err);
    })
  if (type === 1) {
    sysSet.getBlackList(hasToken).then(res => {
        defaultData.value = res.data.map((item: any) => ({
        IP: item.IP,
        Description: item.Description
      }));
    // console.log('res', res.data);
  }).finally(() => {
    loading.value = false; // 加载完成
  });
  }else if (type === 2) {
    sysSet.getWhiteList(hasToken).then(res => {
    // console.log('res', res.data);
    defaultData.value = res.data.map((item: any) => ({
        IP: item.IP,
        Description: item.Description
      }));
  }).finally(() => {
    loading.value = false; // 加载完成
  });
  }else{
    defaultData.value = [];
  }
};
// 监听选中项的变化
watch(selectedListType, (newVal) => {
  getListData(newVal)
});
//初始化
onMounted(() => {
  sysSet.getType(hasToken).then(res => {
          // getListData(res.data)
          selectedListType.value=res.data;
        }).catch(err => {
          console.log('err', err);
          ElMessage({
            message: '初始值获取失败',
            type: 'error'
          })
        })
});

// 添加到列表
const addToList = (type: string) => {
dialogFormVisible.value = true;
dialogTitle.value = type === 'black' ? '添加黑名单' : '添加白名单';
};
const deleteItem = (item: any, index: number) => {
    if(selectedListType.value===1){
        ElMessageBox.confirm('此操作将永久该黑名单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        }).then(() => {
            sysSet.deleteBlack(hasToken,item.IP).then(res => {
            ElMessage({
            message: '删除成功',
            type: 'success'
            })
            getListData(1)
        }).catch(err => {
          console.log('err', err);
          ElMessage({
            message: '删除失败',
            type: 'error'
          })
        })
        })

    }else{
        ElMessageBox.confirm('此操作将永久删除该白名单, 是否继续?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        }).then(() => {
            sysSet.deleteWhite(hasToken,item.IP).then(res => {
            ElMessage({
            message: '删除成功',
            type: 'success'
            })
            getListData(2)
            }).catch(err => {
            console.log('err', err);
            ElMessage({
                message: '删除失败',
                type: 'error'
            })
            })
        })

    }

};
const editItem = (item: any, index: number) => {
dialogFormVisible.value = true;
form.value.ip=item.IP;
form.value.description=item.Description;
if(selectedListType.value===1){
    dialogTitle.value = '编辑黑名单'

}else{
    dialogTitle.value = '编辑白名单'
}
};

//弹窗逻辑
const form = ref({
    ip: '',
    description: '',
  });
const formRef = ref();
//重置表单
const resetForm = () => {
  form.value = {
    ip: '',
    description: ''

  };
  if (formRef.value) {
    formRef.value.clearValidate();
  }
  dialogTitle.value = '';
};
//添加表单检验逻辑
const rules = {
  ip: [
    { required: true, message: '请输入ip', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入描述', trigger: 'blur' }
  ],
};
const blackSave =()=>{
    sysSet.addBlack(hasToken,form.value).then(res => {
        ElMessage({
          message: '添加成功',
          type: 'success'
        })
        dialogFormVisible.value = false;
        getListData(1);
    }).catch(err => {
      console.log('err', err);
      ElMessage({
        message: '添加失败',
        type: 'error'
      })
  }).finally(() => {
    loading.value = false; // 加载完成
  });
}
const whiteSave =()=>{
    // loading.value = true;
    if(!form.value.ip ||!form.value.description){
        ElMessage({
          message: '请输入完整信息',
          type: 'warning'
        })
        return
    }
    sysSet.addWhite(hasToken,form.value).then(res => {
        ElMessage({
          message: '添加成功',
          type: 'success'
        })
        getListData(2);
        resetForm();
        dialogFormVisible.value = false;
    }).catch(err => {
      console.log('err', err);
      ElMessage({
        message: '添加失败',
        type: 'error'
      })
  }).finally(() => {
    loading.value = false; // 加载完成
  });
}
</script>
<style scoped>
/* .w-full {
    width: 100%;
} */
.h-full {
    height: 100%;
}
.text-center {
    text-align: center;
}
.pt-10 {
    padding-top: 10px;
}
.text-2xl {
    font-size: 24px;
}
.sys-box {
    display: flex;
    justify-content: center;
}
.center-box {
    width: 642px;
    border: 2px dashed var(--el-border-color);
    box-shadow: 0px 12px 32px 4px rgba(0, 0, 0, .04), 0px 8px 20px rgba(0, 0, 0, .08);
    padding: 20px;
}
/* .center-box:hover {
    border-color: #409EFF;
} */
.head-title {
    color: #409EFF;
    padding: 10px 0;
}
.dashed {
    border-top: 1px solid var(--el-border-color);
}
.center-box /deep/ .el-radio-group {
    flex-direction: row;
    align-items: baseline;
}
:deep(.el-card__body) {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
}
.card-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
}
.card-left {
    flex: 1;
}
.card-right {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 50px;
}
.btn-box{
    display: flex;
    justify-content: end;
}
:deep(.el-scrollbar__view) {
    height: 505px;
}
.text-blue-500{
    color: #409EFF;
    cursor: pointer;
}
.text-red-500{
    color: #F56C6C;
    cursor: pointer;
}
.el-card{
    margin-bottom: 10px;
}
</style>