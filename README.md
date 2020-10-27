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
- [ ] Gzip データを圧縮するのに使われる？ファイルもできそう？
- [ ] Jaeger 文さんトレーシングシステム？
- [x] JWT
- [ ] Key Auth
- [x] Logger
- [ ] Method Override メソッドのオーバーライド　POSTのみ可能
- [x] Primetheus
- [x] Proxy プロキシの設定ができる
- [x] Recover
- [x] Redirect
- [x] Request ID
- [x] Rewrite
- [x] Secure
- [x] Session
- [x] Static
- [x] Trailing Slash



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
* なぜか認証が失敗する



## JWT

* ログインに使用できる

* デフォルトではユーザー名、パスワードのみを含むがカスタムをすることもできる

  →これをフォーム用にカスタムすることもできそう？



* ログインの手順

  1. loginのルートに入りjwtトークンを受け取る
  2. 認証が必要なルートに対して毎回そのトークンを送り確認、返答をする

  * この時 1.から直接ロブインページに送ることできる？
  * クッキーにjwtのトークンを保存しておけば良さそう？



## Logger

* ミドルウェアに追加するだけでログが流れるようになる

```go
	e.Use(middleware.Logger())
```

* フォーマットの指定もできる

```go
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  Format: "method=${method}, uri=${uri}, status=${status}\n",
}))
```



## Primetheus

* 追加するとhttp requestの詳細を受け取ることができる？

```go
    p := prometheus.NewPrometheus("echo", nil)
```



## Recover

* 追加することでパニックを回避しエラーを出させることができる
* ログを吐くのでエラーが出たときに追跡することができる

```go
	e.Use(middleware.Recover())
```



## Redirect

* ハンドラーから別のルートへリダイレクトするものではなく根本的にurlの一部を変えるために使用される

例 httpをhttpsに置き換え、wwwがついていない場合に追加、それらの組み合わせ



## Request id

* 毎回のリクエストに対してユニークなidを割り振る
* 場所はリスポンスのヘッダー部に追加される

* ヘッダーに追加されるみたいだがヘッダーを出力させても表示が見えない

  →postmanのヘッダーでは確認することができる

* ログインなしで乱数を生成したい時に使えそう？



## Rewrite

* urlを書き直すことができる
* 長いurlを短く変換したり条件を絞る時に使える



## Secure

* インジェクション攻撃から守ってくれる？



## Session

* セッション基本的にクッキーと併用



## Static

* 静的ファイルの場所を指定することができる
* 指定することによってルート指定を短縮させることができる

例 staticを静的ファイルルートに指定→/js/main.jsにリクエストがあった場合static/js/main.jsを取得する



## Trailing Slash

* url末尾の/を取り除くか含むかを指定することができる

* 末尾のスラッシュがあるかないかでルートが変わってくる

  →トレイリングスラッシュをつけてリクエストを送りそのルートがない場合404が帰ってくる

* これを使うことによりそれらを統一することができる