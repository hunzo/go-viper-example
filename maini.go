package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type M struct {
	Staff Model
	Stu   Model
}

type Model struct {
	Server    string
	Attribute []string
	Username  string
	Password  string
	Id        int
}

func main() {
	initConfig()
	user, e := viperExUnmarshal()
	if e != nil {
		log.Printf("%v", e)
	}

	fmt.Printf("%v\n", user.Staff)
	fmt.Printf("%T\n", user.Staff.Attribute)
	fmt.Printf("%v\n", user.Stu)
}

func viperExUnmarshal() (*M, error) {
	m := M{}
	err := viper.Unmarshal(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
