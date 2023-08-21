package dao

import (
	"context"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"server/config"
	"server/core/xerror"
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

func createMC(ctx context.Context) xerror.Error {
	mcClient = memcache.New(strings.Split(config.Config.Memcache.Host, ";")...)
	if err := mcClient.Ping(); err != nil {
		return xerror.Wrap(ctx, nil, &xerror.NewError{
			Code: 120000,
			Err:  err,
		})
	}

	return nil
}

func (d *MCClient) Set(key string, data []byte, expire ...int32) xerror.Error {
	expiration := int32(0)
	if len(expire) > 0 {
		expiration = expire[0]
	}

	if err := d.mc.Set(ItemStruck{
		Item: &memcache.Item{Key: key, Value: data, Expiration: expiration},
	}.Item); err != nil {
		return xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120010,
			Err:  err,
		})
	}

	return nil
}

func (d *MCClient) Get(key string) (*ItemStruck, xerror.Error) {
	if item, err := d.mc.Get(key); err != nil {
		return nil, xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120020,
			Err:  err,
		})
	} else {
		return &ItemStruck{item}, nil
	}
}

func (d *MCClient) MGet(keys []string) (map[string]*ItemStruck, xerror.Error) {
	items, err := d.mc.GetMulti(keys)
	if err != nil {
		return nil, xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120030,
			Err:  err,
		})
	}

	data := make(map[string]*ItemStruck, len(items))
	for _, v := range items {
		data[v.Key] = &ItemStruck{
			Item: v,
		}
	}

	return data, nil
}

func (d *MCClient) Add(key string, data []byte, expire ...int32) xerror.Error {
	expiration := int32(0)
	if len(expire) > 0 {
		expiration = expire[0]
	}

	if err := d.mc.Add(ItemStruck{
		Item: &memcache.Item{Key: key, Value: data, Expiration: expiration},
	}.Item); err != nil {
		return xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120040,
			Err:  err,
		})
	}

	return nil
}

func (d *MCClient) Replace(key string, data []byte, expire ...int32) xerror.Error {
	expiration := int32(0)
	if len(expire) > 0 {
		expiration = expire[0]
	}

	if err := d.mc.Replace(ItemStruck{
		Item: &memcache.Item{Key: key, Value: data, Expiration: expiration},
	}.Item); err != nil {
		return xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120050,
			Err:  err,
		})
	}

	return nil
}

func (d *MCClient) Increment(key string, delta uint64) (uint64, xerror.Error) {
	newValue, err := d.mc.Increment(key, delta)
	if err != nil {
		return 0, xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120060,
			Err:  err,
		})
	}

	return newValue, nil
}

func (d *MCClient) Decrement(key string, delta uint64) (uint64, xerror.Error) {
	newValue, err := d.mc.Increment(key, delta)
	if err != nil {
		return 0, xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120065,
			Err:  err,
		})
	}

	return newValue, nil
}

func (d *MCClient) Expire(key string, seconds int32) xerror.Error {
	if err := d.mc.Touch(key, seconds); err != nil {
		return xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120070,
			Err:  err,
		})
	}

	return nil
}

func (d *MCClient) SetWithExpire(key string, data []byte, expire int32) xerror.Error {
	if err := d.mc.Set(ItemStruck{
		Item: &memcache.Item{Key: key, Value: data, Expiration: expire},
	}.Item); err != nil {
		return xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120075,
			Err:  err,
		})
	}

	return nil
}

func (d *MCClient) Del(key string) xerror.Error {
	if err := d.mc.Delete(key); err != nil {
		return xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120080,
			Err:  err,
		})
	}
	return nil
}

func (d *MCClient) DelAll() xerror.Error {
	if err := d.mc.DeleteAll(); err != nil {
		return xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120085,
			Err:  err,
		})
	}

	return nil
}

func (d *MCClient) Flush() xerror.Error {
	if err := d.mc.FlushAll(); err != nil {
		return xerror.Wrap(d.ctx, nil, &xerror.NewError{
			Code: 120090,
			Err:  err,
		})
	}
	return nil
}
