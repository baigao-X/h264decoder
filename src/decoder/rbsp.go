package decoder

// RBSP: EBSP 去除防竞争码
type RBSP struct {
	data []byte
}

func (n *RBSP) SetData(data []byte, len int) {
	// fmt.Println("RBSP SetData len: ", len)

	n.data = make([]byte, len)
	copy(n.data, data)
	return
}
