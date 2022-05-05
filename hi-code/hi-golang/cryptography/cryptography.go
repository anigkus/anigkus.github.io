/*
Copyright 2022 The https://github.com/anigkus Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cryptography

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Main() {
	fmt.Println("Cryptography...")

	cryptoHash()

	cryptoBase64()
}

/*out:
----------------Small Message----------------
Md5: 0536b60a0ffd59131e2d3aa01877e351

Sha1: cfe6575a2cf9dba32590189092d8c2124fa60be2

Sha256: 221aabb72fd78dc3d4f46833808e676d1c73cdb53cfc4e2b970a5059126a75ad

Sha512: be3ea996c078bfbcbe0899dc5659ed9b1d61b7f0eaa9e69dde9d9f7733d42970ca1d06cdb0721dad606c561fa36b62d6bdbd8d38172f810e5fd59750eb9f4a70


----------------Large Message----------------
Md5: b996257c44da3f271bd725dcee779d3c

Sha1: 13e7bfdc444ece3074136a41181240161558035f

Sha256: d6c50947607da4761120ca411f105fab50de7b9ec98999b6952539526a1ea774

Sha512: 82f130991377b93a9893d51e5da2439e27939db6d58203836734bdf91a3fcdf99e22f9e8703371ccf68bfc89e37428958d7047e948a8c023117ed9bdc53612a2
*/
func cryptoHash() {

	fmt.Println("\n----------------Small Message----------------")
	message := []byte(RandStringRunes(10))

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))

	fmt.Println("\n----------------Large Message----------------")
	message = []byte(RandStringRunes(100))

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))
}

/*out:
base64.RawURLEncoding.EncodeToString(stringEncode1):aHR0cDovL2FuaWdrdXMuZ2l0aHViLmlvP2E9MTAmYj0yOSwxOSZjPXh0
base64.URLEncoding.EncodeToString(stringEncode2):aHR0cDovL2FuaWdrdXMuZ2l0aHViLmlvP2E9MTAmYj0yOSwxOSZjPXh0
base64.RawStdEncoding.DecodeString(stringEncode1):http://anigkus.github.io?a=10&b=29,19&c=xt
*/
func cryptoBase64() {
	//.Encode
	string1 := "http://anigkus.github.io?a=10&b=29,19&c=xt"
	var des_byte1 []byte = make([]byte, 56)
	src_byte1 := []byte(string1)
	stringEncode1 := base64.RawStdEncoding.EncodeToString(src_byte1)
	stringEncode2 := base64.RawURLEncoding.EncodeToString(src_byte1)
	stringEncode3 := base64.StdEncoding.EncodeToString(src_byte1)

	base64.URLEncoding.Encode(des_byte1, src_byte1)
	stringEncode5 := string(des_byte1)
	stringEncode4 := base64.URLEncoding.EncodeToString(src_byte1)

	fmt.Printf("base64.RawStdEncoding.EncodeToString(src_byte1):%v\n", stringEncode1)
	fmt.Printf("base64.RawURLEncoding.EncodeToString(src_byte1):%v\n", stringEncode2)
	fmt.Printf("base64.StdEncoding.EncodeToString(src_byte1):%v\n", stringEncode3)
	fmt.Printf("base64.URLEncoding.EncodeToString(src_byte1):%v\n", stringEncode4)
	fmt.Printf("base64.URLEncoding.Encode(des_byte1, src_byte1):%v\n", stringEncode5)

	//.Decode
	src_byte2, err := base64.URLEncoding.DecodeString(stringEncode5)
	if err != nil {
		fmt.Println(err)
	} else {
		stringDecode1 := string(src_byte2)
		fmt.Printf("base64.RawStdEncoding.DecodeString(stringEncode5):%v\n", stringDecode1)
	}

	var des_byte2 []byte = make([]byte, 57) //Can be greater then length
	base64.URLEncoding.Decode(des_byte2, []byte(stringEncode3))
	stringDecode2 := string(des_byte2)
	fmt.Printf("base64.URLEncoding.Decode(des_byte2,src_byte1):%v\n", stringDecode2)

}
