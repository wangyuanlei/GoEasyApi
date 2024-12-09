<template>
    <div class="basic-information-wrapper">
        <div class="basic-information-title">{{ titleName }}</div>
        <el-scrollbar height="calc(90vh)">
            <el-form :model="form" label-width="auto" v-loading="loading">
                <el-col :span="10">
                    <el-form-item label="名称:" :label-width="formLabelWidth">
                        <el-input v-model="form.Name"/>
                    </el-form-item>
                </el-col>
                <el-col :span="10">
                    <el-form-item label="路径:" :label-width="formLabelWidth">
                        <el-input v-model="form.Path"/>
                    </el-form-item>
                </el-col>
                <el-col :span="10">
                    <el-form-item label="描述:" :label-width="formLabelWidth">
                        <el-input v-model="form.Description"/>
                    </el-form-item>
                </el-col>
                <el-col :span="10">
                    <el-form-item label="请求方式:" :label-width="formLabelWidth">
                        <el-radio-group v-model="form.Method">
                            <el-radio value="get">get</el-radio>
                            <el-radio value="post">post</el-radio>
                        </el-radio-group>
                    </el-form-item>
                </el-col>
                <el-row :gutter="0">
                    <el-col :span="6">
                        <el-form-item label="是否开启缓存:" :label-width="formLabelWidth">
                            <el-radio-group v-model="form.CacheEnabled">
                            <el-radio :value="1">是</el-radio>
                            <el-radio :value="2">否</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="缓存时间:" v-if="form.CacheEnabled === 1">
                            <el-input v-model="form.CacheTime" placeholder="请输入缓存时间（秒）" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="0">
                    <el-col :span="6">
                        <el-form-item label="是否开启请求次数限制:" :label-width="formLabelWidth">
                            <el-radio-group v-model="form.RateLimitEnabled">
                            <el-radio :value="1">是</el-radio>
                            <el-radio :value="2">否</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="请求次数:" v-if="form.RateLimitEnabled === 1">
                            <el-input v-model="form.RateLimitCount" placeholder="请输入请求次数" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="秒数:" v-if="form.RateLimitEnabled === 1">
                            <el-input v-model="form.RateLimitTime" placeholder="请输入秒数" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-col :span="10">
                    <el-form-item label="是否开启token验证:" :label-width="formLabelWidth">
                        <el-radio-group v-model="form.TokenValidationEnabled">
                            <el-radio :value="1">是</el-radio>
                            <el-radio :value="2">否</el-radio>
                        </el-radio-group>
                    </el-form-item>
                </el-col>
                <el-row :gutter="0">
                    <el-col :span="10">
                        <el-form-item label="返回数据类型:" :label-width="formLabelWidth">
                            <el-select v-model="form.ReturnType" placeholder="请选择返回数据类型" @change="form.ReturnValMode=''">
                                <el-option label="insert" value="insert" />
                                <el-option label="update" value="update" />
                                <el-option label="delete" value="delete" />
                                <el-option label="row" value="row" />
                                <el-option label="list" value="list" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="10">
                        <!-- <span>{{ form.ReturnType }}</span> -->
                        <el-form-item label="请选择" label-width="165px" v-if="form.ReturnType === 'insert'">
                            <el-select v-model="form.ReturnValMode" placeholder="请选择">
                                <el-option label="last_id" value="last_id" />
                                <el-option label="uuid" value="uuid" />
                            </el-select>
                        </el-form-item>
                        <el-form-item label="请选择" label-width="165px" v-if="form.ReturnType === 'update'">
                            <el-select v-model="form.ReturnValMode" placeholder="请选择">
                                <el-option label="row" value="row" />
                                <el-option label="bool" value="bool" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-col :span="22">
                    <el-form-item :label-width="formLabelWidth" label="请求参数:">
                        <el-table :data="form.Params" style="width: 100%;" :header-cell-style="{background: '#f5f7fa',color: '#909399'}" empty-text="添加参数">
                            <el-table-column prop="Name" label="参数名称" width="130" header-align="center" align="center">
                                <template #default="scope">
                                    <el-input v-if="scope.row.editing" v-model="scope.row.Name" />
                                    <span v-else>{{ scope.row.Name }}</span>
                                </template>
                            </el-table-column>
                            <el-table-column prop="Type" label="参数类型" width="120" header-align="center" align="center">
                                <template #default="scope">
                                    <!-- <el-input v-if="scope.row.editing" v-model="scope.row.Type" /> -->
                                    <el-select v-model="scope.row.Type"  v-if="scope.row.editing" placeholder="参数类型">
                                        <el-option label="string" value="string" />
                                        <el-option label="int" value="int" />
                                        <el-option label="float" value="float" />
                                        <el-option label="bool" value="bool" />
                                        <el-option label="date" value="date" />
                                        <el-option label="datetime" value="datetime" />
                                    </el-select>
                                    <span v-else>{{ scope.row.Type }}</span>
                                </template>
                            </el-table-column>
                            <el-table-column prop="Description" label="参数描述" header-align="center" align="center">
                                <template #default="scope">
                                    <el-input v-if="scope.row.editing" v-model="scope.row.Description" />
                                    <span v-else>{{ scope.row.Description }}</span>
                                </template>
                            </el-table-column>
                            <el-table-column prop="Required" label="是否必传" width="120" align="center" header-align="center">
                                <template #default="scope">
                                    <el-select v-if="scope.row.editing" v-model="scope.row.Required" placeholder="请选择">
                                        <el-option label="是" :value="1"></el-option>
                                        <el-option label="否" :value="2"></el-option>
                                    </el-select>
                                    <span v-else>
                                        {{ scope.row.Required ===1 ? '是' : '否' }}
                                    </span>
                                </template>
                            </el-table-column>
                            <el-table-column prop="Default" label="默认值" show-overflow-tooltip width="230" align="center" header-align="center">
                                <template #default="scope">
                                    <el-input v-if="scope.row.editing" v-model="scope.row.Default" />
                                    <span v-else>{{ scope.row.Default }}</span>
                                </template>
                            </el-table-column>
                            <el-table-column prop="Example" label="实例值" width="120" header-align="center" align="center">
                                <template #default="scope">
                                    <el-input v-if="scope.row.editing" v-model="scope.row.Example" />
                                    <span v-else>{{ scope.row.Example }}</span>
                                </template>
                            </el-table-column>
                            <el-table-column prop="Regex" label="正则表达式" header-align="center" align="center">
                                <template #default="scope">
                                    <el-input v-if="scope.row.editing" v-model="scope.row.Regex" />
                                    <span v-else>{{ scope.row.Regex }}</span>
                                </template>
                            </el-table-column>
                            <el-table-column label="操作" width="150" align="center" fixed="right" min-width="120" >
                                <template #default="scope">
                                    <el-button v-if="!scope.row.editing && !scope.row.newRow" type="primary" :icon="Edit" @click="editRow(scope.$index)" circle>
                                    </el-button>
                                    <el-button v-if="scope.row.editing || scope.row.newRow" type="success" :icon="Select" @click="saveRow(scope.$index)" circle>
                                    </el-button>
                                    <el-button  type="danger" :icon="Delete" @click="DeltRow(scope.$index)" circle>
                                    </el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                        <el-button type="primary" class="add-btn" :icon="Plus" circle @click="addRequestParam" />
                    </el-form-item>
                </el-col>  
                <el-col :span="15">
                    <el-form-item label="语言" :label-width="formLabelWidth">
                        <div id="editor" style="height: 300px; width: 100%;"></div>
                        <!-- <el-input type="textarea" v-model="form.SqlContent" :rows="6" placeholder="请输入sql语句"/> -->
                    </el-form-item>
                </el-col>  
            </el-form>
            <div class="btn-box">
                <el-button type="default" class="back-btn" @click="back" size="large">返回</el-button>
                <el-button type="primary" class="save-btn" @click="save" size="large">保存</el-button>          
            </div>
        </el-scrollbar>
    </div>
