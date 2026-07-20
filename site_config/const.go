package site_config

// SiteConfigCategory 站点配置分类
type SiteConfigCategory int

const (
	SiteConfigCategoryGame                SiteConfigCategory = 1  // 游戏账号
	SiteConfigCategoryBaseInfo            SiteConfigCategory = 2  // 站点配置
	SiteConfigCategoryFooterInfo          SiteConfigCategory = 3  // 页脚配置
	SiteConfigCategoryDownloadAPP         SiteConfigCategory = 4  // 下载app配置
	SiteConfigCategorySidebarVisualMenu   SiteConfigCategory = 5  // 侧边栏可视化配置
	SiteConfigCategorySidebarLoginYes     SiteConfigCategory = 5  // 侧边栏已登录配置（弃用）
	SiteConfigCategorySidebarLoginNo      SiteConfigCategory = 6  // 侧边栏未登录配置（弃用）
	SiteConfigCategoryFooterMenus         SiteConfigCategory = 7  // 底部主菜单配置
	SiteConfigCategoryUnLoginPromotionImg SiteConfigCategory = 8  // 未登录宣传图
	SiteConfigCategoryLobby               SiteConfigCategory = 9  // 大厅
	SiteConfigCategorySeo                 SiteConfigCategory = 10 // SEO配置
	SiteConfigThirdLogin                  SiteConfigCategory = 11 // 三方登录
	SiteConfigCategoryAgent               SiteConfigCategory = 12 // 代理配置
	SiteConfigCategoryVip                 SiteConfigCategory = 13 // Vip配置
	SiteConfigBackground                  SiteConfigCategory = 14 // 登录/注册背景图
	SiteConfigCurrencyConvert             SiteConfigCategory = 15 // 币种兑换配置
	SiteConfigCategoryAudit               SiteConfigCategory = 16 // 稽核
	SiteConfigCategoryReplenish           SiteConfigCategory = 17 // 审核配置
	SiteConfigCategoryPayoutMonitor       SiteConfigCategory = 18 // 派奖监控
	SiteConfigCategoryPcShowType          SiteConfigCategory = 19 // PC展示样式
	SiteConfigCategoryCountryAccessLimit  SiteConfigCategory = 20 // 区域限制
	SiteConfigCategoryTopDownloadBar      SiteConfigCategory = 21 // 顶部下载条设置
	SiteConfigCategoryMarqueeIcon         SiteConfigCategory = 23 // 跑马灯图标

)

type GameCalcBetAmountType int

const (
	_                                             GameCalcBetAmountType = iota
	GameCalcBetAmountTypeRealBetAmount                                  // 实际投注金额
	GameCalcBetAmountTypeRealBetAmountOrWinAmount                       // min(实际投注金额,abs(用户输赢金额))
)

type PlatformCurrencyType int

const (
	_                             PlatformCurrencyType = iota
	PlatformCurrencyModeCustomize                      // 自定义
	PlatformCurrencyModeFiat                           // 指定法币
)

// SiteConfigKey 站点配置key
type SiteConfigKey string

