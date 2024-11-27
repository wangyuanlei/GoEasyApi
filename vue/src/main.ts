import { createApp } from 'vue'
// Element-plus组件库
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import '@/style/style.css'
import '@/assets/icon/iconfont.css'
import SvgIcon from '@/components/svg/svg.vue';
import 'virtual:svg-icons-register';
// 路由配置
import router from './router'
// pinia
import pinia from './stores/index'
import App from './App.vue'
const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.use(pinia)
app.component('svg-icon',SvgIcon)
// app.component('svg-icon',svg)
app.mount('#app')