package xerr

type CodeGroup string

const (
	CodeGroupToast CodeGroup = "toast"
	CodeGroupModel CodeGroup = "model"
	CodeGroupPage  CodeGroup = "page"
)

var ErrCodeGroupMap = map[ErrCode]CodeGroup{
	ErrCodeActivityEnded:                    CodeGroupModel,
	ErrCodeActivityClosed:                   CodeGroupModel,
	ErrCodeClaimRewardSmsLimit:              CodeGroupModel,
	ErrCodeClaimRewardEmailLimit:            CodeGroupModel,
	ErrCodeClaimRewardSameActivityTypeLimit: CodeGroupModel,
	ErrCodeClaimRewardIPLimit:               CodeGroupModel,
	ErrCodeClaimRewardDeviceLimit:           CodeGroupModel,
	ErrCodeClaimRewardEndpointLimit:         CodeGroupModel,
	ErrCodeUserExists:                       CodeGroupModel,
}

func (g CodeGroup) String() string {
	return string(g)
}
