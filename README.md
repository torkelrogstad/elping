# `elping`

`elping` is a simple tool that allows you to ping Electrum servers.

## Installing

```bash
$ go install github.com/torkelrogstad/elping
```

## Usage

```
elping <target>
  -debug
        turn on debug output
  -timeout duration
        ping timeout (default 3s)
  -tls
        dial with TLS (default true)
```

`target` is mandatory, and must be of the form `host:port` (no protocol,
Electrum operates over raw TCP connections!).
