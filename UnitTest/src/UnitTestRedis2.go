package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func ExampleNewClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	/*
		err22 := rdb.Set(ctx, "Go-Testkey", "Go-Redis-TEST20201005140000003", 60*60*time.Second).Err()
		if err22 != nil {
			panic(err22)
		}
		val22, err22 := rdb.Get(ctx, "Go-Testkey").Result()
		if err22 != nil {
			panic(err22)
		}
		fmt.Println("Go-Testkey", val22)
		fmt.Println("time.Second : ", 60*60*time.Second)
	*/
	/*
		err := rdb.SetNX(ctx, "Go-Testkey", "Go-Redis-TEST20201005140000003", 0).Err()
		if err != nil {
			panic(err)
		}
		val, err := rdb.Get(ctx, "Go-Testkey").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("Go-Testkey", val)
	*/
	/*
		val2, err := rdb.Get(ctx, "Go-Testkey2").Result()
		if err == redis.Nil {
			fmt.Println("Go-Testkey2 does not exist")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("Go-Testkey2", val2)
		}
		// Output: key value
		// key2 does not
	*/
	/*
		fmt.Println("==================================================================")

		// SET key value EX 10 NX
		set, err := rdb.SetNX(ctx, "Go-Testkey3", "Go-Redis-TEST20201005140000000", 10*time.Second).Result()

		val3, err := rdb.Get(ctx, "Go-Testkey3").Result()
		if err == redis.Nil {
			fmt.Println("Go-Testkey3 does not exist")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("Go-Testkey3", val3)
		}
	*/
	/*
		// SET key value keepttl NX
		set, err := rdb.SetNX(ctx, "Go-Testkey3", "value", redis.KeepTTL).Result()
	*/
	/*
		// SORT list LIMIT 0 2 ASC
		vals, err := rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()

		fmt.Println("Go-list", vals)
	*/
	/*
		// ZRANGEBYSCORE zset -inf +inf WITHSCORES LIMIT 0 2
		vals, err := rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{
			Min: "-inf",
			Max: "+inf",
			Offset: 0,
			Count: 2,
		}).Result()

		// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
		vals, err := rdb.ZInterStore(ctx, "out", &redis.ZStore{
			Keys: []string{"zset1", "zset2"},
			Weights: []int64{2, 3}
		}).Result()

		// EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
		vals, err := rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()

		// custom command
		res, err := rdb.Do(ctx, "set", "key", "value").Result()
	*/
}

func main() {
	ExampleClient()
}
