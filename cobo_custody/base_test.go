package cobo_custody

import (
	"flag"
	"testing"
)

var env = flag.String("env", "Develop", "Env Config")
var secret = flag.String("secret", "Demo", "Api Secrect")
var web3Secret = flag.String("web3Secret", "Web3Demo", "Web3 Api Secrect")
var mpcSecret = flag.String("mpcSecret", "MPCDemo", "MPC Api Secrect")

var ConfigData Config
var client Client
var web3Client Web3Client
var mpcClient MPCClient

func GetEnv(env string) Env {
	if env == "Prod" {
		return Prod()
	} else if env == "Develop" {
		return Dev()
	}

	return Dev()
}
func GetData(env string) Config {
	if env == "Prod" {
		return ProdConfig()
	}
	return DevConfig()
}

func TestMain(m *testing.M) {
	flag.Parse()
    if ((secret == nil || mpcSecret == nil) || (*secret == "Demo" || *mpcSecret == "MPCDemo")) {
        panic("secret or mpcSecret should not be empty")
    }
	var localSigner = LocalSigner{
		PrivateKey: *secret,
	}
	ConfigData = GetData(*env)
	client = Client{
		Signer: localSigner,
		Env:    GetEnv(*env),
		Debug:  false,
	}

	var web3LocalSigner = LocalSigner{
		PrivateKey: *web3Secret,
	}
	ConfigData = GetData(*env)
	web3Client = Web3Client{
		Signer: web3LocalSigner,
		Env:    GetEnv(*env),
		Debug:  false,
	}

	var mpcLocalSigner = LocalSigner{
		PrivateKey: *mpcSecret,
	}
	ConfigData = GetData(*env)
	mpcClient = MPCClient{
		Signer: mpcLocalSigner,
		Env:    GetEnv(*env),
		Debug:  false,
	}

	m.Run()
}
