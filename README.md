# file-manager

## Running Locally

```bash
make build-server
./bin/server run --config config.yaml
```

`config.yaml` has the following content:

```yaml
httpPort: 8080
grpcPort: 8081

debug:
  standalone: true
  sqlitePath: /tmp/file_manager.db
```

You can then connect to the DB.

```bash
sqlite3 /tmp/file_manager.db
```

You can then hit the endpoint.

```bash
curl http://localhost:8080/v1/files
```
