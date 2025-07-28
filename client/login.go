package client

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

func LoginInit() (*UUIDInfoResp, error) {
	var (
		req      = httpClient.R().SetQueryParams(clientSuffix())
		respErr  = new(RespErr)
		uuidInfo = new(UUIDInfoResp)
		reqURL   = fmt.Sprintf("%s/family/manage/getQrCodeUUID.action", ApiUrl)
	)

	if _, err := req.
		SetResult(&uuidInfo).
		SetError(&respErr).
		SetHeaders(AppKeySignatureHeader(reqURL, http.MethodGet)).
		Execute(http.MethodGet, reqURL); err != nil {
		return nil, err
	}

	if respErr.HasError() {
		return nil, respErr
	}

	if uuidInfo.UUID == "" {
		return nil, errors.New("uuidInfo is empty")
	}

	return uuidInfo, nil
}

func LoginQuery(uuid string) (*AccessTokenResponse, error) {
	var (
		req              = httpClient.R().SetQueryParams(clientSuffix())
		respErr          = new(RespErr)
		eAccessTokenResp = new(E189AccessTokenResp)
		reqURL           = fmt.Sprintf("%s/family/manage/qrcodeLoginResult.action", ApiUrl)
		tokenInfo        = new(AppSessionResp)
		accessTokenResp  = new(AccessTokenResponse)
	)

	if _, err := req.
		SetResult(eAccessTokenResp).
		SetError(&respErr).
		SetHeaders(AppKeySignatureHeader(reqURL, http.MethodGet)).
		SetQueryParam("uuid", uuid).
		Execute(http.MethodGet, reqURL); err != nil {
		return nil, err
	}

	if respErr.HasError() {
		return nil, respErr
	}

	if eAccessTokenResp.E189AccessToken == "" {
		return nil, errors.New("E189AccessToken is empty")
	}

	reqURL = fmt.Sprintf("%s/family/manage/loginFamilyMerge.action", ApiUrl)
	req = httpClient.R().SetQueryParams(clientSuffix())

	if _, err := req.SetResult(tokenInfo).SetError(&respErr).
		SetHeaders(AppKeySignatureHeader(reqURL, http.MethodGet)).
		SetQueryParam("e189AccessToken", eAccessTokenResp.E189AccessToken).
		Execute(http.MethodGet, reqURL); err != nil {
		return nil, err
	}

	values := url.Values{}
	if _, err := httpClient.R().
		SetForceResponseContentType("application/json").
		SetResult(accessTokenResp).SetError(&respErr).
		SetHeaders(signatureHeaderC(values, "sessionKey", tokenInfo.SessionKey)).
		Get(fmt.Sprintf("%s/open/oauth2/getAccessTokenBySsKey.action?%s", ApiUrl, values.Encode())); err != nil {
		return nil, err
	}

	return accessTokenResp, nil
}
