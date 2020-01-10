package main

import (
	"example/go_cache2"
	"fmt"
	"time"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := go_cache2.New(6*time.Second, 10*time.Minute)
	// Set the value of the key "foo" to "bar", with the default expiration time
	fmt.Println("start time:",time.Now().Format("2006-01-02 15:04:05.000"))
	c.Set("foo", "bar", 1*time.Hour)
	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}
	expire,ok := c.GetExpiredTime("foo")
	if ok{
		fmt.Println(expire.Format("2006-01-02 15:04:05.000"))
	}
	time.Sleep(5*time.Second)
	foo, found = c.Get("foo")
	if !found {
		fmt.Println("not found foo")
	}else{
		fmt.Println("found foo",foo)
	}
	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	c.Set("baz", 42, go_cache2.NoExpiration)
	c.Delete("baz")
	_, found = c.Get("baz")
	if !found {
		fmt.Println("not found baz")
	}

}
