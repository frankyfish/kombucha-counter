# Kombucha Counter

Angular: v20.0.2 (npm v11.3.0)


## Backend

### Redis
By convention `<type-of-document>:<id>`.


Local development with Docker:

```bash
docker run --name some-redis -d -p 6379:6379 redis redis-server --save 60 1 --loglevel warning

docker exec -it some-redis bash
```

#### Basic Redis CLI Commands
```bash
# get all key-value pairs for hash record with id 0
redis-cli HGETALL kombucha:0

redis-cli DEL kombucha:0
```

#### Redis-Go

Example of hash creation(`Differences with hash documents`): https://redis.io/docs/latest/develop/clients/go/queryjson/#differences-with-hash-documents 