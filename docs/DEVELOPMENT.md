# Development

Guide for building, testing, and contributing to git-mirror-action.

<br/>

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Build](#build)
- [Testing](#testing)
- [Docker](#docker)
- [CI/CD Workflows](#cicd-workflows)
- [Conventions](#conventions)
- [Contributing](#contributing)

<br/>

## Prerequisites

- Go 1.24+
- Docker (for container builds)
- Make

<br/>

## Project Structure

```
.
├── cmd/
│   └── main.go                  # Entry point
├── internal/
│   ├── config/
│   │   ├── config.go            # Configuration loading & target parsing
│   │   └── config_test.go       # Config tests
│   ├── mirror/
│   │   ├── mirror.go            # Mirror logic & auth URL injection
│   │   └── mirror_test.go       # Mirror tests
│   └── output/
│       └── output.go            # GitHub Actions output writer
├── docs/                        # Documentation
├── .github/
│   ├── workflows/               # CI/CD workflows (9 files)
│   ├── dependabot.yml           # Dependency updates
│   └── release.yml              # Release note categories
├── action.yml                   # Action metadata (11 inputs, 3 outputs)
├── Dockerfile                   # Multi-stage (golang:1.24-alpine → alpine:3.21)
├── Makefile                     # Build, test, lint commands
├── cliff.toml                   # git-cliff changelog config
├── CODEOWNERS                   # Repository ownership
└── go.mod
```

### Key Directories

| Directory | Description |
|-----------|-------------|
| `cmd/` | Application entry point |
| `internal/config/` | Input parsing from `INPUT_*` env vars, target format parsing, provider auto-detection |
| `internal/mirror/` | Core mirroring logic: remote management, auth URL injection, branch/tag push |
| `internal/output/` | Writes results to `GITHUB_OUTPUT` file (result JSON, mirrored_count, failed_count) |

<br/>

## Build

```bash
make build           # Build binary → ./git-mirror-action
make clean           # Remove build artifacts
```

<br/>

## Testing

```bash
make test            # Run unit tests (alias)
make test-unit       # go test ./internal/... ./cmd/... -v -race -cover
make test-all        # Run all tests
make cover           # Generate coverage report
make cover-html      # Open coverage in browser
make bench           # Run benchmarks
```

### Test Coverage

CI enforces a minimum of **85%** coverage. Current testable areas:

- `internal/config/` — Config loading, target parsing, provider detection, env bool parsing
- `internal/mirror/` — Auth URL building, token injection

> **Note**: Many mirror functions use `exec.Command` for git operations, limiting pure unit test coverage. Integration testing is done via CI workflows.

<br/>

## Docker

```bash
# Build locally
docker build -t git-mirror-action .

# Test with dry run
docker run \
  --env INPUT_TARGETS="generic::https://example.com/repo.git" \
  --env INPUT_DRY_RUN="true" \
  --rm git-mirror-action
```

The Dockerfile uses a multi-stage build:
1. **Builder** — `golang:1.24-alpine` compiles the Go binary
2. **Runtime** — `alpine:3.21` with `git` and `openssh-client` only

<br/>

## CI/CD Workflows

| Workflow | Trigger | Description |
|----------|---------|-------------|
| `ci.yml` | push(main), PR, dispatch | Unit tests → Docker build → Action test → CI result |
| `release.yml` | tag push `v*` | Changelog → GitHub release → Major tag update |
| `changelog-generator.yml` | after release, PR merge | Auto-generate CHANGELOG.md |
| `contributors.yml` | after changelog | Auto-generate CONTRIBUTORS.md |
| `use-action.yml` | after release, dispatch | Smoke test with released action |
| `gitlab-mirror.yml` | push(main) | Backup to GitLab |
| `stale-issues.yml` | daily cron | Auto-close stale issues |
| `dependabot-auto-merge.yml` | PR (dependabot) | Auto-merge minor/patch updates |
| `issue-greeting.yml` | issue opened | Welcome message |

### Workflow Chain

```
tag push v* → Create release
                ├→ Smoke Test (Released Action)
                └→ Generate changelog
                      └→ Generate Contributors
```

<br/>

## Conventions

- **Commits**: Conventional Commits (`feat:`, `fix:`, `docs:`, `refactor:`, `test:`, `ci:`, `chore:`)
- **Secrets**: `PAT_TOKEN` (cross-repo ops), `GITHUB_TOKEN` (releases), `GITLAB_TOKEN` (mirror)
- **Docker**: Multi-stage build, alpine base with git + openssh-client
- **Comments**: English only in code
- **Release**: git-cliff for RELEASE.md, `somaz94/major-tag-action@v1` for major tag
- **cliff.toml**: Skip `^Merge`, `^Update changelog`, `^Auto commit` patterns

<br/>

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feat/my-feature`)
3. Write tests for new functionality
4. Ensure `make test` passes
5. Commit with conventional commit format
6. Submit a Pull Request
