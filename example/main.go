package main

import (
	"github.com/CoboCustody/cobo-go-api/cobo_custody"
	"github.com/google/uuid"
)

var apiSecret = "xxxxxxx" // replace by customer secret

var localSigner = cobo_custody.LocalSigner{
	PrivateKey: apiSecret,
}

var mpcClient = cobo_custody.MPCClient{
	Signer: localSigner,
	Env:    cobo_custody.Sandbox(), // use cobo_custody.Prod() in Prod
}

func main() {
	// create transfer transaction with GETH coin
	coin := "GETH"
	fromAddress := "xxxxx" // replace by customer fromAddress
	toAddress := "xxxxx"   // any ETH address
	requestId := uuid.New().String()
	amount := 10000000
	toAddressDetails := ""
	fee := 0
	operation := 100 // Transfer
	extraParameters := ""
	
	// get estimate fee
	estimateFee, apiError := mpcClient.EstimateFee(coin, amount, toAddress)
	if apiError != nil {
		println(apiError.ErrorCode, apiError.ErrorMessage, apiError.ErrorId)
		return
	}
	str, _ := estimateFee.Encode()
	println(string(str))
	gasPrice, _ := estimateFee.GetPath("average", "gas_price").Int()
	gasLimit, _ := estimateFee.GetPath("average", "gas_limit").Int()
	
	// create transaction
	transactionResp, apiError := mpcClient.CreateTransaction(coin, requestId, amount, fromAddress, toAddress, toAddressDetails, fee, gasPrice, gasLimit, operation, extraParameters)
	if apiError != nil {
		println(apiError.ErrorCode, apiError.ErrorMessage, apiError.ErrorId)
		return
	}
	str, _ = transactionResp.Encode()
	println(string(str))
	coboID, _ := transactionResp.GetPath("cobo_id").String()
	println(coboID)

	// get transaction by request_id
	status := 0
	transactionResp, apiError = mpcClient.TransactionsByRequestIds(requestId, status)
	if apiError != nil {
		println(apiError.ErrorCode, apiError.ErrorMessage, apiError.ErrorId)
		return
	}
	str, _ = transactionResp.Encode()
	println(string(str))
}
