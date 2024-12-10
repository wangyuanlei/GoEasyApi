import{s as b,d as ne,r as d,w as re,o as se,c as de,i as B,j as u,l as ie,u as x,e as l,F as pe,f as i,h as y,m as ce,k as F,b as I,n as me,p as Ee,q as ge,v as $,x as K,E as m,_ as fe}from"./index-IYHzICXp.js";class C{static getUesrList(r,o){let t={page:o.page,page_size:o.page_size,name:o.search};return b({url:"/manger/get_user_list",method:"get",params:t,headers:{token:r}})}static createUser(r,o){const t=new FormData;return t.append("name",o.name),t.append("deptId",o.deptId),t.append("account",o.account),t.append("password",o.password),b({url:"/api/user/register",method:"post",data:t,headers:{token:r}})}static getUser(r,o){return b({url:"/manger/get_user_info",method:"get",params:{user_id:o},headers:{token:r}})}static updateUser(r,o,t){return t.user_id=o,t.is_valid=Number(t.isValid),t.dept_id=t.deptId,delete t.isValid,delete t.deptId,b({url:"/manger/set_user_info",method:"post",data:t,headers:{token:r}})}static deleteUser(r,o){const t=new FormData;return t.append("user_id",o),b({url:"/manger/delete_user",method:"post",data:t,headers:{token:r}})}static updtPsd(r,o,t){const v=new FormData;return v.append("user_id",o),v.append("password",t),console.log("formData",v),b({url:"/manger/set_user_pass",method:"post",data:v,headers:{token:r}})}}var Fe={name:"zh-cn",el:{breadcrumb:{label:"面包屑"},colorpicker:{confirm:"确定",clear:"清空",defaultLabel:"颜色选择器",description:"当前颜色 {color}，按 Enter 键选择新颜色",alphaLabel:"选择透明度的值"},datepicker:{now:"此刻",today:"今天",cancel:"取消",clear:"清空",confirm:"确定",dateTablePrompt:"使用方向键与 Enter 键可选择日期",monthTablePrompt:"使用方向键与 Enter 键可选择月份",yearTablePrompt:"使用方向键与 Enter 键可选择年份",selectedDate:"已选日期",selectDate:"选择日期",selectTime:"选择时间",startDate:"开始日期",startTime:"开始时间",endDate:"结束日期",endTime:"结束时间",prevYear:"前一年",nextYear:"后一年",prevMonth:"上个月",nextMonth:"下个月",year:"年",month1:"1 月",month2:"2 月",month3:"3 月",month4:"4 月",month5:"5 月",month6:"6 月",month7:"7 月",month8:"8 月",month9:"9 月",month10:"10 月",month11:"11 月",month12:"12 月",weeks:{sun:"日",mon:"一",tue:"二",wed:"三",thu:"四",fri:"五",sat:"六"},weeksFull:{sun:"星期日",mon:"星期一",tue:"星期二",wed:"星期三",thu:"星期四",fri:"星期五",sat:"星期六"},months:{jan:"一月",feb:"二月",mar:"三月",apr:"四月",may:"五月",jun:"六月",jul:"七月",aug:"八月",sep:"九月",oct:"十月",nov:"十一月",dec:"十二月"}},inputNumber:{decrease:"减少数值",increase:"增加数值"},select:{loading:"加载中",noMatch:"无匹配数据",noData:"无数据",placeholder:"请选择"},dropdown:{toggleDropdown:"切换下拉选项"},mention:{loading:"加载中"},cascader:{noMatch:"无匹配数据",loading:"加载中",placeholder:"请选择",noData:"暂无数据"},pagination:{goto:"前往",pagesize:"条/页",total:"共 {total} 条",pageClassifier:"页",page:"页",prev:"上一页",next:"下一页",currentPage:"第 {pager} 页",prevPages:"向前 {pager} 页",nextPages:"向后 {pager} 页",deprecationWarning:"你使用了一些已被废弃的用法，请参考 el-pagination 的官方文档"},dialog:{close:"关闭此对话框"},drawer:{close:"关闭此对话框"},messagebox:{title:"提示",confirm:"确定",cancel:"取消",error:"输入的数据不合法!",close:"关闭此对话框"},upload:{deleteTip:"按 delete 键可删除",delete:"删除",preview:"查看图片",continue:"继续上传"},slider:{defaultLabel:"滑块介于 {min} 至 {max}",defaultRangeStartLabel:"选择起始值",defaultRangeEndLabel:"选择结束值"},table:{emptyText:"暂无数据",confirmFilter:"筛选",resetFilter:"重置",clearFilter:"全部",sumText:"合计"},tour:{next:"下一步",previous:"上一步",finish:"结束导览"},tree:{emptyText:"暂无数据"},transfer:{noMatch:"无匹配数据",noData:"无数据",titles:["列表 1","列表 2"],filterPlaceholder:"请输入搜索内容",noCheckedFormat:"共 {total} 项",hasCheckedFormat:"已选 {checked}/{total} 项"},image:{error:"加载失败"},pageHeader:{title:"返回"},popconfirm:{confirmButtonText:"确定",cancelButtonText:"取消"},carousel:{leftArrow:"上一张幻灯片",rightArrow:"下一张幻灯片",indicator:"幻灯片切换至索引 {index}"}}};const ve={class:"search-user-table"},he={class:"search-user-top"},_e={class:"search-left"},De={class:"search-right"},be={class:"search-user-bottom"},Be={class:"dialog-footer"},A="100px",Ce=ne({__name:"UserList",setup(W){const r=d(""),o=d(null),t=()=>{w.value.search=r.value,c()},v=()=>{w.value.search="",c()},E=d(!1),P=d(!1),h=localStorage.getItem("accessToken"),q=d([]),_=d(1),D=d(15),L=d(0),Y=n=>{D.value=n},H=n=>{_.value=n},w=d({page:_.value.toString(),page_size:D.value.toString(),search:r.value});re([_,D],()=>{w.value.page=_.value.toString(),w.value.page_size=D.value.toString(),c()});const c=()=>{P.value=!0,C.getUesrList(h,w.value).then(n=>{q.value=n.data.list,L.value=n.data.total}).catch(n=>{console.log("err",n),P.value=!1})};se(()=>{c()});const T=d("注册用户"),U=d(),s=d({name:"",account:"",password:"",deptId:"",isValid:2}),M=()=>{s.value={name:"",password:"",account:"",deptId:"",isValid:2},U.value&&U.value.clearValidate(),T.value="用户注册",k.value=!1},G={name:[{required:!0,message:"请输入姓名",trigger:"blur"}],deptId:[{required:!0,message:"请输入部门id",trigger:"blur"}],account:[{required:!0,message:"请输入登陆账号",trigger:"blur"}],password:[{required:!0,message:"请输入密码",trigger:"blur"}]},J=n=>{K.prompt("请输入新密码","修改密码",{confirmButtonText:"确定",cancelButtonText:"取消",inputPattern:/^.{6,20}$/,inputErrorMessage:"密码长度6-20位"}).then(({value:e})=>{C.updtPsd(h,n.UserId.toString(),e).then(()=>{m.success("密码修改成功")})})},O=()=>{U.value.validate(n=>{n?k.value&&z.value?C.updateUser(h,z.value,s.value).then(()=>{m.success("修改成功"),E.value=!1,c()}).catch(e=>{console.error(e),m.error("修改失败")}):C.createUser(h,s.value).then(()=>{E.value=!1,m.success("用户添加成功"),M(),c()}):m.error("请填写完整的用户信息")})},k=d(!1),z=d(null),Q=n=>{E.value=!0,k.value=!0,T.value="编辑用户",z.value=n.UserId.toString();const e=n.UserId.toString();C.getUser(h,e).then(p=>{s.value={name:p.data.Name,account:p.data.Account,deptId:p.data.DeptId,password:p.data.Password,isValid:p.data.IsValid}}).catch(p=>{console.error(p),m.error("获取用户信息失败")})},X=n=>{K.confirm("此操作将永久删除该用户, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(()=>{C.deleteUser(h,n.UserId.toString()).then(()=>{m.success("删除成功"),c()}).catch(()=>{m.info("已取消删除")})})};return(n,e)=>{const p=i("el-input"),g=i("el-button"),f=i("el-table-column"),j=i("el-tag"),S=i("el-tooltip"),Z=i("el-button-group"),ee=i("el-table"),ue=i("el-pagination"),te=i("el-config-provider"),V=i("el-form-item"),N=i("el-radio"),ae=i("el-radio-group"),le=i("el-form"),oe=i("el-dialog");return y(),de(pe,null,[B("div",ve,[B("div",he,[B("div",_e,[u(p,{modelValue:r.value,"onUpdate:modelValue":e[0]||(e[0]=a=>r.value=a),style:{width:"240px"},ref_key:"inputRef",ref:o,size:"large",placeholder:"用户姓名 / 登陆账号",clearable:"",onBlur:t,onClear:v,onKeyup:ie(t,["enter"])},null,8,["modelValue"]),u(g,{type:"primary",onClick:c,icon:x(ce),circle:""},null,8,["icon"])]),B("div",De,[u(g,{type:"primary",onClick:e[1]||(e[1]=a=>E.value=!0),size:"large"},{default:l(()=>e[11]||(e[11]=[F("+ 添加新用户")])),_:1})])]),B("div",be,[u(ee,{data:q.value,style:{width:"100%",height:"98%"},border:"","header-cell-style":{background:"#f5f7fa",color:"#909399"},stripe:"","empty-text":"没有数据"},{default:l(()=>[u(f,{type:"index",label:"序号",width:"60",align:"left","header-align":"center"}),u(f,{prop:"Name",label:"用户姓名","header-align":"center"}),u(f,{prop:"IsValid",label:"状态","header-align":"center",align:"center"},{default:l(a=>[a.row.IsValid==2?(y(),I(j,{key:0,type:"success"},{default:l(()=>e[12]||(e[12]=[F("正常")])),_:1})):(y(),I(j,{key:1,type:"danger"},{default:l(()=>e[13]||(e[13]=[F("禁用")])),_:1}))]),_:1}),u(f,{prop:"Account",label:"登录账号","header-align":"center"}),u(f,{prop:"DeptId",label:"部门","show-overflow-tooltip":"",width:"230","header-align":"center"}),u(f,{prop:"RegisterTime",label:"注册时间","show-overflow-tooltip":"",width:"230","header-align":"center"}),u(f,{label:"操作","header-align":"center",align:"center"},{default:l(a=>[u(Z,null,{default:l(()=>[u(S,{class:"box-item",effect:"dark",content:"编辑",placement:"top"},{default:l(()=>[u(g,{type:"primary",icon:x(me),size:"small",onClick:R=>Q(a.row)},null,8,["icon","onClick"])]),_:2},1024),u(S,{class:"box-item",effect:"dark",content:"修改密码",placement:"top"},{default:l(()=>[u(g,{type:"warning",icon:x(Ee),size:"small",style:{"border-radius":"0"},onClick:R=>J(a.row)},null,8,["icon","onClick"])]),_:2},1024),u(S,{class:"box-item",effect:"dark",content:"删除",placement:"top"},{default:l(()=>[u(g,{type:"danger",icon:x(ge),size:"small",onClick:R=>X(a.row)},null,8,["icon","onClick"])]),_:2},1024)]),_:2},1024)]),_:1})]),_:1},8,["data"])]),u(te,{locale:x(Fe)},{default:l(()=>[u(ue,{background:"",style:{"text-align":"center","margin-top":"30px",height:"20px"},"current-page":_.value,"onUpdate:currentPage":e[2]||(e[2]=a=>_.value=a),"page-size":D.value,"onUpdate:pageSize":e[3]||(e[3]=a=>D.value=a),"page-sizes":[15,20,30,50],layout:"total, sizes, prev, pager, next, jumper",total:L.value,onSizeChange:Y,onCurrentChange:H},null,8,["current-page","page-size","total"])]),_:1},8,["locale"])]),u(oe,{modelValue:E.value,"onUpdate:modelValue":e[10]||(e[10]=a=>E.value=a),title:T.value,width:"500",onClose:M},{footer:l(()=>[B("div",Be,[u(g,{onClick:e[9]||(e[9]=a=>E.value=!1)},{default:l(()=>e[16]||(e[16]=[F("取消")])),_:1}),u(g,{type:"primary",onClick:O},{default:l(()=>e[17]||(e[17]=[F(" 确定 ")])),_:1})])]),default:l(()=>[u(le,{model:s.value,rules:G,ref_key:"formRef",ref:U},{default:l(()=>[u(V,{label:"姓名",prop:"name","label-width":A,required:""},{default:l(()=>[u(p,{modelValue:s.value.name,"onUpdate:modelValue":e[4]||(e[4]=a=>s.value.name=a),autocomplete:"off"},null,8,["modelValue"])]),_:1}),k.value?$("",!0):(y(),I(V,{key:0,label:"账号名",prop:"account","label-width":A,required:""},{default:l(()=>[u(p,{modelValue:s.value.account,"onUpdate:modelValue":e[5]||(e[5]=a=>s.value.account=a),autocomplete:"off"},null,8,["modelValue"])]),_:1})),k.value?$("",!0):(y(),I(V,{key:1,label:"密码:",prop:"password","label-width":A,required:""},{default:l(()=>[u(p,{modelValue:s.value.password,"onUpdate:modelValue":e[6]||(e[6]=a=>s.value.password=a),autocomplete:"off","show-password":""},null,8,["modelValue"])]),_:1})),u(V,{label:"部门",prop:"deptId","label-width":A,required:""},{default:l(()=>[u(p,{modelValue:s.value.deptId,"onUpdate:modelValue":e[7]||(e[7]=a=>s.value.deptId=a),autocomplete:"off"},null,8,["modelValue"])]),_:1}),u(V,{label:"状态",prop:"isValid","label-width":A,required:""},{default:l(()=>[u(ae,{modelValue:s.value.isValid,"onUpdate:modelValue":e[8]||(e[8]=a=>s.value.isValid=a)},{default:l(()=>[u(N,{label:2},{default:l(()=>e[14]||(e[14]=[F("开启")])),_:1}),u(N,{label:1},{default:l(()=>e[15]||(e[15]=[F("禁用")])),_:1})]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["modelValue","title"])],64)}}}),ke=fe(Ce,[["__scopeId","data-v-dbf17f1b"]]);export{ke as default};
