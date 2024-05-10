# query ibc channel packet-receipt

Query a packet receipt

### Synopsis

Query a packet receipt

```
zetacored query ibc channel packet-receipt [port-id] [channel-id] [sequence] [flags]
```

### Examples

```
zetacored query ibc channel packet-receipt [port-id] [channel-id] [sequence]
```

### Options

```
      --grpc-addr string   the gRPC endpoint to use for this chain
      --grpc-insecure      allow gRPC over insecure channels, if not TLS the server must use TLS
      --height int         Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help               help for packet-receipt
      --node string        [host]:[port] to Tendermint RPC interface for this chain 
  -o, --output string      Output format (text|json) 
      --prove              show proofs for the query results (default true)
```

### Options inherited from parent commands

```
      --chain-id string     The network chain ID
      --home string         directory for config and data 
      --log_format string   The logging format (json|plain) 
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) 
      --log_no_color        Disable colored logs
      --trace               print out full stack trace on errors
```

### SEE ALSO

* [zetacored query ibc channel](zetacored_query_ibc_channel.md)	 - IBC channel query subcommands
