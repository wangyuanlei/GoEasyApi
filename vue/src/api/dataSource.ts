import request from "@/utils/request";

class AuthAPI {
	static getDatas(AuthorizeCode: string) {
		return request<any>({
			url: `/api/manger/db/get_conf`,
			method: "get",
			headers: {
				token: AuthorizeCode,
			},
		});
	}
	/**
	 * 保存数据库配置
	 */
	static saveData(AuthorizeCode: string, data: any) {
		const formData = new FormData();
        formData.append("name", data.DatabaseName);
		formData.append("description", data.Description);
		formData.append("orm_type", data.OrmType);
		formData.append("dns", data.Dns);
        // console.log('formData', formData);
		return request<any>({
		  url: "/api/manger/db/set_conf",
		  method: "post",
		  data: formData,
		  headers: {
			token: AuthorizeCode,
		  },
		});
	  }
}
export default AuthAPI;