# echo_tutorial


## templates

echoでのハンドラーからテンプレートを返す方法

1. TemplateRenderを作成
2. Render実装　公式の方のifはいらなさそう
3. main関数の時点でテンプレートをプリコンパイル、echoクラスのRenderにプリコンパイルしたものを入れる

実際に使う時はハンドラーでc.Renderを返す。引数は ステータス、テンプレート、返すデータ　となっている。テンプレートにはテンプレート内で定義した名前が使用できる。

```go
{{ define "header" }}
{{ end }}
```

