package gormx

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/sharding"
)

var shardingCount int64

func initShardingDB(db *gorm.DB, tables ...any) {
	db.Use(sharding.Register(sharding.Config{
		ShardingKey: "site_id",
		ShardingAlgorithm: func(columnValue any) (string, error) {
			var (
				siteId int64
			)
			switch v := columnValue.(type) {
			case int64:
				siteId = v
			case uint64:
				siteId = int64(v)
			case int:
				siteId = int64(v)
			default:
				return "", fmt.Errorf("unsupported created_at type: %T", columnValue)
			}
			return ShardingSuffix(siteId), nil
		},
		PrimaryKeyGenerator: sharding.PKCustom,
		PrimaryKeyGeneratorFn: func(tableIdx int64) int64 {
			return 0
		},
	}, tables...))
}

func ShardingSuffix(siteId int64) string {
	return fmt.Sprintf("_%d", siteId%int64(shardingCount))
}

type ShardingInterface interface {
	GetSiteId() int64
}

func ShardingData[T ShardingInterface](data []T) map[string][]T {
	ret := make(map[string][]T)
	for _, item := range data {
		key := ShardingSuffix(item.GetSiteId())
		ret[key] = append(ret[key], item)
	}
	return ret
}
