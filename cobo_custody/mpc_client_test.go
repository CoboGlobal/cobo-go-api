package cobo_custody

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMPCClient_GetSupportedChains(t *testing.T) {
	_, apiError := mpcClient.GetSupportedChains()
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetSupportedCoins(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.GetSupportedCoins(chainCode)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_IsValidAddress(t *testing.T) {
	coin := "ETH"
	address := "0xb2ad1bdf4a1d766e8faeb94689547d3fede5792c"

	_, apiError := mpcClient.IsValidAddress(coin, address)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMainAddress(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.GetMainAddress(chainCode)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GenerateAddressess(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.GenerateAddresses(chainCode, 2)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListAddress(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.ListAddresses(chainCode, "", "", 0, 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetBalance(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.GetBalance("0xb2ad1bdf4a1d766e8faeb94689547d3fede5792c", chainCode, "")
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListBalances(t *testing.T) {
	_, apiError := mpcClient.ListBalances(0, 10, "")
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListSpendable(t *testing.T) {
	_, apiError := mpcClient.ListSpendable("0xb2ad1bdf4a1d766e8faeb94689547d3fede5792c", "ETH")
	assert.Nil(t, apiError, "api error not nil")
}

//func TestMPCClient_CreateTransaction(t *testing.T) {
//	_, apiError := mpcClient.CreateTransaction("VTHO", "test_001", 1,
//		"0xe434c89a6dacc9ceb7e3e94b5f966fa445127ab7", "0x99ea76426bf86f0fe046355606f495b79dd6e180",
//		"", 0, 0, 0, 0, "")
//	assert.Nil(t, apiError, "api error not nil")
//}
//
//func TestMPCClient_DropTransaction(t *testing.T) {
//	_, apiError := mpcClient.DropTransaction("20221213164754000373267000009730", "test_001", 1, 0, 0)
//	assert.Nil(t, apiError, "api error not nil")
//}
//
//func TestMPCClient_SpeedupTransaction(t *testing.T) {
//	_, apiError := mpcClient.SpeedupTransaction("20221213164754000373267000009730", "PressTest-121316-bycaixiao-716894650154221843", 0, 0, 0)
//	assert.Nil(t, apiError, "api error not nil")
//}

func TestMPCClient_TransactionsByRequestIds(t *testing.T) {
	_, apiError := mpcClient.TransactionsByRequestIds("0100", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_TransactionsByCoboIds(t *testing.T) {
	_, apiError := mpcClient.TransactionsByCoboIds("0100", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_TransactionsByTxHash(t *testing.T) {
	_, apiError := mpcClient.TransactionsByTxHash("0100", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListTransactions(t *testing.T) {
	_, apiError := mpcClient.ListTransactions(0, 0, 0, "", "",
		0, "", "", "", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_EstimateFee(t *testing.T) {
	_, apiError := mpcClient.EstimateFee("ETH", 0, "")
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListTssNodeRequests(t *testing.T) {
	_, apiError := mpcClient.ListTssNodeRequests(0, 0)
	assert.Nil(t, apiError, "api error not nil")
}
