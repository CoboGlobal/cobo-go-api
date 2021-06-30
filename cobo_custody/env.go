package cobo_custody

type Env struct {
	Host string
	CoboPub string
}

func Sandbox() Env {
	return Env{Host: "https://api.sandbox.cobo.com", CoboPub: "032f45930f652d72e0c90f71869dfe9af7d713b1f67dc2f7cb51f9572778b9c876"}
}

func Prod() Env {
	return Env{Host: "https://api.custody.cobo.com", CoboPub: "02c3e5bacf436fbf4da78597e791579f022a2e85073ae36c54a361ff97f2811376"}
}