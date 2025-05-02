package decoder

import (
	"fmt"
	"os"
)

// OpenAndParseH264 打开并解析指定的 H.264 文件
// filePath: H.264 文件的路径
// 返回值: error，如果操作过程中发生错误则返回错误信息
func OpenAndParseH264(filePath string) error {
	// 1. 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("无法打开文件 '%s': %w", filePath, err)
	}
	defer file.Close() // 确保函数结束时关闭文件

	fmt.Printf("成功打开文件: %s\n", filePath)

	// 2. 读取和解析文件内容 (此处为占位符)
	// 实际的 H.264 解析逻辑会比较复杂，需要读取 NAL units 等
	// 这里仅作演示，打印文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("无法获取文件信息: %w", err)
	}
	fmt.Printf("文件大小: %d 字节\n", fileInfo.Size())

	buffer := make([]byte, fileInfo.Size())

	n, err := file.Read(buffer)
	if err != nil {
		return fmt.Errorf("读取文件内容失败: %w", err)
	}
	fmt.Printf("成功读取 %d 字节的数据\n", n)

	// TODO: 在这里添加 H.264 码流解析逻辑
	fmt.Println("开始解析 H.264 数据...")
	// ... 实现解析 NAL units 的代码 ...
	reader := AnnexBReader{}
	nalu := Nalu{}
	leftLen := n

	for {
		nalu, err = reader.ReadNalu(buffer[(n-leftLen):], leftLen)
		if err != nil {
			break
		}
		fmt.Printf("nalu data len: %d\n", len(nalu.data))
		leftLen -= len(nalu.data)
	}

	fmt.Println("文件解析完成 (占位符)")

	return nil // 解析成功，返回 nil
}

// 分析AvcC格式 NALU
func parseAvcC(data []byte) error {

	return nil
}

func parseExtradata(data []byte) error {
	return nil
}

// 处理防止竞争码
func processEmulationPreventionBytes(data []byte) error {
	return nil
}
