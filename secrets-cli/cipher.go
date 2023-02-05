package secrets
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)
func newCipherBlock(key string) (cipher.Block, error){
	hasher:=md5.New()
	fmt.Fprint(hasher, key)
	block, err:=aes.NewCipher(hasher.Sum(nil))
	if err!=nil{
		return nil, err
	}
	return block, nil
}
func encryptStream(key string, iv []byte) (cipher.Stream, error){
	block, err:=newCipherBlock(key)
	if err!=nil{
		return nil, err
	}
	return cipher.NewCFBEncrypter(block, iv), nil
}
func decryptStream(key string, iv []byte) (cipher.Stream, error){
	block, err:=newCipherBlock(key)
	if err!=nil{
		return nil, err
	}
	return cipher.NewCFBDecrypter(block, iv),nil
}
func Encrypt(key string, plainText string) (string, error){

	//make fix length key
	//create aes cipher block from this key => block
	//
	//
	//make cipherText slice 16+len(plaintext)
	//
	//first 16 bytes will be used as initialization vector
	//Length of IV and cipherKey must be same
	//fill random data in IV
	//create a stream CFB enc using IV and block
	//
	//Use encryptor to encrypt the plainText and store it into cipherText[BlockSize:]
	cipherText:=make([]byte, aes.BlockSize+len(plainText))
	//initialization vector
	iv:=cipherText[:aes.BlockSize]
	if _, err:=io.ReadFull(rand.Reader, iv);err!=nil{
		return "",err
	}
	stream, err:=encryptStream(key, iv)
	if err!=nil{
		return "", err
	}
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))
	return fmt.Sprintf("%x", cipherText), nil
}

func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error){
	iv:=make([]byte, aes.BlockSize)
	if _,err:=io.ReadFull(rand.Reader, iv);err!=nil{
		return nil, err
	}
	stream, err:=encryptStream(key, iv)
	if err!=nil{
		return nil, err
	}
	n, err:=w.Write(iv)
	if n!=len(iv){
		return nil, errors.New("encrypt: Not able to write iv")
	}
	return &cipher.StreamWriter{S:stream, W:w}, nil
}



func Decrypt(key , cipherHex string)(string, error) {
	cipherText, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", nil
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("Ciphertext too short")
	}

	//iv is still there in text
	iv := cipherText[:aes.BlockSize]
	stream, err := decryptStream(key, iv)
	if err!=nil{
		return "", nil
	}

	stream.XORKeyStream(cipherText, cipherText)
	return fmt.Sprintf("%s", cipherText[aes.BlockSize:]), nil
}
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error){
	iv:=make([]byte, aes.BlockSize)
	n,err:=r.Read(iv)
	if n<len(iv){
		return nil, errors.New("decrypt : Unable to read full IV")
	}
	stream, err:=decryptStream(key, iv)
	if err!=nil{
		return nil, err
	}
	return &cipher.StreamReader{S:stream, R: r}, nil
}

