package dao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/tonly18/xerror"
	"runtime"
	"sync"
	"time"
)

// RedisPoolConn Struct
type RedisPoolConn struct {
	rd  *redis.Client
	ctx context.Context
}

var once sync.Once
var redisConn *redis.Client

// NewRedis
func NewRedis(ctx context.Context, rdConfig *RedisConfig) *RedisPoolConn {
	redis := &RedisPoolConn{
		ctx: ctx,
	}
	once.Do(func() {
		if err := redis.createRedisCluster(rdConfig); err != nil {
			panic(fmt.Errorf(`redis connect happened error:%v`, err))
		}
	})
	if redisConn == nil {
		redis.createRedisCluster(rdConfig)
	}
	redis.rd = redisConn

	//return
	return redis
}

// create redis cluster
func (d *RedisPoolConn) createRedisCluster(rdConfig *RedisConfig) xerror.Error {
	redisConn = redis.NewClient(&redis.Options{
		Addr:     rdConfig.Host,
		Password: rdConfig.Password,
		DB:       rdConfig.DB,

		//连接池容量及闲置连接数量
		PoolSize:     rdConfig.PoolSize,     //链接池最大链接数，默认为cup * 5。
		MinIdleConns: rdConfig.MinIdleConns, //在启动阶段，链接池最小链接数，并长期维持idle状态的链接数不少于指定数量。
		MaxIdleConns: rdConfig.MaxIdleConns,
		//超时设置
		DialTimeout:     5 * time.Second,    //建立链接超时时间，默认为5秒。
		ReadTimeout:     3 * time.Second,    //读超时，默认3秒，-1表示取消读超时。
		WriteTimeout:    3 * time.Second,    //写超时，默认等于读超时。
		PoolTimeout:     5 * time.Second,    //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
		ConnMaxLifetime: 3600 * time.Second, //链接存活时长
		//命令执行失败时的重试策略
		MaxRetries:      3,                      //命令执行失败时，最多重试多少次，默认为0即不重试。
		MinRetryBackoff: 8 * time.Microsecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔。
		MaxRetryBackoff: 512 * time.Microsecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔。

		//仅当客户端执行命令时需要从连接池获取连接时，如果连接池需要新建连接时则会调用此钩子函数。
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			return nil
		},
	})
	if _, err := redisConn.Ping(context.Background()).Result(); err != nil {
		return xerror.Wrap(&xerror.NewError{
			Code:     110000,
			RawError: err,
		}, nil)
	}

	runtime.SetFinalizer(redisConn, func(conn *redis.Client) {
		conn.Close()
	})
	runtime.KeepAlive(redisConn)

	//return
	return nil
}

func (d *RedisPoolConn) HGetRd(key, id string) (string, xerror.Error) {
	data, err := d.rd.HGet(context.Background(), key, id).Result()
	if err == redis.Nil { //Key不存在
		return "", xerror.Wrap(&xerror.NewError{
			Code:     110020,
			RawError: redis.Nil,
		}, nil)
	} else if err != nil { //panic(err)
		return "", xerror.Wrap(&xerror.NewError{
			Code:     110022,
			RawError: err,
		}, nil)
	} else {
		return data, nil
	}
}

func (d *RedisPoolConn) HSetRd(key string, id, value any) xerror.Error {
	_, err := d.rd.HSet(context.Background(), key, id, value).Result()
	if err != nil {
		return xerror.Wrap(&xerror.NewError{
			Code:     110030,
			RawError: err,
		}, nil)
	}

	return nil
}

func (d *RedisPoolConn) HDelRd(key, id string) xerror.Error {
	_, err := d.rd.HDel(context.Background(), key, id).Result()
	if err != nil {
		return xerror.Wrap(&xerror.NewError{
			Code:     110040,
			RawError: err,
		}, nil)
	}

	return nil
}

func (d *RedisPoolConn) GetRd(key string) (string, xerror.Error) {
	data, err := d.rd.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", xerror.Wrap(&xerror.NewError{
			Code:     110050,
			RawError: redis.Nil,
		}, nil)
	} else if err != nil {
		return "", xerror.Wrap(&xerror.NewError{
			Code:     110051,
			RawError: err,
		}, nil)
	} else {
		return data, nil
	}
}

func (d *RedisPoolConn) SetRd(key string, value any) xerror.Error {
	if _, err := d.rd.Set(context.Background(), key, value, time.Hour*24*30).Result(); err != nil {
		return xerror.Wrap(&xerror.NewError{
			Code:     110060,
			RawError: err,
		}, nil)
	}

	return nil
}

func (d *RedisPoolConn) DelRd(key string) xerror.Error {
	if _, err := d.rd.Del(context.Background(), key).Result(); err != nil {
		return xerror.Wrap(&xerror.NewError{
			Code:     110070,
			RawError: err,
		}, nil)
	}

	return nil
}
