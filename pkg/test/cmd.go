package test

import "fmt"

type cmd struct {

}

func (c *cmd)Run()  {
	fmt.Println("tset")
}

func NewCmd() *cmd {
	return &cmd{}
}