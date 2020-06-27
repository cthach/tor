# tor

Go package for interfacing with the Tor network. 

This is a WIP project in *active development* with more features to come. The `master` branch is considered experimental 
and is subject to change. No stable releases have been made so far.

## Features
- Exit node lookup using the offical [Tor Project DNSEL](https://lists.torproject.org/pipermail/tor-project/2020-March/002759.html) service
- No imported third-party dependencies
- Apache 2.0 license

## API

TODO: GoDoc

```
// IsExit returns nil if the IPv4 address belongs to an exit node.
IsExit(ctx context.Context, ip string) (bool, error)
```

## CLI Tools
### check-exit


Check whether an IP address is a known exit node with `check-exit $IP_ADDR`.

Is a known exit node:
```
$ go run cmd/check-exit/main.go 185.220.101.204

PASS! 185.220.101.204 IS a known exit node
```

Is not a known exit node:
```
$ go run cmd/check-exit/main.go 1.2.3.4

FAIL! 1.2.3.4 IS NOT a known exit node
```

## Contact

If you like to say hello, send me a message at [https://chris-thach.com](https://chris-thach.com).