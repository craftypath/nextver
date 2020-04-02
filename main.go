package main

import (
    "fmt"
    "os"
    "verit/pkg/version"
)

const usage = `Usage: verit <current-version> <pattern>

Example: verit 1.0.0 1.x.0 # prints 1.1.0

<current-version>: SemVer denoting the current version of your artifact
<pattern>: SemVer denoting the next version of your artifact. One of <major>.<minor>.<patch> may be set to "x" to increment from current-version.
`

func main() {
    args := os.Args[1:]
    if len(args) != 2 {
        exitWithMessage(usage)
    }
    next, err := version.Next(args[0], args[1])
    if err != nil {
        exitWithMessage(usage)
    }
    fmt.Println(next)
}

func exitWithMessage(message string) {
    _, _ = os.Stderr.WriteString(message)
    os.Exit(1)
}
