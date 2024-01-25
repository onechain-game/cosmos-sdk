package ethaddress

import (
	"encoding/hex"
	secp "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/hashicorp/golang-lru/simplelru"
	"github.com/wealdtech/go-merkletree/keccak256"
)

var (
	ethAddrCache *simplelru.LRU
)

func init() {
	var err error
	if ethAddrCache, err = simplelru.NewLRU(60000, nil); err != nil {
		panic(err)
	}
}
func GetEthAddress(publicKey []byte) (string, error) {
	addr, ok := ethAddrCache.Get(hex.EncodeToString(publicKey))
	if ok {
		return addr.(string), nil
	}
	publicKeyObj, err := secp.ParsePubKey(publicKey)
	if err != nil {
		panic(err)
	}
	publicUncompressed := publicKeyObj.SerializeUncompressed()
	keccak256 := keccak256.New()
	raw := keccak256.Hash(publicUncompressed[1:])
	address := "0x" + hex.EncodeToString(raw[12:])
	ethAddrCache.Add(hex.EncodeToString(publicKey), address)
	return address, nil
}
