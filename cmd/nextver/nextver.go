// Copyright The nextver Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nextver

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/craftypath/nextver/pkg/version"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	GitCommit = "HEAD"
	BuildDate = "unknown"
)

func Execute() {
	command := newNextverCommand()
	if err := command.Execute(); err != nil {
		var exitCode int
		if err, ok := err.(*exec.ExitError); ok {
			exitCode = err.ExitCode()
		} else {
			exitCode = 1
		}
		os.Exit(exitCode)
	}
}

func newNextverCommand() *cobra.Command {
	var currentVersion string
	var pattern string

	fullVersion := fmt.Sprintf("%s (commit=%s, date=%s)", Version, GitCommit, BuildDate)
	command := &cobra.Command{
		Use:     "nextver [flags]",
		Short:   "nextver manages automatic semver versioning",
		Version: fullVersion,
		RunE: func(_ *cobra.Command, _ []string) error {
			next, err := version.Next(currentVersion, pattern)
			if err != nil {
				return err
			}
			fmt.Println(next)
			return nil
		},
	}

	command.Flags().StringVarP(&currentVersion, "current-version", "c", "", "the current version")
	command.Flags().StringVarP(&pattern, "pattern", "p", "", "the versioning pattern")
	command.SetVersionTemplate("{{ .Version }}\n")
	command.SilenceUsage = true
	if err := command.MarkFlagRequired("current-version"); err != nil {
		log.Fatalln(err)
	}
	if err := command.MarkFlagRequired("pattern"); err != nil {
		log.Fatalln(err)
	}
	return command
}
