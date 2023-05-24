package cobo_custody

import (
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

type MPCClient struct {
	Signer ApiSigner
	Env    Env
	Debug  bool
}

func (c MPCClient) GetSupportedChains() (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	return c.Request("GET", "/v1/custody/mpc/get_supported_chains/", params)
}

func (c MPCClient) GetSupportedCoins(chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
	}

	return c.Request("GET", "/v1/custody/mpc/get_supported_coins/", params)
}

func (c MPCClient) GetSupportedNftCollections(chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
	}

	return c.Request("GET", "/v1/custody/mpc/get_supported_nft_collections/", params)
}

func (c MPCClient) GetWalletSupportedCoins() (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	return c.Request("GET", "/v1/custody/mpc/get_wallet_supported_coins/", params)
}

func (c MPCClient) IsValidAddress(coin string, address string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"address": address,
	}

	return c.Request("GET", "/v1/custody/mpc/is_valid_address/", params)
}

func (c MPCClient) GetMainAddress(chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
	}

	return c.Request("GET", "/v1/custody/mpc/get_main_address/", params)
}

func (c MPCClient) GenerateAddresses(chainCode string, count int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
		"count":      strconv.Itoa(count),
	}

	return c.Request("POST", "/v1/custody/mpc/generate_addresses/", params)
}

func (c MPCClient) ListAddresses(chainCode, startId, endId string, limit, sortFlag int) (*simplejson.Json, *ApiError) {
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

func (c MPCClient) GetBalance(address, chainCode, coin string) (*simplejson.Json, *ApiError) {
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

func (c MPCClient) ListBalances(pageIndex, pageLength int, coin string, chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"page_index":  strconv.Itoa(pageIndex),
		"page_length": strconv.Itoa(pageLength),
	}

	if coin != "" {
		params["coin"] = coin
	}
	if chainCode != "" {
		params["chain_code"] = chainCode
	}

	return c.Request("GET", "/v1/custody/mpc/list_balances/", params)
}

func (c MPCClient) ListSpendable(address, coin string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"address": address,
		"coin":    coin,
	}

	return c.Request("GET", "/v1/custody/mpc/list_spendable/", params)
}

func (c MPCClient) CreateTransaction(coin, requestId string, amount *big.Int, fromAddr, toAddr, toAddressDetails string,
	fee *big.Float, gasPrice *big.Int, gasLimit *big.Int, operation int, extraParameters string, maxFee *big.Int,
	maxPriorityFee *big.Int, feeAmount *big.Int) (*simplejson.Json, *ApiError) {
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

	if amount != nil {
		params["amount"] = amount.String()
	}

	if toAddressDetails != "" {
		params["to_address_details"] = toAddressDetails
	}

	if fee != nil {
		params["fee"] = fee.String()
	}

	if gasPrice != nil {
		params["gas_price"] = gasPrice.String()
	}

	if gasLimit != nil {
		params["gas_limit"] = gasLimit.String()
	}

	if operation >= 0 {
		params["operation"] = strconv.Itoa(operation)
	}

	if extraParameters != "" {
		params["extra_parameters"] = extraParameters
	}

	if maxFee != nil {
		params["max_fee"] = maxFee.String()
	}

	if maxPriorityFee != nil {
		params["max_priority_fee"] = maxPriorityFee.String()
	}

	if feeAmount != nil {
		params["fee_amount"] = feeAmount.String()
	}

	return c.Request("POST", "/v1/custody/mpc/create_transaction/", params)
}

func (c MPCClient) SignMessage(chainCode, requestId, fromAddr string, signVersion int, extraParameters string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code":       chainCode,
		"request_id":       requestId,
		"from_address":     fromAddr,
		"sign_version":     strconv.Itoa(signVersion),
		"extra_parameters": extraParameters,
	}

	return c.Request("POST", "/v1/custody/mpc/sign_message/", params)
}

