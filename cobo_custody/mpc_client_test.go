package cobo_custody

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMPCClient_GetMpcSupportedChains(t *testing.T) {
	_, apiError := mpcClient.GetMpcSupportedChains()
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMpcSupportedCoins(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.GetMpcSupportedCoins(chainCode)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMpcMainAddress(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.GetMpcMainAddress(chainCode)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_MpcBatchGenerateAddresses(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.MpcBatchGenerateAddresses(chainCode, 2)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMpcAddressList(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.GetMpcAddressList(chainCode, "", "", 0, 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMpcBalance(t *testing.T) {
	chainCode := "ETH"

	_, apiError := mpcClient.GetMpcBalance("0xb2ad1bdf4a1d766e8faeb94689547d3fede5792c", chainCode, "")
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListMpcBalances(t *testing.T) {
	_, apiError := mpcClient.ListMpcBalances(0, 10, "")
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMpcUnspentInputsList(t *testing.T) {
	_, apiError := mpcClient.GetMpcUnspentInputsList("0xb2ad1bdf4a1d766e8faeb94689547d3fede5792c", "ETH")
	assert.Nil(t, apiError, "api error not nil")
}

//func TestMPCClient_MpcCreateTransaction(t *testing.T) {
//	_, apiError := mpcClient.MpcCreateTransaction("VTHO", "test_001", 1,
//		"0xe434c89a6dacc9ceb7e3e94b5f966fa445127ab7", "0x99ea76426bf86f0fe046355606f495b79dd6e180",
//		"", 0, 0, 0, 0, "")
//	assert.Nil(t, apiError, "api error not nil")
//}
//
//func TestMPCClient_MpcDropTransaction(t *testing.T) {
//	_, apiError := mpcClient.MpcDropTransaction("20221213164754000373267000009730", "test_001", 1, 0, 0)
//	assert.Nil(t, apiError, "api error not nil")
//}
//
//func TestMPCClient_MpcSpeedupTransaction(t *testing.T) {
//	_, apiError := mpcClient.MpcSpeedupTransaction("20221213164754000373267000009730", "PressTest-121316-bycaixiao-716894650154221843", 0, 0, 0)
//	assert.Nil(t, apiError, "api error not nil")
//}

func TestMPCClient_GetMpcTransactionsByRequestIds(t *testing.T) {
	_, apiError := mpcClient.GetMpcTransactionsByRequestIds("0100", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMpcTransactionsByCoboIds(t *testing.T) {
	_, apiError := mpcClient.GetMpcTransactionsByCoboIds("0100", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_GetMpcTransactionsByTxHash(t *testing.T) {
	_, apiError := mpcClient.GetMpcTransactionsByTxHash("0100", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListMpcWalletTransactions(t *testing.T) {
	_, apiError := mpcClient.ListMpcWalletTransactions(0, 0, 0, "", "",
		0, "", "", "", 0)
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_EstimateMpcFee(t *testing.T) {
	_, apiError := mpcClient.EstimateMpcFee("ETH", 0, "")
	assert.Nil(t, apiError, "api error not nil")
}

func TestMPCClient_ListMpcTssNodeRequests(t *testing.T) {
	_, apiError := mpcClient.ListMpcTssNodeRequests(0, 0)
	assert.Nil(t, apiError, "api error not nil")
}
