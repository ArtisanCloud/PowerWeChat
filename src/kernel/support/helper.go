package support
import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"sort"
	"strings"
)

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
