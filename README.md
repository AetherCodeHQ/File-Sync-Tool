# File Sync Tool

![CI](https://github.com/Qyroxen/File-Sync-Tool/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/File-Sync-Tool/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/File-Sync-Tool?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/File-Sync-Tool)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/File-Sync-Tool)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/File-Sync-Tool?style=social)](https://github.com/Qyroxen/File-Sync-Tool/stargazers)

## What is it?

File Sync Tool is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/File-Sync-Tool.git
cd File-Sync-Tool
go build -o filesynctool .

# Run
./filesynctool --help
```

## CLI Usage

```bash
# Basic usage
./filesynctool

# With flags
./filesynctool --verbose --output json

# Get help
./filesynctool --help
```

## Examples

```bash
# Example 1
./filesynctool example1

# Example 2
./filesynctool example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o filesynctool .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/File-Sync-Tool/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/File-Sync-Tool?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/File-Sync-Tool/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/File-Sync-Tool?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/File-Sync-Tool/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/File-Sync-Tool" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/File-Sync-Tool/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/File-Sync-Tool" alt="Pull Requests">
  </a>
</p>
