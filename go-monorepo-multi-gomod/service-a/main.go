package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("test", "TEST_ENV")
	fmt.Println("Hello world from service-a !")
}
