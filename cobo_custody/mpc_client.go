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
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
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

func (c MPCClient) GetCoinInfo(coin string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin": coin,
	}

	return c.Request("GET", "/v1/custody/mpc/coin_info/", params)
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

func (c MPCClient) GenerateAddresses(chainCode string, count int, encoding *int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
		"count":      strconv.Itoa(count),
	}

	if encoding != nil {
		params["encoding"] = strconv.Itoa(*encoding)
	}

	return c.Request("POST", "/v1/custody/mpc/generate_addresses/", params)
}

func (c MPCClient) GenerateAddressMemo(chainCode string, address string, count int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"chain_code": chainCode,
		"address":    address,
		"count":      strconv.Itoa(count),
	}

	return c.Request("POST", "/v1/custody/mpc/generate_address_memo/", params)
}

func (c MPCClient) UpdateAddressDescription(coin string, address string, description string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":        coin,
		"address":     address,
		"description": description,
	}

	return c.Request("POST", "/v1/custody/mpc/update_address_description/", params)
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
	maxPriorityFee *big.Int, feeAmount *big.Int, remark string, autoFuel int, memo string) (*simplejson.Json, *ApiError) {
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

	if remark != "" {
		params["remark"] = remark
	}

	if autoFuel >= 0 {
		params["auto_fuel"] = strconv.Itoa(autoFuel)
	}

	if memo != "" {
		params["memo"] = memo
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

func (c MPCClient) DropTransaction(coboId, requestId string, fee *big.Float, gasPrice *big.Int,
	gasLimit *big.Int, feeAmount *big.Int, autoFuel int, extraParameters string) (*simplejson.Json, *ApiError) {
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

	if autoFuel >= 0 {
		params["auto_fuel"] = strconv.Itoa(autoFuel)
	}

	if extraParameters != "" {
		params["extra_parameters"] = extraParameters
	}

	return c.Request("POST", "/v1/custody/mpc/drop_transaction/", params)
}

func (c MPCClient) SpeedupTransaction(coboId, requestId string, fee *big.Float, gasPrice *big.Int,
	gasLimit *big.Int, feeAmount *big.Int, autoFuel int, extraParameters string) (*simplejson.Json, *ApiError) {
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

	if autoFuel >= 0 {
		params["auto_fuel"] = strconv.Itoa(autoFuel)
	}

	if extraParameters != "" {
		params["extra_parameters"] = extraParameters
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

	return c.Request("POST", "/v1/custody/mpc/estimate_fee/", params)
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

func (c MPCClient) ListTssNode() (*simplejson.Json, *ApiError) {
	var params = map[string]string{}
	return c.Request("GET", "/v1/custody/mpc/list_tss_node/", params)
}

func (c MPCClient) SignMessagesByRequestIds(requestIds string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_ids": requestIds,
	}

	return c.Request("GET", "/v1/custody/mpc/sign_messages_by_request_ids/", params)
}

func (c MPCClient) SignMessagesByCobotIds(CoboIds string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"cobo_ids": CoboIds,
	}

	return c.Request("GET", "/v1/custody/mpc/sign_messages_by_cobo_ids/", params)
}

func (c MPCClient) GetMaxSendAmount(coin string, feeRate big.Float, toAddress string, fromAddress string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":       coin,
		"fee_rate":   feeRate.String(),
		"to_address": toAddress,
	}

	if fromAddress != "" {
		params["from_address"] = fromAddress
	}

	return c.Request("GET", "/v1/custody/mpc/get_max_send_amount/", params)
}

func (c MPCClient) LockSpendable(coin string, txHash string, voutN int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"tx_hash": txHash,
		"vout_n":  strconv.Itoa(voutN),
	}

	return c.Request("POST", "/v1/custody/mpc/lock_spendable/", params)
}

func (c MPCClient) UnlockSpendable(coin string, txHash string, voutN int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"tx_hash": txHash,
		"vout_n":  strconv.Itoa(voutN),
	}

	return c.Request("POST", "/v1/custody/mpc/unlock_spendable/", params)
}

