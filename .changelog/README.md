# Changelog Fragments

Each pull request should include a changelog entry as a separate file in this directory.

## Adding an Entry

Create a new `.md` file with a descriptive name, e.g. `add-bfd-resource.md`:

```markdown
- Add `nxos_bfd` resource and data source
```

- Each line must start with `- `
- Multiple entries per file are allowed for multi-change PRs
- The filename is for identification only (lowercase, hyphens, `.md` extension)

## At Release Time

Run `make release VERSION=X.Y.Z` to collect all fragments into `CHANGELOG.md` and delete the fragment files.

## Previewing

Run `make changelog-preview` to see what the next release section will look like.
