// Adapted from DSDocTemplate by sankaku789.
// Source: https://github.com/sankaku789/DSDocTemplate
// License: CC BY 4.0. Modified for Pandoc and single-file mkiln output.
#import "@preview/codelst:2.0.2": *

#let dads = (
  white: rgb("#ffffff"), gray-50: rgb("#f2f2f2"), gray-100: rgb("#e6e6e6"),
  gray-200: rgb("#cccccc"), gray-536: rgb("#767676"), gray-700: rgb("#4d4d4d"),
  gray-800: rgb("#333333"), gray-900: rgb("#1a1a1a"), blue-50: rgb("#e8f1fe"),
  blue-300: rgb("#9db7f9"), blue-800: rgb("#0031d8"), blue-900: rgb("#0017c1"),
  blue-1100: rgb("#000071"), green-50: rgb("#e6f5ec"), green-800: rgb("#197a4b"),
  yellow-50: rgb("#fbf5e0"), yellow-900: rgb("#927200"), red-50: rgb("#fdeeee"),
  red-900: rgb("#ce0000"),
)
#let s8 = 6pt
#let s16 = 12pt
#let s24 = 18pt
#let s32 = 24pt
#let s64 = 48pt
#let doc-title = [mkiln 表示デバッグ]
#let doc-subtitle = [DADSスタイル・アウトライン・Pandoc要素の確認]
#let doc-date = [2026-08-23]
#let doc-authors = ()
#let doc-affiliation = []
#let sans-font = "Noto Sans CJK JP"
#let mono-font = "Noto Sans Mono CJK JP"

#let rule(color: dads.gray-200, thickness: 0.6pt) = {
  line(length: 100%, stroke: (paint: color, thickness: thickness))
}
#let horizontalrule = rule()
#let endnote(num, contents) = [#super[#num]#contents]

#show terms: it => it.children.map(child => [
  #strong[#child.term]
  #block(inset: (left: 1.5em, top: -0.4em))[#child.description]
]).join()

#let summaryBox(body, title: "概要") = block(
  fill: dads.blue-50, stroke: (paint: dads.blue-300, thickness: 0.8pt),
  inset: s16, radius: 4pt, width: 100%, above: s24, below: s24,
)[
  #text(weight: "bold", fill: dads.blue-1100)[#title]
  #v(4pt)
  #body
]

#let callout(body, kind: "info", title: none) = {
  let styles = (
    info: (bg: dads.blue-50, border: dads.blue-800, text: dads.blue-1100, label: "情報"),
    success: (bg: dads.green-50, border: dads.green-800, text: dads.green-800, label: "完了"),
    warning: (bg: dads.yellow-50, border: dads.yellow-900, text: dads.yellow-900, label: "注意"),
    danger: (bg: dads.red-50, border: dads.red-900, text: dads.red-900, label: "重要"),
  )
  let style = if kind in styles { styles.at(kind) } else { styles.info }
  block(fill: style.bg, stroke: (left: 4pt + style.border), inset: (x: s24, y: s8),
    radius: 4pt, width: 100%, above: s24, below: s24)[
    #text(weight: "bold", fill: style.text)[#if title == none { style.label } else { title }]
    #v(3pt)
    #body
  ]
}

#let kpi(label, value, note: "") = block(
  fill: dads.white, stroke: 0.7pt + dads.gray-200, inset: s16, radius: 4pt, width: 100%,
)[
  #text(size: 8pt, fill: dads.gray-700)[#label]
  #linebreak()
  #text(size: 20pt, weight: "bold", fill: dads.blue-1100)[#value]
  #if note != "" [#linebreak()#text(size: 8pt, fill: dads.gray-700)[#note]]
]

#let codeList(code, caption: none, lang: auto, numbers-side: "left", tab-size: 2,
  showrange: none, highlighted: (), numbers-start: auto) = figure(
  sourcecode(lang: lang, numbers-side: numbers-side, tab-size: tab-size,
    showrange: showrange, highlighted: highlighted, numbers-start: numbers-start)[#code],
  caption: caption, kind: "list",
)
#let includeSrc(filepath: "", lang: "plaintext", caption: none, numbers-side: "left") = {
  figure(sourcefile(read(filepath), lang: lang, numbers-side: numbers-side),
    caption: caption, kind: "list")
}
#let includePDF(file, margin-size: (top: 0mm, bottom: 0mm, left: 0mm, right: 0mm)) = {
  set page(numbering: none, margin: margin-size, columns: 1)
  image(file, format: "pdf")
  pagebreak()
  counter(page).update(1)
}
#let strong_ja(content) = text(weight: "bold", lang: "ja", font: sans-font)[#content]
#let large(content) = text(size: 12pt)[#content]
#let Large(content) = text(size: 14.4pt)[#content]
#let LARGE(content) = text(size: 17.28pt)[#content]
#let huge(content) = text(size: 20.74pt)[#content]
#let Huge(content) = text(size: 24.88pt)[#content]

