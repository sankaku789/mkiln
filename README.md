# mkiln

[English](README.en.md) | [简体中文](README.cn.md)

mkilnは、Pandocを利用してMarkdownを自己完結HTMLまたはプレーンTypstソースへ変換する、薄いGo製CLIです。Markdown解析・数式・コードハイライト・各writerはPandocへ委譲します。

## 目的
MarkdownをHTMLまたはTypstを単純変換するツールが欲しかったですが、Pandocは機能が豊富すぎるゆえに、オプションを設定するのが少々手間でした。そのため、薄いラッパーを作ることで単純なコマンドでMarkdown->HTML、Markdown->Typstができると考え作成しました。

## 機能

- MarkdownからStandalone HTML、MarkdownからTypstを出力する。
- Standalone HTMLは数式にmathmlを使用し、シンプルなビューを提供する。

## 必要環境

- Go 1.24.4以降（ソースからビルドする場合）
- Pandoc

## インストール

```bash
go install github.com/sankaku789/mkiln/cmd/mkiln@latest
```

リポジトリからビルドする場合:

```bash
go build ./cmd/mkiln
```

## 使い方
### セットアップ

Pandocがない場合は`setup`コマンドでPandocをインストールできます。

```bash
mkiln setup
```

なお、インストールには、winget、apt、pacman等OSのデフォルトパッケージマネージャーを使用します。

### 変換

```bash
mkiln note.md  [-o PATH] [-s NAME]
```

`note.html`を生成します。HTMLにはCSS、アウトラインが埋め込まれるため、生成後は単一ファイルで移動・共有・オフライン閲覧できます。  

`-o`, `--output`オプションはoutput先のファイル名とパスを指示します。  

`-s`, `--style`オプションは、使用するCSSを特定ファイルに明示します。オプション非利用時はこちらで設定してるCSSを使用します。

```bash
mkiln FILE --typst [-o PATH]
```

`-t`, `--typst`オプションで、typstソースを出力します。

## 設定ファイル

HTML初回生成時に、ユーザの設定ディレクトリ下へ次を作成します。

```text
mkiln/
├── default.yaml
└── styles/
    └── default.css
```

- `default.yaml`はPandoc defaultsです。mkiln独自形式ではありません。
- `styles/NAME.css`は`--style NAME`で選択できます。
- 既存ファイルは自動上書きしません。
- `--typst`は設定ディレクトリを作成しません。

## 開発

```bash
gofmt -w cmd internal
go test ./...
go vet ./...
go build ./cmd/mkiln
```

## 参考
CSSレイアウトは、[デジタル庁デザインシステム](https://design.digital.go.jp/dads/)を参考にしその[サンプル実装](https://github.com/digital-go-jp/design-system-example-components-html)のコンポーネントを一部利用しています。


## ライセンス
MITライセンスで提供されます。
