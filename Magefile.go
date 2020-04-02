//+build mage

package main

import (
    "github.com/magefile/mage/sh"
)

func Build() error {
    return sh.RunV("go", "build", "-o", "build/verit", "cmd/verit.go")
}

func Test() error {
    return sh.RunV("go", "test", "./...")
}