#set document(title: doc-title, author: doc-authors)
#set page(
  paper: "a4", numbering: "1", number-align: center, fill: dads.white,
  margin: (top: 28mm, bottom: 24mm, left: 24mm, right: 24mm),
  header: context {
    set text(font: sans-font, size: 8pt, fill: dads.gray-700)
    grid(columns: (1fr, auto), gutter: 12pt, [#doc-title], [#doc-affiliation])
    v(3pt)
    rule(color: dads.gray-100, thickness: 0.5pt)
  },
  footer: context {
    set text(font: sans-font, size: 8pt, fill: dads.gray-700)
    rule(color: dads.gray-100, thickness: 0.5pt)
    v(3pt)
    align(center, counter(page).display())
  },
)
#set text(font: sans-font, lang: "ja", size: 10pt, fill: dads.gray-900)
#set par(first-line-indent: (amount: 1em, all: true), leading: 0.7em,
  justify: false, spacing: 1em)
#set footnote(numbering: "1 ")
#set list(indent: 1.1em, body-indent: 0.75em, marker: ([•], [-], [\*], [・]))
#set enum(indent: 1.1em, body-indent: 0.75em, numbering: "(1.a.i.A)")
#set bibliography(style: "sist02", full: true)
#set heading(numbering: "1.1 ", bookmarked: true)
#set table(
  inset: (x: 7pt, y: 6pt),
  stroke: (_, y) => if y == 0 { (bottom: 1pt + dads.blue-800) }
    else { (bottom: 0.5pt + dads.gray-100) },
  fill: (_, y) => if y == 0 { dads.blue-50 },
  align: left + horizon,
)
#show strong: set text(font: sans-font, weight: "bold")
#show raw: set text(font: mono-font, size: 9pt, fill: dads.gray-900)
#show link: it => text(fill: dads.blue-900, underline(it))
#show list: set par(first-line-indent: 0pt)
#show enum: set par(first-line-indent: 0pt)
#show footnote.entry: set par(first-line-indent: 0pt)
#show quote: it => block(fill: dads.gray-50, stroke: (left: 3pt + dads.blue-800),
  inset: (x: s24, y: s8), radius: 3pt, above: s24, below: s24)[
  #set text(fill: dads.gray-800)
  #set par(first-line-indent: 0pt)
  #it
]
#show raw.where(block: true): it => block(fill: dads.gray-50,
  stroke: 0.6pt + dads.gray-200, inset: 9pt, radius: 3pt, width: 100%,
  above: s24, below: s24)[
  #set text(font: mono-font, size: 8.5pt)
  #it
]
#show heading: it => {
  set text(font: sans-font, weight: "bold", fill: dads.gray-900)
  set block(above: s24, below: s8)
  if it.level == 1 {
    block[
      #text(size: 18pt, fill: dads.blue-1100)[#it]
      #v(5pt)
      #rule(color: dads.blue-800, thickness: 1.2pt)
    ]
  } else if it.level == 2 {
    block[
      #text(size: 13pt)[#it]
      #v(3pt)
      #rule(color: dads.gray-200, thickness: 0.7pt)
    ]
  } else {
    text(size: 11pt, fill: dads.gray-800)[#it]
  }
}
#show figure: set block(breakable: true)
#show figure.where(kind: table): set figure(placement: none, supplement: [表], numbering: "1.1")
#show figure.where(kind: table): set figure.caption(position: top, separator: [: ])
#show figure.where(kind: image): set figure(placement: none, supplement: [図], numbering: "1.1")
#show figure.where(kind: image): set figure.caption(position: bottom, separator: [: ])
#show figure.where(kind: "list"): set figure(placement: none, supplement: [リスト], numbering: "1.1")
#show figure.where(kind: "list"): set figure.caption(position: top, separator: [: ])

