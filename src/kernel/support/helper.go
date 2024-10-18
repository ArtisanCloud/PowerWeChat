package support

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

// downloadFileFromURL 从URL下载文件并保存到本地临时文件
func DownloadFileFromURL(fileURL string) (string, error) {
	// 发送HTTP请求获取URL内容
	resp, err := http.Get(fileURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download file: %s", resp.Status)
	}

	// 根据URL获取文件名和扩展名
	fileName := GetFileNameFromURL(fileURL, resp)
	if fileName == "" {
		return "", fmt.Errorf("failed to determine file name")
	}

	// 创建一个临时文件路径，但文件名来自URL
	tempDir := os.TempDir()
	tempFilePath := filepath.Join(tempDir, fileName)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	// 将响应的Body内容写入临时文件
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return "", err
	}

	// 返回临时文件的路径
	return filepath.Abs(tempFilePath)
}

// getFileNameFromURL 从URL中提取并解码文件名，若无扩展名则根据Content-Type推断
func GetFileNameFromURL(fileURL string, resp *http.Response) string {
	// 从URL中提取文件名
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return ""
	}

	// 使用 filepath.Base 提取文件名（包括转码字符）
	encodedFileName := filepath.Base(parsedURL.Path)

	// 使用 QueryUnescape 将文件名中的转码字符还原为原始的中文文件名
	fileName, err := url.QueryUnescape(encodedFileName)
	if err != nil {
		return encodedFileName // 如果解码失败，返回原始转码的文件名
	}

	// 如果文件名没有扩展名，尝试根据Content-Type添加
	if !strings.Contains(fileName, ".") {
		ext := GetFileExtensionFromResponse(resp)
		fileName += ext
	}

	return fileName
}

// getFileExtensionFromResponse 根据响应的Content-Type获取文件扩展名
func GetFileExtensionFromResponse(resp *http.Response) string {
	contentType := resp.Header.Get("Content-Type")
	switch contentType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "audio/amr":
		return ".amr"
	case "video/mp4":
		return ".mp4"
	case "application/octet-stream":
		return "" // 如果是二进制流，不加扩展名
	default:
		return "" // 无法确定扩展名，返回空
	}
}
