package support

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"reflect"
	"sort"
	"strings"
)

func PaymentV2ParamsJoinBackup(params *power.HashMap, key string) string {
	var arr []string
	for k, v := range *params {
		if v == "" {
			continue
		}
		arr = append(arr, k)
	}
	sort.Strings(arr)
	//for i, k := range arr {
	for i := 0; i < len(arr); i++ {
		k := arr[i]
		v := (*params)[k]
		switch v.(type) {
		case string:
			arr[i] = fmt.Sprintf("%s=%s", k, v)
			break
		case int:
		case int8:
		case int16:
		case int32:
		case int64:
			arr[i] = fmt.Sprintf("%s=%d", k, v)
			break
		case float32:
		case float64:
			arr[i] = fmt.Sprintf("%s=%f", k, v)
			break
		}
		//arr[i] = fmt.Sprintf("%s=%x", k, (*params)[k])
	}
	return fmt.Sprintf("%s&key=%s", strings.Join(arr, "&"), key)
}

func PaymentV2ParamsJoin(params *power.StringMap, key string) string {
	var arr []string
	for k, v := range *params {
		if v == "" {
			continue
		}
		arr = append(arr, k)
	}
	sort.Strings(arr)
	for i, k := range arr {
		arr[i] = fmt.Sprintf("%s=%s", k, (*params)[k])
	}
	return fmt.Sprintf("%s&key=%s", strings.Join(arr, "&"), key)
}

// GenerateSignMD5 适用于微信支付V2 MD5签名算法
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=4_3
func GenerateSignMD5(params *power.StringMap, key string) string {
	sign := fmt.Sprintf("%x", md5.Sum([]byte(PaymentV2ParamsJoin(params, key))))
	return strings.ToUpper(sign)
}

// GenerateSignHmacSHA256 适用于微信支付V2 HMAC-SHA256签名算法
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=4_3
func GenerateSignHmacSHA256(params *power.StringMap, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(PaymentV2ParamsJoin(params, key)))

	sign := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(sign)
}

func DeepCopy(src interface{}) (interface{}, error) {
	// 获取源对象的反射值
	v := reflect.ValueOf(src)

	// 如果源对象不可修改或者是空指针，则直接返回源对象
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return src, nil
	}

	// 创建目标对象的反射指针
	vp := reflect.New(v.Elem().Type())

	// 深拷贝源对象的值到目标对象中
	vp.Elem().Set(v.Elem())

	// 返回目标对象的值
	return vp.Elem().Interface(), nil
}
