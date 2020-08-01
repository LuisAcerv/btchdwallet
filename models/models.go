package models

// Wallet struct
type Wallet struct {
	Name    string `json:"name"`
	PubKey  string `json:"pub"`
	PrivKey string `json:"priv"`
}

// EncryptedWallet struct
type EncryptedWallet struct {
	Name      string `json:"name"`
	Encrypted string `json:"encrypted"`
	CreatedAt string `json:"createdAt"`
}
