# nextver

**work in progress**

## Features

- auto-increments major, minor, or patch version when set to `x`:

```
nextver 0.0.0 0.x.0 # prints 0.1.0
nextver 0.1.0 0.x.0 # prints 0.2.0
# ... and so on ...
```

- bump versions explicitly, with the auto-increment being reset to `0`:

```
# time for 1.0.0
nextver 0.487.0 1.x.0 # prints 1.0.0
nextver 1.0.0 1.x.0 # prints 1.1.0

```

- add/keep/remove commonly used `v` prefix:

```
nextver 0.1.0 v0.x.0 # prints v0.2.0
nextver v0.2.0 v0.x.0 # prints v0.3.0
nextver v0.3.0 0.x.0 # prints 0.4.0
```

- add/keep/remove pre-release and/or build info:

```
nextver 2.1.0 2.x.0+abc # prints 2.2.0+abc
nextver 2.2.0+abc 2.x.0+abc # prints 2.3.0+abc
nextver 2.3.0+abc 2.x.0 # prints 2.4.0
```

## Usage
```
nextver <current-version> <pattern>

<current-version>: SemVer denoting the current version of your artifact
<pattern>: SemVer denoting the next version of your artifact. One of <major>.<minor>.<patch> may be set to "x" to increment from current-version.
```

## CI/CD Example

If your artifact should have major version `0`, but you also want to bump the minor version with each build:
```
currentVersion=$(git describe --abbrev=0)
nextVersion=$(nextver "$currentVersion" 0.x.0)
git tag "$nextVersion"
git push --tags
# ... use $version, e.g. to tag a docker image.
```
