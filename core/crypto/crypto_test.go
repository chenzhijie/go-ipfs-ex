package crypto

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDecryptReaderToTar(t *testing.T) {
	decryptPwdStr := "123"
	reader, err := os.Open("/Users/zhijie/Desktop/onchain/ipfs-test/ipfs1/QmbWbv8UJcCE6CQM1JwVtc2jEbq6idWcV7nY8AfyFWw8q5")
	if err != nil {
		t.Fatal(err)
	}
	outReader, err := AESDecryptFileReader(reader, decryptPwdStr)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := ioutil.ReadAll(outReader)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("result: %s\n", buf)
}
