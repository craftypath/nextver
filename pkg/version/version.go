package version

import (
    "fmt"
    "github.com/Masterminds/semver/v3"
    "strings"
)

const positionMajor = 0
const positionMinor = 1
const positionPatch = 2

func Next(current string, incrementPattern string) (string, error) {
    currentVersion, err := semver.StrictNewVersion(current)
    if err != nil {
        return "", err
    }
    // Also, check if incrementPattern was a valid semver version if "x"s were numbers
    incrementPatternWithoutXs := strings.ReplaceAll(incrementPattern, "x", "0")
    incrementVersion, err := semver.StrictNewVersion(incrementPatternWithoutXs)
    if err != nil {
        return "", err
    }

    // only need to consider major.minor.patch (="version core")
    versionCore := incrementPattern
    var versionExtension string
    if indexExtension := strings.IndexAny(incrementPattern, "+-"); indexExtension != -1 {
        versionCore = incrementPattern[:indexExtension]
        versionExtension = incrementPattern[indexExtension:]
    }

    // ensure there is at most one x and determine its index
    positionX := -1
    var countX int
    for i, s := range strings.Split(versionCore, ".") {
        if s == "x" {
            positionX = i
            countX++
        }
        if countX > 1 {
            return "", fmt.Errorf("only one x allowed")
        }
    }

    // set and maybe increment new version
    nextMajor := incrementVersion.Major()
    nextMinor := incrementVersion.Minor()
    nextPatch := incrementVersion.Patch()

    switch positionX {
    case positionMajor:
        if currentVersion.Minor() == incrementVersion.Minor() && currentVersion.Patch() == incrementVersion.Patch() {
            nextMajor = currentVersion.Major() + 1
        }
    case positionMinor:
        if currentVersion.Major() == incrementVersion.Major() && currentVersion.Patch() == incrementVersion.Patch() {
            nextMinor = currentVersion.Minor() + 1
        }
    case positionPatch:
        if currentVersion.Major() == incrementVersion.Major() && currentVersion.Minor() == incrementVersion.Minor() {
            nextPatch = currentVersion.Patch() + 1
        }
    }
    var result string
    newVersion, err := semver.StrictNewVersion(fmt.Sprintf("%v.%v.%v%v", nextMajor, nextMinor, nextPatch, versionExtension))
    if err == nil {
        result = newVersion.String()
    }
    return result, err
}
