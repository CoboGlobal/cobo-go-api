package cobo_custody

import (
	"flag"
	"testing"
)

var env = flag.String("env", "Sandbox", "Env Config")
var secret = flag.String("secret", "Demo", "Api Secrect")
var web3Secret = flag.String("web3Secret", "Web3Demo", "Web3 Api Secrect")
var mpcSecret = flag.String("mpcSecret", "510b3a777ae15ed922136c2aad985789145dea17d906210f66f4bd781a3cfb44", "MPC Api Secrect")

var ConfigData Config
var client Client
var web3Client Web3Client
var mpcClient MPCClient

func GetEnv(env string) Env {
	if env == "Prod" {
		return Prod()
	} else if env == "Develop" {
		return Develop()
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
