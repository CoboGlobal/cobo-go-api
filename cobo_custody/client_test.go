package cobo_custody

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"
)

var env = flag.String("env", "Sandbox", "Env Config")
var secret = flag.String("secret", "Demo", "Api Secrect")

func GetEnv(env string) Env {
	if env == "Prod" {
		return Prod()
	}
	return Sandbox()
}
func GetData(env string) Config {
	if env == "Prod" {
		return ProdConfig()
	}
	return SandboxConfig()
}

var ConfigData Config
var client Client

func TestMain(m *testing.M) {
	flag.Parse()
	var localSigner = LocalSigner{
		PrivateKey: *secret,
	}
	ConfigData = GetData(*env)
	client = Client{
		Signer: localSigner,
		Env:    GetEnv(*env),
		Debug:  false,
	}

	m.Run()
}

func TestClient_GetAccountInfo(t *testing.T) {
	result, apiError := client.GetAccountInfo()
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetAccountInfo")
	fmt.Println(string(str))
}

func TestClient_GetValidCoinInfo(t *testing.T) {
	coins := [...]string{"ETH", "BTC", "ETH_USDT", "XRP"}
	for _, coin := range coins {
		result, apiError := client.GetCoinInfo(coin)
		if apiError != nil {
			t.Fail()
		}
		str, _ := result.Encode()
		fmt.Println("TestClient_GetValidCoinInfo, coin:", coin)
		fmt.Println(string(str))
	}
}

func TestClient_GetInvalidCoinInfo(t *testing.T) {
	coins := [...]string{"BTTB"}
	for _, coin := range coins {
		_, apiError := client.GetCoinInfo(coin)
		if apiError.ErrorCode != 12002 {
			t.Fail()
		}
		fmt.Println("TestClient_GetInvalidCoinInfo, coin:", coin)
	}
}

func TestClient_NewValidDepositAddress(t *testing.T) {
	coins := [...]string{"ETH", "BTC", "ETH_USDT", "XRP"}
	for _, coin := range coins {
		result, apiError := client.NewDepositAddress(coin, false)
		if apiError != nil {
			t.Fail()
		}
		str, _ := result.Encode()
		fmt.Println("TestClient_NewValidDepositAddress coin:", coin)
		fmt.Println(string(str))
	}
}

func TestClient_NewInvalidDepositAddress(t *testing.T) {
	coins := [...]string{"BTTB", "ETTE"}
	for _, coin := range coins {
		_, apiError := client.NewDepositAddress(coin, false)
		if apiError.ErrorCode != 12002 {
			t.Fail()
		}
		fmt.Println("TestClient_NewInvalidDepositAddress coin:", coin)
	}
}

func TestClient_BatchValidNewDepositAddress(t *testing.T) {
	coins := [...]string{"ETH", "BTC"}
	for _, coin := range coins {
		result, apiError := client.BatchNewDepositAddress(coin, 2, true)
		if apiError != nil {
			t.Fail()
		}
		str, _ := result.Encode()
		fmt.Println("TestClient_BatchValidNewDepositAddress coin:", coin)
		fmt.Println(string(str))
	}
}

func TestClient_BatchInvalidNewDepositAddress(t *testing.T) {
	coins := [...]string{"BTTB", "ETTE"}
	for _, coin := range coins {
		_, apiError := client.BatchNewDepositAddress(coin, 2, true)
		if apiError.ErrorCode != 12002 {
			t.Fail()
		}
		fmt.Println("TestClient_BatchInvalidNewDepositAddress coin:", coin)
	}
}

func TestClient_VerifyDepositAddress(t *testing.T) {
	for coin, address := range ConfigData.DeAddress {
		fmt.Println(coin, address)
		result, apiError := client.VerifyDepositAddress(coin, address)
		if apiError != nil {
			t.Fail()
		}
		str, _ := result.Encode()
		fmt.Println("TestClient_VerifyDepositAddress coin:", coin)
		fmt.Println(string(str))
	}
}

