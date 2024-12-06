import request from "@/utils/request";

class AuthAPI {
      static changePsd(AuthorizeCode: string,data: any) {
		const formData = new FormData();
		formData.append("old_pass", data.oldPass);
		formData.append("new_pass", data.newPass);
		return request<any>({
			url: "/manger/reset_pass",
			method: "post",
			data: formData,
			headers: {
				token: AuthorizeCode,
			},
		});
	}
}
export default AuthAPI;