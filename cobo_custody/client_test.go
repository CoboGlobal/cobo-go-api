package cobo_custody

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

var localSigner = LocalSigner{
	PrivateKey: "e7e73fabdd9edb8bddf947954c400a63bf93edc57abf170544ec570757df5453",
}
var client = Client{
	Signer:  localSigner,
	Env: Sandbox(),
	Debug: false,
}

func TestClient_GetAccountInfo(t *testing.T) {
	result, apiError := client.GetAccountInfo()
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetCoinInfo(t *testing.T) {
	result, apiError := client.GetCoinInfo("ETH")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println(string(str))

}

func TestClient_NewDepositAddress(t *testing.T) {
	result, apiError := client.NewDepositAddress("BTC", false)
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_BatchNewDepositAddress(t *testing.T) {
	result, apiError := client.BatchNewDepositAddress("BTC", 3, true)
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_VerifyDepositAddress(t *testing.T) {
	result, apiError := client.VerifyDepositAddress("ETH", "0x05325e6f9d1f0437bd78a72c2ae084fbb8c039ee")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}
func TestClient_BatchVerifyDepositAddress(t *testing.T) {
	result, apiError := client.BatchVerifyDepositAddress("ETH", "0x05325e6f9d1f0437bd78a72c2ae084fbb8c039ee,0x05325e6f9d1f0437bd78a72c2ae084fbb8c039e1")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_VerifyValidAddress(t *testing.T) {
	result, apiError := client.VerifyValidAddress("ETH", "0x05325e6f9d1f0437bd78a72c2ae084fbb8c039ee")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetAddressHistory(t *testing.T) {
	result, apiError := client.GetAddressHistory("ETH")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_CheckLoopAddressDetails(t *testing.T) {
	result, apiError := client.CheckLoopAddressDetails("ETH", "0xe7ebdc5bbb6c99cc8f7f2c1c83ff38aa6647f38a", "")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_VerifyLoopAddressList(t *testing.T) {
	result, apiError := client.VerifyLoopAddressList("ETH", "0xe7ebdc5bbb6c99cc8f7f2c1c83ff38aa6647f38a,0x05325e6f9d1f0437bd78a72c2ae084fbb8c039ee")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetTransactionDetails(t *testing.T) {
	result, apiError := client.GetTransactionDetails("20210422193807000343569000002370")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetTransactionsById(t *testing.T) {
	result, apiError := client.GetTransactionsById(map[string]string{})
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetTransactionsByTime(t *testing.T) {
	result, apiError := client.GetTransactionsByTime(map[string]string{
		"coin": "ETH",
	})
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetPendingTransactions(t *testing.T) {
	result, apiError := client.GetPendingTransactions(map[string]string{
		"coin": "ETH",
	})
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetTransactionHistory(t *testing.T) {
	result, apiError := client.GetTransactionHistory(map[string]string{
		"coin": "ETH",
	})
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetPendingTransaction(t *testing.T) {
	result, apiError := client.GetPendingTransaction("20200604171238000354106000006405")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}
func TestClient_Withdraw(t *testing.T) {
	result, apiError := client.Withdraw("TETH",
		"request_id_"+fmt.Sprintf("%d", time.Now().Unix()*1000),
		"0xb744adc8d75e115eec8e582eb5e8d60eb0972037",
		big.NewInt(1),
		map[string]string{})
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_QueryWithdrawInfo(t *testing.T) {
	result, apiError := client.QueryWithdrawInfo("teth29374893624")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetStakingProductDetails(t *testing.T) {
	result, apiError := client.GetStakingProductDetails("159328", "zh")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetStakingProductList(t *testing.T) {
	result, apiError := client.GetStakingProductList("", "zh")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_Stake(t *testing.T) {
	result, apiError := client.Stake("159328", big.NewInt(1000000))
	if apiError != nil {
		fmt.Println(apiError.ErrorId)
		fmt.Println(apiError.ErrorCode)
		fmt.Println(apiError.ErrorMessage)
	} else {
		str,_ := result.Encode()
		fmt.Println(string(str))
	}
}
func TestClient_Unstake(t *testing.T) {
	result, apiError := client.Unstake("159328", big.NewInt(1000000))
	if apiError != nil {
		fmt.Println(apiError.ErrorId)
		fmt.Println(apiError.ErrorCode)
		fmt.Println(apiError.ErrorMessage)
	} else {
		str,_ := result.Encode()
		fmt.Println(string(str))
	}
}

func TestClient_GetStakings(t *testing.T) {
	result, apiError := client.GetStakings("DASH", "en")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetUnstakings(t *testing.T) {
	result, apiError := client.GetUnstakings("DASH")
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func TestClient_GetStakingHistory(t *testing.T) {
	result, apiError := client.GetStakingHistory()
	if apiError != nil {
		t.Fail()
	}
	str,_ := result.Encode()
	fmt.Println(string(str))
}

func Test_GenerateKeyPair(*testing.T)  {
	apiSecret, apiKey := GenerateKeyPair()
	println("API_SECRET:", apiSecret)
	println("API_KEY:", apiKey)

}
