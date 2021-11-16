package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	student := map[string]string{
		"id":   "st01",
		"name": "namme1",
	}

	set("key1", student, 0)
	get("key1")

}

func set(key string, value map[string]string, ttl int) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := client.Set(key, value, 0).Err()
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
