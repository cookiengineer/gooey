package aesctr

type CryptoKeyType string

const (
	CryptoKeyTypeSecret  CryptoKeyType = "secret"
	CryptoKeyTypePrivate CryptoKeyType = "private"
	CryptoKeyTypePublic  CryptoKeyType = "public"
)

