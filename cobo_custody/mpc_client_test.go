package cobo_custody

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMPCClient_GetSupportedChains(t *testing.T) {
	_, apiError := mpcClient.GetSupportedChains()
	assert.Nil(t, apiError, "api error not nil")
	fmt.Println("TestMPCClient_GetSupportedChains")
}

func TestMPCClient_GetSupportedCoins(t *testing.T) {
	chainCode := "GETH"

	_, apiError := mpcClient.GetSupportedCoins(chainCode)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetWalletSupportedCoins(t *testing.T) {
	_, apiError := mpcClient.GetWalletSupportedCoins()
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_IsValidAddress(t *testing.T) {
	coin := "GETH"
	address := "0x3ede1e59a3f3a66de4260df7ba3029b515337e5c"

	_, apiError := mpcClient.IsValidAddress(coin, address)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMainAddress(t *testing.T) {
	chainCode := "GETH"

	_, apiError := mpcClient.GetMainAddress(chainCode)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GenerateAddressess(t *testing.T) {
	chainCode := "GETH"

	_, apiError := mpcClient.GenerateAddresses(chainCode, 2)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_UpdateAddressDescription(t *testing.T) {
	coin := "GETH"
	address := "0x6a060efe0ff887f4e24dc2d2098020abf28bcce4"
	description := "test"

	_, apiError := mpcClient.UpdateAddressDescription(coin, address, description)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListAddress(t *testing.T) {
	chainCode := "GETH"

	_, apiError := mpcClient.ListAddresses(chainCode, "", "", 0, 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetBalance(t *testing.T) {
	chainCode := "GETH"

	_, apiError := mpcClient.GetBalance("0x6a060efe0ff887f4e24dc2d2098020abf28bcce4", chainCode, "")
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListBalances(t *testing.T) {
	_, apiError := mpcClient.ListBalances(0, 10, "", "")
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListSpendable(t *testing.T) {
	_, apiError := mpcClient.ListSpendable("0x6a060efe0ff887f4e24dc2d2098020abf28bcce4", "GETH")
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_CreateTransaction(t *testing.T) {
	request_id := time.Now().UnixNano() / int64(time.Millisecond)
	_, apiError := mpcClient.CreateTransaction(
		"GETH", fmt.Sprintf("%d", request_id), big.NewInt(9),
		"0x6a060efe0ff887f4e24dc2d2098020abf28bcce4", "0x6a060efe0ff887f4e24dc2d2098020abf28bcce4",
		"", nil, big.NewInt(0), big.NewInt(0), 100, "", big.NewInt(0), big.NewInt(0), nil, "", 0,
	)
	assert.Nil(t, apiError, "api error not nil")
}

// func TestMPCClient_DropTransaction(t *testing.T) {
// 	_, apiError := mpcClient.DropTransaction("20221213164754000373267000009730", "test_001", 1, 0, 0)
// 	assert.Nil(t, apiError, "api error not nil")
// }

// func TestMPCClient_SpeedupTransaction(t *testing.T) {
// 	_, apiError := mpcClient.SpeedupTransaction("20221213164754000373267000009730", "PressTest-121316-bycaixiao-716894650154221843", 0, 0, 0)
// 	assert.Nil(t, apiError, "api error not nil")
// }

func TestMPCClient_TransactionsByRequestIds(t *testing.T) {
	_, apiError := mpcClient.TransactionsByRequestIds("1668678820274", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_TransactionsByCoboIds(t *testing.T) {
	_, apiError := mpcClient.TransactionsByCoboIds("20231219151058000165807000003045", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_TransactionsByTxHash(t *testing.T) {
	_, apiError := mpcClient.TransactionsByTxHash("0x1e14311142db1f5b02e587f0e00643f7fd460c81e73dffff65cf501123fb99dd", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListTransactions(t *testing.T) {
	_, apiError := mpcClient.ListTransactions(0, 0, 0, "", "",
		0, "", "", "", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_EstimateFee(t *testing.T) {
	_, apiError := mpcClient.EstimateFee("GETH", nil, "0x6a060efe0ff887f4e24dc2d2098020abf28bcce4", "", "", "", nil, nil, nil, "")
	assert.Nil(t, apiError, "api error not nil")
}

// func TestMPCClient_ListTssNodeRequests(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("skipping testing in short mode")
// 	}
// 	_, apiError := mpcClient.ListTssNodeRequests(0, 0)
// 	assert.Nil(t, apiError, "api error not nil")
// }

func TestMPCClient_RetryDoubleCheck(t *testing.T) {
	_, apiError := mpcClient.RetryDoubleCheck("123")
	assert.NotNil(t, apiError, "api error not nil")
}

func TestMPCClient_ListTssNode(t *testing.T) {
	_, apiError := mpcClient.ListTssNode()
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_SignMessagesByRequestIds(t *testing.T) {
	res, apiError := mpcClient.SignMessagesByRequestIds("1690349242683,1690268795963,1690187858862")
	fmt.Println(res)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_SignMessagesByCobotIds(t *testing.T) {
	res, apiError := mpcClient.SignMessagesByCobotIds("20230726132723000341052000008222,20230725150636000308867000003494,20230725135301000361318000002480")
	fmt.Println(res)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMaxSendAmount(t *testing.T) {
	res, apiError := mpcClient.GetMaxSendAmount("GETH", *big.NewFloat(0.0), "", "0x6a060efe0ff887f4e24dc2d2098020abf28bcce4")
	fmt.Println(res)
	assert.Nil(t, apiError, "api error not nil")
	fmt.Println("TestMPCClient_GetMaxSendAmount")
}
