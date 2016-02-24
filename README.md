This package contains a convenience function that takes a comma-separated
glob-style string and uses `spacelog.SetLevel` to set the log level
on all the matching loggers.

This is designed so that you can easily enable certain debug levels using
an environment variable or your other method of choice, similar to the NPM
module `debug`.

```go
// From some variable
config := "incoming*,core*"
levelstring.Configure(config, spacelog.Debug)
// Or straight from the environment
levelstring.Configure(os.GetEnv("DEBUG"), spacelog.Debug)
```

See http://godoc.org/github.com/dimfeld/spacelog-levelstring for more examples.