func (c MPCClient) DropTransaction(coboId, requestId string, fee *big.Float, gasPrice *big.Int, gasLimit *big.Int, feeAmount *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"cobo_id":    coboId,
		"request_id": requestId,
	}

	if fee != nil {
		params["fee"] = fee.String()
	}

	if gasLimit != nil {
		params["gas_limit"] = gasLimit.String()
	}

	if gasPrice != nil {
		params["gas_price"] = gasPrice.String()
	}

	if feeAmount != nil {
		params["fee_amount"] = feeAmount.String()
	}

	return c.Request("POST", "/v1/custody/mpc/drop_transaction/", params)
}

func (c MPCClient) SpeedupTransaction(coboId, requestId string, fee *big.Float, gasPrice *big.Int, gasLimit *big.Int, feeAmount *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"cobo_id":    coboId,
		"request_id": requestId,
	}

	if fee != nil {
		params["fee"] = fee.String()
	}

	if gasLimit != nil {
		params["gas_limit"] = gasLimit.String()
	}

	if gasPrice != nil {
		params["gas_price"] = gasPrice.String()
	}

	if feeAmount != nil {
		params["fee_amount"] = feeAmount.String()
	}

	return c.Request("POST", "/v1/custody/mpc/speedup_transaction/", params)
}

func (c MPCClient) TransactionsByRequestIds(requestIds string, status int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_ids": requestIds,
	}

	if status > 0 {
		params["status"] = strconv.Itoa(status)
	}

	return c.Request("GET", "/v1/custody/mpc/transactions_by_request_ids/", params)
}

func (c MPCClient) TransactionsByCoboIds(coboIds string, status int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"cobo_ids": coboIds,
	}

	if status > 0 {
		params["status"] = strconv.Itoa(status)
	}

	return c.Request("GET", "/v1/custody/mpc/transactions_by_cobo_ids/", params)
}

func (c MPCClient) TransactionsByTxHash(txHash string, transactionType int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"tx_hash": txHash,
	}

	if transactionType > 0 {
		params["transaction_type"] = strconv.Itoa(transactionType)
	}

	return c.Request("GET", "/v1/custody/mpc/transactions_by_tx_hash/", params)
}

func (c MPCClient) ListTransactions(startTime, endTime, status int, orderBy, order string, transactionType int,
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

	if orderBy != "" {
		params["order_by"] = orderBy
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

func (c MPCClient) EstimateFee(coin string, amount *big.Int, address string, replace_cobo_id string, from_address string,
	to_address_details string, fee *big.Float, gasPrice *big.Int, gasLimit *big.Int, extra_parameters string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin": coin,
	}

	if amount != nil {
		params["amount"] = amount.String()
	}
	if address != "" {
		params["address"] = address
	}
	if replace_cobo_id != "" {
		params["replace_cobo_id"] = replace_cobo_id
	}
	if from_address != "" {
		params["from_address"] = from_address
	}
	if to_address_details != "" {
		params["to_address_details"] = to_address_details
	}
	if fee != nil {
		params["fee"] = fee.String()
	}
	if gasPrice != nil {
		params["gas_price"] = gasPrice.String()
	}
	if gasLimit != nil {
		params["gas_limit"] = gasLimit.String()
	}
	if extra_parameters != "" {
		params["extra_parameters"] = extra_parameters
	}

	return c.Request("GET", "/v1/custody/mpc/estimate_fee/", params)
}

func (c MPCClient) ListTssNodeRequests(requestType, status int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	if requestType > 0 {
		params["request_type"] = strconv.Itoa(requestType)
	}

	if status > 0 {
		params["status"] = strconv.Itoa(status)
	}

	return c.Request("GET", "/v1/custody/mpc/list_tss_node_requests/", params)
}

func (c MPCClient) RetryDoubleCheck(requestId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id": requestId,
	}

	return c.Request("POST", "/v1/custody/mpc/retry_double_check/", params)
}

func (c MPCClient) request(method string, path string, params map[string]string) (string, error) {
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

func (c MPCClient) Request(method string, path string, params map[string]string) (*simplejson.Json, *ApiError) {
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

func (c MPCClient) VerifyEcc(message string, signature string) bool {
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
