package edgeone

import (
	"fmt"
	"net"
	"net/url"
	"sort"
	"strings"

	common "github.com/og-saas/proto/pb/common/v1"
)

// Domain is the projection of domain_map joined with domain. A missing join must
// stay visible as an empty ZoneID so it cannot be mistaken for a successful purge.
type Domain struct {
	DomainID int64  `gorm:"column:domain_id"`
	Host     string `gorm:"column:domain_url"`
	HTTPS    bool   `gorm:"column:https_enabled"`
	ZoneID   string `gorm:"column:zone_id"`
}

type Batch struct {
	ZoneID  string
	Targets []string
}

func BuildBatches(domains []Domain, paths []string, batchSize int) ([]Batch, error) {
	if len(domains) == 0 {
		return nil, fmt.Errorf("no active CDN domains for site")
	}
	if batchSize == 0 {
		batchSize = 100
	}
	if batchSize < 1 || batchSize > 1000 {
		return nil, fmt.Errorf("invalid EdgeOne batch size")
	}
	groups := map[string]map[string]bool{}
	for _, d := range domains {
		zone := strings.TrimSpace(d.ZoneID)
		if zone == "" {
			return nil, fmt.Errorf("domain_id=%d has no EdgeOne zone_id", d.DomainID)
		}
		host := strings.ToLower(strings.TrimSuffix(strings.TrimSpace(d.Host), "."))
		u, err := url.Parse("https://" + host)
		if err != nil || u.Host != host || u.Hostname() != host || u.User != nil || u.Path != "" || u.RawQuery != "" || u.Fragment != "" || strings.ContainsAny(host, "* /\\?#:@") || !strings.Contains(host, ".") || net.ParseIP(host) != nil {
			return nil, fmt.Errorf("invalid CDN hostname for domain_id=%d", d.DomainID)
		}
		hosts := []string{host}
		schemes := []string{"http"}
		if d.HTTPS {
			schemes = append(schemes, "https")
		}
		if groups[zone] == nil {
			groups[zone] = map[string]bool{}
		}
		for _, h := range hosts {
			for _, scheme := range schemes {
				for _, path := range paths {
					groups[zone][scheme+"://"+h+path] = true
				}
			}
		}
	}
	zones := make([]string, 0, len(groups))
	for z := range groups {
		zones = append(zones, z)
	}
	sort.Strings(zones)
	var result []Batch
	for _, z := range zones {
		targets := make([]string, 0, len(groups[z]))
		for t := range groups[z] {
			targets = append(targets, t)
		}
		sort.Strings(targets)
		for i := 0; i < len(targets); i += batchSize {
			end := i + batchSize
			if end > len(targets) {
				end = len(targets)
			}
			result = append(result, Batch{z, targets[i:end]})
		}
	}
	return result, nil
}

// Directories returns precise API paths for CDN refresh.
// Private user/wallet/agent caches never trigger a CDN refresh.
func Directories(e *common.CacheDeleteEntity) []string {
	if e == nil {
		return nil
	}
	switch e.GetCacheCategory() {
	case common.CacheCategory_CACHE_CATEGORY_SHARED_DICT:
		for _, key := range e.GetEntityElements() {
			if key == "trans_category" {
				return []string{"/api/v1/wallet/getTransactionCategoryList"}
			}
		}
	case common.CacheCategory_CACHE_CATEGORY_SHARED_COUNTRY:
		return []string{"/api/v1/public/getCountryList"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_INFO,
		common.CacheCategory_CACHE_CATEGORY_TENANT_BASIC:
		return []string{"/api/v1/config/site/basic"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_CONFIG,
		common.CacheCategory_CACHE_CATEGORY_TENANT_SKIN,
		common.CacheCategory_CACHE_CATEGORY_TENANT_LANGUAGE,
		common.CacheCategory_CACHE_CATEGORY_TENANT_SYSTEM_CONFIG:
		return []string{"/api/v1/config/getSiteConfig"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_SIDEBAR:
		return []string{"/api/v1/config/getSidebar"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_PTB,
		common.CacheCategory_CACHE_CATEGORY_TENANT_CURRENCY:
		return []string{"/api/v1/config/site/currency"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_AGREEMENT:
		return []string{"/api/v1/config/getAgreement"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_AGENT_CONFIG:
		return []string{"/api/v1/config/getAgentConfig"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_ACCOUNT_POLICY:
		return []string{"/api/v1/config/getLoginRegisterConfig"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_SOCIAL_MEDIA:
		return []string{"/api/v1/config/getSocialMediaList"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_FAQ:
		return []string{"/api/v1/config/getFaqList"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_CUSTOMER_SERVICE:
		return []string{"/api/v1/config/getCustomerServiceList"}
	case common.CacheCategory_CACHE_CATEGORY_TENANT_CURRENCY_RELATION:
		return []string{"/api/v1/config/getConvertRatios"}
	case common.CacheCategory_CACHE_CATEGORY_PROMOTION_VIP_INFO:
		return []string{"/api/v1/vip/getVipConfig"}
	case common.CacheCategory_CACHE_CATEGORY_GAME_CATEGORY:
		return []string{"/api/v1/game/getGameCategoryList"}
	case common.CacheCategory_CACHE_CATEGORY_GAME_PLATFORM:
		return []string{"/api/v1/game/getGamePlatformList"}
	case common.CacheCategory_CACHE_CATEGORY_GAME_SHOW_RULE,
		common.CacheCategory_CACHE_CATEGORY_GAME_INFO,
		common.CacheCategory_CACHE_CATEGORY_GAME_CURRENCY,
		common.CacheCategory_CACHE_CATEGORY_GAME_DIVERSION:
		return []string{"/api/v1/game/getGameInfo"}
	case common.CacheCategory_CACHE_CATEGORY_GAME_WEIGHT:
		return []string{"/api/v1/game/getGameList"}
	}
	return nil
}
