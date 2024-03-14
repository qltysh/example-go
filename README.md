# Qlty Go Coverage Example

[Qlty](https://example.com) is a Code Health Platform with support for code coverage.

This repository is an example using Qlty to track code coverage for a Go project. Coverage data is generated during the test run, and then uploaded to Qlty.

## Requirements

- Go 1.17 or above
- Test run with `go test -coverprofile=coverage.out`
- An account on Qlty (free for open source)
- `QLTY_COVERAGE_TOKEN` is set as a GitHub Actions [repository secret](https://docs.github.com/en/actions/security-guides/using-secrets-in-github-actions#creating-secrets-for-a-repository)

## Set up

See [`.github/workflows/main.yml`](./.github/workflows/main.yml) in this repository for a basic configuration.

## Documentation

- [Advanced code coverage configuration](https://example.com)
- [Alternative supported CI providers](https://example.com)

## Help and feedback

Join the our [Slack Community](https://example.com) for help and to provide feedback that we'll use to improve Qlty.

## License

[MIT License](./LICENSE.md)
