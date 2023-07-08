package dao

import (
	"context"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"server/config"
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
			return nil
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
		return err
	}

	return nil
}

func (d *MCClient) Get(key string) (*ItemStruck, error) {
	if item, err := d.mc.Get(key); err != nil {
		return nil, err
	} else {
		return &ItemStruck{item}, nil
	}
}

func (d *MCClient) MGet(keys []string) (map[string]*ItemStruck, error) {
	items, err := d.mc.GetMulti(keys)
	if err != nil {
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
		return err
	}

	return nil
}

func (d *MCClient) Increment(key string, delta uint64) (newValue uint64, err error) {
	newVAlue, err := d.mc.Increment(key, delta)
	if err != nil {
		return 0, err
	}

	return newVAlue, nil
}

func (d *MCClient) Decrement(key string, delta uint64) (newValue uint64, err error) {
	newVAlue, err := d.mc.Increment(key, delta)
	if err != nil {
		return 0, err
	}

	return newVAlue, nil
}

func (d *MCClient) Expire(key string, seconds int32) error {
	if err := d.mc.Touch(key, seconds); err != nil {
		return err
	}

	return nil
}

func (d *MCClient) SetWithExpire(key string, data []byte, expire int32) error {
	if err := d.mc.Set(ItemStruck{
		Item: &memcache.Item{Key: key, Value: data, Expiration: expire},
	}.Item); err != nil {
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
