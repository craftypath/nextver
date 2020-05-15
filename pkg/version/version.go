package version

import (
	"fmt"
	"strconv"
	"strings"
)

const positionMajor = 0
const positionMinor = 1
const positionPatch = 2

func Next(current string, patternNext string) (string, error) {
	_, current = splitPrefix("v", current)
	currentCore, _ := coreAndExtension(current)
	currentMajor, currentMinor, currentPatch, err := majorMinorPatch(currentCore)
	if err != nil {
		return "", err
	}

	nextPrefix, patternNext := splitPrefix("v", patternNext)
	nextCore, nextExtension := coreAndExtension(patternNext)

	positionX, err := findX(nextCore)
	if err != nil {
		return "", err
	}

	nextCore = strings.ReplaceAll(nextCore, "x", "0")
	nextMajor, nextMinor, nextPatch, err := majorMinorPatch(nextCore)
	if err != nil {
		return "", err
	}

	switch positionX {
	case positionMajor:
		if currentMinor == nextMinor && currentPatch == nextPatch {
			nextMajor = currentMajor + 1
		}
	case positionMinor:
		if currentMajor == nextMajor && currentPatch == nextPatch {
			nextMinor = currentMinor + 1
		}
	case positionPatch:
		if currentMajor == nextMajor && currentMinor == nextMinor {
			nextPatch = currentPatch + 1
		}
	}

	return fmt.Sprintf("%v%v.%v.%v%v", nextPrefix, nextMajor, nextMinor, nextPatch, nextExtension), nil
}

func majorMinorPatch(versionCore string) (int, int, int, error) {
	split := strings.Split(versionCore, ".")
	if len(split) != 3 {
		return 0, 0, 0, fmt.Errorf("version must be <major>.<minor>.<patch>, got %v", versionCore)
	}
	major, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, 0, err
	}
	minor, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, 0, err
	}
	patch, err := strconv.Atoi(split[2])
	if err != nil {
		return 0, 0, 0, err
	}
	return major, minor, patch, nil
}

func splitPrefix(prefix string, str string) (string, string) {
	var foundPrefix string
	if strings.HasPrefix(str, prefix) {
		str = str[len(prefix):]
		foundPrefix = prefix
	}
	return foundPrefix, str
}

// returns 0 if major is x
// returns 1 if minor is x
// returns 2 if patch is x
// returns an error if there is more than one x in major, minor, and patch combined
func findX(versionCore string) (int, error) {
	positionX := -1
	var countX int
	for i, s := range strings.Split(versionCore, ".") {
		if s == "x" {
			positionX = i
			countX++
		}
		if countX > 1 {
			return 0, fmt.Errorf("only one x allowed")
		}
	}
	return positionX, nil
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
