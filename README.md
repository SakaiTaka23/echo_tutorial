# echo_tutorial



## Guide

- [x] Installation 完了
- [x] Customization カスタム
- [x] Context
- [x] Cookies
- [x] Error Handling
- [x] Migration 移行方法
- [x] Request
- [x] Response
- [x] Routing
- [x] Static Files
- [x] Templates
- [x] Testing




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



## ErrorHandling

* あまり挙動がわからない

* デフォルトではエラーが出てきた時メッセージが返ってくるだけ

  →それをカスタムすることができる？




## Request

* 送られてきたリクエストの値をどのように取得するか

1.  専用の構造体を作成しておき、json時、form時、url query時それぞれに関してどのような名前で取得されるかを定義する
2. c (echo.Context)からBindをを使用すると自動で取得してくれる
   * formの場合はFormValueを使うと取得できる



## Response

* ハンドラーからの返り値として文字列、**html、json、xml、csv、png、それ以外のファイルも**返すことができる

*  同じ内容でもいくつかの返信方法がある。
  * Blob:バイナリデータをそれぞれの形に変換して返す？
  * Stream:大きいものに有効 一行で返す
  * Pretty:綺麗に整形させたものを返す
* Context#Redirect(code int, url string) でリダイレクトもできる(ステータスコードは200!?)
* Hooks?リスポンス前後に処理？



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



**パラメーター**

* パラメーターを使ったルーティングもすることができる

* *を使用することによりあらゆる値がマッチするようになる

  /users/1/files/	/users/1/files/hi	/users/1/files/ok/wo/1 ...

```go
// param
e.GET("/users/:id", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users/:id")
})

// static
e.GET("/users/new", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users/new")
})

// match any
e.GET("/users/1/files/*", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users/1/files/*")
})

```

* このような順番で記述した場合であってもechoではstatic→param→match anyの順番で変換される



**URI Building?**

あまりわからなかった



* ルートの出力もできる



## Static Files

echoで静的ファイルを扱う方法

* ルーティングの段階で静的ファイルの場所を定義する

```go
					ルート				実際の場所
e.Static("/static", "assets")

e.File("/favicon.svg", "images/favicon.svg")
e.File("/", "public/index.html")
```

Staticでは使用するファイル（js,css)をFileではファイルを返す場合(html)、faviconを返す場合に使用

**テンプレート側から使用する場合もルートを書いて取得する**

```html
<script src='static/main.js'></script>
```

実際のmain.jsの場所は assets/main.js




## Templates

echoでのハンドラーからテンプレートを返す方法

1. TemplateRenderを作成
2. Render実装　公式の方のifはいらなさそう
3. main関数の時点でテンプレートをプリコンパイル、echoクラスのRenderにプリコンパイルしたものを入れる

実際に使う時はハンドラーでc.Renderを返す。引数は ステータス、テンプレート、返すデータ　となっている。テンプレートにはテンプレート内で定義した名前が使用できる。

```go
{{ define "header" }}
{{ end }}
```



## Testing

流れ自体は他の言語と同じ

* 試しに値を入れそれが変えるべき値と同じかどうかテストする

**実行方法**

1. ハンドラーが*.goの場合テストは\*_test.goとする

2. テストは

   ````shell
   go test -v
   ````

   で実行



## Middleware

- [x] Basic Auth
- [x] Body Dump
- [x] Body Limit
- [ ] Casbin Auth
- [ ] CORS
- [ ] CSRF
- [ ] Gzip
- [ ] Jaeger
- [ ] JWT
- [ ] Key Auth
- [ ] Logger
- [ ] Method Override
- [ ] Primetheus
- [ ] Proxy
- [ ] Recover
- [ ] Redirect
- [ ] Request ID
- [ ] Rewrite
- [ ] Secure
- [ ] Session
- [ ] Static
- [ ] Trailing Slash



* ミドルウェアはグループとして登録することによってそれを使うもの、使わないものを分けることができる



## Basic Auth

* ベーシック認証



## Body Dump

* デバッグ・ログ確認用　
* リクエストが飛んできた時にrequest body,response bodyを表示させることができる
* いずれかの値が大きい時(ファイルや画像を大量に送る時？)は例外処理をしておくこと



## Body Limit

* request bodyの容量を制限することができる
* リミットを超えた場合 413エラーを返す



## Cabin Auth

* わからない
* どのような種類の認証なのかがわかっておらず、sampleも動かない
* みた感じgoのcasbinというパッケージをechoでも使いやすいようにカスタマイズしたもの？



## CORS

* クロスドメインの許可をするもの？



## CSRF

* csrfトークン
* 