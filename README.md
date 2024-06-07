# file-manager

## Running with Docker Compose

Run the following command:

```bash
docker-compose build
docker-compose up
```

## Running Locally

```bash
make build-server
./bin/server run --config config.yaml
```

`config.yaml` has the following content:

```yaml
httpPort: 8080
grpcPort: 8081
workerServiceGrpcPort: 8082

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

curl http://localhost:8080/v1/files \
  --form purpose="fine-tune" \
  --form file="@mydata.jsonl"
```
