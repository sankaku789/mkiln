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
#let doc-title = [$if(title)$$title$$else$無題$endif$]
#let doc-subtitle = [$if(subtitle)$$subtitle$$endif$]
#let doc-date = [$if(date)$$date$$endif$]
#let doc-authors = ($for(author)$$if(author.name)$"$author.name$",$else$"$author$",$endif$$endfor$)
#let doc-affiliation = [$if(affiliation)$$affiliation$$endif$]
#let sans-font = "$if(mainfont)$$mainfont$$else$Noto Sans CJK JP$endif$"
#let mono-font = "$if(monofont)$$monofont$$else$Noto Sans Mono CJK JP$endif$"

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

$body$

$if(citations)$
$if(csl)$#set bibliography(style: "$csl$")$endif$
$if(bibliography)$#bibliography($for(bibliography)$"$bibliography$"$sep$,$endfor$)$endif$
$endif$
