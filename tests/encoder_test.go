package base

import (
	".."
	"testing"
)

func TestAes(t *testing.T) {

	testString := "Hello Password!"

	key := []byte("vpL54DlR2KG{JSAaAX7Tu;*#&DnG`M0o")
	iv := []byte("@z]zv@10-K.5Al0Dm`@foq9k\"VRfJ^~j")
	encrypted := base.EncryptAes(testString, key, iv)
	decrypted := base.DecryptAes(encrypted, key, iv)

	if decrypted != testString {
		t.Error("Decrypt failed", encrypted, decrypted)
	}

}
