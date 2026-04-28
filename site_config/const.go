package site_config

// SiteConfigCategory 站点配置分类
type SiteConfigCategory int

const (
	SiteConfigCategoryGame                SiteConfigCategory = 1  // 游戏账号
	SiteConfigCategoryBaseInfo            SiteConfigCategory = 2  // 站点配置
	SiteConfigCategoryFooterInfo          SiteConfigCategory = 3  // 页脚配置
	SiteConfigCategoryDownloadAPP         SiteConfigCategory = 4  // 下载app配置
	SiteConfigCategorySidebarLoginYes     SiteConfigCategory = 5  // 侧边栏已登录配置
	SiteConfigCategorySidebarLoginNo      SiteConfigCategory = 6  // 侧边栏未登录配置
	SiteConfigCategoryFooterMenus         SiteConfigCategory = 7  // 底部主菜单配置
	SiteConfigCategoryUnLoginPromotionImg SiteConfigCategory = 8  // 未登录宣传图
	SiteConfigCategoryLobby               SiteConfigCategory = 9  // 大厅
	SiteConfigCategorySeo                 SiteConfigCategory = 10 // SEO配置
	SiteConfigThirdLogin                  SiteConfigCategory = 11 // 三方登录
	SiteConfigCategoryAgent               SiteConfigCategory = 12 // 代理配置
	SiteConfigCategoryVip                 SiteConfigCategory = 13 // Vip配置
	SiteConfigBackground                  SiteConfigCategory = 14 // 登录/注册背景图
	SiteConfigCurrencyConvert             SiteConfigCategory = 15 // 币种兑换配置
)

type GameCalcBetAmountType int

const (
	_                                             GameCalcBetAmountType = iota
	GameCalcBetAmountTypeRealBetAmount                                  // 实际投注金额
	GameCalcBetAmountTypeRealBetAmountOrWinAmount                       // min(实际投注金额,abs(用户输赢金额))
)

// SiteConfigKey 站点配置key
type SiteConfigKey string

const (
	SiteConfigKeyGame                       SiteConfigKey = "game"                          // 游戏账号配置
	SiteConfigKeyGameFixedCategory          SiteConfigKey = "fixed_game_category"           // 游戏内置分类
	SiteConfigKeyGameCalcBetAmount          SiteConfigKey = "game_calc_bet_amount"          // 游戏内置分类
	SiteConfigKeyBaseInfo                   SiteConfigKey = "base_info"                     // 站点配置
	SiteConfigKeyFooterInfo                 SiteConfigKey = "footer_info"                   // 页脚配置
	SiteConfigKeyDownloadAPP                SiteConfigKey = "download_app"                  // 下载app配置
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
)
