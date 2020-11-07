package main

//go:generate go run cmd/main.go --gdnative --types --classes

import (
	_ "git.eko.ovh/godoai/pkg/export"
	_ "git.eko.ovh/godoai/pkg/godoai"
)

func main() {
}
