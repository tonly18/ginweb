package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"runtime"
	"server/config"
	"server/core/xerror"
	"time"
)

//RedisPoolConn Struct
type RedisPoolConn struct {
	rd  *redis.Client
	ctx context.Context
}

//redis conn
var redisConn *redis.Client

//NewRedis
func NewRedis(ctx context.Context) *RedisPoolConn {
	if redisConn == nil {
		if err := createRedisCluster(ctx); err != nil {
			return nil
		}
	}

	//return
	return &RedisPoolConn{
		rd:  redisConn,
		ctx: ctx,
	}
}

//init
func init() {
	if err := createRedisCluster(context.TODO()); err != nil {
		panic("[create new redis cluster error: " + err.Error() + "]")
	} else {
		fmt.Println("[redis init successfully] host:", config.Config.Redis.Host)
	}
}

//create redis cluster
func createRedisCluster(ctx context.Context) xerror.Error {
	redisConn = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host,
		Password: config.Config.Redis.Password,

		//连接池容量及闲置连接数量
		PoolSize:     100, //链接池最大链接数，默认为cup * 5。
		MinIdleConns: 50,  //在启动阶段，链接池最小链接数，并长期维持idle状态的链接数不少于指定数量。
		//超时设置
		DialTimeout:  5 * time.Second, //建立链接超时时间，默认为5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒，-1表示取消读超时。
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时。
		PoolTimeout:  5 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
		//闲置链接检查
		IdleCheckFrequency: 60 * time.Second,   //闲置链接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
		IdleTimeout:        1200 * time.Second, //闲置链接超时时长，默认5分钟，-1表示取消闲置超时。
		MaxConnAge:         0 * time.Second,    //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接。
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
		return xerror.Wrap(nil, &xerror.NewError{
			Code: 110000,
			Err:  err,
		})
	}

	runtime.SetFinalizer(redisConn, func(conn *redis.Client) {
		conn.Close()
	})

	//return
	return nil
}

func (d *RedisPoolConn) HGetRd(key, id string) (string, xerror.Error) {
	data, err := d.rd.HGet(context.Background(), key, id).Result()
	if err == redis.Nil { //Key不存在
		return "", xerror.Wrap(nil, &xerror.NewError{
			Code: 110020,
			Err:  redis.Nil,
		})
	} else if err != nil { //panic(err)
		return "", xerror.Wrap(nil, &xerror.NewError{
			Code: 110022,
			Err:  err,
		})
	} else {
		return data, nil
	}
}

func (d *RedisPoolConn) HSetRd(key string, id, value any) xerror.Error {
	_, err := d.rd.HSet(context.Background(), key, id, value).Result()
	if err != nil {
		return xerror.Wrap(nil, &xerror.NewError{
			Code: 110030,
			Err:  err,
		})
	}

	return nil
}

func (d *RedisPoolConn) HDelRd(key, id string) xerror.Error {
	_, err := d.rd.HDel(context.Background(), key, id).Result()
	if err != nil {
		return xerror.Wrap(nil, &xerror.NewError{
			Code: 110040,
			Err:  err,
		})
	}

	return nil
}

func (d *RedisPoolConn) GetRd(key string) (string, xerror.Error) {
	data, err := d.rd.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", xerror.Wrap(nil, &xerror.NewError{
			Code: 110050,
			Err:  redis.Nil,
		})
	} else if err != nil {
		return "", xerror.Wrap(nil, &xerror.NewError{
			Code: 110051,
			Err:  err,
		})
	} else {
		return data, nil
	}
}

func (d *RedisPoolConn) SetRd(key string, value any) xerror.Error {
	if _, err := d.rd.Set(context.Background(), key, value, time.Hour*24*30).Result(); err != nil {
		return xerror.Wrap(nil, &xerror.NewError{
			Code: 110060,
			Err:  err,
		})
	}

	return nil
}

func (d *RedisPoolConn) DelRd(key string) xerror.Error {
	if _, err := d.rd.Del(context.Background(), key).Result(); err != nil {
		return xerror.Wrap(nil, &xerror.NewError{
			Code: 110070,
			Err:  err,
		})
	}

	return nil
}