const (
	SiteConfigKeyGame                       SiteConfigKey = "game"                          // 游戏账号配置
	SiteConfigKeyGameFixedCategory          SiteConfigKey = "fixed_game_category"           // 游戏内置分类
	SiteConfigKeyGameCalcBetAmount          SiteConfigKey = "game_calc_bet_amount"          // 游戏投注金额计算方式
	SiteConfigKeyBaseInfo                   SiteConfigKey = "base_info"                     // 站点基础信息配置
	SiteConfigKeySiteLogo                   SiteConfigKey = "site_logo"                     // 站点logo
	SiteConfigKeyFooterInfo                 SiteConfigKey = "footer_info"                   // 页脚配置
	SiteConfigKeyDownloadAPP                SiteConfigKey = "download_app"                  // 下载app配置
	SiteConfigKeyAPPInstall                 SiteConfigKey = "app_install"                   // APP安装入口
	SiteConfigKeyFooterMenus                SiteConfigKey = "footer_menus"                  // 底部主菜单配置
	SiteConfigKeySidebarUserAccount         SiteConfigKey = "sidebar_user_account"          // 用户账号区域
	SiteConfigKeySidebarRechargeDeposit     SiteConfigKey = "sidebar_recharge_withdraw"     // 充值提现区域
	SiteConfigKeySidebarFeatureOperation    SiteConfigKey = "sidebar_feature_operation"     // 功能操作区域
	SiteConfigKeySidebarOtherLinks          SiteConfigKey = "sidebar_other_links"           // 其他链接区域
	SiteConfigKeySidebarAdPlacement         SiteConfigKey = "sidebar_ad_placement"          // 广告位区域
	SiteConfigKeySidebarAppDownload         SiteConfigKey = "sidebar_app_download"          // APP下载区域
	SiteConfigKeySidebarLanguageSwitch      SiteConfigKey = "sidebar_language_switch"       // 语言切换区域
	SiteConfigKeySidebarVersionInfo         SiteConfigKey = "sidebar_version_info"          // 版本信息区域
	SiteConfigKeySidebarLoginRegister       SiteConfigKey = "sidebar_login_register"        // 登录注册区域
	SiteConfigKeyUnLoginHomeConfig          SiteConfigKey = "home_config"                   // 首页未登录宣传图
	SiteConfigKeyLobbyEnter                 SiteConfigKey = "lobby_enter"                   // 大厅快捷入口
	SiteConfigKeySeoConfig                  SiteConfigKey = "seo_config"                    // SEO配置
	SiteConfigKeyGoogleConfig               SiteConfigKey = "google_config"                 // 谷歌登录配置
	SiteConfigKeyAgentGradeIcons            SiteConfigKey = "agent_grade_icons"             // 代理等级图标
	SiteConfigKeyVipPromoteInfo             SiteConfigKey = "vip_promote_info"              // Vip宣传图
	SiteConfigKeyVipCustomerService         SiteConfigKey = "vip_customer_service"          // Vip专属客服
	SiteConfigKeyPlatformCurrencySafetyRisk SiteConfigKey = "platform_currency_safety_risk" // 币种兑换规则配置
	SiteConfigKeyPayoutMonitor              SiteConfigKey = "payout_monitor"                // 派奖监控
	SiteConfigKeyPlatformCurrencyMode       SiteConfigKey = "platform_currency_mode"        // 平台币模式
	SiteConfigKeyRegisterBg                 SiteConfigKey = "register_backgroud_img"        // 注册背景图配置
	SiteConfigKeyLoginBg                    SiteConfigKey = "login_backgroud_img"           // 登录背景图配置
	SiteConfigKeyAuditMode                  SiteConfigKey = "audit_mode"                    // 稽核模式
	SiteConfigKeyAuditWeight                SiteConfigKey = "audit_weight"                  // 稽核权重
	SiteConfigKeyWithdrawAuditConfig        SiteConfigKey = "withdraw_audit_config"         // 提现审核配置
	SiteConfigKeyManualReviewProcess        SiteConfigKey = "manual_review_process"         // 提款自动转人工审核配置
	SiteConfigKeyPcShowType                 SiteConfigKey = "site_pc_show_type"             // PC展示样式
	SiteConfigKeyRegionControl              SiteConfigKey = "region_access_control"         // 地区访问控制
	SiteConfigKeyInterceptionPage           SiteConfigKey = "interception_page"             // 拦截页设置
	SiteConfigKeyMaintenance                SiteConfigKey = "site_maintenance"              // 站点维护设置
	SiteConfigKeyTopDownloadBar             SiteConfigKey = "top_download_bar"              // 顶部下载条设置
	SiteConfigKeySidebarVisualMenu          SiteConfigKey = "sidebar_visual_menu"           // 侧边栏可视化菜单
	SiteConfigKeySAppConfig                 SiteConfigKey = "app_config"                    // app打包配置
	SiteConfigKeyServerUrl                  SiteConfigKey = "server_url"                    // app api服务地址
	SiteConfigKeyMarqueeIcon                SiteConfigKey = "marquee_config"                // 跑马灯图标
)

func (k SiteConfigKey) String() string {
	return string(k)
}

func (c SiteConfigCategory) Int32() int32 {
	return int32(c)
}

// MemberScope 侧边栏菜单会员可见范围
type MemberScope int

const (
	MemberScopeAll          MemberScope = 1 // 全部会员
	MemberScopeCustomMember MemberScope = 2 // 自定义会员
	MemberScopeVipLevel     MemberScope = 3 // VIP等级
	MemberScopeChannel      MemberScope = 4 // 指定渠道
)
