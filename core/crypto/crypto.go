package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/IPFS-eX/go-ipfs-ex/core/prefix"
	"golang.org/x/crypto/scrypt"
)

// ScryptParam contains the parameters used in scrypt function
type ScryptParam struct {
	P     int `json:"p"`
	N     int `json:"n"`
	R     int `json:"r"`
	DKLen int `json:"dkLen,omitempty"`
}

const (
	DEFAULT_N                  = 16384
	DEFAULT_R                  = 8
	DEFAULT_P                  = 8
	DEFAULT_DERIVED_KEY_LENGTH = 64
)

// Return the default parameters used in scrypt function
func GetScryptParameters() *ScryptParam {
	return &ScryptParam{
		N:     DEFAULT_N,
		R:     DEFAULT_R,
		P:     DEFAULT_P,
		DKLen: DEFAULT_DERIVED_KEY_LENGTH,
	}
}

type SymmetricScheme byte

const (
	AES SymmetricScheme = iota
)

var names []string = []string{
	"AES",
}

func GetScheme(name string) (SymmetricScheme, error) {
	for i, v := range names {
		if strings.ToUpper(v) == strings.ToUpper(name) {
			return SymmetricScheme(i), nil
		}
	}
	return 0, errors.New("unknown symmetric scheme " + name)
}
func kdf(pwd []byte, salt []byte) (dKey []byte, err error) {
	param := GetScryptParameters()
	if param.DKLen < 32 {
		err = errors.New("derived key length too short")
		return nil, err
	}
	// Derive the encryption key
	dKey, err = scrypt.Key([]byte(pwd), salt, param.N, param.R, param.P, param.DKLen)
	return dKey, err
}
func randomBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// AESEncryptFile encrypt file and save file locally.
// The file include first 32 bytes of salt data, and the remains are encrypted data
func AESEncryptFile(file string, password string, out string) error {
	salt, err := randomBytes(16)
	if err != nil {
		return errors.New("salt generate error")
	}
	dKey, err := kdf([]byte(password), salt)
	if err != nil {
		return err
	}
	nonce := dKey[:16]
	eKey := dKey[len(dKey)-16:]

	inFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer inFile.Close()

	block, err := aes.NewCipher(eKey)
	if err != nil {
		return err
	}
	stream := cipher.NewOFB(block, nonce)
	outFile, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer outFile.Close()

	tag := []byte{}
	for _, v := range salt {
		tag = append(tag, v)
	}
	encAlg := make([]byte, 16)
	encAlg[15] = byte(AES)
	for _, v := range encAlg {
		tag = append(tag, v)
	}
	_, err = outFile.WriteAt(tag, 0)
	outFile.Seek(int64(len(tag)), io.SeekStart)

	writer := &cipher.StreamWriter{S: stream, W: outFile}
	if _, err := io.Copy(writer, inFile); err != nil {
		return err
	}
	return err
}

func AESEncryptFileReader(inFile io.Reader, password string) (io.Reader, error) {
	salt, err := randomBytes(16)
	if err != nil {
		return nil, errors.New("salt generate error")
	}
	dKey, err := kdf([]byte(password), salt)
	if err != nil {
		return nil, err
	}
	nonce := dKey[:16]
	eKey := dKey[len(dKey)-16:]

	block, err := aes.NewCipher(eKey)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, nonce)
	reader := &cipher.StreamReader{S: stream, R: inFile}
	tag := []byte{}
	for _, v := range salt {
		tag = append(tag, v)
	}
	encAlg := make([]byte, 16)
	encAlg[15] = byte(AES)
	for _, v := range encAlg {
		tag = append(tag, v)
	}
	r := io.MultiReader(bytes.NewReader(tag), reader)
	return r, nil
}