func (c MPCClient) GetRareSatoshis(coin string, txHash string, voutN int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"tx_hash": txHash,
		"vout_n":  strconv.Itoa(voutN),
	}

	return c.Request("GET", "/v1/custody/mpc/get_rare_satoshis/", params)
}

func (c MPCClient) GetUTXOAssets(coin string, txHash string, voutN int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"coin":    coin,
		"tx_hash": txHash,
		"vout_n":  strconv.Itoa(voutN),
	}

	return c.Request("GET", "/v1/custody/mpc/get_utxo_assets/", params)
}

func (c MPCClient) GetOrdinalsInscription(inscriptionId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"inscription_id": inscriptionId,
	}

	return c.Request("GET", "/v1/custody/mpc/get_ordinals_inscription/", params)
}

func (c MPCClient) BabylonPrepareStaking(requestId string, stakeInfo string, feeRate big.Float, maxStakingFee *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id": requestId,
		"stake_info": stakeInfo,
		"fee_rate":   feeRate.String(),
	}
	if maxStakingFee != nil {
		params["max_staking_fee"] = maxStakingFee.String()
	}

	return c.Request("POST", "/v1/custody/mpc/babylon/prepare_staking/", params)
}

func (c MPCClient) BabylonReplaceStakingFee(requestId string, relatedRequestId string, feeRate big.Float, maxStakingFee *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id":         requestId,
		"related_request_id": relatedRequestId,
		"fee_rate":           feeRate.String(),
	}
	if maxStakingFee != nil {
		params["max_staking_fee"] = maxStakingFee.String()
	}

	return c.Request("POST", "/v1/custody/mpc/babylon/replace_staking_fee/", params)
}

func (c MPCClient) BabylonDropStaking(requestId string, relatedRequestId string, feeRate big.Float, maxStakingFee *big.Int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id":         requestId,
		"related_request_id": relatedRequestId,
		"fee_rate":           feeRate.String(),
	}
	if maxStakingFee != nil {
		params["max_staking_fee"] = maxStakingFee.String()
	}

	return c.Request("POST", "/v1/custody/mpc/babylon/drop_staking/", params)
}

func (c MPCClient) BabylonUnbonding(requestId string, stakingRequestId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id":         requestId,
		"staking_request_id": stakingRequestId,
	}

	return c.Request("POST", "/v1/custody/mpc/babylon/unbonding/", params)
}

func (c MPCClient) BabylonWithdraw(requestId string, feeRate big.Float, maxFeeAmount *big.Int, unbondingRequestId *string, stakingRequestId *string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id": requestId,
		"fee_rate":   feeRate.String(),
	}
	if maxFeeAmount != nil {
		params["max_fee_amount"] = maxFeeAmount.String()
	}
	if unbondingRequestId != nil {
		params["unbonding_request_id"] = *unbondingRequestId
	}
	if stakingRequestId != nil {
		params["staking_request_id"] = *stakingRequestId
	}
	return c.Request("POST", "/v1/custody/mpc/babylon/withdraw/", params)
}

func (c MPCClient) BabylonBroadcastStakingTransaction(requestId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id": requestId,
	}

	return c.Request("POST", "/v1/custody/mpc/babylon/broadcast_staking_transaction/", params)
}

func (c MPCClient) BabylonBatchBroadcastStakingTransaction(requestIds []string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_ids": strings.Join(requestIds, ","),
	}

	return c.Request("POST", "/v1/custody/mpc/babylon/batch_broadcast_staking_transaction/", params)
}

func (c MPCClient) BabylonGetStakingInfo(requestId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id": requestId,
	}

	return c.Request("GET", "/v1/custody/mpc/babylon/get_staking_info/", params)
}

func (c MPCClient) BabylonListWaitingBroadcastTransactions(coin string, address string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"asset_coin": coin,
		"address":    address,
	}

	return c.Request("GET", "/v1/custody/mpc/babylon/list_waiting_broadcast_transactions/", params)
}