</template>

<script lang="ts" setup>
import { ref, watch,onMounted } from 'vue';
import { Plus,Edit,Select,Delete } from '@element-plus/icons-vue';
import { ElMessage, ElMessageBox } from 'element-plus'
import SetApi from '@/api/setApi';
const hasToken = <string>localStorage.getItem('accessToken');
const formLabelWidth = '220px';
import { defineProps } from 'vue';



import ace from 'ace-builds';


import { defineEmits } from 'vue';

const emit = defineEmits(['switchToActuator']);
const back= ()=>{
    // 返回上一页
    emit('switchToActuator', null,'apiList');
    // 清空数据
    resetForm();
}
const loading = ref(false);

const titleName = ref('');

    onMounted(() => {
        const editor = ace.edit('editor');
      editor.setTheme('ace/theme/dracula'); // 黑色主题
      editor.session.setMode('ace/mode/sql'); // SQL 模式
      editor.setOptions({
        enableBasicAutocompletion: true,  // 启用基础自动补全
        enableLiveAutocompletion: true,   // 启用实时补全
      });
    });
    const editor = ref<any>(); // 用 ref 保存编辑器实例
    // 监听 editor 实例的变化
    watch(() => editor.value, (newEditor) => {
      if (newEditor) {
        const cmInstance = newEditor.cminstance;  // 获取 CodeMirror 实例
        cmInstance.on("inputRead", () => {
          cmInstance.showHint();  // 显示代码提示
        });
      }
    });
    // }