func TestClient_BatchVerifyDepositAddress(t *testing.T) {
	for coin, addresses := range ConfigData.DeAddresses {
		result, apiError := client.BatchVerifyDepositAddress(coin, addresses)
		if apiError != nil {
			t.Fail()
		}
		str, _ := result.Encode()
		fmt.Println("TestClient_BatchVerifyDepositAddress")
		fmt.Println(string(str))
	}
}

func TestClient_VerifyValidAddress(t *testing.T) {
	result, apiError := client.VerifyValidAddress("ETH", "0x05325e6f9d1f0437bd78a72c2ae084fbb8c039ee")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_VerifyValidAddress")
	fmt.Println(string(str))
}

func TestClient_GetAddressHistory(t *testing.T) {
	coins := [...]string{"ETH", "BTC"}
	for _, coin := range coins {
		result, apiError := client.GetAddressHistory(coin)
		if apiError != nil {
			t.Fail()
		}
		str, _ := result.Encode()
		fmt.Println("TestClient_GetAddressHistory coin:", coin)
		fmt.Println(string(str))
	}
}

func TestClient_CheckLoopAddressDetails(t *testing.T) {
	for coin, address := range ConfigData.LpAddress {
		var memo = ""
		if strings.Contains(address, "|") {
			string_slice := strings.Split(address, "|")
			address = string_slice[0]
			memo = string_slice[1]

		}
		result, apiError := client.CheckLoopAddressDetails(coin, address, memo)
		if apiError != nil {
			t.Fail()
		}
		str, _ := result.Encode()
		fmt.Println("TestClient_CheckLoopAddressDetails")
		fmt.Println(string(str))
	}
}

func TestClient_VerifyLoopAddressList(t *testing.T) {
	for coin, addresses := range ConfigData.LpAddresses {
		result, apiError := client.VerifyLoopAddressList(coin, addresses)
		if apiError != nil {
			t.Fail()
		}
		str, _ := result.Encode()
		fmt.Println("TestClient_VerifyLoopAddressList")
		fmt.Println(string(str))
	}
}

func TestClient_GetTransactionDetails(t *testing.T) {
	result, apiError := client.GetTransactionDetails(ConfigData.CoboId)
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetTransactionDetails")
	fmt.Println(string(str))
}

func TestClient_GetTransactionsById(t *testing.T) {
	result, apiError := client.GetTransactionsById(map[string]string{})
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetTransactionsById")
	fmt.Println(string(str))
}

func TestClient_GetTransactionsByTxid(t *testing.T) {
	result, apiError := client.GetTransactionsByTxid(ConfigData.TxId)
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetTransactionsByTxid")
	fmt.Println(string(str))
}

func TestClient_GetTransactionsByTime(t *testing.T) {
	result, apiError := client.GetTransactionsByTime(map[string]string{})
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetTransactionsByTime")
	fmt.Println(string(str))
}

func TestClient_GetPendingTransactions(t *testing.T) {
	result, apiError := client.GetPendingTransactions(map[string]string{})
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetPendingTransactions")
	fmt.Println(string(str))
}

func TestClient_GetTransactionHistory(t *testing.T) {
	result, apiError := client.GetTransactionHistory(map[string]string{})
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetTransactionHistory")
	fmt.Println(string(str))
}

