package main

import (
	"encoding/base32"
	"encoding/json"
	"fmt"
	"os"

	"crypto/sha512"
	"encoding/base64"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
)

// Payment tx
type Transaction struct {
	Type string `codec:"type"`

	Sender      []byte `codec:"snd"`
	Fee         uint64 `codec:"fee"`
	FirstValid  uint64 `codec:"fv"`
	LastValid   uint64 `codec:"lv"`
	Note        []byte `codec:"note"`
	GenesisID   string `codec:"gen"`
	GenesisHash []byte `codec:"gh"`
	Group       []byte `codec:"grp"`
	Receiver    []byte `codec:"rcv"`
	Amount      uint64 `codec:"amt"`
}

func main() {
	data, _ := os.ReadFile("v2.json")

	var dict map[string]interface{}
	_ = json.Unmarshal(data, &dict)

	// decode base64 encoded fields
	keys := []string{"grp", "note", "gh"}
	for _, key := range keys {
		if val, ok := dict[key]; ok {
			decoded, _ := base64.StdEncoding.DecodeString(val.(string))
			dict[key] = decoded
		}
	}

	// decode base32 encoded addresses
	keys = []string{"snd", "rcv"}
	for _, key := range keys {
		decoded, _ := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(dict[key].(string))
		pubkey := decoded[:32]
		dict[key] = pubkey
	}

	tx := Transaction{
		Type:        "pay",
		Sender:      dict["snd"].([]byte),
		Fee:         uint64(dict["fee"].(float64)),
		FirstValid:  uint64(dict["fv"].(float64)),
		LastValid:   uint64(dict["lv"].(float64)),
		Note:        dict["note"].([]byte),
		GenesisID:   "mainnet-v1.0",
		GenesisHash: dict["gh"].([]byte),
		Group:       dict["grp"].([]byte),
		Receiver:    dict["rcv"].([]byte),
		Amount:      uint64(dict["amt"].(float64)),
	}

	buf := msgpack.Encode(tx)

	sign := []byte(string("TX"))
	sign = append(sign, buf...)

	fmt.Printf("msgpack encoded:\n%x\n", buf)

	sha_512_256 := sha512.Sum512_256(sign)
	fmt.Printf("sha512_256 hash:\n%x\n", sha_512_256)

	txId := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(sha_512_256[:])
	fmt.Printf("base32 id: %s\n", txId)
	fmt.Println("expected : ZRWRJSHSMUCP7Y3HBUDKNCR32GSZLEHDLLJMBEXUTK4ZXHVGE22Q")
}
