# go-ipfs-ex

English | [中文](README_CN.md)
## What is IPFS-eX

IPFS-eX is a implementation that uses [DSP] (https://github.com/DSP-Labs/dsp) protocol to upgrade [IPFS] (https://github.com/ipfs/ipfs/) protocol network Storage protocol, the goal is to bring more key protocol capabilities such as distribute data actively, file encryption, and permission control to the IPFS network, and ultimately establish a more complete distributed storage alternative network.

# Table of Contents

- [Backgroud](#backgroud)
- [Features](#features)
- [Intall Guide](#intall-guide)
  - [System Requirment](#system-requirment)
  - [Install From Source](#install-from-source)
    - [Intall Go](#intall-go)
    - [Download and Build](#download-and-build)
- [Getting Started](#getting-started)
    - [Usage](#usage)
    - [Examples](#examples)
- [Development](#development)
  - [API](#api)
  - [Test](#test)
- [Contributing](#contributing)
- [License](#license)

## Backgroud

[DSP(Distributed Storage Protocol)](https://github.com/DSP-Labs/dsp) protocol is a next generation Internet protocol paradigm based on multi-dimensional data file encryption, distribution, storage, sharing, etc. The goal of [DSP](https://github.com/DSP-Labs/dsp) is to become the core infrastructure of Web3.0. What we need is a new set of free and open network protocols-based on this set of protocols, people hope that they can achieve true freedom of data and freedom of interconnection. At the same time, it is necessary to form a user data file sharing library (database) under a free network. This database is neutral and jointly owned by the users themselves.

As an open source protocol that began to explore distributed storage protocols earlier, [IPFS](https://github.com/ipfs/ipfs/) has done a lot of work in the development of distributed storage protocols, and has also completed phased progress, and has won the favor of a large number of developers, with a large number of storage nodes distributed in the network, it has an ideal distributed storage environment foundation.

However, we have higher expectations for the future Web3.0 storage network world. Various types of data files should be able to be actively distributed and freely shared according to the actual business needs. Any level of client should be able to become a distribution node, any member should have the right to fully participate in the distribution, sharing, and storage of files on the network. But at the same time, the privacy and data security of all participants should be fully protected. All data in the network will be the most valuable asset. 

[IPFS-eX](https://github.com/IPFS-eX/IPFS-eX) is a open source project that use [DSP](https://github.com/DSP-Labs/dsp) to expand [IPFS](https://github.com/ipfs/ipfs/) protocol, using [DSP](https://github.com/DSP-Labs/dsp) protocol to ensure data is distributed efficiently, highly concurrently and safely on the network. At the same time, it supports end-to-end file data sharing, and through the support of encryption technology, file sharing and stable storage can be safely performed in a decentralized storage network. Through file fragment encryption and multi-point storage, the availability and stability of network data transmission can be greatly improved. With the support of applied cryptography, control of the file's own access rights is also achieved.

`go-ipfs-ex`is the golang implementation of `IPFS-eX`.

>
>We write a set of documents, posts, and tutorials to explain what DSP is, what IPFS-eX is, how go-ipfs-ex works, and how it can help your existing and new projects.
>
>- [**DSP spec**](https://github.com/DSP-Labs/specs)
>- [**IPFS-eX design**](https://github.com/IPFS-eX/IPFS-eX)
>- [**IPFS-eX technical docs**](https://github.com/IPFS-eX/docs)
>- [**go-ipfs-ex**](https://github.com/IPFS-eX/go-ipfs-ex)
>- [**carrier**](https://github.com/IPFS-eX/carrier)

## Features

- active distribution
- privacy protection
- access control

coming soon...

## Intall Guide 
### System Requirment
[Golang](https://golang.org/doc/install) 1.13 or later.

`go-ipfs-ex` use [Go modules](https://github.com/golang/go/wiki/Modules) manage dependency

`go-ipfs-ex` can run on most Linux, macOS, and Windows systems. We recommend running it on a machine with at least 2 GB of RAM and 2 CPU cores (go-ipfs is highly parallel).

If your system is resource-constrained, we recommend:
1. manually rebuild `go-ipfs-ex`
```
make build GOTAGS=OpenSSL
```
2. Initializing your daemon with go-ipfs-ex init --profile=lowpower

### Install From Source
#### Intall Go
`go-ipfs-ex`need [Golang](https://golang.org/doc/install) 1.13 or later，If you don't have it: [Download Go 1.13+](https://golang.org/dl/)。

You'll need to add Go's bin directories to your `$PATH` environment variable e.g., by adding these lines to your `/etc/profile` (for a system-wide installation) or `$HOME/.profile`:

```
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOPATH/bin
```
(If you run into trouble, see the [Go install instructions](https://golang.org/doc/install)).
#### Download and Build
```
$ git clone https://github.com/IPFS-eX/go-ipfs-ex.git

$ cd go-ipfs-ex
$ make install
```
## Getting Started
### Usage
*continue use ipfs command*
```
  ipfs - Global p2p merkle-dag filesystem.

  ipfs [<flags>] <command> [<arg>] ...

SUBCOMMANDS
  BASIC COMMANDS
    init          Initialize ipfs local configuration
    add <path>    Add a file to ipfs
    cat <ref>     Show ipfs object data
    get <ref>     Download ipfs objects
    ls <ref>      List links from an object
    refs <ref>    List hashes of links from an object

  DATA STRUCTURE COMMANDS
    block         Interact with raw blocks in the datastore
    object        Interact with raw dag nodes
    files         Interact with objects as if they were a unix filesystem

  ADVANCED COMMANDS
    daemon        Start a long-running daemon process
    mount         Mount an ipfs read-only mount point
    resolve       Resolve any type of name
    name          Publish or resolve IPNS names
    dns           Resolve DNS links
    pin           Pin objects to local storage
    repo          Manipulate an IPFS repository

  NETWORK COMMANDS
    id            Show info about ipfs peers
    bootstrap     Add or remove bootstrap peers
    swarm         Manage connections to the p2p network
    dht           Query the DHT for values or peers
    ping          Measure the latency of a connection
    diag          Print diagnostics

  TOOL COMMANDS
    config        Manage configuration
    version       Show ipfs version information
    update        Download and apply go-ipfs updates
    commands      List all available commands

  Use 'ipfs <command> --help' to learn more about each command.

  ipfs uses a repository in the local file system. By default, the repo is located at
  ~/.ipfs. To change the repo location, set the $IPFS_PATH environment variable:

    export IPFS_PATH=/path/to/ipfsrepo
```
### Examples
coming soon...
## Development
### API
[![GoDoc](https://godoc.org/github.com/IPFS-eX/go-ipfs-ex?status.svg)](https://godoc.org/github.com/IPFS-eX/go-ipfs-ex)
### Test
```
`go test ./...`
```
## Contributing
Thank you for considering contributing source code to IPFS-eX. We welcome any individuals and organizations on the Internet to participate in this open source project.

We have the following code contribution rules that need to be followed by everyone:

- Code needs to be followed Golang [formatting](https://golang.org/doc/effective_go.html%23formatting) guidance document(e.g.use `gofmt`);
- Documents need to be followed Golang [commentary](https://golang.org/doc/effective_go.html%23commentary) suggestion; 
- The commit information needs to reflect the modified source package name;
- You can also follow the [development Manual](https://github.com/IPFS-eX/docs) for more detailed development information.
## License
`go-ipfs-ex` is available under MIT/Apache-2.0。