const props = defineProps({
    data: Object
});

const fetchDetailData = (data: any) => {
    loading.value = true;
    SetApi.getInterfaceInfo(hasToken, data).then(res => {
        loading.value = false;
        // console.log('res=', res);
        form.value = res.data;
    });
}
const id = ref<any>('');




// 定义请求参数和参数检验的类型
interface RequestParam {
    Default: string;
    Description: string;
    Example: string;
    Name: string;
    Regex: string;
    Required: string;
    Type: string;
    editing: boolean;
    newRow: boolean;
}

interface formType {
    Name: string;
    Path: string;
    Description: string;
    Method: string;
    CacheEnabled: number;
    RateLimitEnabled: number;
    RateLimitCount: number | string;
    CacheTime: string | number;
    RateLimitTime: string | number;
    TokenValidationEnabled: number;
    ReturnValMode: string;
    ReturnType: string;
    SqlContent: string;
    Params: RequestParam[];
}
const form = ref<formType>({
    Name: '',
    Path: '',
    Description: '',
    Method: 'post',
    CacheEnabled: 2,
    RateLimitEnabled: 2,
    RateLimitCount: '',
    CacheTime: '',
    RateLimitTime: '',
    TokenValidationEnabled: 1,
    ReturnValMode:'',
    ReturnType:'',
    SqlContent:'',
    Params: [] as RequestParam[],
});
const resetForm = () => {
    form.value = {
        Name: '',
        Path: '',
        Description: '',
        Method: 'post',
        CacheEnabled: 2,
        RateLimitEnabled: 2,
        RateLimitCount: '',
        CacheTime: '',
        RateLimitTime: '',
        TokenValidationEnabled: 1,
        ReturnValMode: '',
        ReturnType: '',
        SqlContent: '',
        Params: [] as RequestParam[],
    };
}

const addRequestParam = () => {
    // verifyParams();
    if(!verifyParams()){
        return;
    }
    form.value.Params.push({
        Name: '',
        Type: 'string',
        Description: '',
        Required: '',
        Default: '',
        Example: '',
        Regex: '',
        editing: true,
        newRow: true
    });
}

