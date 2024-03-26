// Package shift implements shift cipher. It encrypts a plaintext message by adding
// a key to each byte. It decrypts a ciphertext message by subtracting a key from
// each byte. Start with tests. Adapted from https://github.com/bitfield/eg-crypto.
package shift

func Encrypt(plaintext []byte, key byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b + key
	}
	return ciphertext
}

func Decrypt(ciphertext []byte, key byte) []byte {
	return Encrypt(ciphertext, -key)
}
