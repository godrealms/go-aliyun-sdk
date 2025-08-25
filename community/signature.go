package community

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

// SignatureHelper 签名助手
type SignatureHelper struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewSignatureHelper 创建签名助手
func NewSignatureHelper(privateKeyPEM string) (*SignatureHelper, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return &SignatureHelper{
		privateKey: privateKey,
	}, nil
}

// GenerateSignature 生成签名
func (sh *SignatureHelper) GenerateSignature(params interface{}) (string, error) {
	// 1. 过滤参数
	filteredParams := sh.filterParams(params)

	// 2. 排序并构建待签名字符串
	signContent := sh.buildSignContent(filteredParams)
	log.Println("signContent: ", signContent)
	// 3. 生成签名
	signature, err := sh.sign(signContent)
	if err != nil {
		return "", fmt.Errorf("failed to generate signature: %w", err)
	}

	return signature, nil
}

// filterParams 过滤参数，支持 map 和结构体
func (sh *SignatureHelper) filterParams(params interface{}) map[string]string {
	filtered := make(map[string]string)

	// 使用反射获取参数值
	val := reflect.ValueOf(params)

	// 处理指针类型
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		// 处理结构体
		sh.handleStruct(val, filtered)
	case reflect.Map:
		// 处理 map
		sh.handleMap(val, filtered)
	default:
		// 不支持的类型
		return filtered
	}

	return filtered
}

// handleStruct 处理结构体类型
func (sh *SignatureHelper) handleStruct(val reflect.Value, filtered map[string]string) {
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// 获取 json 标签
		tag := fieldType.Tag.Get("json")
		if tag == "" {
			tag = fieldType.Name
		} else {
			// 处理 json 标签中的选项
			if idx := strings.Index(tag, ","); idx != -1 {
				tag = tag[:idx]
			}
			// 跳过 "-" 标记的字段
			if tag == "-" {
				continue
			}
		}

		// 跳过 sign 字段
		if tag == "sign" {
			continue
		}

		// 处理嵌套结构体
		if field.Kind() == reflect.Struct {
			sh.handleStruct(field, filtered)
			continue
		}

		// 获取字段值
		value := sh.getFieldValue(field)
		if value != "" {
			filtered[tag] = value
		}
	}
}

// handleMap 处理 map 类型
func (sh *SignatureHelper) handleMap(val reflect.Value, filtered map[string]string) {
	for _, key := range val.MapKeys() {
		keyStr := fmt.Sprint(key.Interface())
		if keyStr == "sign" {
			continue
		}

		value := val.MapIndex(key)
		strValue := sh.getFieldValue(value)
		if strValue != "" {
			filtered[keyStr] = strValue
		}
	}
}

// getFieldValue 获取字段值的字符串表示
func (sh *SignatureHelper) getFieldValue(value reflect.Value) string {
	// 处理接口类型
	if value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	// 处理指针类型
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return ""
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.String:
		str := value.String()
		if str == "" {
			return ""
		}
		return str

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", value.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", value.Uint())

	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%.2f", value.Float())

	case reflect.Bool:
		return fmt.Sprintf("%t", value.Bool())

	case reflect.Slice, reflect.Array:
		if value.Len() == 0 {
			return ""
		}
		// 对于切片类型，我们可以选择将其转换为 JSON 字符串
		bytes, err := json.Marshal(value.Interface())
		if err != nil {
			return ""
		}
		return string(bytes)

	case reflect.Map, reflect.Struct:
		// 对于 Map 和 Struct 类型，转换为 JSON 字符串
		bytes, err := json.Marshal(value.Interface())
		if err != nil {
			return ""
		}
		return string(bytes)

	default:
		// 其他类型尝试直接转字符串
		return fmt.Sprint(value.Interface())
	}
}

// buildSignContent 构建待签名内容
func (sh *SignatureHelper) buildSignContent(params map[string]string) string {
	// 获取所有键并排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建待签名字符串
	var pairs []string
	for _, key := range keys {
		pairs = append(pairs, fmt.Sprintf("%s=%s", key, params[key]))
	}

	return strings.Join(pairs, "&")
}

