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

package version

import (
	"fmt"
	"strconv"
	"strings"
)

const positionMajor = 0
const positionMinor = 1
const positionPatch = 2

const incrementChar = "x"
const wildcardChar = "?"
const vPrefix = "v"

var errVersionDecrease = fmt.Errorf("version decrease in combination with auto increment is not supported")

func Next(current string, patternNext string) (string, error) {
	_, current = splitPrefix(vPrefix, current)
	currentCore, _ := coreAndExtension(current)

	nextPrefix, patternNext := splitPrefix(vPrefix, patternNext)
	nextCore, nextExtension := coreAndExtension(patternNext)

	positionIncrement, err := findIncrementChar(nextCore)
	if err != nil {
		return "", err
	}

	currentMajorMinorPatchList, err := majorMinorPatchList(currentCore)
	if err != nil {
		return "", err
	}
	nextMajorMinorPatchList, err := majorMinorPatchList(nextCore)
	if err != nil {
		return "", err
	}
	nextMajorMinorPatchList, err = replaceWildcardsAndIncrements(nextMajorMinorPatchList, currentMajorMinorPatchList)
	if err != nil {
		return "", err
	}
	currentMajor, currentMinor, currentPatch, err := majorMinorPatchListToInt(currentMajorMinorPatchList)
	if err != nil {
		return "", err
	}
	nextMajor, nextMinor, nextPatch, err := majorMinorPatchListToInt(nextMajorMinorPatchList)
	if err != nil {
		return "", err
	}

	switch positionIncrement {
	case positionMajor:
		nextMajor = currentMajor + 1
	case positionMinor:
		if currentMajor == nextMajor {
			nextMinor = currentMinor + 1
		} else if currentMajor > nextMajor {
			return "", errVersionDecrease
		}
	case positionPatch:
		if currentMajor == nextMajor && currentMinor == nextMinor {
			nextPatch = currentPatch + 1
		} else if currentMajor > nextMajor {
			return "", errVersionDecrease
		} else if currentMajor == nextMajor && currentMinor > nextMinor {
			return "", errVersionDecrease
		}
	}

	return fmt.Sprintf("%v%v.%v.%v%v", nextPrefix, nextMajor, nextMinor, nextPatch, nextExtension), nil
}

func replaceWildcardsAndIncrements(nextCore []string, currentCore []string) ([]string, error) {
	for i := 0; i < 3; i++ {
		if nextCore[i] == wildcardChar {
			nextCore[i] = currentCore[i]
		} else if nextCore[i] == incrementChar {
			nextCore[i] = "0"
		}
	}
	return nextCore, nil
}

func majorMinorPatchList(versionCore string) ([]string, error) {
	split := strings.Split(versionCore, ".")
	if len(split) != 3 {
		return split, fmt.Errorf("version must be <major>.<minor>.<patch>, got %v", versionCore)
	}
	return split, nil
}

func majorMinorPatchListToInt(versionCore []string) (int, int, int, error) {
	major, err := strconv.Atoi(versionCore[positionMajor])
	if err != nil {
		return 0, 0, 0, err
	}
	minor, err := strconv.Atoi(versionCore[positionMinor])
	if err != nil {
		return 0, 0, 0, err
	}
	patch, err := strconv.Atoi(versionCore[positionPatch])
	if err != nil {
		return 0, 0, 0, err
	}
	return major, minor, patch, nil
}

func splitPrefix(prefix string, str string) (string, string) {
	if !strings.HasPrefix(str, prefix) {
		return "", str
	}
	return prefix, str[len(prefix):]
}

// Returns 0 if major is incrementChar
// Returns 1 if minor is incrementChar
// Returns 2 if patch is incrementChar
// Returns an error if there is more than one incrementChar in major, minor, and patch combined
func findIncrementChar(versionCore string) (int, error) {
	positionIncrement := -1
	var countIncrement int
	for i, s := range strings.Split(versionCore, ".") {
		if s == incrementChar {
			positionIncrement = i
			countIncrement++
		}
		if countIncrement > 1 {
			return 0, fmt.Errorf("only one %s allowed", incrementChar)
		}
	}
	return positionIncrement, nil
}

func coreAndExtension(version string) (string, string) {
	versionCore := version
	var versionExtension string
	if indexExtension := strings.IndexAny(version, "+-"); indexExtension != -1 {
		versionCore = version[:indexExtension]
		versionExtension = version[indexExtension:]
	}
	return versionCore, versionExtension
}
