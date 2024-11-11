import request from "@/utils/request";

class AuthAPI {
      static getUesrLogin(data: any) {
		const formData = new FormData();
        console.log('data=',data);
        
		formData.append("account", data.account);
		formData.append("pass", data.pwd);
		return request<any>({
			url: "/api/manger/login",
			method: "post",
			data: formData,
			headers: {
			},
		});
	}
    //测试接口
    static heath() {
		return request<any>({
			url: "/health",
			method: "get",
			headers: {
			},
		});
	}
}
export default AuthAPI;