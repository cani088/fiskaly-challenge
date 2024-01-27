package crypto

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
}

// TODO: implement RSA and ECDSA signing ...
func Sign(dataToBeSigned []byte) ([]byte, error) {
	return nil, nil
}
