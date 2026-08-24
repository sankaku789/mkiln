# DADS / browser print CSS

## 状態
`unfinished`

## 背景
`internal/assets/default.css` の文書表示や印刷挙動を変更するときに必要な設計境界。

## 決定
- DADS由来ruleの参照元は、作業時に取得済みだった隣接ローカルcopy `../design-system-example-components-html` に固定する。外部から最新版や別versionを取得して置換しない。
- DADSのfoundation / prose / TOCだけをPandocのHTML構造へadaptし、Storybook専用ruleやcomponent library全体、npm build chainは持ち込まない。
- screenとbrowser printは同じembedded `default.css` で管理する。print専用themeやGo側のtarget abstractionは作らない。
- `outline.js` はprogressive enhancementに留め、JSなしでも静的TOCを利用可能にする。
- 長いtableは印刷時に分割可能とし、table全体ではなくrowの途中分割だけを避ける。

## 理由
mkilnをPandocの薄いwrapperのまま保ち、DADS upstream更新やfrontend toolchainへのruntime/build依存を避けながら、由来が追跡可能な文書CSSとブラウザ印刷を提供するため。

## コード起点
- `internal/assets/default.css` — DADS-adapted foundations/prose/TOC、mkiln screen layout、print profile
- `internal/assets/outline.js` — TOC progressive enhancement
- `THIRD_PARTY_NOTICES` — DADS MIT attribution
- `internal/mkiln_test.go` — embedded print profile contract

## 制約・罠
- `#TOC > ul` をscreen既定で非表示にすると、JS無効時にTOCが利用不能になる。折り畳みはJSが追加する `.mobile-doc-header` 配下だけへ適用する。
- 印刷時に `table { break-inside: avoid; }` を指定しない。1ページを超えるtableの空白やlayout破綻につながる。
- native MathMLとplain Typst経路はCSS変更から独立させる。

## 次
- 日本語本文、MathML、横長code/table/imageを含むsampleを複数browserの通常表示と印刷previewで目視確認する。
- GitHub ActionsのUbuntu/Windows matrixを実際に通し、Windows上でfake executable testを確認する。

## 検証
- `go test ./...`、`go vet ./...`、`go build ./cmd/mkiln`、`git diff --check` は成功。
- Windows向けtest binaryとCLIのcross compileは成功。
- browser print previewとWindows runner上の実行は未検証。
