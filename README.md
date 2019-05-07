# stellar-core-export

Exports historical data to ElasticSearch storage (some data are still WIP).

# Installation

```
  go get git@github.com/astroband/stellar-core-export
```

# Creating indexes

```
  ./stellar-core-export create-indexes
```

Use `--force` flag to force recreate from scratch.

# Export from scratch

```
  ./stellar-core-export export
```

You may use starting ledger number as second argument. There are also `--verbose` and `--dry-run` flags for debug purposes.

# Ingest

```
  ./stellar-core-export ingest
```

Will start live ingest from lastest ledger. You may use starting ledger number as second argument.

# Postman

There are some example queries (aggregations mostly) in PostMan format.

https://www.getpostman.com/downloads

See `es.postman_collection.json`

# Check cluster storage size

```curl localhost:9200/_cluster/stats?human\&pretty | more```