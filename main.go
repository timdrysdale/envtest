package main

import (
	"fmt"

	"github.com/ory/viper"
)

func main() {
	viper.SetEnvPrefix("ROOTCMD")
	viper.AutomaticEnv()

	fmt.Println("Set these two environment variables, e.g")
	fmt.Println("export ROOTCMD_TEST=bar")
	fmt.Println("export SUBCMD_TEST=pong")
	fmt.Println()
	fmt.Println("------------RESULTS--------------------")

	fmt.Printf("\nmain()::viper.Get(\"test\")                   = %s\n", viper.Get("test"))

	c := make(chan struct{})

	go func() {
		fmt.Printf("main()::go func()::viper.Get(\"test\")        = %s\n", viper.Get("test"))
		close(c)
	}()
	<-c

	tryagain()
	trydiff()
}

func tryagain() {
	fmt.Printf("main()::func()::viper.Get(\"test\")           = %s\n", viper.Get("test"))

}

func trydiff() {
	fmt.Println("main()::func()::viper.SetEnvPrefix(\"SUBCMD\")")
	viper.SetEnvPrefix("SUBCMD")
	fmt.Printf("main()::func()::viper.Get(\"test\")           = %s\n", viper.Get("test"))
	viper.AutomaticEnv()

}
