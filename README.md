# dict-go [![go](https://github.com/lescactus/dict-go/actions/workflows/go.yml/badge.svg)](https://github.com/lescactus/dict-go/actions/workflows/go.yml) [![release](https://github.com/lescactus/dict-go/actions/workflows/release.yml/badge.svg)](https://github.com/lescactus/dict-go/actions/workflows/release.yml)

This repository contains a simple cli written in go to lookup for the definition of a given word.
It uses the [`dictionaryapi.dev`](https://dictionaryapi.dev/) API to fetch the words definitions.
## Usage

```sh
./dict-go -h      
dict-go is a simple cli used to lookup for the definition of a given word. 
	You need to provide the word you are looking for and the language (optional - default is "en").

	The source code is available at https://github.com/lescactus/dict-go.

Usage:
  dict-go [flags]

Flags:
  -h, --help          help for dict-go
  -l, --lang string   Lang of the word (optional) (default "en")
  -w, --word string   Word to lookup

```

## Installation

Prebuilt binaries can be downloaded from the GitHub Releases [section](https://github.com/lescactus/dict-go/releases), or using a Docker image from the Github Container Registry.

### Running with Docker

```bash
docker run --rm -it --name dict-go ghcr.io/lescactus/dict-go -w <word>
```

## Building

<details>

### Requirements

* Golang 1.13 or higher


### From source with go

You need a working [go](https://golang.org/doc/install) toolchain (It has been developped and tested with go >= 1.13). Refer to the official documentation for more information (or from your Linux/Mac/Windows distribution documentation to install it from your favorite package manager).

```sh
# Clone this repository
git clone https://github.com/lescactus/dict-go.git && cd dict-go/

# Build from sources. Use the '-o' flag to change the compiled binary name
go build

# Default compiled binary is dict-go
# You can optionnaly move it somewhere in your $PATH to access it shell wide
./dict-go -help
```

### From source with docker

If you don't have [go](https://golang.org/) installed but have docker, run the following command to build inside a docker container:

```sh
# Build from sources inside a docker container. Use the '-o' flag to change the compiled binary name
# Warning: the compiled binary belongs to root:root
docker run --rm -it -v "$PWD":/app -w /app golang:1.21 go build

# Default compiled binary is dict-go
# You can optionnaly move it somewhere in your $PATH to access it shell wide
./dict-go -help
```

### From source with docker but built inside a docker image

If you don't want to pollute your computer with another program, this cli comes with its own docker image:

```sh
docker build -t dict-go .

docker run --rm dict-go --word hello
```

</details>