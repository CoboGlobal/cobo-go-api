package main

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

var localSigner = LocalSigner{
	"e7e73fabdd9edb8bddf947954c400a63bf93edc57abf170544ec570757df5453",
}
var client = Client{
	"0397ef0d81938bcf9587466ee33ab93caa77677416ada3297e70e92aa42245d99e",
	localSigner,
	"032f45930f652d72e0c90f71869dfe9af7d713b1f67dc2f7cb51f9572778b9c876",
	"https://api.sandbox.cobo.com",
}

func TestClient_GetAccountInfo(t *testing.T) {
	res := client.GetAccountInfo()
	fmt.Println(res)
}

func TestClient_GetCoinInfo(t *testing.T) {
	res := client.GetCoinInfo("ETH")
	fmt.Println(res)
}

func TestClient_NewDepositAddress(t *testing.T) {
	res := client.NewDepositAddress("BTC", false)
	fmt.Println(res)
}

func TestClient_BatchNewDepositAddress(t *testing.T) {
	res := client.BatchNewDepositAddress("BTC", 3, true)
	fmt.Println(res)
}

func TestClient_VerifyDepositAddress(t *testing.T) {
	res := client.VerifyDepositAddress("ETH", "0x05325e6f9d1f0437bd78a72c2ae084fbb8c039ee")
	fmt.Println(res)
}
func TestClient_BatchVerifyDepositAddress(t *testing.T) {
	res := client.BatchVerifyDepositAddress("ETH", "0x05325e6f9d1f0437bd78a72c2ae084fbb8c039ee,0x05325e6f9d1f0437bd78a72c2ae084fbb8c039e1")
	fmt.Println(res)
}

func TestClient_VerifyValidAddress(t *testing.T) {
	res := client.VerifyValidAddress("ETH", "0x05325e6f9d1f0437bd78a72c2ae084fbb8c039ee")
	fmt.Println(res)
}

func TestClient_GetAddressHistory(t *testing.T) {
	res := client.GetAddressHistory("ETH")
	fmt.Println(res)
}

func TestClient_CheckLoopAddressDetails(t *testing.T) {
	res := client.CheckLoopAddressDetails("ETH", "0xe7ebdc5bbb6c99cc8f7f2c1c83ff38aa6647f38a", "")
	fmt.Println(res)
}

func TestClient_VerifyLoopAddressList(t *testing.T) {
	res := client.VerifyLoopAddressList("ETH", "0xe7ebdc5bbb6c99cc8f7f2c1c83ff38aa6647f38a,0x05325e6f9d1f0437bd78a72c2ae084fbb8c039ee")
	fmt.Println(res)
}

func TestClient_GetTransactionDetails(t *testing.T) {
	res := client.GetTransactionDetails("20210422193807000343569000002370")
	fmt.Println(res)
}

func TestClient_GetTransactionsById(t *testing.T) {
	res := client.GetTransactionsById(map[string]string{})
	fmt.Println(res)
}

func TestClient_GetTransactionsByTime(t *testing.T) {
	res := client.GetTransactionsByTime(map[string]string{})
	fmt.Println(res)
}

func TestClient_GetPendingTransactions(t *testing.T) {
	res := client.GetPendingTransactions(map[string]string{})
	fmt.Println(res)
}

func TestClient_GetTransactionHistory(t *testing.T) {
	res := client.GetTransactionHistory(map[string]string{})
	fmt.Println(res)
}

func TestClient_GetPendingTransaction(t *testing.T) {
	res := client.GetPendingTransaction("20200604171238000354106000006405")
	fmt.Println(res)
}
func TestClient_Withdraw(t *testing.T) {
	res := client.Withdraw("TETH",
		"request_id_"+fmt.Sprintf("%d", time.Now().Unix()*1000),
		"0xb744adc8d75e115eec8e582eb5e8d60eb0972037",
		big.NewInt(1),
		map[string]string{})
	fmt.Println(res)
}

func TestClient_QueryWithdrawInfo(t *testing.T) {
	res := client.QueryWithdrawInfo("teth29374893624")
	fmt.Println(res)
}

func TestClient_GetStakingProductDetails(t *testing.T) {
	res := client.GetStakingProductDetails("159328", "zh")
	fmt.Println(res)
}

func TestClient_GetStakingProductList(t *testing.T) {
	res := client.GetStakingProductList("", "zh")
	fmt.Println(res)
}

func TestClient_Stake(t *testing.T) {
	res := client.Stake("159328", big.NewInt(1000000))
	fmt.Println(res)
}
func TestClient_Unstake(t *testing.T) {
	res := client.Unstake("159328", big.NewInt(1000000))
	fmt.Println(res)
}

func TestClient_GetStakings(t *testing.T) {
	res := client.GetStakings("DASH", "en")
	fmt.Println(res)
}

func TestClient_GetUnstakings(t *testing.T) {
	res := client.GetUnstakings("DASH")
	fmt.Println(res)
}

func TestClient_GetStakingHistory(t *testing.T) {
	res := client.GetStakingHistory()
	fmt.Println(res)
}
