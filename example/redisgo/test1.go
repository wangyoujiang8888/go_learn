package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main()  {
	server := "127.0.0.1:6379"

	option := redis.DialPassword("123456")
	c, err := redis.Dial("tcp", server, option)
	if err != nil {
		log.Println("connect server failed:", err)
		return
	}

	defer c.Close()


	v, err := redis.Int64(c.Do("SADD", "myset", "10.8.37.98"))
	if err != nil {
		log.Println("SADD failed:", err)
		return
	}

	log.Println("value:", v)

}
