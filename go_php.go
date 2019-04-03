package main

import (
	"github.com/deuill/go-php"
	"os"
)

func main() {
	engine, _ := php.New()

	context, _ := engine.NewContext()
	context.Output = os.Stdout

	context.Exec("index.php")
	engine.Destroy()
}