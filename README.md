

## Purpose

This repository is a proof of concept for integrating the compose file support in the devworkspace operator repository by reusing the pkgs of kubernetes kompose.

- Currently there does not exist a way to utilize the komopse as a library (only as CLI)

## Work in progress

- Addition of unit testing for testing different compose files.
- Conversion of kubernetes files to components and their deployment.

## Installation

This repo is written in pure golang, so make sure you have golang setup (go 1.16+) on your system before setting up the development environment.

To run use the following command: 


```sh
go run main.go
```
