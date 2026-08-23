# HTML math renderer decision

## 状態
`active`

## 背景
HTML math must remain fast to generate, offline to view, and compatible with mkiln's thin-Pandoc-wrapper boundary.

## 決定
Use Pandoc native MathML for the default HTML math method. Do not add a client-side math renderer, CDN dependency, resource vendoring, or renderer cache unless requirements materially change.

## 理由
KaTeX with Pandoc `embed-resources` preserved offline output but downloaded and embedded JS/CSS/fonts on each generation. In the observed sample this took about 12 seconds versus about 0.3 seconds for MathML. Runtime-CDN KaTeX would make generation fast but violate offline viewing. MathML needs no renderer assets or JavaScript and keeps the wrapper thin.

## コード起点
- `internal/assets/default.yaml` — `html-math-method`
- `internal/config.go` — `ensureUserConfig`

## 失敗・却下
- KaTeX plus `embed-resources` — generation latency from repeated CDN resource fetching and embedding.
- KaTeX loaded when opening HTML — generated files would require network access.
- Local KaTeX asset caching/vendor — adds cache lifecycle or bundled renderer complexity contrary to the thin-wrapper goal.

## 制約・罠
- Existing user config is intentionally never overwritten. Changing embedded defaults does not migrate `~/.config/mkiln/default.yaml`; an old KaTeX/`embed-resources` config can silently preserve the slow path and must be updated explicitly.
- Plain Typst conversion must remain independent of HTML defaults.
- `outline.js` is an intentional UX enhancement and is unrelated to the math-renderer decision.

## 検証
- `go test ./...`: 11 tests passed.
- `go vet ./...`: passed.
- Pandoc integration output contained native `<math>` markup and no KaTeX/MathJax/WebTeX resource references.