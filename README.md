# nextver

**work in progress**

## Features

- auto-increments major, minor, or patch version when set to `x`:

```
nextver -c 0.0.0 -p 0.x.0 # prints 0.1.0
nextver -c 0.1.0 -p 0.x.0 # prints 0.2.0
# ... and so on ...
```

- bump versions explicitly, with the auto-increment being reset to `0`:

```
# time for 1.0.0
nextver -c 0.487.0 -p 1.x.0 # prints 1.0.0
nextver -c 1.0.0 -p 1.x.0 # prints 1.1.0

```

- add/keep/remove commonly used `v` prefix:

```
nextver -c 0.1.0 -p v0.x.0 # prints v0.2.0
nextver -c v0.2.0 -p v0.x.0 # prints v0.3.0
nextver -c v0.3.0 -p 0.x.0 # prints 0.4.0
```

- add/keep/remove pre-release and/or build info:

```
nextver -c 2.1.0 -p 2.x.0+abc # prints 2.2.0+abc
nextver -c 2.2.0+abc -p 2.x.0+abc # prints 2.3.0+abc
nextver -c 2.3.0+abc -p 2.x.0 # prints 2.4.0
```

## Usage

```console
$ nextver -h
nextver manages automatic semver versioning

Usage:
  nextver [flags]

Flags:
  -c, --current-version string   the current version
  -h, --help                     help for nextver
  -p, --pattern string           the versioning pattern
  -v, --version                  version for nextver
```

## CI/CD Example

If your artifact should have major version `0`, but you also want to bump the minor version with each build:
```
currentVersion=$(git describe --abbrev=0)
nextVersion=$(nextver -c "$currentVersion" -p 0.x.0)
git tag "$nextVersion"
git push --tags
# ... use $version, e.g. to tag a docker image.
```
