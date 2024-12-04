// 配置axios
import router from '@/router';
import axios from 'axios'
import type { InternalAxiosRequestConfig, AxiosResponse } from "axios";
import { ElMessage } from "element-plus";
// 创建axios实例

const service = axios.create({
	// baseURL: 'http://127.0.0.1:8008',
	baseURL: import.meta.env.VITE_APP_BASE_API,
	timeout: 500000,
	headers: { "Content-Type": "application/json;charset=utf-8" },
});

// 请求拦截器
service.interceptors.request.use(
	(config: InternalAxiosRequestConfig) => {
		const accessToken = localStorage.getItem('accessToken');
		
		if (accessToken) {
			config.headers.Authorization = accessToken;
		}
		return config;
	},
	(error: any) => {
		return Promise.reject(error);
	}
);
// 响应拦截器
service.interceptors.response.use(
	(response: AxiosResponse) => {
		// 检查配置的响应类型是否为二进制类型（'blob' 或 'arraybuffer'）, 如果是，直接返回响应对象
		if (response.config.responseType === "blob" || response.config.responseType === "arraybuffer") {
			return response;
		}

		
		if (response.data.code*1==200||response.data.code*1==10000) {
			return response.data;
		}
		if (response.data.code=='401') {
			// const accessToken = <string>localStorage.getItem(TOKEN_KEY);
			// 清空存储的token值 
			localStorage.clear()
			// 调用windows刷新页面方法
			location.reload();
			return
		}
		if (response.data.code=='501'|| response.data.code=='502') {
			router.push('/')
		}
		ElMessage.error(response.data.message);
		return Promise.reject(new Error(response.data || "Error"));
	},
	(error: any) => {
		// 异常处理
		if (error.response.data) {
			const { Msg } = error.response.data;
			ElMessage.error(Msg);
		}
		return Promise.reject(error.message);
	}
);
 
 
export default service;