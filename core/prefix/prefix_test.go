package prefix

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGenPrefix(t *testing.T) {
	addr := "QmQicWZ5oeqaz5YbLi1pTMyBJQ6WHmqhVwbpHMmYHeN4MH"
	prefix := NewPrefix([]byte(addr), " @@%^#*)!?><|}{}>~!.txt ", "1234")
	buf := prefix.Serialize()
	fmt.Printf("prefix-len: %v, buf %v, str %s, addr %d\n", len(buf), buf, buf, []byte(addr))

	prefix2 := &FilePrefix{}
	prefix2.Deserialize(buf)
	fmt.Printf("version: %d\n", prefix2.Version)
	fmt.Printf("encrypt: %t\n", prefix2.Encrypt)
	fmt.Printf("salt: %v\n", prefix2.EncryptSalt)
	fmt.Printf("hash: %v\n", prefix2.EncryptHash)
	fmt.Printf("owner: %s\n", prefix2.Owner)
	fmt.Printf("fileSize: %d\n", prefix2.FileSize)
	fmt.Printf("fileNameLen: %d\n", prefix2.FileNameLen)
	fmt.Printf("fileName: %s\n", prefix2.FileName)
	verify := VerifyEncryptPassword("123224", prefix2.EncryptSalt, prefix2.EncryptHash)
	fmt.Printf("verify : %t\n", verify)
}

func TestGetPrefixFromFile(t *testing.T) {
	prefix2, prefixBuf, err := GetPrefixFromFile("QmbWbv8UJcCE6CQM1JwVtc2jEbq6idWcV7nY8AfyFWw8q5")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("version: %d\n", prefix2.Version)
	fmt.Printf("encrypt: %t\n", prefix2.Encrypt)
	fmt.Printf("salt: %v\n", prefix2.EncryptSalt)
	fmt.Printf("hash: %v\n", prefix2.EncryptHash)
	fmt.Printf("owner: %s\n", prefix2.Owner)
	fmt.Printf("fileSize: %d\n", prefix2.FileSize)
	fmt.Printf("fileNameLen: %d\n", prefix2.FileNameLen)
	fmt.Printf("fileName: %s\n", prefix2.FileName)
	fmt.Printf("prefix str %s, len %d\n", prefixBuf, len(prefixBuf))
	verify := VerifyEncryptPassword("123", prefix2.EncryptSalt, prefix2.EncryptHash)
	fmt.Printf("verify : %t\n", verify)
}

func TestReader(t *testing.T) {
	reader := bytes.NewReader([]byte("123456"))
	buf1 := make([]byte, 3)
	buf2 := make([]byte, 3)
	reader.Read(buf1)
	reader.Read(buf2)
	fmt.Printf("buf1 %v, buf2 %v\n", buf1, buf2)
}
