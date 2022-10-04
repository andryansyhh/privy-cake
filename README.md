# privy-cake
## Run/Test Project

Copy `.env.sample` for working configuration
```bash
cp .env.sample .env
```

Run application:
```bash
docker-compose up -d --build
```

Test application:
```bash
make generate-mocks
make test_coverage
```