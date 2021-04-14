# Zap to Hashicorp logger adapter

Package implements `hclog.Logger` interface from [Hashicorp's `go-hclog` project](https://github.com/hashicorp/go-hclog).  
It wraps [Uber's zap logger](https://github.com/uber-go/zap) and adapts it to be used as a logger in Hashicorp's software (like Consul, etc.).