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
    currentMajor, currentMinor, currentPatch, _, err := components(current)
    if err != nil {
        return "", err
    }

    nextMajor, nextMinor, nextPatch, nextExtension, err := components(patternNext)
    if err != nil {
        return "", err
    }

    positionX, err := findX(patternNext)
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

    return fmt.Sprintf("%v.%v.%v%v", nextMajor, nextMinor, nextPatch, nextExtension), err
}

func components(version string) (uint, uint, uint, string, error) {
    core, extension := getCoreAndExtension(version)
    core = strings.ReplaceAll(core, "x", "0")
    split := strings.Split(core, ".")
    if len(split) != 3 {
        return 0, 0, 0, "", fmt.Errorf("version must be <major>.<minor>.<patch>, got %v", version)
    }
    major, err := strconv.Atoi(split[0])
    if err != nil {
        return 0, 0, 0, "", err
    }
    minor, err := strconv.Atoi(split[1])
    if err != nil {
        return 0, 0, 0, "", err
    }
    patch, err := strconv.Atoi(split[2])
    if err != nil {
        return 0, 0, 0, "", err
    }
    if major < 0 || minor < 0 || patch < 0 {
        return 0, 0, 0, "", fmt.Errorf("version must only contain positive numbers got %v", version)
    }
    return uint(major), uint(minor), uint(patch), extension, nil
}

// returns 0 if major is x
// returns 1 if minor is x
// returns 2 if patch is x
// returns an error if there is more than one x in major, minor, and patch combined
func findX(version string) (int, error) {
    versionCore, _ := getCoreAndExtension(version)
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

func getCoreAndExtension(version string) (string, string) {
    versionCore := version
    var versionExtension string
    if indexExtension := strings.IndexAny(version, "+-"); indexExtension != -1 {
        versionCore = version[:indexExtension]
        versionExtension = version[indexExtension:]
    }
    return versionCore, versionExtension
}
