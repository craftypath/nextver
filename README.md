# VerIt

**work in progress**

## Usage
```
verit <current-version> <pattern>

<current-version>: SemVer denoting the current version of your artifact
<pattern>: SemVer denoting the next version of your artifact. One of <major>.<minor>.<patch> may be set to "x" to increment from current-version.
```

## Examples
```
verit 1.0.0 1.x.0 # prints 1.1.0
```
```
verit 0.9.0 1.x.0 # prints 1.0.0
```

## CI/CD Example

If your artifact should have major version `0`, but you also want to bump the minor version with each build:
```
currentVersion=$(git describe --abbrev=0)
nextVersion=$(verit "$currentVersion" 0.x.0)
git tag "$nextVersion"
git push --tags
# ... use $version, e.g. to tag a docker image.
```
