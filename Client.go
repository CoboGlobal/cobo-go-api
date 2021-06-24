package main

import (
	"encoding/hex"
	"fmt"
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

type Client struct {
	apiKey  string
	signer  ApiSigner
	coboPub string
	host    string
}

func (c Client) GetAccountInfo() string {
	return c.Request("GET", "/v1/custody/org_info/", map[string]string{})
}

func (c Client) GetCoinInfo(coin string) string {
	return c.Request("GET", "/v1/custody/coin_info/", map[string]string{
		"coin": coin,
	})
}

func (c Client) NewDepositAddress(coin string, nativeSegwit bool) string {
	var params = map[string]string{
		"coin": coin,
	}
	if nativeSegwit {
		params["native_segwit"] = "true"
	}
	return c.Request("POST", "/v1/custody/new_address/", params)
}

func (c Client) BatchNewDepositAddress(coin string, count int,nativeSegwit bool) string {
	var params = map[string]string{
		"coin":  coin,
		"count": strconv.Itoa(count),
	}
	if nativeSegwit {
		params["native_segwit"] = "true"
	}
	return c.Request("POST", "/v1/custody/new_addresses/", params)
}

func (c Client) VerifyDepositAddress(coin string, address string) string {
	var params = map[string]string{
		"coin":  coin,
		"address": address,
	}
	return c.Request("GET", "/v1/custody/address_info/", params)
}

func (c Client) BatchVerifyDepositAddress(coin string, addresses string) string {
	var params = map[string]string{
		"coin":  coin,
		"address": addresses,
	}
	return c.Request("GET", "/v1/custody/addresses_info/", params)
}

func (c Client) VerifyValidAddress(coin string, addresses string) string {
	var params = map[string]string{
		"coin":  coin,
		"address": addresses,
	}
	return c.Request("GET", "/v1/custody/is_valid_address/", params)
}

func (c Client) GetAddressHistory(coin string) string {
	var params = map[string]string{
		"coin":  coin,
	}
	return c.Request("GET", "/v1/custody/address_history/", params)
}

func (c Client) CheckLoopAddressDetails(coin string, address string, memo string) string {
	var params = map[string]string{
		"coin":  coin,
		"address": address,
	}
	if memo !="" {
		params["memo"]=memo
	}
	return c.Request("GET", "/v1/custody/internal_address_info/", params)
}

func (c Client) VerifyLoopAddressList(coin string, addresses string) string {
	var params = map[string]string{
		"coin":  coin,
		"address": addresses,
	}

	return c.Request("GET", "/v1/custody/internal_address_info_batch/", params)
}

func (c Client) GetTransactionDetails(txId string,) string {
	var params = map[string]string{
		"id":  txId,
	}

	return c.Request("GET", "/v1/custody/transaction/", params)
}

func (c Client) GetTransactionsById(params map[string]string) string {
	return c.Request("GET", "/v1/custody/transactions_by_id/", params)
}

func (c Client) GetTransactionsByTime(params map[string]string) string {
	return c.Request("GET", "/v1/custody/transactions_by_time/", params)
}

func (c Client) GetPendingTransactions(params map[string]string) string {
	return c.Request("GET", "/v1/custody/pending_transactions/", params)
}

func (c Client) GetPendingTransaction(id string) string {
	return c.Request("GET", "/v1/custody/pending_transactions/", map[string]string {
		"id":  id,
	})
}

func (c Client) GetTransactionHistory(params map[string]string) string {
	return c.Request("GET", "/v1/custody/transaction_history/", params)
}

func (c Client) Withdraw(coin string, requestId string, address string, amount *big.Int, options map[string]string) string {
	var params =  map[string]string{
		"coin":  coin,
		"request_id":requestId,
		"address":address,
		"amount":amount.String(),
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

func (c Client) QueryWithdrawInfo(requestId string) string {
	return c.Request("GET", "/v1/custody/withdraw_info_by_request_id/", map[string]string{"request_id":requestId})
}

func (c Client) GetStakingProductDetails(productId string, language string) string{
	var params =  map[string]string{
		"product_id":  productId,
		"language":language,
	}
    return c.Request("GET", "/v1/custody/staking_product/", params)
}

func (c Client) GetStakingProductList(coin string, language string) string {
	var params =map[string]string{
		"language":language,
	}
	if coin !="" {
		params["coin"]=coin
	}

	return c.Request("GET", "/v1/custody/staking_products/",params)
}

func (c Client) Stake(productId string, amount *big.Int) string {
	var params = map[string]string {
		"product_id":productId,
		"amount":amount.String(),
	}
	return c.Request("POST", "/v1/custody/staking_stake/",params)
}

func (c Client) Unstake(productId string, amount *big.Int) string {
	var params = map[string]string {
		"product_id":productId,
		"amount":amount.String(),
	}
	return c.Request("POST", "/v1/custody/staking_unstake/",params)
}

func (c Client) GetStakings(coin string, language string) string {
	var params =map[string]string{
		"language":language,
	}
	if coin !="" {
		params["coin"]=coin
	}
	return c.Request("GET", "/v1/custody/stakings/",params)
}

func (c Client) GetUnstakings(coin string)string {
	var params =map[string]string{}
	if coin !="" {
		params["coin"]=coin
	}
	return c.Request("GET", "/v1/custody/unstakings/", params)
}

func (c Client) GetStakingHistory() string {
	return c.Request("GET", "/v1/custody/staking_history/", map[string]string{})
}


func (c Client) Request(method string, path string, params map[string]string) string {
	httpClient := &http.Client{}
	nonce := fmt.Sprintf("%d", time.Now().Unix()*1000)
	sorted := SortParams(params)
	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest(method, c.host+path, strings.NewReader(sorted))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, _ = http.NewRequest(method, c.host+path+"?"+sorted, strings.NewReader(""))
	}
	content := strings.Join([]string{method, path, nonce, sorted}, "|")

	req.Header.Set("Biz-Api-Key", c.apiKey)
	req.Header.Set("Biz-Api-Nonce", nonce)
	req.Header.Set("Biz-Api-Signature", c.signer.Sign(content))

	resp, _ := httpClient.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	timestamp := resp.Header["Biz-Timestamp"][0]
	signature := resp.Header["Biz-Resp-Signature"][0]
	success := c.VerifyEcc(string(body)+"|"+timestamp, signature)
	fmt.Println("verify success?", success)
	return string(body)
}

func SortParams(params map[string]string) string {
	keys := make([]string, len(params))
	i := 0
	for k, _ := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	sorted := make([]string, len(params))
	i = 0
	for _, k := range keys {
		sorted[i] = k + "=" + url.QueryEscape(params[k])
		i++
	}
	return strings.Join(sorted, "&")
}

func (c Client) VerifyEcc(message string, signature string) bool {
	pubKeyBytes, _ := hex.DecodeString(c.coboPub)
	pubKey, _ := btcec.ParsePubKey(pubKeyBytes, btcec.S256())

	sigBytes, _ := hex.DecodeString(signature)
	sigObj, _ := btcec.ParseSignature(sigBytes, btcec.S256())

	verified := sigObj.Verify([]byte(Hash256x2(message)), pubKey)
	return verified
}
