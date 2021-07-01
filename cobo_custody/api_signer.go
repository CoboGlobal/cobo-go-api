package cobo_custody

type ApiSigner interface {
	Sign(message string) string
	GetPublicKey() string
}
