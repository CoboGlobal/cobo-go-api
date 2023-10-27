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
		CoboId:     "20220314181458000331767000003732",
		TxId:       "0x1c4d137bc2a2ee8f22cbdf9e90405974e72e65d922f42eb81d9f7a05d0f64fc6",
		WithdrawId: "web_send_by_user_915_1647252768642",
		DeAddress: map[string]string{
			"BTC": "3JBYNrbB4bHtGWHTEa3ZPuRK9kwTiEUo4D",
			"XRP": "rfKyCMyoV6Ln2GZ7YDbrBrnXCbAyBbxRqB|2047482901",
		},
		DeAddresses: map[string]string{
			"BTC": "3JBYNrbB4bHtGWHTEa3ZPuRK9kwTiEUo4D,bc1qf22hpu33u2tkyy528mdvpnre45n8lu5s3ycatu",
			"XRP": "rfKyCMyoV6Ln2GZ7YDbrBrnXCbAyBbxRqB|2047482901,rfKyCMyoV6Ln2GZ7YDbrBrnXCbAyBbxRqB|3752417374",
		},
		LpAddress: map[string]string{
			"BTC": "35eXJPLRTSp4Wn8n2f6pkQF4t3KdU2cuhz",
			"XRP": "rfKyCMyoV6Ln2GZ7YDbrBrnXCbAyBbxRqB|477817505",
		},
		LpAddresses: map[string]string{
			"BTC": "35eXJPLRTSp4Wn8n2f6pkQF4t3KdU2cuhz,34R4JHecUwGNEFVGKz1vR8R6BHGi5FUqPt",
			"XRP": "rfKyCMyoV6Ln2GZ7YDbrBrnXCbAyBbxRqB|477817505,rfKyCMyoV6Ln2GZ7YDbrBrnXCbAyBbxRqB|2874421071",
		},
		Withdraw: map[string]string{
			"COBO_ETH": "0xE410157345be56688F43FF0D9e4B2B38Ea8F7828",
			"XLM":      "GBJDU6TPWHKGV7HRLNTIBA46MG3MB5DUG6BISHX3BF7I75H2HLPV6RJX|4e73f03b",
		},
	}
}

func ProdConfig() Config {
	return Config{
		CoboId:     "20220311154108000184408000002833",
		TxId:       "4041A888C9966BE8916FE65F2FEE7AE9A9DC3F49D0F1643A768C842CA95FA736",
		WithdrawId: "sdk_request_id_fe80cc5f_1647068483396",
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
