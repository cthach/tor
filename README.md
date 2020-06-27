# tor

Go package for interfacing with the Tor network. 

This is a WIP project in *active development* with more features to come. The `master` branch is considered experimental 
and is subject to change. No releases have been made so far.

## Features
- Tiny as possible
- No imported third-party dependencies
- Apache 2.0 license
- Exit node lookup via the offical [Tor Project DNSEL](https://lists.torproject.org/pipermail/tor-project/2020-March/002759.html) service

## API

TODO: GoDoc

```
// IsExit returns nil if the IPv4 address belongs to an exit node.
IsExit(ctx context.Context, ip string) (bool, error)
```

## CLI
### check-exit
#### Compile
```
$ make build/check-exit
```

#### Run
Check whether an IP address is a known exit node with `check-exit $IP_ADDR`.

Is a known exit node:
```
$ ./check-exit 185.220.101.204

PASS! 185.220.101.204 IS a known exit node
```

Is not a known exit node:
```
$ ./check-exit 1.2.3.4

FAIL! 1.2.3.4 IS NOT a known exit node
```

## Contact

If you like to say hello, send me a message at [https://chris-thach.com](https://chris-thach.com).