// DSDocTemplate-style cover
#set page(numbering: none, header: none, footer: none,
  margin: (top: 27mm, bottom: 25mm, left: 20mm, right: 20mm))
#v(s64)
#rule(color: dads.blue-800, thickness: 3pt)
#v(s24)
#text(size: 30pt, weight: "bold", fill: dads.blue-1100)[#doc-title]
#v(s8)
#text(size: 14pt, fill: dads.gray-700)[#doc-subtitle]
#v(1fr)
#rule()
#v(s8)
#grid(columns: (1fr, 1fr), gutter: s24,
  [#text(size: 8pt, fill: dads.gray-700)[作成者]#linebreak()#doc-authors.join("、")],
  [#text(size: 8pt, fill: dads.gray-700)[日付]#linebreak()#doc-date],
)
#pagebreak()
#counter(page).update(1)
#set page(numbering: "1", header: context {
  set text(font: sans-font, size: 8pt, fill: dads.gray-700)
  grid(columns: (1fr, auto), [#doc-title], [#doc-affiliation])
  v(3pt); rule(color: dads.gray-100, thickness: 0.5pt)
}, footer: context {
  set text(font: sans-font, size: 8pt, fill: dads.gray-700)
  rule(color: dads.gray-100, thickness: 0.5pt); v(3pt)
  align(center, counter(page).display())
})
#outline(title: [目次], depth: 3)
#pagebreak()

= 基本タイポグラフィ
<基本タイポグラフィ>
この文書は、mkilnが生成するHTMLの表示を確認するためのデバッグ用Markdownです。本文、見出し、アウトライン、リンク、表、コード、数式などを一度に確認できます。

通常の本文は16px以上、行間1.7で表示します。日本語が長く続いた場合でも、一行が広がりすぎず、ブラウザの文字拡大時に内容が欠落しないことを確認します。英数字を含む文章として、Digital
Agency Design SystemとMarkdown renderingの折り返しも確認します。

== 見出しレベル2
<見出しレベル2>
見出しレベル2は大きなセクションを表します。デスクトップとモバイルで文字サイズと上下余白が変化します。

=== 見出しレベル3
<見出しレベル3>
見出しレベル3はアウトラインにも表示されます。現在位置へスクロールすると、左側の対応項目が青い背景で強調されます。

==== 見出しレベル4
<見出しレベル4>
レベル4以降は本文には表示されますが、既定のアウトライン深度3では目次に含まれません。

===== 見出しレベル5
<見出しレベル5>
小さな見出しでも本文との区別が保たれることを確認します。

====== 見出しレベル6
<見出しレベル6>
最下位の見出しです。見出しレベルは視覚サイズではなく文書構造として使用します。

= インライン要素
<インライン要素>
== 文字装飾
<文字装飾>
これは#strong[太字];、これは#emph[強調];、これは#strike[取り消し線];です。インラインコードは`go test ./...`のように表示されます。キーボード入力や識別子として`Ctrl+Shift+R`、`HTMLBuild`、`--include-in-header`も確認します。

== リンク
<リンク>
#link("https://design.digital.go.jp/dads/")[デジタル庁デザインシステム];への通常リンクです。一度開いたリンクは訪問済みの紫色になり、キーボードフォーカス時は黄色背景と黒いアウトラインが表示されます。

長いURLの折り返し確認:
#link("https://example.com/documents/debug/very-long-path/that-must-wrap-without-expanding-the-document-beyond-the-viewport")

== 脚注
<脚注>
本文中に脚注を置きます。#footnote[Markdownの解析とHTML生成はPandocへ委譲しています。]
脚注リンクをキーボードで操作でき、本文へ戻れることを確認します。

= リスト
<リスト>
== 箇条書き
<箇条書き>
- 第1階層は黒丸です
- 長い項目は複数行へ折り返され、2行目が項目本文の開始位置に揃います。日本語の文章が増えてもアウトラインや本文領域を押し広げません
  - 第2階層は白丸です
  - 関連する項目をまとめます
    - 第3階層は四角です
    - 深い階層でも余白を維持します
- 最後の項目です

== 番号付きリスト
<番号付きリスト>
+ 入力Markdownを解析する
+ Pandoc defaultsを適用する
+ CSSとアウトライン用JSをHTMLへ埋め込む
+ 単体で移動可能なHTMLを生成する

== タスクリスト
<タスクリスト>
- ☒ HTMLを生成する
- ☒ CSSを埋め込む
- ☐ 複数の画面幅で目視確認する

= 引用と区切り
<引用と区切り>
== 引用
<引用>
#quote(block: true)[
デザインシステムは見た目を統一するだけでなく、情報構造、操作方法、アクセシビリティを一貫させるために使用します。

引用が複数段落でも、灰色の左罫線と適切な内側余白を維持します。
]

