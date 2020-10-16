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



## Request

* 送られてきたリクエストの値をどのように取得するか

1.  専用の構造体を作成しておき、json時、form時、url query時それぞれに関してどのような名前で取得されるかを定義する
2. c (echo.Context)からBindをを使用すると自動で取得してくれる
   * formの場合はFormValueを使うと取得できる



## Routing

* ルーティングの方法

e(echo class)にメソッド(get,post,put,delete,updateなど)をつけるとルートができる

```go
e.GET("/hello",some_handler)
e.method(route,handler)

e.Any(route,handler)
e.Match([methods],route,handler)
```

* echoではそれ以外にもどのメソッドでも受け付けるAnyや幾つかのメソッドを受け付けるMatchを使うこともできる



* echoでのハンドラー関数は echo.Contextを受け取りerrorを返すものが定義されている
* 





















