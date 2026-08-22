# Conventions

- Standard library only unless a demonstrated requirement cannot be met.
- Preserve concrete HTML/Typst branches; no target enum, writer registry, or generic conversion pipeline.
- Execute subprocesses with `exec.Command(name, args...)`, never a shell string.
- Setup never invokes sudo; Linux privilege errors direct the user to `sudo mkiln setup`.
- Existing user config files are never overwritten.