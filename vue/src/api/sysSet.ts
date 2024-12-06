import request from "@/utils/request";

class AuthAPI {
	//获取黑名单
	static getBlackList(AuthorizeCode: string) {

		return request<any>({
			url: `/manger/black_list/get_list`,
			method: "get",
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	//获取黑名单
	static getWhiteList(AuthorizeCode: string) {

		return request<any>({
			url: `/manger/whilt_list/get_list`,
			method: "get",
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	/**
	 * 新增黑名单
	 */
	static addBlack(AuthorizeCode: string, data: any) {
		// console.log('data', data);
		const formData = new FormData();
		formData.append("ip", data.ip);
		formData.append("description", data.description);

		return request<any>({
			url: "/manger/black_list/add",
			method: "post",
			data: formData,
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	/**
	 * 新增白名单
	 */
	static addWhite(AuthorizeCode: string, data: any) {
		// console.log('data', data);
		const formData = new FormData();
		formData.append("ip", data.ip);
		formData.append("description", data.description);

		return request<any>({
			url: "/manger/whilt_list/add",
			method: "post",
			data: formData,
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	/**
	 * 编辑保存
	 */
	static updateUser(AuthorizeCode: string,id:string,data:any) {
		const formData = new FormData();
		// console.log('data',data);
		
		formData.append("name", data.name);
		formData.append("dept_id",data.deptId);
		formData.append("is_valid",data.isValid);
		formData.append("user_id",id);

		return request<any>({
			url: "/manger/set_user_info",
			method: "post",
			data: formData,
			headers: {
				token: AuthorizeCode,
			  },
		});
	}
	/**
	 * 删除黑名单
	 */
	static deleteBlack(AuthorizeCode: string, data: string) {
		// console.log('blackdata',data);
		const formData = new FormData();
		formData.append("ip", data);
		return request<any>({
		  url: "/manger/black_list/del",
		  method: "post",
		  data: formData,
		  headers: {
			token: AuthorizeCode,
		  },
		});
	  }
	/**
	 * 删除白名单
	 */
	static deleteWhite(AuthorizeCode: string, data: string) {
		// console.log(data);
		const formData = new FormData();
		formData.append("ip", data);
		return request<any>({
		  url: "/manger/whilt_list/del",
		  method: "post",
		  data: formData,
		  headers: {
			token: AuthorizeCode,
		  },
		});
	  }
	  /**
	 * 获取名单类型
	 * 
	 */
	  static getType(AuthorizeCode: string) {
		return request<any>({
		  url: "/manger/list/get_type",
		  method: "get",
		  headers: {
			token: AuthorizeCode,
		  },
		});
	  
	  }
	/**
	 * 设置名单类型
	 */
	static SetType(AuthorizeCode: string, data: string) {
		// console.log('datadata111',data);
		const formData = new FormData();
		formData.append("list_type", data);
		return request<any>({
		  url: "/manger/list/set_type",
		  method: "post",
		  data: formData,
		  headers: {
			token: AuthorizeCode,
		  },
		});
	  }
	}
export default AuthAPI;
