package cobo_custody

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

type MPCPrimeBrokerClient struct {
	Signer ApiSigner
	Env    Env
	Debug  bool
}

func (c MPCPrimeBrokerClient) CreateBinding(userId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"user_id": userId,
	}

	return c.Request("POST", "/v1/custody/guard/create_binding/", params)
}

func (c MPCPrimeBrokerClient) QueryBinding(binderId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"binder_id": binderId,
	}

	return c.Request("GET", "/v1/custody/guard/query_binding/", params)
}

func (c MPCPrimeBrokerClient) QueryUserAuth(userId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"user_id": userId,
	}

	return c.Request("GET", "/v1/custody/guard/query_user_auth/", params)
}

func (c MPCPrimeBrokerClient) BindAddresses(userId string, addresses string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"user_id":   userId,
		"addresses": addresses,
	}

	return c.Request("POST", "/v1/custody/guard/bind_addresses/", params)
}

func (c MPCPrimeBrokerClient) ChangeBinding(userId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"user_id": userId,
	}

	return c.Request("POST", "/v1/custody/guard/change_binding/", params)
}

func (c MPCPrimeBrokerClient) UnbindBinding(userId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"user_id": userId,
	}

	return c.Request("POST", "/v1/custody/guard/unbind_binding/", params)
}

func (c MPCPrimeBrokerClient) QueryStatement(statementId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"statement_id": statementId,
	}

	return c.Request("GET", "/v1/custody/guard/query_statement/", params)
}

func (c MPCPrimeBrokerClient) request(method string, path string, params map[string]string) (string, error) {
	httpClient := &http.Client{}
	nonce := fmt.Sprintf("%d", time.Now().UnixNano()/1000)
	sorted := SortParams(params)
	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest(method, c.Env.Host+path, strings.NewReader(sorted))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, _ = http.NewRequest(method, c.Env.Host+path+"?"+sorted, strings.NewReader(""))
	}
	content := strings.Join([]string{method, path, nonce, sorted}, "|")

	req.Header.Set("Biz-Api-Key", c.Signer.GetPublicKey())
	req.Header.Set("Biz-Api-Nonce", nonce)
	req.Header.Set("Biz-Api-Signature", c.Signer.Sign(content))

	if c.Debug {
		fmt.Println("request >>>>>>>>")
		fmt.Println(method, "\n", path, "\n", params, "\n", content, "\n", req.Header)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("http request err", err.Error())
		return "", err
	}
	if resp.Header == nil {
		return "", errors.New("http resp header is nil")
	}
	if len(resp.Header["Biz-Timestamp"]) <= 0 {
		return "", errors.New("http resp header timestamp is illegal")
	}
	if len(resp.Header["Biz-Resp-Signature"]) <= 0 {
		return "", errors.New("http resp header signature is illegal")
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	timestamp := resp.Header["Biz-Timestamp"][0]
	signature := resp.Header["Biz-Resp-Signature"][0]
	if c.Debug {
		fmt.Println("response <<<<<<<<")
		fmt.Println(string(body), "\n", timestamp, "\n", signature)
	}
	success := c.VerifyEcc(string(body)+"|"+timestamp, signature)
	if !success {
		panic("response signature verify failed")
	}
	return string(body), nil
}

func (c MPCPrimeBrokerClient) Request(method string, path string, params map[string]string) (*simplejson.Json, *ApiError) {
	jsonString, err := c.request(method, path, params)
	if err != nil {
		return nil, &ApiError{ErrorMessage: err.Error()}
	}

	json, _ := simplejson.NewJson([]byte(jsonString))
	success, _ := json.Get("success").Bool()
	if !success {
		errorId, _ := json.Get("error_id").String()
		errorMessage, _ := json.Get("error_message").String()
		errorCode, _ := json.Get("error_code").Int()
		apiError := ApiError{
			ErrorId:      errorId,
			ErrorMessage: errorMessage,
			ErrorCode:    errorCode,
		}
		return nil, &apiError
	}

	result := json.Get("result")
	return result, nil
}

func (c MPCPrimeBrokerClient) VerifyEcc(message string, signature string) bool {
	pubKeyBytes, err := hex.DecodeString(c.Env.CoboPub)
	if err != nil {
		fmt.Println("decode pubkey error ", err)
		return false
	}

	pubKey, err := btcec.ParsePubKey(pubKeyBytes)
	if err != nil {
		fmt.Println("parse pubkey error ", err)
		return false
	}

	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		fmt.Println("decode signature error ", err)
		return false
	}

	sigObj, err := ecdsa.ParseSignature(sigBytes)
	if err != nil {
		fmt.Println("parse signature error ", err)
		return false
	}

	verified := sigObj.Verify([]byte(Hash256x2(message)), pubKey)
	return verified
}
