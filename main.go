package main

import (
	"fmt"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Detail struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Detail{Name: "kkb", Age: 20})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
