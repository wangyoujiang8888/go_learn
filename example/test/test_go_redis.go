package main

import (
	UserCacheTest "example/go_redis_test"
	"log"
)

func main() {
	user, _ := UserCacheTest.GetUserById(4530)
	const (
		Ldate         = 1 << iota     // 1 << 0 当地时区的日期: 2009/01/23
		Ltime                         // 1 << 1 当地时区的时间: 01:23:23
		Lmicroseconds                 // 1 << 2 显示精度到微秒: 01:23:23.123123 (应该和Ltime一起使用)
		Llongfile                     // 1 << 3 显示完整文件路径和行号: /a/b/c/d.go:23
		Lshortfile                    // 1 << 4 显示当前文件名和行号: d.go:23 (如果与Llongfile一起出现，此项优先)
		LUTC                          // 1 << 5如果设置了Ldata或者Ltime, 最好使用 UTC 时间而不是当地时区
		LstdFlags     = Ldate | Ltime // 标准日志器的初始值
	)
	var b []byte
	Itoa(&b,2020,4)
	log.Println('0'+2)

	log.Println(Ldate)

	log.Print(user)
}

func Itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}
