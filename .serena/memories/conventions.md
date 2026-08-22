# Conventions

- Standard library only unless a demonstrated requirement cannot be met.
- Preserve concrete HTML/Typst branches; no target enum, writer registry, or generic conversion pipeline.
- The Typst template is adapted from DSDocTemplate under CC BY 4.0. Preserve source URI, attribution, and modification notice when changing/distributing it.
- Keep document metadata in Markdown/Pandoc variables, not a sidecar `config.yaml`, so `.typ` remains self-contained apart from referenced document resources and the codelst package.
- Execute subprocesses with `exec.Command(name, args...)`, never a shell string.
- Setup never invokes sudo; Linux privilege errors direct the user to `sudo mkiln setup`.
- Existing user config files are never overwritten.