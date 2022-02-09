package qianchuanSDK

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"

	"github.com/CriarBrand/qianchuanSDK/auth"
)

type transport struct {
	http.RoundTripper
	credentials *auth.Credentials
}

func newTransport(credentials *auth.Credentials, tr http.RoundTripper) *transport {
	if tr == nil {
		tr = http.DefaultTransport
	}
	return &transport{tr, credentials}
}

// Base64Encode Base64编码
func Base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

// Base64Decode Base64解码
func Base64Decode(encodeString string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodeString)
}

// PKCS5UnPadding PKCS5填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesDecrypt AES解密
func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != block.BlockSize() {
		var keySizeError aes.KeySizeError
		return nil, errors.New(keySizeError.Error())
	}
	encrypter := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	encrypter.CryptBlocks(origData, crypted)
	return PKCS5UnPadding(origData), nil
}

// BuildQuery 组建query参数，noInclude是不需要写入到query的参数（填的是结构体json属性的值）
func BuildQuery(reqUrl string, param interface{}, noInclude []string) (string, error) {
	//利用omitempty的特性，用序列化后再反序列化到map去除零值
	marshal, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	var paramMap = make(map[string]interface{})
	err = json.Unmarshal(marshal, &paramMap)
	if err != nil {
		return "", err
	}
	//拼接sql
	var paramSlice []string
	for k, v := range paramMap {
		var omit bool
		for _, v2 := range noInclude {
			if k == v2 {
				omit = true
				break
			}
		}
		if !omit {
			//整型序列化后再反序列化到map[string]interface{}会变成float64
			fmt.Printf("%T,%v\n", v, v)
			if float64Value, ok := v.(float64); ok {
				//判断是不是真的是浮点数
				if float64Value > math.Floor(float64Value) {
					paramSlice = append(paramSlice, k+"="+fmt.Sprintf("%f", float64Value))
				} else {
					paramSlice = append(paramSlice, k+"="+fmt.Sprintf("%+v", int64(float64Value)))
				}
			} else if mapValue, ok := v.(map[string]interface{}); ok {
				bytes, err := json.Marshal(mapValue)
				if err != nil {
					return "", err
				}
				paramSlice = append(paramSlice, k+"="+url.QueryEscape(fmt.Sprintf("%v", string(bytes))))
			} else if sliceValue, ok := v.([]interface{}); ok {
				bytes, err := json.Marshal(sliceValue)
				if err != nil {
					return "", err
				}
				paramSlice = append(paramSlice, k+"="+url.QueryEscape(fmt.Sprintf("%v", string(bytes))))
			} else {
				paramSlice = append(paramSlice, k+"="+url.QueryEscape(fmt.Sprintf("%v", v)))
			}
		}
	}
	paramStr := strings.Join(paramSlice, "&")
	if paramStr != "" {
		reqUrl += "?" + paramStr
	}
	fmt.Println("请求", reqUrl)
	return reqUrl, nil
}
