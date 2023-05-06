/*
 * error code: 30003000 ` 30003999
 */

package dao

import (
	"context"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"server/config"
	"server/core/logger"
	"strings"
)

type ItemStruck struct {
	*memcache.Item
}

type MCClient struct {
	mc  *memcache.Client
	ctx context.Context
}

var mcClient *memcache.Client

func NewMCClient(ctx context.Context) *MCClient {
	if mcClient == nil {
		if err := createMC(ctx); err != nil {
			logger.Error(ctx, fmt.Sprintf(`[1100600] create new memcached error: %v`, err))
		}
	}

	//return
	return &MCClient{
		mc:  mcClient,
		ctx: ctx,
	}
}

func init() {
	if err := createMC(context.TODO()); err != nil {
		panic("[create new memcached error: " + err.Error() + "]")
	} else {
		fmt.Println("[memcache init successfully] host:", config.Config.Memcache.Host)
	}
}

func createMC(ctx context.Context) error {
	mcClient = memcache.New(strings.Split(config.Config.Memcache.Host, ";")...)
	if err := mcClient.Ping(); err != nil {
		logger.Error(ctx, fmt.Sprintf(`[1100605] memcached ping error: %v`, err.Error()))
		return err
	}

	return nil
}

func (d *MCClient) Set(key string, data []byte, expire ...int32) error {
	expiration := int32(0)
	if len(expire) > 0 {
		expiration = expire[0]
	}

	if err := d.mc.Set(ItemStruck{
		Item: &memcache.Item{Key: key, Value: data, Expiration: expiration},
	}.Item); err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1100610] set key: %v, value: %v, error: %v`, key, string(data), err))
		return err
	}

	return nil
}

func (d *MCClient) Get(key string) (*ItemStruck, error) {
	if item, err := d.mc.Get(key); err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1100615] get key: %v, error: %v`, key, err))
		return nil, err
	} else {
		return &ItemStruck{item}, nil
	}
}

func (d *MCClient) MGet(keys []string) (map[string]*ItemStruck, error) {
	items, err := d.mc.GetMulti(keys)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1100620] mget key: %v, error: %v`, strings.Join(keys, ","), err))
		return nil, err
	}

	data := make(map[string]*ItemStruck, len(items))
	for _, v := range items {
		data[v.Key] = &ItemStruck{
			Item: v,
		}
	}

	return data, nil
}

func (d *MCClient) Add(key string, data []byte, expire ...int32) error {
	expiration := int32(0)
	if len(expire) > 0 {
		expiration = expire[0]
	}

	if err := d.mc.Add(ItemStruck{
		Item: &memcache.Item{Key: key, Value: data, Expiration: expiration},
	}.Item); err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1100630] add key: %v, value: %v, error: %v`, key, data, err))
		return err
	}

	return nil
}

func (d *MCClient) Replace(key string, data []byte, expire ...int32) error {
	expiration := int32(0)
	if len(expire) > 0 {
		expiration = expire[0]
	}

	if err := d.mc.Replace(ItemStruck{
		Item: &memcache.Item{Key: key, Value: data, Expiration: expiration},
	}.Item); err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1100640] replace key: %v, value: %v, error: %v`, key, data, err))
		return err
	}

	return nil
}

func (d *MCClient) Increment(key string, delta uint64) (newValue uint64, err error) {
	newVAlue, err := d.mc.Increment(key, delta)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1100650] expire key: %v, delta: %v, error: %v`, key, delta, err))
		return 0, err
	}

	return newVAlue, nil
}

func (d *MCClient) Decrement(key string, delta uint64) (newValue uint64, err error) {
	newVAlue, err := d.mc.Increment(key, delta)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1100660] expire key: %v, delta: %v, error: %v`, key, delta, err))
		return 0, err
	}

	return newVAlue, nil
}

func (d *MCClient) Expire(key string, seconds int32) error {
	if err := d.mc.Touch(key, seconds); err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1100670] expire key: %v, seconds: %v, error: %v`, key, seconds, err))
		return err
	}

	return nil
}

func (d *MCClient) SetWithExpire(key string, data []byte, expire int32) error {
	if err := d.mc.Set(ItemStruck{
		Item: &memcache.Item{Key: key, Value: data, Expiration: expire},
	}.Item); err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1100680] expire key: %v, value: %v, error: %v`, key, data, err))
		return err
	}

	return nil
}

func (d *MCClient) Del(key string) error {
	return d.mc.Delete(key)
}

func (d *MCClient) DelAll() error {
	return d.mc.DeleteAll()
}

func (d *MCClient) Flush() error {
	return d.mc.FlushAll()
}
