package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

////16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法

func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpaading := int(data[length-1])
	return data[:(length - unpaading)]
}

func PKCS7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padingByte := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padingByte...)
}

//aes 加密
func AesCBCEncrypt(data, key, iv []byte) string {
	//创建加密实例
	block, _ := aes.NewCipher(key)
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充数据
	data = PKCS7Padding(data, blockSize)
	fmt.Println("padding后的值是", data)
	//使用cbc加密方式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//初始化加密数据接收切片
	encryptData := make([]byte, len(data))
	//执行加密
	blockMode.CryptBlocks(encryptData, data)

	return base64.StdEncoding.EncodeToString(encryptData)
}

//aes 解密
func AesCBCDecrypt(data, key, iv []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	blockSize := block.BlockSize()

	if len(data) < blockSize {
		return nil, errors.New("ciphertext too short")
	}

	if len(data)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	decryptData := make([]byte, len(data))
	blockMode.CryptBlocks(decryptData, data)
	decryptData = PKCS7UnPadding(decryptData)
	//解密之后的串
	return decryptData, nil
}

//bcrypt 加密密码
func PasswordHash(password string) ([]byte, error) {
	pByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return pByte, err
}

//bcrypt 密码比对
func PasswordVerify(hashPass, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPass, password)
	fmt.Println(err)
	return err == nil
}
