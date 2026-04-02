package uniqueid

import (
	"context"
	"testing"

	"github.com/og-saas/framework/stores/redisx"
)

func TestNewRedisSnowflakeGenerator(t *testing.T) {
	redisx.Must(redisx.Config{
		Addrs: []string{"127.0.0.1:6379"},
	})

	snow := NewRedisSnowflakeGenerator(redisx.Engine.RDB(context.Background()))

	for i := 0; i < 10; i++ {
		id, err := snow.NextId()
		if err != nil {
			t.Errorf("next id err:%v", err)
			return
		}
		t.Log(id)
	}

}
