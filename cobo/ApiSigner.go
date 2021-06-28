package cobo
type ApiSigner interface {
	Sign(message string) string
}
