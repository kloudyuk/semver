# semver

A simple CLI based on [masterminds/semver](https://github.com/Masterminds/semver) for dealing with semantic versions

## Install

```sh
go install github.com/kloudyuk/semver@latest
```

## Usage

```text
Usage:
  semver [command]

Available Commands:
  bump        Takes a given version and bumps it by the desired level (major|minor|patch)
  help        Help about any command
  parse       Parses a version in accordance with the SemVer 2 specification and returns the full version or an error
  sort        Sort semantic versions in accordance with the SemVer 2 specification

Flags:
  -h, --help   help for semver

Use "semver [command] --help" for more information about a command.
```
