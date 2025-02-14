import request from "@/utils/request";

class AuthAPI {
	//获取用户列表
	static getUesrList(AuthorizeCode: string, data: any) {
		let formobj = {
			page: data.page,
			page_size: data.page_size,
			name: data.search,
		}
		return request<any>({
			url: `/manger/get_user_list`,
			method: "get",
			params: formobj,
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	/**
	 * 用户注册
	 */
	static createUser(AuthorizeCode: string, data: any) {
		const formData = new FormData();
		formData.append("name", data.name);
		formData.append("deptId", data.deptId);
		formData.append("account", data.account);
		formData.append("password", data.password);
		return request<any>({
			url: "/api/user/register",
			method: "post",
			data: formData,
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	/**
 * 获取用户信息
 */
	static getUser(AuthorizeCode: string, data: any) {
		let formobj = {
			user_id: data,
		}
		return request<any>({
			url: "/manger/get_user_info",
			method: "get",
			params: formobj,
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	/**
	 * 编辑保存
	 */
	static updateUser(AuthorizeCode: string,id:string,data:any) {
		data.user_id=id;
		data.is_valid=Number(data.isValid);
		data.dept_id=data.deptId
		delete data.isValid;
		delete data.deptId;

		return request<any>({
			url: "/manger/set_user_info",
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
	static deleteUser(AuthorizeCode: string, data: string) {
		// console.log('data=',data);
		const formData = new FormData();
		formData.append("user_id", data);
		return request<any>({
		  url: "/manger/delete_user",
		  method: "post",
		  data: formData,
		  headers: {
			token: AuthorizeCode,
		  },
		});
	  }
	/**
	 * 用户修改密码
	 */
	static updtPsd(AuthorizeCode: string, id: string, password:string) {
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
