package decoder

type NaluType int

const (
	NALU_TYPE_UNDEFINE NaluType = iota
	NALU_TYPE_NO_IDR
	NALU_TYPE_NO_IDR_A
	NALU_TYPE_NO_IDR_B
	NALU_TYPE_NO_IDR_C
	NALU_TYPE_IDR
	NALU_TYPE_SEI
	NALU_TYPE_SPS
	NALU_TYPE_PPS
	NALU_TYPE_AUD
	NALU_TYPE_EOSEQ
	NALU_TYPE_EOSTREAM
	NALU_TYPE_FILLER //10000000
	NALU_TYPE_SPSE   //10000001
)

func (n NaluType) String() string {
	return [...]string{
		"UNDEFINE",
		"NO_IDR",
		"NO_IDR_A",
		"NO_IDR_B",
		"NO_IDR_C",
		"IDR",
		"SEI",
		"SPS",
		"PPS",
		"AUD",
		"EOSEQ",
		"EOSTREAM",
		"FILLER",
		"SPSE",
	}[n]
}

type NalRefIdc int

const (
	NALU_REF_IDC_DISPOSABLE NalRefIdc = iota
	NALU_REF_IDC_LOW
	NALU_REF_IDC_HIGH
	NALU_REF_IDC_HIGHESt
)

func (n NalRefIdc) String() string {
	return [...]string{
		"DISPOSABLE",
		"LOW",
		"HIGH",
		"HIGHESt",
	}[n]
}

type Nalu struct {
	startCodeLen   int
	nalu_unit_type NaluType
	forbidden_bit  int
	nal_ref_idc    NalRefIdc
	data           []byte
}

func (n *Nalu) GetNaluType() NaluType {
	return n.nalu_unit_type
}

func (n *Nalu) SetNaluType(t NaluType) {
	n.nalu_unit_type = t
	return
}

func (n *Nalu) GetForbiddenBit() int {
	return n.forbidden_bit
}

func (n *Nalu) SetForbiddenBit(t int) {
	n.forbidden_bit = t
	return
}

func (n *Nalu) GetNalRefIdc() NalRefIdc {
	return n.nal_ref_idc
}

func (n *Nalu) SetNalRefIdc(t NalRefIdc) {
	n.nal_ref_idc = t
	return
}

func (n *Nalu) SetData(data []byte, len int) {
	// fmt.Println("Nalu SetData len: ", len)

	n.data = make([]byte, len)
	copy(n.data, data)
	return
}

func (n *Nalu) GetEBSP() EBSP {
	ebsp := EBSP{}
	ebsp.SetData(n.data[n.startCodeLen:], len(n.data)-n.startCodeLen)
	return ebsp
}
