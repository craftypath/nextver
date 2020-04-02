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
    // check if patternNext was a valid semver version if "x"s were numbers
    versionNext, err := semver.StrictNewVersion(strings.ReplaceAll(patternNext, "x", "0"))
    if err != nil {
        return "", err
    }

    // only need to consider major.minor.patch (="version core")
    patternNextCore := patternNext
    var patternNextExtension string
    if indexExtension := strings.IndexAny(patternNext, "+-"); indexExtension != -1 {
        patternNextCore = patternNext[:indexExtension]
        patternNextExtension = patternNext[indexExtension:]
    }

    // ensure there is at most one x and determine its position
    positionX := -1
    var countX int
    for i, s := range strings.Split(patternNextCore, ".") {
        if s == "x" {
            positionX = i
            countX++
        }
        if countX > 1 {
            return "", fmt.Errorf("only one x allowed")
        }
    }

    // set and maybe increment new version
    nextMajor := versionNext.Major()
    nextMinor := versionNext.Minor()
    nextPatch := versionNext.Patch()

    switch positionX {
    case positionMajor:
        if currentVersion.Minor() == versionNext.Minor() && currentVersion.Patch() == versionNext.Patch() {
            nextMajor = currentVersion.Major() + 1
        }
    case positionMinor:
        if currentVersion.Major() == versionNext.Major() && currentVersion.Patch() == versionNext.Patch() {
            nextMinor = currentVersion.Minor() + 1
        }
    case positionPatch:
        if currentVersion.Major() == versionNext.Major() && currentVersion.Minor() == versionNext.Minor() {
            nextPatch = currentVersion.Patch() + 1
        }
    }
    var result string
    nextVersion, err := semver.StrictNewVersion(fmt.Sprintf("%v.%v.%v%v", nextMajor, nextMinor, nextPatch, patternNextExtension))
    if err == nil {
        result = nextVersion.String()
    }
    return result, err
}