// AESDecryptFile. use AES algorithm to decrypt a file
// The file is include first 32 bytes of salt data, the prefix data if exists, and remains are encrypted data
func AESDecryptFile(file, password, out string) error {
	inFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer inFile.Close()
	encodeSizeLen := base64.StdEncoding.EncodedLen(prefix.PAYLOAD_SIZE_LEN)
	payloadSizeEncodeBuf := make([]byte, encodeSizeLen)
	payloadSizeDecodeBuf := make([]byte, encodeSizeLen)
	if _, err := inFile.Read(payloadSizeEncodeBuf); err != nil {
		return err
	}
	if _, err := base64.StdEncoding.Decode(payloadSizeDecodeBuf, payloadSizeEncodeBuf); err != nil {
		return err
	}
	payloadSize := prefix.GetPayloadLenFromBuf(payloadSizeDecodeBuf)
	prefixLen := base64.StdEncoding.EncodedLen(int(payloadSize))
	prefix := make([]byte, prefixLen)
	// read salt data after skip first N bytes of prefix
	readPrefixLen, err := inFile.Read(prefix)
	if err != nil {
		return err
	}
	if readPrefixLen != prefixLen {
		return errors.New("prefix length not match")
	}

	extData := make([]byte, 32)
	// read salt data after skip first N bytes of prefix
	inFile.Read(extData)

	salt := extData[:16]
	dKey, err := kdf([]byte(password), salt)
	if err != nil {
		return err
	}
	nonce := dKey[:16]
	eKey := dKey[len(dKey)-16:]
	encAlg := extData[31]
	if encAlg != byte(AES) {
		return errors.New("unknown algorithm")
	}
	block, err := aes.NewCipher(eKey)
	if err != nil {
		return err
	}

	stream := cipher.NewOFB(block, nonce)
	outFile, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer outFile.Close()
	reader := &cipher.StreamReader{S: stream, R: inFile}
	// Copy the input file to the output file, decrypting as we go.
	if _, err := io.Copy(outFile, reader); err != nil {
		return err
	}
	return nil
}

func AESDecryptFileWriter(inFile *os.File, password string) (io.Writer, error) {
	extData := make([]byte, 32)
	inFile.ReadAt(extData, 0)
	_, err := inFile.Seek(int64(len(extData)), io.SeekStart)
	if err != nil {
		return nil, err
	}
	salt := extData[:16]
	dKey, err := kdf([]byte(password), salt)
	if err != nil {
		return nil, err
	}
	nonce := dKey[:16]
	eKey := dKey[len(dKey)-16:]
	encAlg := extData[31]
	if encAlg != byte(AES) {
		return nil, errors.New("unknown algorithm")
	}
	block, err := aes.NewCipher(eKey)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, nonce)
	writer := &cipher.StreamWriter{S: stream}
	return writer, nil
}

// AESDecryptFile. use AES algorithm to decrypt a file
// The file is include first 32 bytes of salt data, the prefix data if exists, and remains are encrypted data
func AESDecryptFileReader(file io.Reader, password string) (io.Reader, error) {
	encodeSizeLen := base64.StdEncoding.EncodedLen(prefix.PAYLOAD_SIZE_LEN)
	payloadSizeEncodeBuf := make([]byte, encodeSizeLen)
	payloadSizeDecodeBuf := make([]byte, encodeSizeLen)
	if _, err := file.Read(payloadSizeEncodeBuf); err != nil {
		return nil, err
	}
	if _, err := base64.StdEncoding.Decode(payloadSizeDecodeBuf, payloadSizeEncodeBuf); err != nil {
		return nil, err
	}
	payloadSize := prefix.GetPayloadLenFromBuf(payloadSizeDecodeBuf)
	prefixLen := base64.StdEncoding.EncodedLen(int(payloadSize))
	prefix := make([]byte, prefixLen)
	// read salt data after skip first N bytes of prefix
	readPrefixLen, err := file.Read(prefix)
	if err != nil {
		return nil, err
	}
	if readPrefixLen != prefixLen {
		return nil, errors.New("prefix length not match")
	}
	extData := make([]byte, 32)
	extLen, err := file.Read(extData)
	if extLen != len(extData) {
		return nil, errors.New("ext length not match")
	}

	salt := extData[:16]
	dKey, err := kdf([]byte(password), salt)
	if err != nil {
		return nil, err
	}
	nonce := dKey[:16]
	eKey := dKey[len(dKey)-16:]
	encAlg := extData[31]
	if encAlg != byte(AES) {
		return nil, errors.New("unknown algorithm")
	}
	block, err := aes.NewCipher(eKey)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, nonce)
	reader := &cipher.StreamReader{S: stream, R: file}
	return reader, nil
}
