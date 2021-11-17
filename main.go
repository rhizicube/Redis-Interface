package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {

	student := map[string]string{
		"id":   "st01",
		"name": "namme1",
	}

	bs, _ := json.Marshal(student)

	set("key1", bs, 0)
	get("key1")

}

func set(key string, value interface{}, ttl int) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := client.Set(key, value, time.Duration(ttl)).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func get(key string) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(val)
	return true
}
