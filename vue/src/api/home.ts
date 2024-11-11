import request from "@/utils/request";

class AuthAPI {
      static changePsd(AuthorizeCode: string,data: any) {
		const formData = new FormData();
		formData.append("oldPass", data.oldPass);
		formData.append("newPass", data.newPass);
		return request<any>({
			url: "/api/manger/reset_pass",
			method: "post",
			data: formData,
			headers: {
				token: AuthorizeCode,
			},
		});
	}
}
export default AuthAPI;