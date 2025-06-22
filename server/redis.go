package server

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

// func init() {
// 	rdb = redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379", // todo: make configurable
// 		Password: "",               // no password set
// 		DB:       0,                // use default DB
// 	})
// }

type KombuchaStorage interface {
	GetCurrentCount(ctx context.Context) (*string, error) // using pointer to be able to return no value (nil)
	GetCurrentStats(ctx context.Context) (*map[string]string, error)
	IncCount(ctx context.Context) error
}

type RedisKombuchaStorage struct{}

const RedisHashName string = "kombucha:0"
const RedisCountKey string = "count"
const RedisMlKey string = "ml"
const RedisSavedCurrencyKey string = "saved"

//todo: there is a way to read redis-hash to struct: https://redis.uptrace.dev/guide/scanning-hash-fields.html

// attempts to preconfigure redis Set befor using it - not sure it is needed
func NewRedisKombuchaStorage() *RedisKombuchaStorage {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // todo: make configurable
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	ctx := context.Background()
	if !hashExists(ctx) {
		rdb.HSet(ctx, RedisHashName, RedisCountKey, 0) // initializing the value in redis
	}

	return &RedisKombuchaStorage{}
}

// returns current count of Viver Kombucha bottles
func (ks *RedisKombuchaStorage) GetCurrentCount(ctx context.Context) (*string, error) {
	val, err := rdb.HGet(ctx, RedisHashName, RedisCountKey).Result()

	switch {
	case err == redis.Nil:
		log.Printf("key=%s does not exist /n", RedisCountKey)
		return nil, err
	case err != nil:
		log.Println("Get failed", err)
		panic(err)
	case val == "":
		log.Println("value is empty")
		return nil, nil
	}

	return &val, nil
}

func (ks *RedisKombuchaStorage) GetCurrentStats(ctx context.Context) (*map[string]string, error) {
	val, err := rdb.HGetAll(ctx, RedisHashName).Result()

	switch {
	case err == redis.Nil:
		log.Printf("hash=%s does not exist \n", RedisHashName)
		return nil, err
	case err != nil:
		log.Println("Failed to get all values from hash", err)
		panic(err)
	case len(val) == 0:
		log.Printf("hash=%s value is empty \n", RedisHashName)
		return nil, nil
	}

	return &val, err
}

// increases current counter value by one bottle of Viver Kombucha
// values of ml, price are predefined
func (ks *RedisKombuchaStorage) IncCount(ctx context.Context) error {
	// curCount := ks.getCurrentCountAsInt(ctx)
	// // by default increasing by 330ml which is the size of a standard Viver kombucha bottle
	// newCount, err := rdb.HSet(
	// 	ctx,
	// 	RedisHashName,
	// 	RedisCountKey, curCount+330).Result() // todo: make configurable

	newCount, err := rdb.HIncrBy(ctx, RedisHashName, RedisCountKey, 1).Result()
	if err != nil {
		log.Println("Increase count failed", err)
		return err
	}
	_, err = rdb.HIncrBy(ctx, RedisHashName, RedisMlKey, 330).Result() // todo: make configurable
	if err != nil {
		log.Println("Increase Ml failed", err)
		return err
	}
	// todo: currency limitation by int type of Redis Increase function
	_, err = rdb.HIncrByFloat(ctx, RedisHashName, RedisSavedCurrencyKey, 2.85).Result() // todo: make configurable
	if err != nil {
		log.Println("Inc currency save failed", err)
		return err
	}

	log.Printf("Set updated current count to %d\n", newCount)
	return nil
}

// func handleRedisError(err error) {
// 	switch {
// 	case err == redis.Nil:
// 		log.Println("key does not exist")
// 		return nil, err
// 	case err != nil:
// 		log.Println("Get failed", err)
// 		panic(err)
// 	}
// }

func hashExists(ctx context.Context) bool {
	_, err := rdb.HGet(ctx, RedisHashName, RedisCountKey).Result()
	if err == nil {
		return true
	}
	return false
}
