package ethaddress

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetEthAddress(t *testing.T) {
	privateKey, _ := hex.DecodeString("6368868e8b5778af1f6a34a6b49479d1b463efbfb80ea11f8f1091d9951db5ac")
	prvKey := secp256k1.PrivKey{Key: privateKey}
	publicKeyCompress := prvKey.PubKey()
	address, _ := GetEthAddress(publicKeyCompress.Bytes())
	fmt.Println("address is :" + address)
	require.Equal(t, address, "0x243b5d459359bc6f5dd63758b24d0647192cc9d3")

	address2, _ := GetEthAddress(publicKeyCompress.Bytes())
	fmt.Println("address is :" + address2)
	require.Equal(t, address2, "0x243b5d459359bc6f5dd63758b24d0647192cc9d3")
}
