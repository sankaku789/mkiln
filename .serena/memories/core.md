# mkiln core

- Thin Go CLI around Pandoc: Markdown to standalone HTML is primary; DADS-templated Typst source plus PDF is secondary.
- Keep HTML and Typst flows explicit. `-t`/`--typst` is a bool, never a generic target option. Typst emits sibling `.typ` and `.pdf` files.
- Pandoc owns parsing/rendering/highlighting. Do not add Markdown/YAML parsers or expose arbitrary writers.
- Source map: `cmd/mkiln` entry point; `internal` CLI/config/Pandoc/setup logic; embedded defaults under `internal/assets` (Go embed cannot traverse to a root assets directory).
- HTML conversion creates missing user config files only; Typst conversion must not create/read mkiln config.
- Detailed stack and commands: `mem:tech_stack`, `mem:suggested_commands`; conventions: `mem:conventions`; completion checks: `mem:task_completion`.