package main

import "github.com/kyokomi/emoji"

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}

func main() {
	emoji.Print(GetMessage())
}
