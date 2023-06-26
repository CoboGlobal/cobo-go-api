package cobo_custody

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/btcsuite/btcd/btcec"
)

type Client struct {
	Signer ApiSigner
	Env    Env
	Debug  bool
}

func (c Client) GetAccountInfo() (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/custody/org_info/", map[string]string{})
}

func (c Client) GetCoinInfo(coin string) (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/custody/coin_info/", map[string]string{
		"coin": coin,
	})
}

func (c Client) NewDepositAddress(coin string, nativeSegwit bool) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin": coin,
	}
	if nativeSegwit {
		params["native_segwit"] = "true"
	}
	return c.Request("POST", "/v1/custody/new_address/", params)
}

func (c Client) BatchNewDepositAddress(coin string, count int, nativeSegwit bool) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":  coin,
		"count": strconv.Itoa(count),
	}
	if nativeSegwit {
		params["native_segwit"] = "true"
	}
	return c.Request("POST", "/v1/custody/new_addresses/", params)
}

func (c Client) VerifyDepositAddress(coin string, address string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"address": address,
	}
	return c.Request("GET", "/v1/custody/address_info/", params)
}

func (c Client) BatchVerifyDepositAddress(coin string, addresses string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"address": addresses,
	}
	return c.Request("GET", "/v1/custody/addresses_info/", params)
}

func (c Client) VerifyValidAddress(coin string, addresses string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"address": addresses,
	}
	return c.Request("GET", "/v1/custody/is_valid_address/", params)
}

func (c Client) GetAddressHistory(coin string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin": coin,
	}
	return c.Request("GET", "/v1/custody/address_history/", params)
}

// @param coin  string "ETH"
// @param page_index int start with 0 page
// @param page_length int page size <= 50
// @param sort_flag int 0:DESCENDING 1:ASCENDING
func (c Client) GetAddressHistoryWithPage(params map[string]string) (*simplejson.Json, *ApiError) {

	return c.Request("GET", "/v1/custody/address_history/", params)
}

func (c Client) CheckLoopAddressDetails(coin string, address string, memo string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"address": address,
	}
	if memo != "" {
		params["memo"] = memo
	}
	return c.Request("GET", "/v1/custody/internal_address_info/", params)
}

func (c Client) VerifyLoopAddressList(coin string, addresses string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"address": addresses,
	}

	return c.Request("GET", "/v1/custody/internal_address_info_batch/", params)
}

func (c Client) GetTransactionDetails(txId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"id": txId,
	}

	return c.Request("GET", "/v1/custody/transaction/", params)
}

func (c Client) GetTransactionsById(params map[string]string) (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/custody/transactions_by_id/", params)
}

func (c Client) GetTransactionsByTxid(txid string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"txid": txid,
	}
	return c.Request("GET", "/v1/custody/transaction_by_txid/", params)
}

func (c Client) GetTransactionsByTime(params map[string]string) (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/custody/transactions_by_time/", params)
}

func (c Client) GetPendingTransactions(params map[string]string) (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/custody/pending_transactions/", params)
}

func (c Client) GetPendingTransaction(id string) (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/custody/pending_transaction/", map[string]string{
		"id": id,
	})
}

func (c Client) GetTransactionHistory(params map[string]string) (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/custody/transaction_history/", params)
}

func (c Client) Withdraw(coin string, requestId string, address string, amount *big.Int, options map[string]string) (*simplejson.Json, *ApiError) {
	if requestId == "" {
		hashResult := sha256.Sum256([]byte(address))
		requestId = fmt.Sprintf("sdk_request_id_%s_%d", fmt.Sprintf("%x", hashResult)[0:8], time.Now().Unix()*1000)
	}
	var params = map[string]string{
		"coin":       coin,
		"request_id": requestId,
		"address":    address,
		"amount":     amount.String(),
	}
	if options["memo"] != "" {
		params["memo"] = options["memo"]
	}

	if options["force_external"] != "" {
		params["force_external"] = options["force_external"]
	}

	if options["force_internal"] != "" {
		params["force_internal"] = options["force_internal"]
	}
	return c.Request("POST", "/v1/custody/new_withdraw_request/", params)
}

func (c Client) QueryWithdrawInfo(requestId string) (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/custody/withdraw_info_by_request_id/", map[string]string{"request_id": requestId})
}

func (c Client) GetStakingProductDetails(productId string, language string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"product_id": productId,
		"language":   language,
	}
	return c.Request("GET", "/v1/custody/staking_product/", params)
}

func (c Client) GetStakingProductList(coin string, language string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"language": language,
	}
	if coin != "" {
		params["coin"] = coin
	}

	return c.Request("GET", "/v1/custody/staking_products/", params)
}

func (c Client) Stake(productId string, amount *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"product_id": productId,
		"amount":     amount.String(),
	}
	return c.Request("POST", "/v1/custody/staking_stake/", params)
}

func (c Client) Unstake(productId string, amount *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"product_id": productId,
		"amount":     amount.String(),
	}
	return c.Request("POST", "/v1/custody/staking_unstake/", params)
}

func (c Client) GetStakings(coin string, language string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"language": language,
	}
	if coin != "" {
		params["coin"] = coin
	}
	return c.Request("GET", "/v1/custody/stakings/", params)
}

func (c Client) GetUnstakings(coin string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{}
	if coin != "" {
		params["coin"] = coin
	}
	return c.Request("GET", "/v1/custody/unstakings/", params)
}

func (c Client) GetStakingHistory() (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/custody/staking_history/", map[string]string{})
}

func (c Client) request(method string, path string, params map[string]string) (string, error) {
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

func (c Client) Request(method string, path string, params map[string]string) (*simplejson.Json, *ApiError) {
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

func (c Client) VerifyEcc(message string, signature string) bool {
	pubKeyBytes, err := hex.DecodeString(c.Env.CoboPub)
	if err != nil {
		fmt.Println("decode pubkey error ", err)
		return false
	}

	pubKey, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	if err != nil {
		fmt.Println("parse pubkey error ", err)
		return false
	}

	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		fmt.Println("decode signature error ", err)
		return false
	}

	sigObj, err := btcec.ParseSignature(sigBytes, btcec.S256())
	if err != nil {
		fmt.Println("parse signature error ", err)
		return false
	}

	verified := sigObj.Verify([]byte(Hash256x2(message)), pubKey)
	return verified
}
