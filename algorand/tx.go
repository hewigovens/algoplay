package algorand

// Payment tx
type Transaction struct {
	_struct struct{} `codec:",omitempty,omitemptyarray"`
	Type    string   `codec:"type"`

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
