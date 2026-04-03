package uniqueid

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
)

type RedisSnowflakeGenerator struct {
	rdb    redis.UniversalClient
	node   *sonyflake.Sonyflake
	workId int64
}

var redisSnowflakeWorkerIdCacheKey = "snowflake:worker_id_counter"

func NewRedisSnowflakeGenerator(rdb redis.UniversalClient) *RedisSnowflakeGenerator {
	snow := &RedisSnowflakeGenerator{}
	snow.rdb = rdb
	snow.node = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Date(1997, 1, 14, 0, 0, 0, 0, time.UTC),
		MachineID: snow.machineID,
	})
	return snow
}

func (g *RedisSnowflakeGenerator) NextId() (int64, error) {
	id, err := g.node.NextID()
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

func (g *RedisSnowflakeGenerator) machineID() (uint16, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	workerID, err := g.rdb.Incr(ctx, redisSnowflakeWorkerIdCacheKey).Result()
	if err != nil {
		return 0, err
	}
	workerID = workerID % 1024
	if workerID < 0 {
		workerID = 0
	}
	logx.Infof("✅ 本节点分配到的 workerId = %d", workerID)
	return uint16(workerID), nil
}
