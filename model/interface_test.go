package model

import (
	"GoEasyApi/database"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInterface_AddInterface(t *testing.T) {
	m := &Interface{}
	info := database.Interface{
		InterfaceName:          "test",
		Description:            "test",
		DatabaseId:             "OpenSource",
		Path:                   "/test",
		Method:                 "get",
		CacheEnabled:           1,
		CacheTime:              10,
		RateLimitEnabled:       1,
		RateLimitCount:         10,
		RateLimitTime:          10,
		SqlContent:             "select * from user",
		TokenValidationEnabled: 1,
		ReturnType:             "json",
	}

	id, err := m.AddInterface(info)
	assert.Nil(t, err)
	assert.NotNil(t, id)
}

func TestInterface_UpdateInterface(t *testing.T) {
	m := &Interface{}
	info := database.Interface{
		InterfaceId:            uuid.New().String(),
		InterfaceName:          "test",
		Description:            "test",
		DatabaseId:             "test",
		Path:                   "/test",
		Method:                 "get",
		CacheEnabled:           1,
		CacheTime:              10,
		RateLimitEnabled:       1,
		RateLimitCount:         10,
		RateLimitTime:          10,
		SqlContent:             "test",
		TokenValidationEnabled: 1,
		ReturnType:             "json",
	}

	err := m.UpdateInterface(info)
	assert.Nil(t, err)
}

func TestInterface_DeleteInterface(t *testing.T) {
	m := &Interface{}
	info := database.Interface{
		InterfaceId: uuid.New().String(),
	}

	err := m.DeleteInterface(info.InterfaceId)
	assert.Nil(t, err)
}

func TestInterface_AddParams(t *testing.T) {
	m := &Interface{}
	params := database.Params{
		Name:        "test",
		Type:        "string",
		Description: "test",
		Required:    1,
		Default:     "test",
		Example:     "test",
		Regex:       "test",
	}

	err := m.AddParams(uuid.New().String(), params)
	assert.Nil(t, err)
}

func TestInterface_UpdateParams(t *testing.T) {
	m := &Interface{}
	params := database.Params{
		InterfaceId: uuid.New().String(),
		ParamsId:    uuid.New().String(),
		Name:        "test",
		Type:        "string",
		Description: "test",
		Required:    1,
		Default:     "test",
		Example:     "test",
		Regex:       "test",
	}

	err := m.UpdateParams(params.InterfaceId, params.ParamsId, params)
	assert.Nil(t, err)
}

func TestInterface_DeleteParams(t *testing.T) {
	m := &Interface{}
	params := database.Params{
		InterfaceId: uuid.New().String(),
		ParamsId:    uuid.New().String(),
	}

	err := m.DeleteParams(params.InterfaceId, params.ParamsId)
	assert.Nil(t, err)
}

func TestInterface_GetList(t *testing.T) {
	m := &Interface{}

	list, err := m.GetList()
	assert.Nil(t, err)
	assert.NotNil(t, list)
}

func TestInterface_GetInfo(t *testing.T) {
	m := &Interface{}
	info, err := m.GetInfo(uuid.New().String())
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func TestInterface_CheckMethod(t *testing.T) {
	m := &Interface{}

	err := m.CheckMethod("get")
	assert.Nil(t, err)

	err = m.CheckMethod("post")
	assert.Nil(t, err)

	err = m.CheckMethod("put")
	assert.NotNil(t, err)
}

func TestInterface_CheckEnabled(t *testing.T) {
	m := &Interface{}

	err := m.CheckEnabled(1)
	assert.Nil(t, err)

	err = m.CheckEnabled(2)
	assert.Nil(t, err)

	err = m.CheckEnabled(3)
	assert.NotNil(t, err)
}

func TestInterface_CheckStringFormat(t *testing.T) {
	m := &Interface{}

	err := m.CheckStringFormat("test")
	assert.Nil(t, err)

	err = m.CheckStringFormat("test!")
	assert.NotNil(t, err)
}

func TestInterface_CheckParamType(t *testing.T) {
	m := &Interface{}

	err := m.CheckParamType("string")
	assert.Nil(t, err)

	err = m.CheckParamType("int")
	assert.Nil(t, err)

	err = m.CheckParamType("float")
	assert.Nil(t, err)

	err = m.CheckParamType("bool")
	assert.Nil(t, err)

	err = m.CheckParamType("date")
	assert.Nil(t, err)

	err = m.CheckParamType("datetime")
	assert.Nil(t, err)

	err = m.CheckParamType("test")
	assert.NotNil(t, err)
}
