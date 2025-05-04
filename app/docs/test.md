# testの方法

- Unit test
- go test -v -coverprofile=coverage.out .でcoverageが100%になったら終了

```
coverageを確認
make test TEST_DIR=./domain/hoge

coverできてないとこの確認
make serve TEST_DIR=./domain/hoge

coverageファイルの削除
make clean TEST_DIR=./domain/hoge
```

- カレントディレクトリが app/domain/hoge の場合、app/Makefile を実行するには make の -C オプションでmake -C ../../ testこのように実行可能