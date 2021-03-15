package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func rClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return client
}

func ping(client *redis.Client) error {
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}

func set(client *redis.Client) error {
	err := client.Set("name", "risa", 0).Err()
	if err != nil {
		return err
	}

	err = client.Set("country", "Philippines", 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func get(client *redis.Client) error {
	nameVal, err := client.Get("name").Result()
	if err != nil {
		return (err)
	}
	fmt.Println("name", nameVal)

	countryVal, err := client.Get("country").Result()
	if err == redis.Nil {
		fmt.Println("no value found")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("country", countryVal)
	}

	return nil
}

func main() {

	// creates a client
	client := rClient()

	// check connection status
	err := ping(client)
	if err != nil {
		fmt.Println(err)
	}

	// Using the SET command to set Key-value pair
	err = set(client)
	if err != nil {
		fmt.Println(err)
	}

	// Using the GET command to get values from keys
	err = get(client)
	if err != nil {
		fmt.Println(err)
	}

}

//https://pkg.go.dev/github.com/go-redis/redis
// go mod init github.com/my/repo
// go mod tidy
