import request from "@/utils/request";

class AuthAPI {
	//获取接口列表
	static getUesrList(AuthorizeCode: string) {
		// console.log('formobj=',formobj);

		return request<any>({
			url: `/manger/interface/list`,
			method: "get",
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	/**
 * 获取接口详情
 */
	static getInterfaceInfo(AuthorizeCode: string, data: any) {
		let formobj = {
			id: data,
		}
		return request<any>({
			url: "/manger/interface/info",
			method: "get",
			//   data: formData,
			params: formobj,
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	/**
	 * 编辑保存
	 */
	static update(AuthorizeCode: string,data:any) {
		// console.log('data',data);
		return request<any>({
			url: "/manger/interface/update",
			method: "post",
			data: data,
			headers: {
				token: AuthorizeCode,
			  },
		});
	}
	/**
	 * 新增
	 */
	static addInterface(AuthorizeCode: string, data: any) {
		console.log('data',data);
		return request<any>({
			url: "/manger/interface/add",
			method: "post",
			data: data,
			headers: {
				token: AuthorizeCode,
			  },
		});
	}
	/**
	 * 删除
	 */
	static deleteInterface(AuthorizeCode: string, data: string) {
		// console.log(data);
		const formData = new FormData();
		formData.append("id", data);
		return request<any>({
		  url: "/manger/interface/delete",
		  method: "post",
		  data: formData,
		  headers: {
			token: AuthorizeCode,
			// "Content-Type":"multipart/form-data"
		  },
		});
	  }
	/**
	 * 用户修改密码
	 */
	static updtPsd(AuthorizeCode: string, id: string, password:string) {
		// console.log(data);
		const formData = new FormData();
		formData.append("user_id", id);
		formData.append("password", password);
		console.log('formData', formData);
		return request<any>({
		  url: "/manger/set_user_pass",
		  method: "post",
		  data: formData,
		  headers: {
			token: AuthorizeCode,
		  },
		});
	  }
}
export default AuthAPI;
