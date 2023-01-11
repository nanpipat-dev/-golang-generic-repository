package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nanpipat-dev/gorm-generic-repository/config"
	"github.com/nanpipat-dev/gorm-generic-repository/db"
	"github.com/nanpipat-dev/gorm-generic-repository/repository"
)

const dev = "development"

type ExampleGormModel struct {
	ID   string `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func main() {
	cfg, err := config.NewConfig(dev)
	if err != nil {
		fmt.Println("Error reading config file", err)
		os.Exit(1)
	}

	//connect db
	db := db.Connect(cfg)
	ctx := context.Background()

	// example use repository
	newValues := ExampleGormModel{
		ID:   "test",
		Name: "test",
	}
	err = repository.NewRepository[ExampleGormModel](db, ctx).Create(&newValues)
	if err != nil {
		panic(err)
	}

	result, err := repository.NewRepository[ExampleGormModel](db, ctx).FindOne("id = ?", newValues.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
