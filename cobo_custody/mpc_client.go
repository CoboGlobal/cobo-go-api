package cobo_custody

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/btcsuite/btcd/btcec"
)

type MPCClient struct {
	Signer ApiSigner
	Env    Env
	Debug  bool
}

func (c MPCClient) GetMpcSupportedChains() (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	return c.Request("GET", "/v1/custody/mpc/get_supported_chains/", params)
}

func (c MPCClient) GetMpcSupportedCoins(chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
	}

	return c.Request("GET", "/v1/custody/mpc/get_supported_coins/", params)
}

func (c MPCClient) GetMpcMainAddress(chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
	}

	return c.Request("GET", "/v1/custody/mpc/get_main_address/", params)
}

func (c MPCClient) MpcBatchGenerateAddresses(chainCode string, count int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
		"count":      strconv.Itoa(count),
	}

	return c.Request("POST", "/v1/custody/mpc/generate_addresses/", params)
}

func (c MPCClient) GetMpcAddressList(chainCode, startId, endId string, limit, sortFlag int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
	}

	if startId != "" {
		params["start_id"] = startId
	}

	if endId != "" {
		params["end_id"] = endId
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	if sortFlag > 0 {
		params["sort"] = strconv.Itoa(sortFlag)
	}

	return c.Request("GET", "/v1/custody/mpc/list_addresses/", params)
}

func (c MPCClient) GetMpcBalance(address, chainCode, coin string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"address": address,
	}

	if chainCode != "" {
		params["chain_code"] = chainCode
	}

	if coin != "" {
		params["coin"] = coin
	}

	return c.Request("GET", "/v1/custody/mpc/get_balance/", params)
}

func (c MPCClient) ListMpcBalances(pageIndex, pageLength int, coin string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"page_index":  strconv.Itoa(pageIndex),
		"page_length": strconv.Itoa(pageLength),
	}

	if coin != "" {
		params["coin"] = coin
	}

	return c.Request("GET", "/v1/custody/mpc/list_balances/", params)
}

func (c MPCClient) GetMpcUnspentInputsList(address, coin string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"address": address,
		"coin":    coin,
	}

	return c.Request("GET", "/v1/custody/mpc/list_spendable/", params)
}

func (c MPCClient) MpcCreateTransaction(coin, requestId string, amount int, fromAddr, toAddr, toAddressDetails string,
	fee, gasPrice, gasLimit, operation int, extraParameters string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":       coin,
		"request_id": requestId,
	}

	if fromAddr != "" {
		params["from_address"] = fromAddr
	}

	if toAddr != "" {
		params["to_address"] = toAddr
	}

	if amount > 0 {
		params["amount"] = strconv.Itoa(amount)
	}

	if toAddressDetails != "" {
		params["to_address_details"] = toAddressDetails
	}

	if fee > 0 {
		params["fee"] = strconv.Itoa(fee)
	}

	if gasPrice > 0 {
		params["gas_price"] = strconv.Itoa(gasPrice)
	}

	if gasLimit > 0 {
		params["gas_limit"] = strconv.Itoa(gasLimit)
	}

	if operation >= 0 {
		params["operation"] = strconv.Itoa(operation)
	}

	if extraParameters != "" {
		params["extra_parameters"] = extraParameters
	}

	return c.Request("POST", "/v1/custody/mpc/create_transaction/", params)
}

func (c MPCClient) MpcDropTransaction(coboId string, gasPrice, gasLimit int, requestId string, fee int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"cobo_id":   coboId,
		"gas_price": strconv.Itoa(gasPrice),
	}

	if gasLimit > 0 {
		params["gas_limit"] = strconv.Itoa(gasLimit)
	}

	if requestId != "" {
		params["request_id"] = requestId
	}

	if fee > 0 {
		params["fee"] = strconv.Itoa(fee)
	}

	return c.Request("POST", "/v1/custody/mpc/drop_transaction/", params)
}

func (c MPCClient) MpcSpeedupTransaction(coboId, requestId string, gasPrice, gasLimit, fee int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"cobo_id":    coboId,
		"request_id": requestId,
	}

	if gasPrice > 0 {
		params["gas_price"] = strconv.Itoa(gasPrice)
	}

	if gasLimit > 0 {
		params["gas_limit"] = strconv.Itoa(gasLimit)
	}

	if fee > 0 {
		params["fee"] = strconv.Itoa(fee)
	}

	return c.Request("POST", "/v1/custody/mpc/speedup_transaction/", params)
}

func (c MPCClient) GetMpcTransactionsByRequestIds(requestIds string, status int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_ids": requestIds,
	}

	if status > 0 {
		params["status"] = strconv.Itoa(status)
	}

	return c.Request("GET", "/v1/custody/mpc/transactions_by_request_ids/", params)
}

func (c MPCClient) GetMpcTransactionsByCoboIds(coboIds string, status int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"cobo_ids": coboIds,
	}

	if status > 0 {
		params["status"] = strconv.Itoa(status)
	}

	return c.Request("GET", "/v1/custody/mpc/transactions_by_cobo_ids/", params)
}

func (c MPCClient) GetMpcTransactionsByTxHash(txHash string, transactionType int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"tx_hash": txHash,
	}

	if transactionType > 0 {
		params["transaction_type"] = strconv.Itoa(transactionType)
	}

	return c.Request("GET", "/v1/custody/mpc/transactions_by_tx_hash/", params)
}

func (c MPCClient) ListMpcWalletTransactions(startTime, endTime, status int, order string, transactionType int,
	coins, fromAddress, toAddress string, limit int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	if startTime > 0 {
		params["start_time"] = strconv.Itoa(startTime)
	}

	if endTime > 0 {
		params["end_time"] = strconv.Itoa(endTime)
	}

	if status > 0 {
		params["status"] = strconv.Itoa(status)
	}

	if order != "" {
		params["order"] = order
	}

	if transactionType > 0 {
		params["transaction_type"] = strconv.Itoa(transactionType)
	}

	if coins != "" {
		params["coins"] = coins
	}

	if fromAddress != "" {
		params["from_address"] = fromAddress
	}

	if toAddress != "" {
		params["to_address"] = toAddress
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	return c.Request("GET", "/v1/custody/mpc/list_transactions/", params)
}

func (c MPCClient) EstimateMpcFee(coin string, amount int, address string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"amount":  strconv.Itoa(amount),
		"address": address,
	}

	return c.Request("GET", "/v1/custody/mpc/estimate_fee/", params)
}

func (c MPCClient) ListMpcTssNodeRequests(requestType, status int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	if requestType > 0 {
		params["request_type"] = strconv.Itoa(requestType)
	}

	if status > 0 {
		params["status"] = strconv.Itoa(status)
	}

	return c.Request("GET", "/v1/custody/mpc/list_tss_node_requests/", params)
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
