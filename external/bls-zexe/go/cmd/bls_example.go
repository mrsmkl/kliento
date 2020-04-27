// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/celo-org/bls-zexe/go/bls"
)

func main() {
	bls.InitBLSCrypto()
	privateKey, _ := bls.GeneratePrivateKey()
	defer privateKey.Destroy()
	privateKeyBytes, _ := privateKey.Serialize()
	fmt.Printf("Private key: %x\n", privateKeyBytes)
	publicKey, _ := privateKey.ToPublic()
	publicKeyBytes, _ := publicKey.Serialize()
	fmt.Printf("Public key: %x\n", publicKeyBytes)
	message := []byte("test")
	extraData := []byte("extra")
	directHashNoPoP, _ := bls.HashDirect(message, false)
	fmt.Printf("Direct Hash: %x\n", directHashNoPoP)
	directHashForPoP, _ := bls.HashDirect(message, true)
	fmt.Printf("Direct Hash for PoP: %x\n", directHashForPoP)
	compositeHash, _ := bls.HashComposite(message, extraData)
	fmt.Printf("Composite Hash: %x\n", compositeHash)
	compressedG1Elem, _ := bls.CompressSignature(directHashForPoP)
	fmt.Printf("Compressed G1 Elem: %x\n", compressedG1Elem)
	fmt.Printf("Compressed G1 Elem length: %d\n", len(compressedG1Elem))
	fmt.Printf("Hash length: %d\n", len(directHashNoPoP))
	uncompressedG2Elem := []byte{0x8a, 0x50, 0x02, 0x7a, 0x01, 0xe4, 0xa1, 0xd4, 0xdb, 0xa6, 0xdb, 0x5d, 0x40, 0xf5, 0x1c, 0xf0, 0xb5, 0xac, 0x54, 0x4f, 0x05, 0xd1, 0xce, 0xf6, 0x83, 0x2a, 0xcf, 0x9e, 0x6f, 0x7c, 0x7c, 0x9f, 0x29, 0x27, 0xdf, 0x52, 0x4a, 0xad, 0x32, 0x19, 0x8b, 0xf5, 0x34, 0x62, 0xfc, 0x39, 0xd6, 0x00, 0x2a, 0xc2, 0x11, 0x33, 0x3d, 0x78, 0x9b, 0x95, 0x26, 0x4f, 0xfc, 0x48, 0x73, 0x84, 0xe7, 0xa3, 0x42, 0x18, 0x44, 0x6c, 0xf9, 0xff, 0x81, 0x19, 0x98, 0xe1, 0xcd, 0x7f, 0x87, 0x27, 0xb5, 0x43, 0x4b, 0xde, 0x4c, 0x48, 0x59, 0x20, 0x36, 0x9d, 0x76, 0x74, 0x33, 0xa4, 0xb3, 0x9c, 0x48, 0x01, 0xe2, 0x29, 0x68, 0xe1, 0xc7, 0x67, 0xb0, 0x92, 0xda, 0xdc, 0xb5, 0x39, 0xf9, 0x88, 0xd1, 0x4d, 0xe3, 0x91, 0x55, 0x2f, 0xb2, 0x42, 0xe6, 0x7c, 0xec, 0x43, 0xb9, 0xa5, 0x55, 0x7c, 0x4c, 0x28, 0x6c, 0x7c, 0x80, 0x70, 0x32, 0xfe, 0x59, 0x66, 0x10, 0x56, 0x3b, 0x84, 0x81, 0xfc, 0xc3, 0x00, 0x89, 0x0f, 0xd4, 0xcd, 0xb3, 0xfa, 0xb8, 0xa9, 0x1c, 0xba, 0xad, 0x86, 0xfe, 0xd4, 0x84, 0x19, 0xac, 0xc3, 0x0f, 0xb2, 0x01, 0x29, 0xf1, 0xa0, 0xcc, 0xa0, 0x9e, 0xab, 0xca, 0x78, 0x2c, 0x6a, 0x75, 0x30, 0x0c, 0x13, 0x47, 0xe8, 0x03, 0x32, 0x23, 0x1f, 0xdd, 0x94, 0x70, 0x64, 0x34, 0x00}
	compressedG2Elem, _ := bls.CompressPublickey(uncompressedG2Elem)
	fmt.Printf("Compressed G2 Elem: %x\n", compressedG2Elem)
	fmt.Printf("Compressed G2 Elem Length: %d\n", len(compressedG2Elem))
	signature, _ := privateKey.SignMessage(message, extraData, true)
	signatureBytes, _ := signature.Serialize()
	fmt.Printf("Signature: %x\n", signatureBytes)
	err := publicKey.VerifySignature(message, extraData, signature, true)
	fmt.Printf("Verified: %t\n", err == nil)

	privateKey2, _ := bls.GeneratePrivateKey()
	defer privateKey2.Destroy()
	privateKeyBytes2, _ := privateKey2.Serialize()
	fmt.Printf("Private key 2: %x\n", privateKeyBytes2)
	publicKey2, _ := privateKey2.ToPublic()
	publicKeyBytes2, _ := publicKey2.Serialize()
	fmt.Printf("Public key 2: %x\n", publicKeyBytes2)
	signature2, _ := privateKey2.SignMessage(message, extraData, true)
	signatureBytes2, _ := signature2.Serialize()
	fmt.Printf("Signature 2: %x\n", signatureBytes2)
	err = publicKey2.VerifySignature(message, extraData, signature2, true)
	fmt.Printf("Verified 2: %t\n", err == nil)

	aggergatedPublicKey, _ := bls.AggregatePublicKeys([]*bls.PublicKey{publicKey, publicKey2})
	aggregatedPublicKeyBytes, _ := aggergatedPublicKey.Serialize()
	fmt.Printf("Aggregated public key: %x\n", aggregatedPublicKeyBytes)
	aggergatedSignature, _ := bls.AggregateSignatures([]*bls.Signature{signature, signature2})
	aggregatedSignatureBytes, _ := aggergatedSignature.Serialize()
	fmt.Printf("Aggregated signature: %x\n", aggregatedSignatureBytes)
	err = aggergatedPublicKey.VerifySignature(message, extraData, aggergatedSignature, true)
	fmt.Printf("Aggregated verified: %t\n", err == nil)
	err = publicKey.VerifySignature(message, extraData, aggergatedSignature, true)
	fmt.Printf("Aggregated verified (with wrong pk): %t\n", err == nil)
}