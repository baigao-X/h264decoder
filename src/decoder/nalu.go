package decoder

import (
	"fmt"
)

type Nalu struct {
	startCodeLen int
	naluType     byte
	data         []byte
}

func (n *Nalu) GetNaluType() byte {
	return n.naluType
}

func (n *Nalu) SetData(data []byte, len int) {
	fmt.Println("Nalue SetData len: ", len)

	n.data = make([]byte, len)
	copy(n.data, data)
	return
}
