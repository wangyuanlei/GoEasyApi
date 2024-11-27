// 路由配置
import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
 
// 首页
import HomeView from '@/view/Home/index.vue'
import LoginView from '@/view/Login/index.vue'
 
 
const routes: Array<RouteRecordRaw> = [
  {
    path:'/',
    name:'LoginView',
    component:LoginView,
  },
  {
    path:'/Home',
    name:'Home',
    component:HomeView,
    redirect:'/Home/DataSource',
    children:[
        {
            path:'/Home/DataSource',
            name:'DataSource',
            component:()=>import('@/view/InfoView/DataSource.vue')
        },
        {
          path:'/Home/UserList',
          name:'UserList',
          component:()=>import('@/view/InfoView/UserList.vue')
       },
       {
         path:'/Home/SystemSetting',
         name:'SystemSetting',
         component:()=>import('@/view/InfoView/SystemSetting.vue')
       },
       {
        path:'/Home/SetApi',
        name:'SetApi',
        component:()=>import('@/view/InfoView/SetApi.vue')
      },
    ]
  },
]
 
const router = createRouter({
  history:createWebHistory(),
  routes
})
 
// 路由守卫
router.beforeEach((to,from,next) => {
  // next()
  const hasToken = localStorage.getItem('accessToken');
  if (!hasToken && to.name !== 'LoginView') {
    next({ name: 'LoginView' });
  } else {
    next();
  }
})
router.afterEach((to, from) => {
  document.title = 'XXX系统';
})
 
export default router