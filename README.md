<p align="center">
  <img alt="golangci-lint logo" src="static/images/logo.png" height="150" />
  <h3 align="center">thredl.it</h3>
  <p align="center">The thread between ChatOps and Support</p>
</p>
<a name="readme-top"></a>

---

## Overview

**`thredl`** is a simple, configurable, ChatOps platform.

## To start using `thredle`

Head over to [thredle.isnotregisteryet.io](http://sorry)

## To start contributing to `thredle`
### Prerequisites

Ensure the folllowing is installed and working...

- [Go](https://go.dev/doc/install)
- Make
- Brew

To install other prequesites you can simply run

```
make local
```

The above will install the following...

- OrbStack
- KubeCtl
- Tilt
- K9s
- goclangci-lint

### Getting Started

1. Ensure you have all the prerequisites setup
2. Clone the repo
    ```sh
    git clone git@github.com:brettmostert/thredle.it.git
    ```
3. **Optional**. Setup your IDE by following the instructions at
    - https://github.com/mvdan/gofumpt
    - https://golangci-lint.run/usage/integrations
4. To run (locally) as a monolith
```sh
make dev
```
---

## Components

| Name          | Description          |    Language    |    Type     |
| ------------- | -------------------- | :------------: | :---------: |
| **thredl.it** | The platform as a monolith   |       Go       |     service     |
| **thredle.it-api**     | The platforms API |       Go       |   service   |
| **thredle-it-ui** | The platforms UI | Flutter / Dart / HTMX | ui / mobile |

---

## Technology

Below is a list of the technology primarily used in this project.

### Backend

- Go
- GRPC, REST
- K8s & Containers

### FrontEnd

- Flutter & Dart
- HTMX

## Project Structure

Below is a brief description of the project folder structure.

### `/cmd`

Contains the main executeable applications, namely `cli`

### `/internal`

Contains all the private code for this application to work.

### `/package`

Contains all libs/modules which may be shared.

### `/scripts`

Scripts to perform various tasks such as `build` or `sca` (static code analysis), etc. This allows us to keep our makefile in the root folder small.

## Project Setup

Below is a brief description of how to setup this project.

### BootStrap

The tools listed below are required for the build and/or packaging process.

| Category | Description                        | URL                       |
| -------- | ---------------------------------- | ------------------------- |
| Linter   | `golangci` is our linter of choice | https://golangci-lint.run |
| Formatter| `gofumpt` replaces gofmt           | https://github.com/mvdan/gofumpt |

# Giving Thanks

A Big Shout Out! To the people who worked on the following, the work you have done has aided in my learning of the go language and eco-system.

- project layout inspired by <https://github.com/golang-standards/project-layout> and <https://github.com/stellar/go>
- linting from <https://golangci-lint.run> and <https://golangci-lint.run/usage/integrations/#go-for-visual-studio-code>
- project logo inspired by my horrific graphic design and created with <https://krita.org/en/>
- scany <https://github.com/georgysavva/scany>
- pgx <https://github.com/jackc/pgx>

# Other

## Emoji Legend

| meaning | emoji              | text                 |
| ------- | ------------------ | -------------------- |
| done    | :heavy_check_mark: | `:heavy_check_mark:` |
| wip     | :construction:     | `:construction:`     |
| note    | :memo:             | `:memo:`             |
