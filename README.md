goAppSetter
======================
AppSetter をたたくプログラム

必要なもの
------
追加では必要ない

ファイルリスト
------
- README.md
    - このファイル
- config.yml
    - 設定ファイル
    - AppSetter.aspへのパス等の情報
- goAppSetter.go
    - ソースコード
- pkg
    - パッケージフォルダ
    - darwin_386
        - goAppSetter
        - 32bit Mac OS X 用バイナリ
    - darwin_amd64
        - goAppSetter
        - 64bit Max OS X 用バイナリ
    - linux_386
        - goAppSetter
        - 32bit Linux 用バイナリ
    - linux_amd64
        - goAppSetter
        - 64bit Linux 用バイナリ
    - windows_386
        - goAppSetter.exe
        - 32bit Windows 用バイナリ
    - windows_amd64
        - goAppSetter.exe
        − 64bit Windows 用バイナリ


設定ファイル
------

```
servers:
    - server:
        scheme: http
        host: 192.168.0.1
        port: 80
        path: /asp/admin/AppSetter.asp
    - server:
        scheme: http
        host: dev-web.example.com
        port: 80
        path: /api/asp/admin/AppSetter.asp
    - server:
        scheme: http
        host: dev-api.example.com
        port: 80
        path: /asp/admin/AppSetter.asp
```

- `servers:` サーバ情報
    - `- server:` サーバ情報（配列）
        - `scheme:` スキーム
        - `host:` ホスト
        - `post:` ポート
        - `path:` AppSetterへのパス


使い方
------
    % ./goAppSetter -h
    Usage:
      goAppSetter [OPTIONS] TARGET_FILE

    Help Options:
      -h, --help  Show this help message


