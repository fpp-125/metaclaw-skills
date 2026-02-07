# metaclaw-skills

Reusable skill modules for MetaClaw.

## Repository Contract

Every skill directory must include:

- `SKILL.md`
- `capability.contract.yaml`

Example: `skills/obsidian.search/`

## Validation

Run skill linter:

```bash
go run ./cmd/skilllint ./skills
```

Run tests:

```bash
go test ./...
```

## Publish Flow

1. Implement skill and contract in this repo.
2. Tag a version.
3. Build/publish OCI artifact for the skill bundle.
4. Register metadata in `metaclaw-registry`.
