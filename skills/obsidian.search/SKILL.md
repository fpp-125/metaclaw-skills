# obsidian.search

Searches markdown notes from a mounted Obsidian vault and returns ranked snippets.

## Inputs

- `query`: natural language query
- `vault_path`: mounted vault directory

## Outputs

- ranked snippets with source note path

## Security Notes

- Requires read-only mount to `/vault`.
- Does not write to source notes.
