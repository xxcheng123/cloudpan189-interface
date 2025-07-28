package client

import (
	"encoding/xml"
	"fmt"
)

type E189AccessTokenResp struct {
	E189AccessToken string `json:"accessToken"`
	ExpiresIn       int64  `json:"expiresIn"`
}

type AccessTokenResponse struct {
	ExpiresIn   int64  `json:"expiresIn"`
	AccessToken string `json:"accessToken"`
}

type UserSessionResp struct {
	ResCode    int    `json:"res_code"`
	ResMessage string `json:"res_message"`

	LoginName string `json:"loginName"`

	KeepAlive       int `json:"keepAlive"`
	GetFileDiffSpan int `json:"getFileDiffSpan"`
	GetUserInfoSpan int `json:"getUserInfoSpan"`

	// 个人云
	SessionKey    string `json:"sessionKey"`
	SessionSecret string `json:"sessionSecret"`
	// 家庭云
	FamilySessionKey    string `json:"familySessionKey"`
	FamilySessionSecret string `json:"familySessionSecret"`
}

// AppSessionResp 登录返回
type AppSessionResp struct {
	UserSessionResp

	IsSaveName string `json:"isSaveName"`

	// 会话刷新Token
	AccessToken string `json:"accessToken"`
	//Token刷新
	RefreshToken string `json:"refreshToken"`
}

// RespErr 居然有四种返回方式
type RespErr struct {
	ResCode    any    `json:"res_code"` // int or string
	ResMessage string `json:"res_message"`

	Error_ string `json:"error"`

	XMLName xml.Name `xml:"error"`
	Code    string   `json:"code" xml:"code"`
	Message string   `json:"message" xml:"message"`
	Msg     string   `json:"msg"`

	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

func (e *RespErr) HasError() bool {
	switch v := e.ResCode.(type) {
	case int, int64, int32:
		return v != 0
	case string:
		return e.ResCode != ""
	}
	return (e.Code != "" && e.Code != "SUCCESS") || e.ErrorCode != "" || e.Error_ != ""
}

func (e *RespErr) Error() string {
	switch v := e.ResCode.(type) {
	case int, int64, int32:
		if v != 0 {
			return fmt.Sprintf("res_code: %d ,res_msg: %s", v, e.ResMessage)
		}
	case string:
		if e.ResCode != "" {
			return fmt.Sprintf("res_code: %s ,res_msg: %s", e.ResCode, e.ResMessage)
		}
	}

	if e.Code != "" && e.Code != "SUCCESS" {
		if e.Msg != "" {
			return fmt.Sprintf("code: %s ,msg: %s", e.Code, e.Msg)
		}
		if e.Message != "" {
			return fmt.Sprintf("code: %s ,msg: %s", e.Code, e.Message)
		}
		return "code: " + e.Code
	}

	if e.ErrorCode != "" {
		return fmt.Sprintf("err_code: %s ,err_msg: %s", e.ErrorCode, e.ErrorMsg)
	}

	if e.Error_ != "" {
		return fmt.Sprintf("error: %s ,message: %s", e.ErrorCode, e.Message)
	}
	return ""
}

type UUIDInfoResp struct {
	UUID string `json:"uuid"`
}

type FileInfo struct {
	CreateDate  string `json:"createDate"`
	FileCata    int    `json:"fileCata"`
	Id          String `json:"id"`
	LastOpTime  string `json:"lastOpTime"`
	Md5         string `json:"md5"`
	MediaType   int    `json:"mediaType"`
	Name        string `json:"name"`
	Rev         string `json:"rev"`
	Size        int64  `json:"size"`
	StarLabel   int    `json:"starLabel"`
	Orientation int    `json:"orientation,omitempty"`
}

type FolderInfo struct {
	CreateDate   string `json:"createDate"`
	FileCata     int    `json:"fileCata"`
	FileCount    int    `json:"fileCount"`
	FileListSize int    `json:"fileListSize"`
	Id           String `json:"id"`
	LastOpTime   string `json:"lastOpTime"`
	Name         string `json:"name"`
	ParentId     String `json:"parentId"`
	Rev          string `json:"rev"`
	StarLabel    int    `json:"starLabel"`
}

type ListFilesResponse struct {
	ResCode    int    `json:"res_code"`
	ResMessage string `json:"res_message"`
	FileListAO struct {
		Count        int64        `json:"count"`
		FileList     []FileInfo   `json:"fileList"`
		FileListSize int64        `json:"fileListSize"`
		FolderList   []FolderInfo `json:"folderList"`
	} `json:"fileListAO"`
	LastRev int64 `json:"lastRev"`
}
