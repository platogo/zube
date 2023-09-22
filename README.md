# About

This is a small library that implements a Zube API client in Go.

## Usage

You can take a look at [the examples folder](examples/) to see how to use the library.

```go
go run examples/epics/fetch_epics.go
```

Remember to set your own credential values first.

## Authentication

You can read the details of the auth mechanism in the [official docs](https://zube.io/docs/api#verifying-a-private-key).

In short, you will need:

- Your Zube `clientId`
- Your private key as a `.pem` file
- Your access token

By default, the library will look for the private key in `~/.ssh/zube_api_key.pem`, but this will be configurable in the future.

This is a barebones implementation of creating a client. If the access token passed is empty, it will be automatically fetched from Zube.

```go
 clientId := "fd59a5e4-2505-4be0-8818-4f25466fded3"
 client, _ := zube.NewClient(clientId, "")

 client.FetchSources()
```

If you don't want to go through the access token regeneration after every call to the constructor, you can cache the access token and re-use it (note: as of the time of writing it is valid for only up to 24 hours).
