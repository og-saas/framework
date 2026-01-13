package consts

// StatusType 状态
// StatusTypeEnable 启用
// StatusTypeDisable 禁用
type StatusType uint32

const (
	StatusTypeEnable  StatusType = 1 // 启用
	StatusTypeDisable StatusType = 2 // 禁用
)

// Bool 状态转换成bool
func (t StatusType) Bool() bool {
	return t == StatusTypeEnable
}
