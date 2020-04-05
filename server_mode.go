package sip

// 当前服务器运行模式
var ServerMode = 0

// 服务器模式
const (
	ModeRegistrar   = 1 << 1 // 注册模式
	ModeProxy       = 1 << 2 // 代理模式
	ModeMiddleMan   = 1 << 3 // 中间人模式
	ModeRecordAudio = 1 << 4 // (基于ModeMiddleMan) 录音模式
)

// 判断是否包含某种模式
func InMode(mode int) bool {
	return mode&ServerMode == mode
}