== 水平線
<水平線>
前の内容です。

#horizontalrule

水平線の後の内容です。セクション見出しの代わりには使用しません。

= コード
<コード>
== コードブロック
<コードブロック>
```go
package main

import "fmt"

func main() {
    message := "mkiln debug"
    fmt.Println(message)
}
```

== 長いコード行
<長いコード行>
コードブロックだけが横スクロールし、ページ全体には横スクロールが発生しないことを確認します。

```text
GET /api/v1/documents/example?include=metadata,sections,links,footnotes&format=standalone-html&style=default HTTP/1.1
```

= 表
<表>
== 基本テーブル
<基本テーブル>
#figure(
align(center)[#table(
  columns: 3,
  align: (col, row) => (auto,right,auto,).at(col),
  inset: 6pt,
  [項目], [既定値], [確認内容],
  [本文],
  [16px / 1.7],
  [日本語の可読性と文字拡大],
  [アウトライン],
  [256px],
  [左端固定と独立スクロール],
  [本文最大幅],
  [960px],
  [デスクトップでの可変幅],
  [ブレークポイント],
  [1024px],
  [1カラムへの切り替え],
)]
)

== 内容量の多いテーブル
<内容量の多いテーブル>
#figure(
align(center)[#table(
  columns: 3,
  align: (col, row) => (auto,auto,auto,).at(col),
  inset: 6pt,
  [環境], [動作], [備考],
  [モバイル],
  [アウトラインを本文上部へ配置],
  [横幅を圧迫せず、通常の文書順で読み上げる],
  [デスクトップ],
  [アウトラインを画面左端へ固定],
  [現在の見出しを選択状態として表示する],
  [印刷],
  [固定配置を解除],
  [コード、引用、表の途中改ページを可能な範囲で避ける],
)]
)

= 数式
<数式>
== インライン数式
<インライン数式>
質量とエネルギーの関係は $E = m c^2$
です。インライン数式が本文の行高を不自然に崩さないことを確認します。

== ディスプレイ数式
<ディスプレイ数式>
$ f (x) = integral_(- oo)^oo hat(f) (xi) e^(2 pi i x xi) thin d xi $

KaTeXで中央表示され、狭い画面でも内容へアクセスできることを確認します。

= レスポンシブ確認
<レスポンシブ確認>
== モバイル幅
<モバイル幅>
ブラウザの開発者ツールで幅を768px未満へ変更します。本文には左右16pxの余白が残り、見出し、表、コードがビューポート外へ本文全体を押し出さないことを確認します。

== タブレット幅
<タブレット幅>
768pxから1023pxでは1カラムを維持し、本文幅は最大48remです。アウトラインは本文上部のカードとして表示されます。

== デスクトップ幅
<デスクトップ幅>
1024px以上ではアウトラインが画面左端へ固定され、本文領域は画面幅に応じて最大960pxまで広がります。

== 文字拡大
<文字拡大>
ブラウザを200%へ拡大します。実効ビューポートが狭くなるため1カラムへ切り替わり、情報の欠落や重なりが発生しないことを確認します。

= アウトライン追従
<アウトライン追従>
== スクロール位置の検出
<スクロール位置の検出>
ページを上下へスクロールすると、現在位置の直前にある見出しが選択されます。高速スクロール時にも複数の更新処理が重複しないよう、`requestAnimationFrame`で調整しています。

== キーボード操作
<キーボード操作>
Tabキーでアウトライン内のリンクへ移動し、Enterキーで対応する見出しへ移動します。フォーカス表示と現在位置表示が色だけに依存せず、太字とアウトラインでも区別できることを確認します。

== 最終セクション
<最終セクション>
ページ末尾までスクロールしたとき、この項目が選択状態になります。ここまで確認できれば、アウトラインの長さ、独立スクロール、現在位置表示の基本動作をまとめて検証できます。
