package main

import (
	"fmt"
	"github.com/dengpju/higo-config/config"
)

func main() {

	config.Set("hh.gg.tt.hh", "1")
	config.Set("hh.jj.tt1.hh2", "2")
	config.Set("hh.jj.tt1.hh3", 3)
	config.Set("hh1", 3)
	fmt.Println(config.All())
	fmt.Println(config.Get("hh"))
	fmt.Println(config.String("hh.jj.tt1.hh2"))
	fmt.Println(config.Int("hh.jj.tt1.hh3"))
	fmt.Println(config.Get("1"))
	fmt.Println(config.Get("hh.jj.tt1.1"))

	/**

	gg := configure.New()
	gg.Set("gg", configure.New().Set("gg1", "ggg"))
	fmt.Println(gg)

	tt := configure.New()
	tt.Set("tt", gg)
	fmt.Println(tt)

	configure.Set("hh", tt)

	fmt.Println(configure.All())
	fmt.Println(configure.Get("hh"))
	fmt.Println(configure.String("hh.tt.gg.gg1"))

	 */
}
