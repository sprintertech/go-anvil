# `go-anvil`: Go Bindings for Anvil

[![Go Reference](https://pkg.go.dev/badge/github.com/sprintertech/go-anvil.svg)](https://pkg.go.dev/github.com/sprintertech/go-anvil)
[![Go Report Card](https://goreportcard.com/badge/github.com/sprintertech/go-anvil)](https://goreportcard.com/report/github.com/sprintertech/go-anvil)

`go-anvil` provides an easy way to use [Anvil](https://book.getfoundry.sh/anvil/) nodes from Go.

> [!NOTE]
> `go-anvil` requires the `foundy` binary to be installed. See the [foundry](https://github.com/foundry-rs/foundry) repository for installation instructions.

```
go get github.com/sprintertech/go-anvil
```

## Getting Started

```go
cli := anvil.New(
    anvil.WithPort(8545),
)

err := cli.Run()
if err != nil {
    print(err)	
} 
```

> [!WARNING]
> This package is pre-1.0. There might be breaking changes between minor versions.