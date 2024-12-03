<template>
    <div class="overall-plugin-table">
      <div class="overall-plugin-top">
        <div class="search-left">
            <el-form-item label="名单类型:">
                <el-radio-group v-model="selectedListType">
                    <el-radio :label="0">无</el-radio>
                    <el-radio :label="1">黑名单</el-radio>
                    <el-radio :label="2">白名单</el-radio> 
                </el-radio-group>
            </el-form-item>
        </div>
        <div class="search-right">
          <el-button type="primary" @click="addToList('black')" v-if="selectedListType === 1" size="large">+ 添加黑名单</el-button>
          <el-button type="primary" @click="addToList('white')" v-if="selectedListType === 2" size="large">+ 添加白名单</el-button>
        </div>
      </div>
      <div class="overall-plugin-bottom">
        <el-table :data="tableData" style="width: 100%; height: 100%;" v-loading="loading" border stripe :header-cell-style="{background: '#f5f7fa',color: '#909399'}" empty-text="没有数据">
          <el-table-column type="index" label="序号" min-width="30" align="center" header-align="center"></el-table-column>
          <el-table-column prop="IP" label="Ip" header-align="center"></el-table-column>
          <el-table-column prop="Description" label="描述" show-overflow-tooltip header-align="center"></el-table-column>
          <el-table-column label="操作" header-align="center" align="center">
            <template #default="scope">
            <el-button-group>
              <el-tooltip class="box-item" effect="dark" content="删除" placement="top">
                <el-button type="danger" :icon="Delete" size="small" @click="deleteItem(scope.row)" />
              </el-tooltip>
            </el-button-group>
            </template>
          </el-table-column>
        </el-table>
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
import { ref, watch ,onMounted} from 'vue'
import { ElMessageBox,ElMessage } from 'element-plus'
import { Delete} from '@element-plus/icons-vue';
import  sysSet  from '@/api/sysSet'

const loading = ref(false);
const hasToken = <string>localStorage.getItem('accessToken');
const tableData = ref([]);


const selectedListType = ref<number>(0) // 默认选中黑名单
const dialogTitle = ref('');

const formLabelWidth = '110px';

// 弹窗开启关闭
const dialogFormVisible = ref(false);

// 监听选中项的变化
watch(selectedListType, (newVal) => {
  getListData(newVal)
});
//加载列表接口
const getListData = (type: number) => {
  loading.value = true; // 开始加载

  sysSet.SetType(hasToken,type.toString()).then(res => {
    }).catch(err => {
      console.log('err', err);
    })
   // 开始加载
  if (type === 1) {
    sysSet.getBlackList(hasToken).then(res => {
        loading.value = false;
    // console.log('res', res.data);
        tableData.value = res.data;
  })
//   .finally(() => {
//     loading.value = false; // 加载完成
//   });
  }else if (type === 2) {
    sysSet.getWhiteList(hasToken).then(res => {
        loading.value = false;
        tableData.value = res.data;

  })
//   .finally(() => {
//     loading.value = false; // 加载完成
//   });
  }else{
    loading.value = false;
    tableData.value = [];
  }
};
onMounted(()=>{
    sysSet.getType(hasToken).then(res => {
          selectedListType.value=res.data;
            if(res.data==1){
                selectedListType.value=1;
            }else if(res.data==2){
                selectedListType.value=2;
            }
            getListData(res.data)
        }).catch(err => {
          console.log('err', err);
          ElMessage({
            message: '初始值获取失败',
            type: 'error'
          })
        })
    // getTableData();
});

const deleteItem = (item: any) => {
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

const formRef = ref();
//弹窗逻辑
const form = ref({
    ip: '',
    description: '',
  });
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
  })
//   .finally(() => {
//     loading.value = false; // 加载完成
//   });
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
  })
//   .finally(() => {
//     loading.value = false; // 加载完成
//   });
}
// 添加到列表
const addToList = (type: string) => {
dialogFormVisible.value = true;
dialogTitle.value = type === 'black' ? '添加黑名单' : '添加白名单';
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
  /* margin-bottom: 10px; */
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