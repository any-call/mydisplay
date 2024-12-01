package mydisplay

import (
	"fmt"
	"github.com/mattn/go-runewidth"
)

func AlignString(s string, width int, leftAlign bool) string {
	// 计算字符串的实际显示宽度
	strWidth := runewidth.StringWidth(s)
	// 计算填充的空格数量
	pad := width - strWidth
	if pad > 0 {
		if leftAlign {
			// 左对齐，填充在右边
			return s + fmt.Sprintf("%*s", pad, "")
		} else {
			// 右对齐，填充在左边
			return fmt.Sprintf("%*s", pad, "") + s
		}
	}
	// 如果字符串宽度超过指定宽度，则直接返回原字符串
	return s
}
