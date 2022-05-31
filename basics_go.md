## VSCode での環境構築

```console
go install -v golang.org/x/lint/golint@latest
go install -v github.com/sqs/goreturns@latest
go install -v github.com/stamblerre/gocode@latest
```

settings.json に下記を追加

```json
{
  "[go]": {
    "editor.tabSize": 4,
    "editor.insertSpaces": false,
    "editor.formatOnSave": true,
    "editor.formatOnPaste": false,
    "editor.formatOnType": false
  },
  "go.formatTool": "goreturns",
  "go.formatFlags": ["-w"],
  "go.lintTool": "golint",
  "go.lintOnSave": "package",
  "go.vetOnSave": "package",
  "go.buildOnSave": "package",
  "go.testOnSave": false,
  "go.gocodeAutoBuild": true,
  "go.installDependenciesWhenBuilding": true
}
```

## Hello World

```Go

// 一番初めに実行される特別な関数
// パッケージの初期化に用いられる
// main 関数よりも先に呼ばれる
func init() {
	fmt.Println("Init!")
}

func main(){
	fmt.Println("Hello","World")
}

```

### Go Doc の使い方

```console
go doc fmt Println
```

## 変数の宣言

```Go
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)
fmt.Println(i, f64, s, t, f)

// 下記の書き方は関数内でしか使えない
xi := 1
xf64 := 1.2
xs := "test"
xt, xf := true, false
```