func (c MPCClient) BabylonListTransactionsByStatus(status int, address *string, minCoboId *string, limit *int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"status": strconv.Itoa(status),
	}
	if address != nil {
		params["address"] = *address
	}
	if minCoboId != nil {
		params["min_cobo_id"] = *minCoboId
	}
	if limit != nil {
		params["limit"] = strconv.Itoa(*limit)
	}

	return c.Request("GET", "/v1/custody/mpc/babylon/list_transactions_by_status/", params)
}

func (c MPCClient) GetApprovalDetails(requestId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"request_id": requestId,
	}

	return c.Request("GET", "/v1/custody/mpc/get_approval_details/", params)
}

func (c MPCClient) BabylonListEligibles(status string, minId string, limit *int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	if status != "" {
		params["status"] = status
	}

	if minId != "" {
		params["min_id"] = minId
	}

	if limit != nil {
		params["limit"] = strconv.Itoa(*limit)
	}

	return c.Request("GET", "/v1/custody/mpc/babylon/airdrops/list_eligibles/", params)
}

func (c MPCClient) BabylonSubmitRegistration(btcAddress string, babylonAddress string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"btc_address":     btcAddress,
		"babylon_address": babylonAddress,
	}

	return c.Request("POST", "/v1/custody/mpc/babylon/airdrops/submit_registration/", params)
}

func (c MPCClient) BabylonListRegistrations(status string, btcAddress string, minId string, limit *int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	if status != "" {
		params["status"] = status
	}

	if btcAddress != "" {
		params["btc_address"] = btcAddress
	}

	if minId != "" {
		params["min_id"] = minId
	}

	if limit != nil {
		params["limit"] = strconv.Itoa(*limit)
	}

	return c.Request("GET", "/v1/custody/mpc/babylon/airdrops/list_registrations/", params)
}

func (c MPCClient) BabylonGetRegistration(registrationId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"registration_id": registrationId,
	}

	return c.Request("GET", "/v1/custody/mpc/babylon/airdrops/get_registration/", params)
}

func (c MPCClient) BabylonListEligibleStakings(status string, minId string, limit *int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	if status != "" {
		params["status"] = status
	}

	if minId != "" {
		params["min_id"] = minId
	}

	if limit != nil {
		params["limit"] = strconv.Itoa(*limit)
	}

	return c.Request("GET", "/v1/custody/mpc/babylon/stakings/list_eligibles/", params)
}

func (c MPCClient) BabylonSubmitStakingRegistration(stakingId string, babylonAddress string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"staking_id":      stakingId,
		"babylon_address": babylonAddress,
	}

	return c.Request("POST", "/v1/custody/mpc/babylon/stakings/submit_registration/", params)
}

func (c MPCClient) BabylonListStakingRegistrations(stakingId string, status string, minId string, limit *int) (*simplejson.Json, *ApiError) {
	var params = map[string]string{}

	if stakingId != "" {
		params["staking_id"] = stakingId
	}

	if status != "" {
		params["status"] = status
	}

	if minId != "" {
		params["min_id"] = minId
	}

	if limit != nil {
		params["limit"] = strconv.Itoa(*limit)
	}

	return c.Request("GET", "/v1/custody/mpc/babylon/stakings/list_registrations/", params)
}

func (c MPCClient) BabylonGetStakingRegistration(registrationId string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"registration_id": registrationId,
	}

	return c.Request("GET", "/v1/custody/mpc/babylon/stakings/get_registration/", params)
}

func (c MPCClient) BabylonClaimBabylonRewards(babylonAddress string) (*simplejson.Json, *ApiError) {
	var params = map[string]string{
		"babylon_address": babylonAddress,
	}

	return c.Request("POST", "/v1/custody/mpc/babylon/claim_rewards/", params)
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
		return "", errors.New("response signature verify failed")
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
