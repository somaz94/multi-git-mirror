# Changelog

All notable changes to this project will be documented in this file.

## Unreleased (2026-04-16)

### Documentation

- remove duplicate rules covered by global CLAUDE.md ([d100f8e](https://github.com/somaz94/multi-git-mirror/commit/d100f8e2428892996da19f622286fa2420731f53))

### Chores

- **deps:** bump dependabot/fetch-metadata from 2 to 3 ([6972347](https://github.com/somaz94/multi-git-mirror/commit/697234790673be30647005c74aa1754692d9037e))
- **deps:** bump softprops/action-gh-release from 2 to 3 ([1f264df](https://github.com/somaz94/multi-git-mirror/commit/1f264df595a75603a4382e82226ce389fac0aa5c))
- remove duplicate rules from CLAUDE.md (moved to global) ([5567a87](https://github.com/somaz94/multi-git-mirror/commit/5567a87d229c5e879307ce92012283996ba83f2e))
- add git config protection to CLAUDE.md ([548b81e](https://github.com/somaz94/multi-git-mirror/commit/548b81e659d8a14be1405bf6f692564785c84d90))

### Contributors

- somaz

<br/>

## [v1.1.1](https://github.com/somaz94/multi-git-mirror/compare/v1.1.0...v1.1.1) (2026-03-25)

### Bug Fixes

- correct license badge to MIT, optimize Docker build ([57deca5](https://github.com/somaz94/multi-git-mirror/commit/57deca5b43e87567de4351083d1318ced1e983b6))

### Documentation

- add no-push rule to CLAUDE.md ([094881a](https://github.com/somaz94/multi-git-mirror/commit/094881a2d7aed45e1e681bc173dbc091cbf8ecab))

### Continuous Integration

- skip auto-generated changelog and contributors commits in release notes ([f1d6abf](https://github.com/somaz94/multi-git-mirror/commit/f1d6abf96297f2bfffd7d006f6a0628c7bda3245))
- revert to body_path RELEASE.md in release workflow ([8005744](https://github.com/somaz94/multi-git-mirror/commit/80057449b84e57bb402cd8b8cbd2aa82a1011120))
- use generate_release_notes instead of body_path in release workflow ([3f2b874](https://github.com/somaz94/multi-git-mirror/commit/3f2b87431cf705f270e6f82990e74bed0f4b627d))
- add auto-generated PR body script for make pr ([211ea5c](https://github.com/somaz94/multi-git-mirror/commit/211ea5cb9b809d3d50a3f05c5420affd1f52a521))

### Chores

- add workflow Makefile targets (check-gh, branch, pr) ([fcb15ce](https://github.com/somaz94/multi-git-mirror/commit/fcb15cef90ee6c58f8a2666acb4c3be1b16bf4c5))
- upgrade Go version to 1.26 ([516524e](https://github.com/somaz94/multi-git-mirror/commit/516524ee003e23d1e916947964fd99380e21e9d4))

### Contributors

- somaz

<br/>

## [v1.1.0](https://github.com/somaz94/multi-git-mirror/compare/v1.0.1...v1.1.0) (2026-03-18)

### Features

- change action name ([76a413e](https://github.com/somaz94/multi-git-mirror/commit/76a413e5ac3baca0d6809fb2fd74a09e789702b6))

### Code Refactoring

- rename git-mirror-action to multi-git-mirror across entire project ([2e4c6c4](https://github.com/somaz94/multi-git-mirror/commit/2e4c6c4d23fc6a82ab3f888beb9617a8ce5533b8))

### Documentation

- README.md ([9cce1f3](https://github.com/somaz94/multi-git-mirror/commit/9cce1f3ea49dabcdf42b57ce784bedae0b569ae5))

### Contributors

- somaz

<br/>

## [v1.0.1](https://github.com/somaz94/multi-git-mirror/compare/v1.0.0...v1.0.1) (2026-03-18)

### Bug Fixes

- use secrets for Bitbucket username and document x-token-auth usage ([0903b9e](https://github.com/somaz94/multi-git-mirror/commit/0903b9ed8b4e3c3182621404afff2c4ffc9d5da5))
- use local action in bitbucket-mirror.yml until next release ([a5d0211](https://github.com/somaz94/multi-git-mirror/commit/a5d0211bed0a3efe08eb1133a4a0fd5a6acf7ee2))

### Code Refactoring

- rename bitbucket_password to bitbucket_api_token ([6d523b4](https://github.com/somaz94/multi-git-mirror/commit/6d523b4f3a7e835b488bb5597ec6d0d9c6286f6f))

### Documentation

- note Workspace Access Token requires Bitbucket Premium plan ([bec56b9](https://github.com/somaz94/multi-git-mirror/commit/bec56b9cea5e08ec8fbafb815d35bf2edeee8cc4))
- update Bitbucket auth guide with Repository/Workspace Access Tokens ([4369237](https://github.com/somaz94/multi-git-mirror/commit/4369237470d747bf4623a51a504d14cfd8b6a460))

### Continuous Integration

- disable Bitbucket tests until plan issue is resolved ([c07baa4](https://github.com/somaz94/multi-git-mirror/commit/c07baa471b77f1478dcf6d5699d2b93b2f2f199d))
- add Bitbucket mirror workflow and Bitbucket CI tests ([221443e](https://github.com/somaz94/multi-git-mirror/commit/221443e44af6159571255fe7268650abb451412d))

### Contributors

- somaz

<br/>

## [v1.0.0](https://github.com/somaz94/multi-git-mirror/releases/tag/v1.0.0) (2026-03-18)

### Features

- add retry, pre-check, exclude branches, and parallel mirroring ([7165fea](https://github.com/somaz94/multi-git-mirror/commit/7165fea86181e48f994b3ab9fb04a5005e0f72cb))
- scaffold Go-based git-mirror-action project structure ([5529b82](https://github.com/somaz94/multi-git-mirror/commit/5529b826a354f67829bd3692df46c4c631760641))

### Bug Fixes

- resolve parallel race condition, add mutex and prevent credential prompts ([d2802e5](https://github.com/somaz94/multi-git-mirror/commit/d2802e543f59d8b71d85a14c0e86521ec3aa126c))
- ensure git repo exists for Docker dry-run and improve CI test targets ([7aadbb3](https://github.com/somaz94/multi-git-mirror/commit/7aadbb3bbc7edf66a390656d81fb652d1cf1ef76))
- **ci:** use exit code validation instead of grep on annotations ([f749cbe](https://github.com/somaz94/multi-git-mirror/commit/f749cbec8d37cd98c7a3f44d117d1e60564d8780))
- improve security for all providers ([a1276f2](https://github.com/somaz94/multi-git-mirror/commit/a1276f20c04b598b5a47c846e2f2164089c91b02))
- add git-lfs to Docker image to prevent pre-push hook failure ([8a1ef8d](https://github.com/somaz94/multi-git-mirror/commit/8a1ef8de32480b8af30b6c81ffa02720ba666d49))
- **ci:** remove deprecated buildx install input and add missing env vars ([423456e](https://github.com/somaz94/multi-git-mirror/commit/423456e101d417e7e0a7ac5bf6031f45444cc596))
- add safe.directory config for Docker workspace ownership ([915f76d](https://github.com/somaz94/multi-git-mirror/commit/915f76d2b8b1a3b34e4c4095b61db58786c75ae9))

### Code Refactoring

- add SSH support, credential encoding, log masking, and config validation ([b09eb66](https://github.com/somaz94/multi-git-mirror/commit/b09eb660354c921208782479f6b7fedca60fdd0e))

### Documentation

- update documentation for new features and add authentication guides ([d0b4ef7](https://github.com/somaz94/multi-git-mirror/commit/d0b4ef7fa0d97ac6194718d66c6a0d2bcdbf06aa))
- update AUTHENTICATION.md and CLAUDE.md for new features ([230e8f8](https://github.com/somaz94/multi-git-mirror/commit/230e8f88c85a319a8597da3a60ab2c47f82ab1e0))
- add documentation, workflows, and GitHub configs ([f9db271](https://github.com/somaz94/multi-git-mirror/commit/f9db271e885629641a040a8ace7eda3f7414f274))

### Tests

- improve SSH test coverage with configurable sshDir ([a337e4c](https://github.com/somaz94/multi-git-mirror/commit/a337e4cf52aa754d93c1a90127cd9722b27aa781))
- add tests for SSH, credential encoding, masking, and validation ([4a7d7d6](https://github.com/somaz94/multi-git-mirror/commit/4a7d7d6919896cc7b9cc0c14d0a95015431cd689))
- improve test coverage with mock-based tests ([7f3b1d3](https://github.com/somaz94/multi-git-mirror/commit/7f3b1d3d5a4aec5469522dd0deb8559212403be2))
- add unit tests for config, mirror, and output packages ([ec5a8a2](https://github.com/somaz94/multi-git-mirror/commit/ec5a8a2255b8da486fdcfabe7d65c27f57168bb2))

### Continuous Integration

- improve CI and smoke test coverage ([87191c9](https://github.com/somaz94/multi-git-mirror/commit/87191c9b5cf3b89c460f550373628a67c9f7f74a))

### Contributors

- somaz

<br/>

