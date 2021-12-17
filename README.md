# Gochain -  A Simple Blockchain in Go

![stability-wip](https://img.shields.io/badge/stability-work_in_progress-lightgrey.svg)

[![On Push](https://github.com/JoseRodrigues443/miguel-blockchain-golang/actions/workflows/go.yml/badge.svg)](https://github.com/JoseRodrigues443/miguel-blockchain-golang/actions/workflows/go.yml)

A blockchain implementation in golang.
For better understanding how the core of a blockchain would be implemented.

In teh future will have a consensus system, with proof of work and proof of stake to show the differences in implementation.

## Requirements

- make
- golang

## How to run

`make run`

Should show a output like:

```bash

Genesis is done
______________
Block number ==>  0
Previous ==>
Data ==>  Genesis
Hash ==> 81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5
______________
Block number ==>  1
Previous ==>  81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5
Data ==>  Block One
Hash ==> 5a94fa6b1698092f166e6b9b92aa0745d8b14008f0ffa4a8cf74d0ac2c51a3e1
______________
Block number ==>  2
Previous ==>  5a94fa6b1698092f166e6b9b92aa0745d8b14008f0ffa4a8cf74d0ac2c51a3e1
Data ==>  Block Two
Hash ==> 8f8f556e94227e3f906f91962724409664fce56abe48db2fbc57f3b9feb6475d
______________
Block number ==>  3
Previous ==>  8f8f556e94227e3f906f91962724409664fce56abe48db2fbc57f3b9feb6475d
Data ==>  Block Three
Hash ==> 1a8e02750937a379dfba064f9380c89cdb521a53a783a8558e73b74384990e4d

```
