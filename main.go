package main

import (
	"embed"
	"encoding/json"
	"fmt"

	bootstrap "test/src"
)

//go:embed assets/*
var assets_embed embed.FS

//go:embed build/*
var ui_embed embed.FS

type Math struct {
}

func (ctx *Math) Addition(a int32, b int32) int32 {
	return a + b
}

type Movement struct {
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (ctx *Movement) Change(position string) bool {
	var coordinate Coordinate

	err := json.Unmarshal([]byte(position), &coordinate)

	if err != nil {
		fmt.Println(err, ":", position)
		return false
	}

	fmt.Println("Position Change:", coordinate)

	return true
}

func main() {
	configuration := bootstrap.Configuration{
		UI_EMBED:     &ui_embed,
		ASSETS_EMBED: &assets_embed,
		DEBUG:        true,
	}

	configuration.Bootstrap(func(application bootstrap.Application) {
		application.Bind("Math2", &Math{})
		application.Bind("Movement", &Movement{})
	})
}