func TestClient_GetPendingTransaction(t *testing.T) {
	result, apiError := client.GetPendingTransaction("20211214231857000374360000005692")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_Withdraw(t *testing.T) {
	for coin, address := range ConfigData.Withdraw {
		hashResult := sha256.Sum256([]byte(address))
		var requestId = fmt.Sprintf("sdk_request_id_%s_%d", fmt.Sprintf("%x", hashResult)[0:8], time.Now().Unix()*1000)
		var options = make(map[string]string)
		if strings.Contains(address, "|") {
			string_slice := strings.Split(address, "|")
			address = string_slice[0]
			options["memo"] = string_slice[1]

		}
		result, apiError := client.Withdraw(
			coin,
			requestId,
			address,
			big.NewInt(1),
			options)
		if apiError != nil {
			t.Fail()
		}
		str, _ := result.Encode()
		fmt.Println("TestClient_Withdraw coin:", coin)
		fmt.Println(string(str))
	}
}

func TestClient_QueryWithdrawInfo(t *testing.T) {
	result, apiError := client.QueryWithdrawInfo(ConfigData.WithdrawId)
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_QueryWithdrawInfo")
	fmt.Println(string(str))
}

func TestClient_GetStakingProductDetails(t *testing.T) {
	re, err := client.GetStakingProductList("", "zh")
	if err != nil {
		t.Fail()
	}
	str, _ := re.Encode()
	var jsonSlice []map[string]interface{}
	json.Unmarshal(str, &jsonSlice)
	var product_id = strings.Split(fmt.Sprintf("%f", jsonSlice[0]["product_id"]), ".")[0]
	result, apiError := client.GetStakingProductDetails(product_id, "zh")
	if apiError != nil {
		t.Fail()
	}
	strDetail, _ := result.Encode()
	fmt.Println("TestClient_GetStakingProductDetails")
	fmt.Println(string(strDetail))
}

func TestClient_GetStakingProductList(t *testing.T) {
	result, apiError := client.GetStakingProductList("", "zh")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetStakingProductList")
	fmt.Println(string(str))
}

func TestClient_Stake(t *testing.T) {
	re, err := client.GetStakingProductList("TETH", "zh")
	if err != nil {
		t.Fail()
	}
	str, _ := re.Encode()
	var jsonSlice []map[string]interface{}
	json.Unmarshal(str, &jsonSlice)
	if len(jsonSlice) == 0 {
		t.Skip("no TETH staking product")
	}
	var product_id = strings.Split(fmt.Sprintf("%f", jsonSlice[0]["product_id"]), ".")[0]
	result, apiError := client.Stake(product_id, big.NewInt(1000000000000000000))
	if apiError != nil {
		fmt.Println(apiError.ErrorId)
		fmt.Println(apiError.ErrorCode)
		fmt.Println(apiError.ErrorMessage)
	} else {
		str, _ := result.Encode()
		fmt.Println("TestClient_Stake")
		fmt.Println(string(str))
	}
}

func TestClient_Unstake(t *testing.T) {
	re, err := client.GetStakingProductList("TETH", "zh")
	if err != nil {
		t.Fail()
	}
	str, _ := re.Encode()
	var jsonSlice []map[string]interface{}
	json.Unmarshal(str, &jsonSlice)
	if len(jsonSlice) == 0 {
		t.Skip("no TETH staking product")
	}
	var product_id = strings.Split(fmt.Sprintf("%f", jsonSlice[0]["product_id"]), ".")[0]
	result, apiError := client.Unstake(product_id, big.NewInt(1000000000000000000))
	if apiError != nil {
		fmt.Println(apiError.ErrorId)
		fmt.Println(apiError.ErrorCode)
		fmt.Println(apiError.ErrorMessage)
	} else {
		str, _ := result.Encode()
		fmt.Println("TestClient_Unstake")
		fmt.Println(string(str))
	}
}

func TestClient_GetStakings(t *testing.T) {
	result, apiError := client.GetStakings("TETH", "en")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetStakings")
	fmt.Println(string(str))
}

func TestClient_GetUnstakings(t *testing.T) {
	result, apiError := client.GetUnstakings("TETH s")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetUnstakings")
	fmt.Println(string(str))
}

func TestClient_GetStakingHistory(t *testing.T) {
	result, apiError := client.GetStakingHistory()
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetStakingHistory")
	fmt.Println(string(str))
}

func Test_GenerateKeyPair(*testing.T) {
	apiSecret, apiKey := GenerateKeyPair()
	println("API_SECRET:", apiSecret)
	println("API_KEY:", apiKey)
}
