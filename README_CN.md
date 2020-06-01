# go-ipfs-ex

中文 | [English](README.md)
## IPFS-eX是什么?

IPFS-eX是一个利用[DSP](https://github.com/DSP-Labs/dsp)协议思想对[IPFS](https://github.com/ipfs/ipfs/)协议网络进行升级的分布式存储协议，目标是给IPFS网络带来数据文件主动分发，文件加密，权限控制等更完善的核心协议能力，最终建立一个更完善的分布式存储底层网络。


# 目录

- [背景](#背景)
- [特性](#特性)
- [安装指南](#安装指南)
  - [系统需求](#系统需求)
  - [从源码安装](#从源码安装)
    - [安装go](#安装go)
    - [下载并编译](#下载并编译)
- [入门指南](#入门指南)
    - [使用](#使用)
    - [示例](#示例)
- [开发](#开发)
  - [API](#API)
  - [单元测试](#单元测试)
- [贡献](#贡献)
- [许可证](#许可证)

## 背景

[DSP(Distributed Storage Protocol)](https://github.com/DSP-Labs/dsp)协议是基于数据文件加密，分发，存储，共享等多维度的新一代互联网协议范式，[DSP](https://github.com/DSP-Labs/dsp)协议的目标是成为Web3.0的核心基础设施。我们需要的是一组全新自由开放的网络协议——基于这套协议，人们希望可以实现真正意义上的数据自由，互联自由。同时在自由的网络下需要可以构成一个用户数据文件共享库(数据库)，这个数据库是中立的，由用户自己共同拥有。

[IPFS](https://github.com/ipfs/ipfs/)作为较早开始进行分布式存储协议探索的开源协议，在分布式存储协议的发展中做了大量的工作，也完成了阶段性的进展，并且获得了大量开发者的青睐，有大量的存储节点分布在网络中，拥有一个较为理想的分布式存储环境基础。

但是我们对于未来Web3.0的存储网络世界有着更高的期许，各类型的数据文件，应该是可以按照实际业务的需要进行文件的主动分发和自由分享，任何级别的客户端都应该能成为网络中的一个分发节点，任何一个节点成员都应该得到充分参与网络的文件分发，分享，存储的权利。但同时所有的参与者的隐私，数据安全也都应该得到充分的保护。网络中的所有数据都会是最具价值的资产。

[IPFS-eX](https://github.com/IPFS-eX/IPFS-eX)是一个基于[DSP](https://github.com/DSP-Labs/dsp)协议对[IPFS](https://github.com/ipfs/ipfs/)进行扩展的开源项目，通过[DSP](https://github.com/DSP-Labs/dsp)协议来确保数据高效、高并发、安全的进行分发。同时支持端到端的文件数据共享，并且通过加密技术的支持，可以在去中心化存储网络中，安全的进行文件的分享和稳定存储。通过文件分片加密及多点存储，可极大提高了网络数据传输的可用性及稳定性。在应用密码学的支持下同时实现了对文件自身访问权限的控制。

`go-ipfs-ex`是`IPFS-eX`的go语言实现。

>
>我们编写一组文档、帖子、教程来解释什么是DSP,什么是IPFS-eX，go-ipfs-ex是怎么运行的，以及它如何帮助您现有的和新的项目。
>
>- [**DSP协议规范**](https://github.com/DSP-Labs/specs)
>- [**IPFS-eX协议设计**](https://github.com/IPFS-eX/IPFS-eX)
>- [**IPFS-eX技术文档**](https://github.com/IPFS-eX/docs)
>- [**go-ipfs-ex**](https://github.com/IPFS-eX/go-ipfs-ex)
>- [**carrier**](https://github.com/IPFS-eX/carrier)

## 特性

- 主动分发
- 隐私保护
- 权限控制
- coming soon

## 安装指南
### 系统需求
[Golang](https://golang.org/doc/install) 1.13及以后的版本

`go-ipfs-ex`使用[Go modules](https://github.com/golang/go/wiki/Modules)进行依赖管理。

`go-ipfs-ex` 可以在大多数Linux、macOS和Windows系统上运行。建议在至少有2GB RAM和2个CPU内核的机器上运行它。
如果系统资源有限，建议：
1. 手动重建`go-ipfs-ex`
```
make build GOTAGS=OpenSSL
```
2. 使用go-ipfs-ex init --profile=lowpower

### 从源码安装
#### 安装go
`go-ipfs-ex`需要[Golang](https://golang.org/doc/install) 1.13及以后的版本，如果没有安装，可以[Download Go 1.13+](https://golang.org/dl/)。

可能需要添加Go的bin目录到系统`$PATH`下，可以在`/etc/profile`或者`$HOME/.profile`中添加下列代码
```
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOPATH/bin
```
(如有疑问，参考[Go install instructions](https://golang.org/doc/install)).
#### 下载并编译
```
$ git clone https://github.com/IPFS-eX/go-ipfs-ex.git
$ cd go-ipfs-ex
$ make install
```
## 入门指南
### 使用
*继续使用ipfs的命令行*
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
### 示例
coming soon...
## 开发
### API
[![GoDoc](https://godoc.org/github.com/IPFS-eX/go-ipfs-ex?status.svg)](https://godoc.org/github.com/IPFS-eX/go-ipfs-ex)
### 单元测试
```
`go test ./...`
```
## 贡献
感谢您考虑为IPFS-eX贡献源码，我们欢迎互联⽹网上的任何个体及组织参与到这一开源项目中来。

我们有如下的代码贡献规则需要大家共同遵守:

- 代码需要坚持遵循Golang[官方格式](https://golang.org/doc/effective_go.html%23formatting)指导文档(例如:使用`gofmt`);
- ⽂档需要坚持遵循Golang[官方注释](https://golang.org/doc/effective_go.html%23commentary)建议格式; 
- commit信息需要体现出修改的源码包名称; 
- 也可以通过进一步阅读[开发⼿册](https://github.com/IPFS-eX/docs)，获得更详细的开发信息。
## 许可证
`go-ipfs-ex` 同`go-ipfs`一样遵循MIT/Apache-2.0。