// sign 签名
func (sh *SignatureHelper) sign(content string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(content))
	digest := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, sh.privateKey, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifySignature 验证签名
func (sh *SignatureHelper) VerifySignature(params interface{}, signature string) error {
	// 1. 过滤参数
	filteredParams := sh.filterParams(params)

	// 2. 构建待验证的内容
	signContent := sh.buildSignContent(filteredParams)
	log.Println("signContent: ", signContent)
	// 3. 解码签名
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("failed to decode signature: %w", err)
	}

	// 4. 计算哈希
	hash := sha256.New()
	hash.Write([]byte(signContent))
	digest := hash.Sum(nil)

	// 5. 验证签名
	err = rsa.VerifyPKCS1v15(&sh.privateKey.PublicKey, crypto.SHA256, digest, signatureBytes)
	if err != nil {
		return fmt.Errorf("signature verification failed: %w", err)
	}

	return nil
}

// AlipayVerifyService alipay visa verification service
type AlipayVerifyService struct {
	AlipayPublicKey string // 支付宝公钥
}

func NewAlipayVerifyService(alipayPublicKey string) *AlipayVerifyService {
	return &AlipayVerifyService{
		AlipayPublicKey: alipayPublicKey,
	}
}

// VerifyNotifySign Verify Alipay asynchronous notification signature
func (s *AlipayVerifyService) VerifyNotifySign(values url.Values) (bool, error) {
	// 1. 获取签名
	sign := values.Get("sign")
	if sign == "" {
		return false, fmt.Errorf("signature is empty")
	}

	// 2. 获取签名类型
	signType := values.Get("sign_type")
	if signType == "" {
		signType = "RSA2" // 默认RSA2
	}

	// 3. 构建待签名字符串
	signStr, err := s.buildSignString(values)
	if err != nil {
		return false, fmt.Errorf("failed to build signature string: %v", err)
	}

	log.Printf("signContent: %s\n", signStr)

	// 4. 验证签名
	return s.verifyRSA2Sign(signStr, sign)
}

// build a string to be signed
func (s *AlipayVerifyService) buildSignString(values url.Values) (string, error) {
	var keys []string

	// 收集所有参数名（除了sign和sign_type）
	for key := range values {
		if key != "sign" && key != "sign_type" {
			keys = append(keys, key)
		}
	}

	// 按字典序排序
	sort.Strings(keys)

	// 构建签名字符串
	var signParts []string
	for _, key := range keys {
		value := values.Get(key)
		if value != "" { // 空值不参与签名
			signParts = append(signParts, fmt.Sprintf("%s=%s", key, value))
		}
	}

	return strings.Join(signParts, "&"), nil
}

// rsa2 signature verification
func (s *AlipayVerifyService) verifyRSA2Sign(data, sign string) (bool, error) {
	// 1. 解析公钥
	publicKey, err := s.parsePublicKey(s.AlipayPublicKey)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %v", err)
	}

	// 2. Base64解码签名
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, fmt.Errorf("signature base64 decoding failed: %v", err)
	}

	// 3. 计算数据的SHA256哈希
	hash := sha256.Sum256([]byte(data))

	// 4. 验证签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signBytes)
	if err != nil {
		return false, fmt.Errorf("signature verification failed: %v", err)
	}

	return true, nil
}

// analyze the alipay public key
func (s *AlipayVerifyService) parsePublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
	// 如果公钥不包含PEM头尾，则添加
	if !strings.Contains(publicKeyStr, "-----BEGIN") {
		publicKeyStr = "-----BEGIN PUBLIC KEY-----\n" + publicKeyStr + "\n-----END PUBLIC KEY-----"
	}

	// 解析PEM格式
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, fmt.Errorf("public key format error")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	publicKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key type error not rsa public key")
	}

	return publicKey, nil
}
