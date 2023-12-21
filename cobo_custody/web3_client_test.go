//go:build skip

package cobo_custody

import (
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeb3Client_GetWeb3SupportedChains(t *testing.T) {
	_, apiError := web3Client.GetWeb3SupportedChains()
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3SupportedCoins(t *testing.T) {
	chainCode := "RETH"

	_, apiError := web3Client.GetWeb3SupportedCoins(chainCode)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3SupportedNftCollections(t *testing.T) {
	_, apiError := web3Client.GetWeb3SupportedNftCollections()
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3SupportedContracts(t *testing.T) {
	chainCode := "RETH"

	_, apiError := web3Client.GetWeb3SupportedContracts(chainCode)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3SupportedContractMethods(t *testing.T) {
	chainCode := "RETH"
	contractAddress := "0x7851dcc90e79f3f2c59915e7f4d6fabd8d3d305b"

	_, apiError := web3Client.GetWeb3SupportedContractMethods(chainCode, contractAddress)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_BatchWeb3NewAddress(t *testing.T) {
	chainCode := "ETH"
	count := 2

	_, apiError := web3Client.BatchWeb3NewAddress(chainCode, count)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3AddressList(t *testing.T) {
	chainCode := "ETH"
	pageIndex := 0
	pageLength := 40
	sortFlag := 0

	_, apiError := web3Client.GetWeb3AddressList(chainCode, pageIndex, pageLength, sortFlag)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3WalletAssetList(t *testing.T) {
	address := "0xd387292d5be73c8b9d6d3a4dcdd49e00edf75b6a"
	chainCode := "RETH"

	_, apiError := web3Client.GetWeb3WalletAssetList(address, chainCode)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3WalletNftList(t *testing.T) {
	nftCode := "NFT_RETH_PROOF_MOONBIRDS"
	address := "0xd387292d5be73c8b9d6d3a4dcdd49e00edf75b6a"

	_, apiError := web3Client.GetWeb3WalletNftList(nftCode, address)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3WalletNftDetail(t *testing.T) {
	nftCode := "NFT_RETH_PROOF_MOONBIRDS"
	tokenId := "148"

	_, apiError := web3Client.GetWeb3WalletNftDetail(nftCode, tokenId)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_Web3Withdraw(t *testing.T) {
	coin := "ETH"
	requestId := strconv.FormatInt(time.Now().Unix(), 10)
	fromAddr := "0xd2176409a1ac767824921e45b7ee300745cb1e3f"
	toAddr := "0xD2176409a1Ac767824921e45B7Ee300745cB1e3f"
	amount := big.NewInt(101)

	_, apiError := web3Client.Web3Withdraw(coin, requestId, fromAddr, toAddr, amount)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3WithdrawTransaction(t *testing.T) {
	requestId := "1665303298935"

	_, apiError := web3Client.GetWeb3WithdrawTransaction(requestId)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_Web3Contract(t *testing.T) {
	chainCode := "ETH"
	requestId := strconv.FormatInt(time.Now().Unix(), 10)
	walletAddr := "0xd2176409a1ac767824921e45b7ee300745cb1e3f"
	contractAddr := "0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0"
	methodId := "0xa9059cbb"
	methodName := "transfer"
	args := "[\"0x040149e133077aebcfe4594e00638135eb4bc77f\", 1]"
	amount := big.NewInt(1)

	_, apiError := web3Client.Web3Contract(chainCode, requestId, walletAddr, contractAddr, methodId, methodName, args, amount, nil)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_GetWeb3ContractTransaction(t *testing.T) {
	requestId := "1664239624441"

	_, apiError := web3Client.GetWeb3ContractTransaction(requestId)
	assert.Nil(t, apiError, "api error not nil")
}

func TestWeb3Client_ListWeb3WalletTransactions(t *testing.T) {
	address := "0xd2176409a1ac767824921e45b7ee300745cb1e3f"
	chainCode := "ETH"
	maxId := ""
	minId := "20221009161459000368403000001228"
	limit := 10

	_, apiError := web3Client.ListWeb3WalletTransactions(address, chainCode, maxId, minId, limit)
	assert.Nil(t, apiError, "api error not nil")
}
