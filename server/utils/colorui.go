package utils

// 定义背景颜色常量及其对应的十六进制值，以便在应用程序中使用。
const (
	// 实色系
	BgGreen     = "#39b54a" // 绿色
	BgCyan      = "#1cbbb4" // 青色
	BgBlue      = "#0081ff" // 蓝色
	BgPurple    = "#6739b6" // 紫色
	BgMauve     = "#9c26b0" // 淡紫色
	BgPink      = "#e03997" // 粉色
	BgBrown     = "#a5673f" // 棕色
	BgGrey      = "#8799a3" // 灰色
	BgGrayLight = "#f0f0f0" // 浅灰色

	// 渐变色 - 从上至下
	BgShadeTop    = "linear-gradient(rgba(0, 0, 0, 1), rgba(0, 0, 0, 0.01))" // 顶部深色渐变至透明
	BgShadeBottom = "linear-gradient(rgba(0, 0, 0, 0.01), rgba(0, 0, 0, 1))" // 底部透明渐变至深色

	// 浅色变体
	BgRedLight    = "#fadbd9"   // 浅红色背景
	BgOrangeLight = "#fde6d2"   // 浅橙色背景
	BgYellowLight = "#fef2ce"   // 浅黄色背景
	BgOliveLight  = "#e8f4d9"   // 浅橄榄绿背景
	BgGreenLight  = "#d7f0dbff" // 浅绿色背景
	BgCyanLight   = "#d2f1f0"   // 浅青色背景
	BgBlueLight   = "#cce6ff"   // 浅蓝色背景
	BgPurpleLight = "#e1d7f0"   // 浅紫色背景
	BgMauveLight  = "#ebd4ef"   // 浅淡紫色背景
	BgPinkLight   = "#f9d7ea"   // 浅粉色背景
	BgBrownLight  = "#ede1d9"   // 浅棕色背景
	BgGreyLight   = "#e7ebed"   // 浅灰色背景

	// 渐变背景
	BgGradualRed    = "linear-gradient(45deg, #f43f3b, #ec008c)" // 红色渐变
	BgGradualOrange = "linear-gradient(45deg, #ff9700, #ed1c24)" // 橙色渐变
	BgGradualGreen  = "linear-gradient(45deg, #39b54a, #8dc63f)" // 绿色渐变
	BgGradualPurple = "linear-gradient(45deg, #9000ff, #5e00ff)" // 紫色渐变
	BgGradualPink   = "linear-gradient(45deg, #ec008c, #6739b6)" // 粉色渐变
	BgGradualBlue   = "linear-gradient(45deg, #0081ff, #1cbbb4)" // 蓝色渐变
)

// LightColors 切片包含所有浅色背景颜色
var LightColors = []string{
	BgRedLight,
	BgOrangeLight,
	BgYellowLight,
	BgOliveLight,
	BgGreenLight,
	BgCyanLight,
	BgBlueLight,
	BgPurpleLight,
	BgMauveLight,
	BgPinkLight,
	BgBrownLight,
	BgGreyLight,

	BgGreen,
	BgCyan,
	BgBlue,
	BgPurple,
	BgMauve,
	BgPink,
	BgBrown,
	BgGrey,
	BgGrayLight,
}
