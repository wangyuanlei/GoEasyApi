package controller

import (
	"GoEasyApi/helper"
	"GoEasyApi/model"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

var UserModel = model.User{}

// 用户信息注册
func Register(ctx *gin.Context) {
	var params struct {
		Name     string `json:"name"`
		Account  string `json:"account"`
		Password string `json:"password"`
		DeptId   string `json:"deptId"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	err := UserModel.RegisterUser(params.Name, params.Account, params.Password, params.DeptId)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}

// 修改用户信息
func ChangeUserInfo(ctx *gin.Context) {
	var params struct {
		UserId string `json:"user_id"`
		Name   string `json:"name"`
		DeptId string `json:"dept_id"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	err := UserModel.ChangeInfo(params.UserId, params.Name, params.DeptId)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}

// 修改密码
func ChangeUserPassword(ctx *gin.Context) {
	var params struct {
		UserId      string `json:"user_id"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	if params.OldPassword == params.NewPassword {
		helper.ApiError(ctx, 602, "新密码不能和旧密码一样", nil)
		return
	}

	err := UserModel.ChangePassword(params.UserId, params.OldPassword, params.NewPassword)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}

// 管理员设置用户密码
func AdminChangeUserPassord(ctx *gin.Context) {
	var params struct {
		UserId   string `json:"user_id"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		helper.ApiError(ctx, 601, "请求数据格式错误", nil)
		return
	}

	err := UserModel.AdminChangePassword(params.UserId, params.Password)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, true)
}

// 获得用户信息
func GetUserInfo(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	if userId == "" {
		helper.ApiError(ctx, 601, "用户id信息错误", nil)
		return
	}

	info, err := UserModel.GetCurrentUserInfo(userId)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, info)
}

// 用户列表筛选
func GetUserList(ctx *gin.Context) {
	PageNo, _ := strconv.Atoi(ctx.Query("page_no"))
	PageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	DeptId := ctx.Query("dept_id")
	Name := ctx.Query("name")
	IsValid, _ := strconv.Atoi(ctx.Query("is_valid"))
	if PageNo < 1 {
		PageNo = 1
	}
	if PageSize < 1 {
		PageSize = 15
	}

	list, total, err := UserModel.GetUserList(PageNo, PageSize, DeptId, Name, IsValid)
	if err != nil {
		ShowModelError(ctx, err)
		return
	}

	helper.ApiSuccess(ctx, map[string]interface{}{
		"page_no":    PageNo,
		"page_size":  PageSize,
		"page_count": int(math.Ceil(float64(total) / float64(PageSize))),
		"total":      total,
		"list":       list,
	})
}