const editRow = (index: number) => {
    // verifyParams();
    if(!verifyParams()){
        return;
    }
    form.value.Params[index].editing = true;
}

const saveRow = (index: number) => {
    form.value.Params[index].editing = false;
    form.value.Params[index].newRow = false;
}
const DeltRow=(index: number) =>{
    ElMessageBox.confirm("确定删除该行数据吗？", {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    form.value.Params.splice(index, 1);
  })
};
function toLowerCaseCamelCase(obj: any): any {
    if (Array.isArray(obj)) {
        return obj.map(item => toLowerCaseCamelCase(item));
    }
    if (typeof obj === 'object' && obj !== null) {
        const newObj: any = {};
        for (const key in obj) {
            if (obj.hasOwnProperty(key)) {
                const newKey = key.split(/(?=[A-Z])/).join('').toLowerCase();
                newObj[newKey] = toLowerCaseCamelCase(obj[key]);
            }
        }
        return newObj;
    }
    return obj;
}

// 验证参数
const verifyParams =()=>{
    const hasUnsavedRows = form.value.Params.some(param => param.editing || param.newRow);
    if (hasUnsavedRows) {
        ElMessage({
            type: 'error',
            message: '有参数行未保存',
        });
        return false;
    }else{
        return true;
    }
}
const save = () => {
    // console.log(id.value);

    // 验证参数
    if (!verifyParams()) {
        return;
    }
    const formData = toLowerCaseCamelCase(form.value);

    // 移除不需要参数字段
    if (formData.params && formData.params.length > 0) {
        formData.params = formData.params.map((item: any) => {
            const { interfaceid, editing, newrow, ...rest } = item;
            return rest;
        });
    }
    // if (formData.Name) {
    //     formData.interfacename=formData.name;
    //     delete formData.name;
    // }
    if (formData.cacheenabled === 2) {
        delete formData.cachetime;
    }

    if (formData.ratelimittime) {
        formData.ratelimittime = Number(formData.ratelimittime);
    }
    if (formData.ratelimitcount) {
        formData.ratelimitcount = Number(formData.ratelimitcount);
    }
    if (formData.cachetime) {
        formData.cachetime=Number(formData.cachetime);
    }
    if (formData.ratelimitenabled === 2) {
        delete formData.ratelimitcount;
        delete formData.ratelimittime;
    }

    if (formData.returntype !== 'insert' && formData.returntype !== 'update' || formData.returntype === '') {
        delete formData.returnvalmode;
    }

    if (id.value) {
        // 编辑
        loading.value = true;
        SetApi.update(hasToken, formData).then(_ => {
            loading.value = false;
            ElMessage({
                type: 'success',
                message: '编辑成功',
            });
        });
    } else {
        // 新增
        // formData.interfacename=formData.name;
        // delete formData.name;

        loading.value = true;
        SetApi.addInterface(hasToken, formData).then(_ => {
            loading.value = false;
            ElMessage({
                type: 'success',
                message: '新增成功',
            });
        });
    }
}
// 监听 props.data 的变化
watch(() => props.data, (newData) => {
    // console.log('newData', newData);
    if (newData) {
        titleName.value = '接口详情';
        id.value = newData; // 修改这里
        // 调用接口获取详情数据
        fetchDetailData(newData);
    }else {
        // 清空数据
        titleName.value = '新增接口';
        resetForm();
    }
}, { immediate: true });
</script>

<style scoped lang="scss">
.basic-information-wrapper {
    width: 100%;
    height: 100%;
    overflow: hidden;
    display: flex;
    justify-content: center;
    flex-direction: column;
}
.basic-information-title{
    font-size: 18px;
    font-weight: bold;
    margin-bottom: 10px;
    color: #333;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;

}
.btn-box{
    margin: 0 135px 60px 0;
    display: flex;;
    justify-content: end;

}
.add-btn{
    margin-top: 20px;
}
</style>