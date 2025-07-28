package client

import (
	"bytes"
	md52 "crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/xml"
	"resty.dev/v3"
	"time"
)

// timestamp 时间戳
func timestamp() int64 {
	return time.Now().UTC().UnixNano() / 1e6
}

func md5(data string) string {
	h := md52.New()
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}

func newRestyClient() *resty.Client {
	c := resty.New().
		SetHeader("User-Agent", UserAgent).
		SetHeader("Accept", Accept).
		SetRetryCount(3).
		SetTimeout(DefaultTimeout).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return c
}

type String string

func (s *String) UnmarshalJSON(b []byte) error { return s.Unmarshal(b) }
func (s *String) UnmarshalXML(e *xml.Decoder, ee xml.StartElement) error {
	b, err := e.Token()
	if err != nil {
		return err
	}
	if b, ok := b.(xml.CharData); ok {
		if err = s.Unmarshal(b); err != nil {
			return err
		}
	}
	return e.Skip()
}
func (s *String) Unmarshal(b []byte) error {
	*s = String(bytes.Trim(b, "\""))
	return nil
}
