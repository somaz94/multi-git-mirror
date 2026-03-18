# Configuration

Complete reference for all inputs, outputs, and target configuration.

<br/>

## Table of Contents

- [Inputs](#inputs)
- [Outputs](#outputs)
- [Target Format](#target-format)
- [Supported Providers](#supported-providers)
- [Branch Configuration](#branch-configuration)
- [Default Values](#default-values)

<br/>

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `targets` | Mirror target URLs (newline-separated, `provider::url` or auto-detect) | Yes | - |
| `gitlab_token` | GitLab personal access token | No | `''` |
| `github_token` | GitHub personal access token | No | `''` |
| `bitbucket_username` | Bitbucket username for app password auth | No | `''` |
| `bitbucket_password` | Bitbucket app password | No | `''` |
| `ssh_private_key` | SSH private key for SSH-based authentication | No | `''` |
| `mirror_branches` | Branches to mirror (comma-separated, or `all`) | No | `all` |
| `mirror_tags` | Mirror tags | No | `true` |
| `force_push` | Use force push | No | `true` |
| `dry_run` | Log actions without pushing | No | `false` |
| `debug` | Enable debug logging | No | `false` |

<br/>

## Outputs

| Output | Description | Example |
|--------|-------------|---------|
| `result` | JSON array with mirror results per target | `[{"target":{...},"success":true,"message":"mirrored successfully"}]` |
| `mirrored_count` | Number of successfully mirrored targets | `2` |
| `failed_count` | Number of failed mirror targets | `0` |

### Result JSON Structure

Each element in the `result` array has the following structure:

```json
{
  "target": {
    "Provider": "gitlab",
    "URL": "https://gitlab.com/org/repo.git"
  },
  "success": true,
  "message": "mirrored successfully"
}
```

<br/>

## Target Format

Targets are specified one per line in the `targets` input. Two formats are supported:

```
provider::url          # explicit provider
url                    # auto-detect from URL
```

### Examples

```yaml
targets: |
  gitlab::https://gitlab.com/myorg/myrepo.git
  codecommit::https://git-codecommit.us-east-1.amazonaws.com/v1/repos/myrepo
  https://bitbucket.org/myorg/myrepo.git
```

In the example above, the third target will auto-detect `bitbucket` as the provider from the URL.

<br/>

## Supported Providers

| Provider | Auth Method | Auto-detect Pattern |
|----------|-------------|---------------------|
| `gitlab` | OAuth2 token (`oauth2:<token>`) | URL contains `gitlab` |
| `github` | x-access-token (`x-access-token:<token>`) | URL contains `github` |
| `bitbucket` | Username + App password (`user:pass`) | URL contains `bitbucket` |
| `codecommit` | IAM / credential-helper (URL as-is) | URL contains `codecommit` |
| `generic` | SSH key or URL as-is | Default fallback |

<br/>

## Branch Configuration

### Mirror All Branches

```yaml
mirror_branches: 'all'    # default
```

### Mirror Specific Branches

```yaml
mirror_branches: 'main,develop,release'
```

Branches are specified as a comma-separated list. Each branch is pushed using the refspec `refs/heads/<branch>:refs/heads/<branch>`.

<br/>

## Default Values

| Setting | Default | Notes |
|---------|---------|-------|
| `mirror_branches` | `all` | Pushes all branches with `--all` flag |
| `mirror_tags` | `true` | Pushes tags with `--tags` flag |
| `force_push` | `true` | Uses `-f` flag for force push |
| `dry_run` | `false` | When `true`, skips actual push |
| `debug` | `false` | When `true`, logs git commands |

Boolean inputs accept: `true`, `1`, `yes` (truthy) or `false`, `0`, `no` (falsy).
