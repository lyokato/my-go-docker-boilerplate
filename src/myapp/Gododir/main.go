package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("run_web", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/web"})
	}).Src("*.go", "**/*.go").Debounce(3000)
}

func main() {
	do.Godo(tasks)
}
