package xerr

type CodeGroup string

const (
	CodeGroupToast CodeGroup = "toast"
	CodeGroupModel CodeGroup = "model"
	CodeGroupPage  CodeGroup = "page"
)

var ErrCodeGroupMap = map[ErrCode]CodeGroup{
	ErrCodeClaimRewardIPLimit:       CodeGroupModel,
	ErrCodeClaimRewardDeviceLimit:   CodeGroupModel,
	ErrCodeClaimRewardEndpointLimit: CodeGroupModel,
}

func (g CodeGroup) String() string {
	return string(g)
}
