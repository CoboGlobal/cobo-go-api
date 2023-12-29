package cobo_custody

type Env struct {
	Host    string
	CoboPub string
}

func Dev() Env {
	return Env{Host: "https://api.dev.cobo.com", CoboPub: "03596da539963fb1dd29d5859e25903eb76b9f7ed2d58516e29c9f80c201ff2c1b"}
}

func Prod() Env {
	return Env{Host: "https://api.custody.cobo.com", CoboPub: "02c3e5bacf436fbf4da78597e791579f022a2e85073ae36c54a361ff97f2811376"}
}

// type DepositAddress struct {
// 	BTC string
// 	XRP string
// }
// type DepositAddresses struct {
// 	BTC string
// 	XRP string
// }
// type LoopAddress struct {
// 	BTC string
// 	XRP string
// }
// type LoopAddresses struct {
// 	BTC string
// 	XRP string
// }

type Config struct {
	CoboId      string
	TxId        string
	WithdrawId  string
	DeAddress   map[string]string
	DeAddresses map[string]string
	LpAddress   map[string]string
	LpAddresses map[string]string
	Withdraw    map[string]string
}

func DevConfig() Config {
	return Config{
		CoboId:     "20231213152104000114035000006167",
		TxId:       "332d0377c0cc08bc9f9d5b07320add949e30d8da0b5fea5140de63e3779101a0",
		WithdrawId: "82ddd375-901a-4d0f-81a4-36d04fbc69a4",
		DeAddress: map[string]string{
			"BTC": "38kcymiNQXk8WTWX9tPLRZP9wxvXPXcsFy",
			"XRP": "rBphERztHKga1cyMgWiDen7WDkbkfn1iPE|3414236551",
		},
		DeAddresses: map[string]string{
			"BTC": "38kcymiNQXk8WTWX9tPLRZP9wxvXPXcsFy,3ApTsekq5XpUtM5CzAKqntHkvoSpYdCDHw",
			"XRP": "rfKyCMyoV6Ln2GZ7YDbrBrnXCbAyBbxRqB|20474829019,rfKyCMyoV6Ln2GZ7YDbrBrnXCbAyBbxRqB|3752417374",
		},
		LpAddress: map[string]string{
			"BTC": "3FKpEfhsULvsnutcbX8gXPpTo4ewXy7jWJ",
			"XRP": "rBphERztHKga1cyMgWiDen7WDkbkfn1iPE|2284746463",
		},
		LpAddresses: map[string]string{
			"BTC": "3FKpEfhsULvsnutcbX8gXPpTo4ewXy7jWJ,3FhponzJguuN2nvoKkdb5bJJMT1zyZvH8w",
			"XRP": "rBphERztHKga1cyMgWiDen7WDkbkfn1iPE|2284746463,rBphERztHKga1cyMgWiDen7WDkbkfn1iPE|2446372187",
		},
		Withdraw: map[string]string{
			"COBO_ETH": "0x00a70fa1125e336afc22a641b015c878f44c1c1d",
			"XLM":      "GCXMPEHKXQQIZIAGBB67HX55PSN35M2XWVTBNQWLABXS5T3UY42LBJGS|481247198",
		},
	}
}

func ProdConfig() Config {
	return Config{
		CoboId:     "20231213152104000114035000006167",
		TxId:       "332d0377c0cc08bc9f9d5b07320add949e30d8da0b5fea5140de63e3779101a0",
		WithdrawId: "82ddd375-901a-4d0f-81a4-36d04fbc69a4",
		DeAddress: map[string]string{
			"BTC": "36xYx7vf7DUKpJDixpY3EoV2jchFwYSNCb",
			"XRP": "rBWpYJhuJWBPAkzJ4kYQqHShSkkF3rgeD|3992922539",
		},
		DeAddresses: map[string]string{
			"BTC": "36xYx7vf7DUKpJDixpY3EoV2jchFwYSNCb,bc1q0l24tf5sjdu9t7l6hrlhxz9aq9yeej9h2sc7tk",
			"XRP": "rBWpYJhuJWBPAkzJ4kYQqHShSkkF3rgeD|3992922539,rBWpYJhuJWBPAkzJ4kYQqHShSkkF3rgeD|1492154866",
		},
		LpAddress: map[string]string{
			"BTC": "34WLjtk9ta96BVxc1jRF7j5eVvehoftsVV",
			"XRP": "rBWpYJhuJWBPAkzJ4kYQqHShSkkF3rgeD|633829231",
		},
		LpAddresses: map[string]string{
			"BTC": "34WLjtk9ta96BVxc1jRF7j5eVvehoftsVV,33P1kjMfDCKipR58S7XbsCqbmPT5YGrhUo",
			"XRP": "rBWpYJhuJWBPAkzJ4kYQqHShSkkF3rgeD|633829231,rBWpYJhuJWBPAkzJ4kYQqHShSkkF3rgeD|935940214",
		},
		Withdraw: map[string]string{
			"COBO_ETH": "0xE410157345be56688F43FF0D9e4B2B38Ea8F7828",
			"XLM":      "GBJDU6TPWHKGV7HRLNTIBA46MG3MB5DUG6BISHX3BF7I75H2HLPV6RJX|4e73f03b",
		},
	}
}
