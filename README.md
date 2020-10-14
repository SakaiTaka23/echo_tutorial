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



## Cookies

echoでのクッキーの設定・読み取り方法

1. http.cookieクラスを作っておいて中に値を入れる

2. echo.contextからSetCookieを使うとクッキーを設定することができる



入れることのできる値

* 最低限name,valueは必要でその他はオプション

| Attribute | Optional |
| --------- | -------- |
| Name      | No       |
| Value     | No       |
| Path      | Yes      |
| Domain    | Yes      |
| Expires   | Yes      |
| Secure    | Yes      |
| HttpOnly  | Yes      |



削除

クッキーにMaxAge -1を設定しそのクッキーを書き込むと削除される