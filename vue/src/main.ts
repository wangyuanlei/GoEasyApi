import { createApp } from 'vue'
// Element-plus组件库
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import '@/style/style.css'
import '@/assets/icon/iconfont.css'
// 路由配置
import router from './router'
// pinia
import pinia from './stores/index'
import App from './App.vue'
const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.use(pinia)
app.mount('#app')