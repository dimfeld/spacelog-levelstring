This package contains a convenience function that takes a comma-separated
glob-style string and uses `spacelog.SetLevel` to set the log level
on all the matching loggers.

This is designed so that you can easily enable certain debug levels using
methods such as an environment variable or an HTTP endpoint.


```go
// From some variable
config := "incoming*,core*"
levelstring.Configure(config, spacelog.Debug)
// Or straight from the environment
levelstring.Configure(os.Getenv("DEBUG"), spacelog.Debug)
```

See http://godoc.org/github.com/dimfeld/spacelog-levelstring for more examples.

### Limitations

Calls to `Configure` only affect logger names that exist at the time of the call. New
loggers created afterward will use the logger collection's default level.
