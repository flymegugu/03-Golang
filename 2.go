/*
 * @Author: [MaxGu]
 * @Date: 2022-07-19 17:24:09
 * @LastEditors: [MaxGu]
 * @LastEditTime: 2022-07-19 17:43:04
 * @Description:
 */
package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "test check")
}
func main() {
	flag.Parse()
	fmt.Println("hello,%s!\n", name)
}
