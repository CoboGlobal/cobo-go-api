package cobo_custody

import (
	"fmt"
	"math/big"
	"testing"
)

var localSigner = LocalSigner{
	PrivateKey: "apiSecret",
}
var client = Client{
	Signer: localSigner,
	Env:    Sandbox(),
	Debug:  false,
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

func TestClient_GetCoinInfo(t *testing.T) {
	result, apiError := client.GetCoinInfo("ETH")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetCoinInfo")
	fmt.Println(string(str))
}

func TestClient_NewDepositAddress(t *testing.T) {
	result, apiError := client.NewDepositAddress("BTC", false)
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_NewDepositAddress")
	fmt.Println(string(str))
}

func TestClient_BatchNewDepositAddress(t *testing.T) {
	result, apiError := client.BatchNewDepositAddress("BTC", 3, true)
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_BatchNewDepositAddress")
	fmt.Println(string(str))
}

func TestClient_VerifyDepositAddress(t *testing.T) {
	result, apiError := client.VerifyDepositAddress("BTC", "3JBYNrbB4bHtGWHTEa3ZPuRK9kwTiEUo4D")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_VerifyDepositAddress")
	fmt.Println(string(str))
}

func TestClient_BatchVerifyDepositAddress(t *testing.T) {
	result, apiError := client.BatchVerifyDepositAddress("BTC", "3JBYNrbB4bHtGWHTEa3ZPuRK9kwTiEUo4D,bc1qf22hpu33u2tkyy528mdvpnre45n8lu5s3ycatu")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_BatchVerifyDepositAddress")
	fmt.Println(string(str))
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
	result, apiError := client.GetAddressHistory("ETH")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetAddressHistory")
	fmt.Println(string(str))
}

func TestClient_CheckLoopAddressDetails(t *testing.T) {
	result, apiError := client.CheckLoopAddressDetails("BTC", "35eXJPLRTSp4Wn8n2f6pkQF4t3KdU2cuhz", "")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_CheckLoopAddressDetails")
	fmt.Println(string(str))
}

func TestClient_VerifyLoopAddressList(t *testing.T) {
	result, apiError := client.VerifyLoopAddressList("BTC", "35eXJPLRTSp4Wn8n2f6pkQF4t3KdU2cuhz,34R4JHecUwGNEFVGKz1vR8R6BHGi5FUqPt")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_VerifyLoopAddressList")
	fmt.Println(string(str))
}

func TestClient_GetTransactionDetails(t *testing.T) {
	result, apiError := client.GetTransactionDetails("20220314181458000331767000003732")
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
	result, apiError := client.Withdraw("COBO_ETH",
		"",
		"0xE410157345be56688F43FF0D9e4B2B38Ea8F7828",
		big.NewInt(1),
		map[string]string{})
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_Withdraw")
	fmt.Println(string(str))
}

func TestClient_QueryWithdrawInfo(t *testing.T) {
	result, apiError := client.QueryWithdrawInfo("web_send_by_user_915_1647252768642")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_QueryWithdrawInfo")
	fmt.Println(string(str))
}

func TestClient_GetStakingProductDetails(t *testing.T) {
	result, apiError := client.GetStakingProductDetails("184100", "zh")
	if apiError != nil {
		t.Fail()
	}
	str, _ := result.Encode()
	fmt.Println("TestClient_GetStakingProductDetails")
	fmt.Println(string(str))
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
	result, apiError := client.Stake("184100", big.NewInt(1000000))
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
	result, apiError := client.Unstake("184100", big.NewInt(1000000))
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
	result, apiError := client.GetUnstakings("TETH")
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
