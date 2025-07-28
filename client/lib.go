package client

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func signatureHeader(values url.Values, accessToken string) map[string]string {
	tt := timestamp()
	signType := "1"

	values.Add("AccessToken", accessToken)
	values.Add("Timestamp", strconv.FormatInt(tt, 10))
	values.Add("AppKey", TVAppKey)

	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	query := ""

	for _, k := range keys {
		vals := values[k]
		sort.Strings(vals) //
		for _, v := range vals {
			if len(query) > 0 {
				query += "&"
			}
			query += k + "=" + v
		}
	}

	mp := map[string]string{
		"Signature":   md5(query),
		"Sign-Type":   signType,
		"Timestamp":   strconv.FormatInt(tt, 10),
		"AppKey":      TVAppKey,
		"AccessToken": accessToken,
	}

	return mp
}

func clientSuffix() map[string]string {
	return map[string]string{
		"clientType":        AndroidTV,
		"version":           TvVersion,
		"channelId":         TvChannelId,
		"clientSn":          "unknown",
		"model":             "PJX110",
		"osFamily":          "Android",
		"osVersion":         "35",
		"networkAccessMode": "WIFI",
		"telecomsOperator":  "46011",
	}
}

// AppKeySignatureOfHmac HMAC签名
func AppKeySignatureOfHmac(sessionSecret, appKey, operate, fullUrl string, timestamp int64) string {
	urlpath := regexp.MustCompile(`://[^/]+((/[^/\s?#]+)*)`).FindStringSubmatch(fullUrl)[1]
	mac := hmac.New(sha1.New, []byte(sessionSecret))
	data := fmt.Sprintf("AppKey=%s&Operate=%s&RequestURI=%s&Timestamp=%d", appKey, operate, urlpath, timestamp)
	mac.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}

func AppKeySignatureHeader(url, method string) map[string]string {
	tempTime := timestamp()
	header := map[string]string{
		"Timestamp":    strconv.FormatInt(tempTime, 10),
		"X-Request-ID": uuid.NewString(),
		"AppKey":       TVAppKey,
		"AppSignature": AppKeySignatureOfHmac(TVAppSignatureSecret, TVAppKey, method, url, tempTime),
	}
	return header
}

func signatureHeaderC(values url.Values, tokenType string, token string) map[string]string {
	tt := timestamp()
	signType := "1"

	// AccessToken
	values.Add(tokenType, token)

	values.Add("Timestamp", strconv.FormatInt(tt, 10))
	values.Add("AppKey", TVAppKey)

	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	query := ""

	for _, k := range keys {
		vals := values[k]
		sort.Strings(vals) // 可选：多个值排序，避免顺序影响
		for _, v := range vals {
			if len(query) > 0 {
				query += "&"
			}
			query += k + "=" + v
		}
	}

	mp := map[string]string{
		"Signature": md5(query),
		"Sign-Type": signType,
		"Timestamp": strconv.FormatInt(tt, 10),
		"AppKey":    TVAppKey,
	}

	mp[tokenType] = token

	return mp
}
