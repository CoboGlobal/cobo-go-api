package cobo_custody

import (
	"flag"
	"testing"
)

var env = flag.String("env", "Sandbox", "Env Config")
var secret = flag.String("secret", "Demo", "Api Secrect")
var web3Secret = flag.String("web3Secret", "Web3Demo", "Web3 Api Secrect")

var ConfigData Config
var client Client
var web3Client Web3Client

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

	var web3LocalSigner = LocalSigner{
		PrivateKey: *web3Secret,
	}
	ConfigData = GetData(*env)
	web3Client = Web3Client{
		Signer: web3LocalSigner,
		Env:    GetEnv(*env),
		Debug:  false,
	}

	m.Run()
}
