package cobo_custody

import (
	"encoding/hex"
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

type Web3Client struct {
	Signer ApiSigner
	Env    Env
	Debug  bool
}

func (c Web3Client) GetWeb3SupportedChains() (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	return c.Request("GET", "/v1/custody/web3_supported_chains/", params)
}

func (c Web3Client) GetWeb3SupportedCoins(chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
	}

	return c.Request("GET", "/v1/custody/web3_supported_coins/", params)
}

func (c Web3Client) GetWeb3SupportedNftCollections() (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	return c.Request("GET", "/v1/custody/web3_supported_nft_collections/", params)
}

func (c Web3Client) GetWeb3SupportedContracts(chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
	}

	return c.Request("GET", "/v1/custody/web3_supported_contracts/", params)
}

func (c Web3Client) GetWeb3SupportedContractMethods(chainCode string, contractAddress string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code":       chainCode,
		"contract_address": contractAddress,
	}

	return c.Request("GET", "/v1/custody/web3_supported_contract_methods/", params)
}

func (c Web3Client) BatchWeb3NewAddress(chainCode string, count int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
		"count":      strconv.Itoa(count),
	}

	return c.Request("POST", "/v1/custody/web3_add_addresses/", params)
}

func (c Web3Client) GetWeb3AddressList(chainCode string, pageIndex int, pageLength int, sortFlag int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code":  chainCode,
		"page_index":  strconv.Itoa(pageIndex),
		"page_length": strconv.Itoa(pageLength),
	}
	if sortFlag > 0 {
		params["sort_flag"] = strconv.Itoa(sortFlag)
	}

	return c.Request("GET", "/v1/custody/web3_list_wallet_address/", params)
}

func (c Web3Client) GetWeb3WalletAssetList(address string, chainCode string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{}
	if address != "" {
		params["address"] = address
	}
	if chainCode != "" {
		params["chain_code"] = chainCode
	}

	return c.Request("GET", "/v1/custody/web3_list_wallet_assets/", params)
}

func (c Web3Client) GetWeb3WalletNftList(nftCode string, address string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"nft_code": nftCode,
	}
	if address != "" {
		params["address"] = address
	}

	return c.Request("GET", "/v1/custody/web3_list_wallet_nfts/", params)
}

func (c Web3Client) GetWeb3WalletNftDetail(nftCode string, tokenId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"nft_code": nftCode,
		"token_id": tokenId,
	}

	return c.Request("GET", "/v1/custody/web3_wallet_nft_detail/", params)
}

func (c Web3Client) Web3Withdraw(coin string, requestId string, fromAddr string, toAddr string, amount *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":       coin,
		"request_id": requestId,
		"from_addr":  fromAddr,
		"to_addr":    toAddr,
	}
	if amount != nil {
		params["amount"] = amount.String()
	}

	return c.Request("POST", "/v1/custody/web3_withdraw/", params)
}

func (c Web3Client) GetWeb3WithdrawTransaction(requestId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id": requestId,
	}

	return c.Request("GET", "/v1/custody/web3_get_withdraw_transaction/", params)
}

func (c Web3Client) Web3Contract(chainCode string, requestId string, walletAddr string, contractAddr string, methodId string,
	methodName string, args string, amount *big.Int, gasLimit *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code":    chainCode,
		"request_id":    requestId,
		"wallet_addr":   walletAddr,
		"contract_addr": contractAddr,
		"method_id":     methodId,
		"method_name":   methodName,
		"args":          args,
	}
	if amount != nil {
		params["amount"] = amount.String()
	}
	if gasLimit != nil {
		params["gas_limit"] = gasLimit.String()
	}

	return c.Request("POST", "/v1/custody/web3_contract/", params)
}

func (c Web3Client) GetWeb3ContractTransaction(requestId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id": requestId,
	}

	return c.Request("GET", "/v1/custody/web3_get_contract_transaction/", params)
}

func (c Web3Client) ListWeb3WalletTransactions(address string, chainCode string, maxId string, minId string, limit int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"address": address,
	}

	if chainCode != "" {
		params["chain_code"] = chainCode
	}
	if maxId != "" {
		params["max_id"] = maxId
	}
	if minId != "" {
		params["min_id"] = minId
	}
	if limit != 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	return c.Request("GET", "/v1/custody/web3_list_wallet_transactions/", params)
}

func (c Web3Client) request(method string, path string, params map[string]string) string {
	httpClient := &http.Client{}
	nonce := fmt.Sprintf("%d", time.Now().UnixMicro())
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

func (c Web3Client) Request(method string, path string, params map[string]string) (*simplejson.Json, *ApiError) {
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

func (c Web3Client) VerifyEcc(message string, signature string) bool {
	pubKeyBytes, _ := hex.DecodeString(c.Env.CoboPub)
	pubKey, _ := btcec.ParsePubKey(pubKeyBytes, btcec.S256())

	sigBytes, _ := hex.DecodeString(signature)
	sigObj, _ := btcec.ParseSignature(sigBytes, btcec.S256())

	verified := sigObj.Verify([]byte(Hash256x2(message)), pubKey)
	return verified
}
