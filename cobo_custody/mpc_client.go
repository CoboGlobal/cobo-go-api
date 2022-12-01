package cobo_custody

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/btcsuite/btcd/btcec"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type MPCClient struct {
	Signer ApiSigner
	Env    Env
	Debug  bool
}

func (c MPCClient) GetSupportedChains() (*simplejson.Json, *ApiError) {
	var params = map[string]string{
	}
	return c.Request("GET", "/v1/custody/mpc/get_supported_chains/", params)
}

func (c MPCClient) GetSupportedCoins(chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
	}
	return c.Request("GET", "/v1/custody/mpc/get_supported_coins/", params)
}

func (c MPCClient) BatchGenerateNewAddresses(chainCode string, count int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
		"count": strconv.Itoa(count),
	}
	return c.Request("POST", "/v1/custody/mpc/generate_new_addresses/", params)
}

func (c MPCClient) GetAddressList(chainCode string, pageIndex int, pageLength int, options map[string]string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
		"page_index": strconv.Itoa(pageIndex),
		"page_length": strconv.Itoa(pageLength),
	}
	if options["sort_flag"] != "" {
		params["sort_flag"] = options["sort_flag"]
	}

	return c.Request("GET", "/v1/custody/mpc/list_addresses/", params)
}

func (c MPCClient) GetAssetList(options map[string]string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
	}

	if options["address"] != "" {
		params["address"] = options["address"]
	}
	if options["chain_code"] != "" {
		params["chain_code"] = options["chain_code"]
	}
	
	return c.Request("GET", "/v1/custody/mpc/list_assets/", params)
}

func (c MPCClient) CreateTransaction(coin string, requestId string, fromAddr string, toAddr string, amount *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin": coin,
		"request_id": requestId,
		"from_address": fromAddr,
		"to_address": toAddr,
		"amount": amount.String(),
	}
	return c.Request("POST", "/v1/custody/mpc/create_transaction/", params)
}

func (c MPCClient) GetTransaction(requestId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id": requestId,
	}
	return c.Request("GET", "/v1/custody/mpc/transaction_info/", params)
}

func (c MPCClient) GetTransactionByTxId(txId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"tx_id": txId,
	}
	return c.Request("GET", "/v1/custody/mpc/transaction_info_by_tx_id/", params)
}

func (c MPCClient) GetWalletTransactions(address string, options map[string]string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"address": address, 
	}

	if options["coin"] != ""{
		params["coin"] = options["coin"]
	}
	if options["max_id"] != "" {
		params["max_id"] = options["max_id"]
	}
	if options["min_id"] != "" {
		params["min_id"] = options["min_id"]
	}
	if options["limit"] != "" {
		params["limit"] = options["limit"]
	}

	return c.Request("GET", "/v1/custody/mpc/list_transactions/", params)
}

func (c MPCClient) request(method string, path string, params map[string]string) string {
	httpClient := &http.Client{}
	nonce := fmt.Sprintf("%d", time.Now().Unix()*1000)
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
	resp, _ := httpClient.Do(req)

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
	return string(body)
}

func (c MPCClient) Request(method string, path string, params map[string]string) (*simplejson.Json, *ApiError) {
	jsonString := c.request(method, path, params)
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

func (c MPCClient) VerifyEcc(message string, signature string) bool {
	pubKeyBytes, _ := hex.DecodeString(c.Env.CoboPub)
	pubKey, _ := btcec.ParsePubKey(pubKeyBytes, btcec.S256())

	sigBytes, _ := hex.DecodeString(signature)
	sigObj, _ := btcec.ParseSignature(sigBytes, btcec.S256())

	verified := sigObj.Verify([]byte(Hash256x2(message)), pubKey)
	return verified
}
