# Go-APIServer-Study

GolangでAPIサーバを作成してみる

## APIサーバーの起動方法

### docker-compose を使う

1. `docker-compose up -d` を実行する
1. `http://localhost:8080/hello` にアクセス
1. `hello, golang!` と表示される

### go を使う

1. `go run server.go` を実行する
1. `http://localhost:8080/hello` にアクセス
1. `hello, golang!` と表示される