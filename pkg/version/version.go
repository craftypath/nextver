package version

import (
    "fmt"
    "github.com/Masterminds/semver/v3"
    "strings"
)

const positionMajor = 0
const positionMinor = 1
const positionPatch = 2

func Next(current string, patternNext string) (string, error) {
    currentVersion, err := semver.StrictNewVersion(current)
    if err != nil {
        return "", err
    }

    versionNext, err := semver.StrictNewVersion(strings.ReplaceAll(patternNext, "x", "0"))
    if err != nil {
        return "", err
    }

    patternNextCore, patternNextExtension := splitVersionExtension(patternNext)

    // set and maybe increment new version
    nextMajor := versionNext.Major()
    nextMinor := versionNext.Minor()
    nextPatch := versionNext.Patch()

    currentMajor := currentVersion.Major()
    currentMinor := currentVersion.Minor()
    currentPatch := currentVersion.Patch()

    positionX, err := findX(patternNextCore)
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

    var result string
    nextVersion, err := semver.StrictNewVersion(fmt.Sprintf("%v.%v.%v%v", nextMajor, nextMinor, nextPatch, patternNextExtension))
    if err == nil {
        result = nextVersion.String()
    }
    return result, err
}

func splitVersionExtension(version string) (string, string) {
    versionCore := version
    var versionExtension string
    if indexExtension := strings.IndexAny(version, "+-"); indexExtension != -1 {
        versionCore = version[:indexExtension]
        versionExtension = version[indexExtension:]
    }
    return versionCore, versionExtension
}

func findX(version string) (int, error) {
    positionX := -1
    var countX int
    for i, s := range strings.Split(version, ".") {
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
