# go-twitter-proxy

Twitter Proxy

## Dependencies

- github.com/ChimeraCoder/anaconda
- github.com/gorilla/mux

## Run

```sh
go run main.go
```

## API

1. Get tweets by screen_name

   ```sh
   GET http://localhost:5000/tweets/{screen_name}
   ```

2. Get top 10 tweets by screen_name

   ```sh
   GET http://localhost:5000/tweets/{screen_name}/top-10
   ```
