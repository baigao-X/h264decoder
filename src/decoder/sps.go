package decoder

import "fmt"

// 第0个字节 type

// 第一个字节
type ProfileIdc uint8

const (
	PROFILE_BASELINE = 66
	PROFILE_MAIN     = 77
	PROFILE_EXTENDE  = 88
	PROFILE_HIGH     = 100
	PROFILE_HIGH_10  = 110
	PROFILE_HIGH_422 = 122
	PROFILE_HIGH_444 = 244
)

func (p ProfileIdc) String() string {
	return [...]string{
		"BASELINE",
		"MAIN",
		"EXTENDED",
		"HIGH",
		"HIGH_10",
		"HIGH_422",
		"HIGH_444",
	}[p]
}

// 第三个字节
type LevelIdc uint8

const (
	LEVEL_1_0 = 10 //0x0A, 最大分辨率176x144(QCIF),最大帧数16FPS,最大码率64kbps,针对早期移动设备
	LEVEL_1_1 = 11 //0x0B, 最大分辨率352x288(CIF),最大帧数30FPS,最大码率192kbps,针对移动设备标清流媒体
	LEVEL_3_0 = 30 //0x1E, 最大分辨率720x480      ,最大帧数30FPS,最大码率10Mbps,标清流媒体(SD)
	LEVEL_3_1 = 31 //0x1F, 最大分辨率1280x720     ,最大帧数30FPS,最大码率14Mbps,高清流媒体(720P)
	LEVEL_4_0 = 40 //0x28, 最大分辨率1920x1080    ,最大帧数30FPS,最大码率20Mbps,高清蓝光/直播（1080P)
	LEVEL_4_1 = 41 //0x29, 最大分辨率1920x1080    ,最大帧数60FPS,最大码率50Mbps,针对高清1080P视频
	LEVEL_5_0 = 50 //0x32, 最大分辨率4096x2160    ,最大帧数30FPS,最大码率135Mbps,针对4K超高清
	LEVEL_5_1 = 51 //0x33, 最大分辨率4096x2304    ,最大帧数60FPS,最大码率240Mkbps,针对高帧率4K超高清
)

func (l LevelIdc) String() string {
	return [...]string{
		"1.0",
		"1.1",
		"3.0",
		"3.1",
		"4.0",
		"4.1",
		"5.0",
		"5.1",
	}[l]
}

type SPS struct {
	nalu Nalu

	profile_idc             ProfileIdc
	level_idc               LevelIdc
	seq_parameter_set_id    uint8
	chroma_format_idc       uint8
	separate_colour_plane   uint8
	bit_depth_luma_minus8   uint8
	bit_depth_chroma_minus8 uint8
}

func (s *SPS) Parse() {
	fmt.Println("SPS Parse")
	s.profile_idc = (ProfileIdc)(s.nalu.rbsp[1])
	fmt.Println("profile_idc: ", (uint8)(s.profile_idc))
	s.level_idc = (LevelIdc)(s.nalu.rbsp[3])
	fmt.Println("level_idc: ", (uint8)(s.level_idc))
}
