# Go Struct Stream Generator

## 環境変数

|変数|説明|
|---|---|
|STREAMER_DIRECTORY|読み取りたいディレクトリ|
|STREAMER_RECURSION|再帰的に処理をするかどうか(何にかセットするだけで再帰を辞める)|

## Build

```shell
docker build . -t go-streamer
docker run --rm -v ${PWD}:/file go-streamer
```
