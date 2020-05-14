//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	return sh.RunV("go", "build", "-o", "build/nextver", "main.go")
}

func Test() error {
	return sh.RunV("go", "test", "./...")
}
