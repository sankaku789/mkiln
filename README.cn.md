# mkiln

[日本語](README.md) | [English](README.en.md)

mkiln 是一个轻量的 Go 命令行工具，使用 Pandoc 将 Markdown 转换为自包含 HTML 或纯 Typst 源码。Markdown 解析、数学公式、代码高亮和各类 writer 均由 Pandoc 负责。

## 目的

我希望有一个能够简单地将 Markdown 转换为 HTML 或 Typst 的工具，但 Pandoc 功能非常丰富，配置选项略显繁琐。因此，我创建了这个轻量包装器，以便通过简单的命令完成 Markdown 到 HTML、Markdown 到 Typst 的转换。

## 功能

- 将 Markdown 转换为 standalone HTML 或 Typst。
- Standalone HTML 使用 MathML 显示数学公式，并提供简洁的阅读界面。

## 环境要求

- 从源码构建时需要 Go 1.24.4 或更高版本
- Pandoc

## 安装

```bash
go install github.com/sankaku789/mkiln/cmd/mkiln@latest
```

从仓库构建：

```bash
go build ./cmd/mkiln
```

## 使用方法

### 设置

如果未安装 Pandoc，可以使用 `setup` 命令进行安装。

```bash
mkiln setup
```

安装时会使用操作系统默认的包管理器，例如 winget、apt 或 pacman。

### 转换

```bash
mkiln note.md [-o PATH] [-s NAME]
```

该命令生成 `note.html`。HTML 中嵌入了 CSS 和大纲，因此生成后的单个文件可以移动、共享并离线浏览。

`-o`, `--output` 选项用于指定输出文件名和路径。

`-s`, `--style` 选项用于选择 CSS 文件。未指定时使用随附的默认 CSS。

```bash
mkiln FILE --typst [-o PATH]
```

`-t`, `--typst` 选项用于输出 Typst 源码。

## 配置文件

首次生成 HTML 时，mkiln 会在用户配置目录下创建：

```text
mkiln/
├── default.yaml
└── styles/
    └── default.css
```

- `default.yaml` 是 Pandoc defaults 文件，不是 mkiln 自定义格式。
- 使用 `--style NAME` 选择 `styles/NAME.css`。
- 已存在的文件不会被自动覆盖。
- `--typst` 不会创建配置目录。

## 开发

```bash
gofmt -w cmd internal
go test ./...
go vet ./...
go build ./cmd/mkiln
```

## 参考

CSS 布局参考了[日本数字厅设计系统](https://design.digital.go.jp/dads/)，并复用了其[示例实现](https://github.com/digital-go-jp/design-system-example-components-html)中的部分组件。

## 许可证

本项目采用 MIT 许可证。
