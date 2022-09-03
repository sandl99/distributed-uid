package main

import (
	"fmt"

	"github.com/sandl99/distrbuted-uid/snowflake"
)

func main() {
	snowflake.InfoLogger.Println("Starting the application...")

	constant, err := snowflake.InitConstant()
	if err != nil {
		snowflake.ErrorLogger.Fatal(err)
	}

	snowflaker, err := snowflake.NewNode(1, constant)
	if err != nil {
		snowflake.ErrorLogger.Fatal(err)
	}

	fmt.Println(snowflaker.Generate().Int64())
	snowflake.InfoLogger.Println("Stopping the application...")
}
