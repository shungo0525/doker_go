- zipファイル生成
```
% GOOS=linux go build -o handler handler.go
% zip function.zip line_bot
% aws lambda update-function-code --function-name LineBot --zip-file fileb://function.zip
```

- zipファイルをlambdaにアップロードする
- ref
  - https://www.cview.co.jp/cvcblog/2021.01.15.Pu1pTwLHNUXVe6dnNHrHc