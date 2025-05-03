package decoder

// EBSP: nalu 去除startcode
type EBSP struct {
	data []byte
}

func (n *EBSP) SetData(data []byte, len int) {
	// fmt.Println("EBSP SetData len: ", len)

	n.data = make([]byte, len)
	copy(n.data, data)
	return
}

func (n *EBSP) GetRBSP() RBSP {
	rbsp := RBSP{}
	buffer := make([]byte, len(n.data))
	copylen := 0
	len := len(n.data)
	for i := 0; i < len; i++ {
		if n.data[i] == 0x03 && i > 2 && i < len-1 {
			if n.data[i-1] == 0x00 && n.data[i-2] == 0x00 && (n.data[i+1] == 0x00 || n.data[i+1] == 0x01 || n.data[i+1] == 0x02 || n.data[i+1] == 0x03) {
				continue
			}
		}
		buffer[i] = n.data[i]
		copylen++
	}

	rbsp.SetData(buffer, copylen)
	return rbsp
}
