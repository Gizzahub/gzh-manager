gzh-manager
===========

<div style="text-align: center;">
A general purpose project for easy development.
<br>
<br>
This template serves as a starting point for Golang command-line applications. It is based on high-quality Golang projects and various useful blog posts that helped me understand Golang better.
<br>
<br>
<img src="https://github.com/gizzahub/gzh-manager-go/actions/workflows/test.yml/badge.svg" alt="Test Status"/>
<img src="https://github.com/gizzahub/gzh-manager-go/actions/workflows/lint.yml/badge.svg" alt="Lint Status"/>
<img src="https://pkg.go.dev/badge/github.com/gizzahub/gzh-manager-go.svg" alt="GoDoc"/>
<img src="https://codecov.io/gh/Gizzahub/gzh-manager-go/branch/main/graph/badge.svg" alt="Code Coverage"/>
<img src="https://img.shields.io/github/v/release/Gizzahub/gzh-manager-go" alt="Latest Release"/>
<img src="https://img.shields.io/docker/pulls/Gizzahub/gzh-manager-go" alt="Docker Pulls"/>
<img src="https://img.shields.io/github/downloads/Gizzahub/gzh-manager-go/total.svg" alt="Total Downloads"/>
</div>

Comprehensive CLI Tool

# Table of Contents
<!--ts-->
  * [Usage](#usage)
  * [Features](#features)
  * [Project Layout](#project-layout)
  * [How to use this template](#how-to-use-this-template)
  * [Demo Application](#demo-application)
  * [Makefile Targets](#makefile-targets)
  * [Contribute](#contribute)

<!-- Added by: morelly_t1, at: Tue 10 Aug 2021 08:54:24 AM CEST -->

<!--te-->

# Usage

## Setclone
Clone repositories by GitHub account (user, org) or GitLab group.

- setclone
  - git
  - gitea
  - github
  - gitlab
  - gogs

### CLI

### Setclone config

```yaml
# setclone.yaml 
github:
  ScriptonBasestar:
   auth: token
   proto: https
   targetPath: $HOME/mywork/ScriptonBasestar
   default:
    strategy: include
    branch: develop
   include:
    proxynd:
      branch: develop
    devops-minim-engine:
      branch: dev
   exclude:
    - sb-wp-*
   override:
    include:
  nginxinc:
   targetPath: $HOME/mywork/nginxinc
```

```bash
gzh setclone -o nginxinc
gzh setclone -o nginxinc -t $HOME/mywork/nginxinc
gzh setclone -o nginxinc -t $HOME/mywork/nginxinc --auth token
```

## Trigger

와이파이 변경.. 등

# Features
- [goreleaser](https://goreleaser.com/) with `deb.` and `.rpm` packer and container (`docker.hub` and `ghcr.io`) releasing including `manpages` and `shell completions` and grouped Changelog generation.
- [golangci-lint](https://golangci-lint.run/) for linting and formatting
- [Github Actions](.github/worflows) Stages (Lint, Test (`windows`, `linux`, `mac-os`), Build, Release) 
- [Gitlab CI](.gitlab-ci.yml) Configuration (Lint, Test, Build, Release)
- [cobra](https://cobra.dev/) example setup including tests
- [Makefile](Makefile) - with various useful targets and documentation (see Makefile Targets)
- [Github Pages](_config.yml) using [jekyll-theme-minimal](https://github.com/pages-themes/minimal) (checkout [https://Gizzahub.github.io/gzh-manager-go/](https://Gizzahub.github.io/gzh-manager-go/))
- Useful `README.md` badges
- [pre-commit-hooks](https://pre-commit.com/) for formatting and validating code before committing

## Project Layout
* [assets/](https://pkg.go.dev/github.com/gizzahub/gzh-manager-go/assets) => docs, images, etc
* [cmd/](https://pkg.go.dev/github.com/gizzahub/gzh-manager-go/cmd)  => command-line configurations (flags, subcommands)
* [pkg/](https://pkg.go.dev/github.com/gizzahub/gzh-manager-go/pkg)  => packages that are okay to import for other projects
* [internal/](https://pkg.go.dev/github.com/gizzahub/gzh-manager-go/pkg)  => packages that are only for project internal purposes
- [`tools/`](tools/) => for automatically shipping all required dependencies when running `go get` (or `make bootstrap`) such as `golang-ci-lint` (see: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)
- [`scripts/`](scripts/) => build scripts 

## How to use this template
```sh
```

In order to make the CI work you will need to have the following Secrets in your repository defined:

Repository  -> Settings -> Secrets & variables -> `CODECOV_TOKEN`, `DOCKERHUB_TOKEN` & `DOCKERHUB_USERNAME`

## Demo Application

```sh
$> gzh-manager -h
golang-cli cli application by managing gzh-manager

Usage:
  gzh [flags]
  gzh [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  setclone    Clone repositories in bulk
  help        Help about any command
  version     gzh-manager version

Flags:
  -h, --help   help for gzh-manager

Use "gzh-manager [command] --help" for more information about a command.
```

First, create a configuration file in the desired path. Refer to
[setclone.yaml](pkg/setclone/setclone.yaml)

```sh
$> gzh setclone -t $HOME/mywork

This won't work:
$> gzh setclone -t ./mywork
$> gzh setclone -t $HOME/mywork
$> gzh setclone -t ~/mywork
```

# Makefile Targets
```sh
$> make
bootstrap                      install build deps
build                          build golang binary
clean                          clean up environment
cover                          display test coverage
docker-build                   dockerize golang application
fmt                            format go files
help                           list makefile targets
install                        install golang binary
lint                           lint go files
pre-commit                     run pre-commit hooks
run                            run the app
test                           display test coverage
```

# Contribute
If you find issues in this setup or have some nice features/improvements, I would welcome an issue or a PR :)

# Notes
I used the template from `https://github.com/FalcoSuessgott/golang-cli-template`, but `https://github.com/create-go-app/cli` seems better due to its simplicity.
It has become somewhat complex. 
