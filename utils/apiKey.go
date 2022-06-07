package utils

import (
	"math/big"
	"crypto/rand"
	"encoding/hex"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)
var (
	b58Alphabet 		=		 	[]byte("0123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
)
//创建公私钥
func CreateAppKeySecret(username string) (string, string, error) {
	b 					:= 		make([]byte, 20)
	_, err 				:= 		rand.Read(b)
	if err != nil {
		return "", "", err
	}
	var data []byte 	= 		[]byte(username)
	bytes			 	:= 		append(b, data...)
	key 				:= 		Ripemd160Hash(bytes)
	key 				= 		Base58Encode(key)
	keys 				:= 		string(key)
	bytes 				=		append(b, bytes...)
	secret 				:= 		Ripemd160Hash(bytes)
	secrets 			:= 		hex.EncodeToString(secret)
	return keys, secrets, nil
}
//将传入的公钥进行256运算，返回256位hash值
func Ripemd160Hash(publicKey []byte) []byte {
	hash256 			:= 		sha256.New()
	hash256.Write(publicKey)
	hash 				:=	 	hash256.Sum(nil)
	//将上面的256位hash值进行160运算，返回160位的hash值
	ripemd160 			:=	 	ripemd160.New()
	ripemd160.Write(hash)
	//返回Pub Key hash
	return ripemd160.Sum(nil)
}
//字节数组转 Base58,加密
func Base58Encode(input []byte) []byte {
	var result []byte
	x 					:= 		big.NewInt(0).SetBytes(input)
	base	 			:= 		big.NewInt(int64(len(b58Alphabet)))
	zero 				:= 		big.NewInt(0)
	mod 				:= 		&big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}
	return result
}