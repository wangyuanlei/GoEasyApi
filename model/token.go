package model

import (
	"GoEasyApi/cron"
	"GoEasyApi/database"
	"GoEasyApi/libraries"
	"time"

	"github.com/google/uuid"
)

type Token struct{}

// 获得一个新的token
func (t *Token) NewId() string {
	token := uuid.New().String()
	var existing database.Token
	if err := DB.Where("token = ?", token).First(&existing).Error; err == nil {
		return t.NewId()
	}

	return token
}

// 给用户创建一个token
func (t *Token) CreateToken(userId string) (string, error) {
	token := t.NewId()

	err := DB.Create(&database.Token{
		Token:     token,
		UserId:    userId,
		ValidTime: time.Now().Add(2 * time.Hour),
	}).Error

	// uuid 存储到 cache 类
	libraries.AddCache(token, userId, 2*time.Hour)
	return token, err
}

// 删除token
func (t *Token) DeleteToken(token string) error {
	err := DB.Delete(&database.Token{Token: token}).Error
	libraries.DeleteCache(token) //从 cache 删除token
	return err
}

// 根据token 获得用户信息
func (t *Token) GetTokenInfo(token string) (string, error) {
	userId, found := libraries.GetCache(token)
	if found {
		return userId.(string), nil
	}

	var dbToken database.Token
	err := DB.Where("token = ?", token).First(&dbToken).Error
	if err != nil {
		return "", err
	}

	if dbToken.ValidTime.Before(time.Now()) {
		return "", cron.CreateCustomError(601, "token 已经过期")
	}

	return dbToken.UserId, nil
}

// 给 token 续时间(2小时)
func (t *Token) TokenExtendTime(userId string, token string) error {
	err := DB.Model(&database.Token{}).Where("token = ?", token).Update("valid_time", time.Now().Add(2*time.Hour)).Error
	if err != nil {
		return err
	}
	libraries.UpdateCache(token, userId, 2*time.Hour)
	return nil
}

// 清理过期的token
func (t *Token) ClearToken() {
	var expiredTokens []database.Token
	DB.Where("valid_time < ?", time.Now()).Find(&expiredTokens)
	for _, token := range expiredTokens {
		t.DeleteToken(token.Token)
	}
}
