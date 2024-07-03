# Cozy - light configuration library for Go.

#### Suitable for use during application initialization, singleton allows you to use from any module after initialization



```go
// Init requires specifying the path to the file. The default setting is nil. By default, config.cozy.
if err := cozy.Init(nil); err != nil {
    panic(err)
}

var port = new(any)

if err := cozy.Get("port", *port); err != nil {
    panic(err)
}
```