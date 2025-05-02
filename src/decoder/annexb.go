package decoder

import (
	"fmt"
)

type AnnexBReader struct {
}

func (reader *AnnexBReader) ReadNalu(data []byte, len int) (Nalu, error) {
	nalu := Nalu{}
	fmt.Println("ReadNalu")

	startCodeLen, err := reader.checkStartCode(data, len)
	if err != nil {
		return nalu, fmt.Errorf("checkStartCode fail: %w", err)
	}

	nalu.startCodeLen = startCodeLen
	// fmt.Println("startCodeLen: ", startCodeLen)

	//find nalu end
	var endPos int = 0
	for i := startCodeLen; i < len; i++ {
		startCodeLen, err = reader.checkStartCode(data[i:], len-i)
		if err == nil {
			endPos = i
			break
		}
	}

	if endPos > 0 {
		//find next start code
		nalu.SetData(data, endPos)
	} else {
		//not find next start code
		nalu.SetData(data, len)
	}

	return nalu, nil
}

func (reader *AnnexBReader) checkStartCode(data []byte, len int) (int, error) {
	if len <= 2 {
		return 0, fmt.Errorf("len <= 2")
	}

	if (len >= 3) && (data[0] == 0x00) && (data[1] == 0x00) && (data[2] == 0x01) {
		return 3, nil
	}

	if (len >= 4) && (data[0] == 0x00) && (data[1] == 0x00) && (data[2] == 0x00) && (data[3] == 0x01) {
		return 4, nil
	}
	return 0, fmt.Errorf("not found start code")
}
