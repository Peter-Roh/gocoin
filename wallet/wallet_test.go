package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"testing"
)

const (
	testKey     string = "30770201010420aa55d9530d70e79f9003cb6614553af619408f0aed5b9736dafc136f701bbc61a00a06082a8648ce3d030107a14403420004eb1a6a150fffb4178627e396e73b473512e555c8362ab925c718faf25c626b06c685baad60166d92b51e6df4be87bd1cb566896630bed67c3927227bc7921a74"
	testPayload string = "00d432c1446b8e1a6c2b35c5fb69ba41b12d2ca69ef27bb7b438fdb7ce9903b6"
	testSig     string = "1ba1619f7ee79a7dfe2ef80769f1fa9b556e581046a0d08ba4f8239d33615ed6a4658d5c676aed8f3c7d417755c4cf73cb312789d0325c0fea5c75e7e18f9aca"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = addressFromKey(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}

	tests := []test{
		{
			input: testPayload,
			ok:    true,
		},
		{
			input: "04d432c1446b8e1a6c2b35c5fb69ba41b12d2ca69ef27bb7b438fdb7ce9903b6",
			ok:    false,
		},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify could not verify test signature and payload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("restoreBigInts should return error when payload is not hex.")
	}
}
