# Remove Label Tool

A simple Go utility to remove a specific label from all issues and pull requests in a GitHub repository.

## Features

- **Dry-run by default**: Safely preview changes before applying them.
- **Pagination**: Handles repositories with a large number of issues.
- **Supports Issues & PRs**: Correctly identifies and processes both.
- **No dependencies**: Uses only the Go standard library.

## Prerequisites

- [Go](https://go.dev/doc/install) installed.
- A GitHub Personal Access Token (PAT) with `repo` scope.

## Usage

You can run the tool directly using `go run`:

```bash
go run tools/remove_label.go \
    -owner <REPOSITORY_OWNER> \
    -repo <REPOSITORY_NAME> \
    -token <YOUR_GITHUB_TOKEN>
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-owner` | GitHub repository owner (user or organization) | (Required) |
| `-repo` | GitHub repository name | (Required) |
| `-token` | GitHub Personal Access Token | (Required) |
| `-label` | The label to remove | `lfx-mentorship/docs-analysis` |
| `-dry-run` | Whether to simulate removal or perform it. Set to `false` to apply changes. | `true` |

### Examples

**Preview changes (Dry-run):**
```bash
go run tools/remove_label.go -owner envoyproxy -repo gateway -token your_token
```

**Perform actual removal:**
```bash
go run tools/remove_label.go -owner envoyproxy -repo gateway -token your_token -dry-run=false
```

**Remove a different label:**
```bash
go run tools/remove_label.go -owner envoyproxy -repo gateway -token your_token -label "stale" -dry-run=false
```
