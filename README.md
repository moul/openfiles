# openfiles

:smile: openfiles

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/openfiles)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/openfiles/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/openfiles.svg)](https://github.com/moul/openfiles/releases)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/openfiles.svg)](https://microbadger.com/images/moul/openfiles)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

[![Go](https://github.com/moul/openfiles/workflows/Go/badge.svg)](https://github.com/moul/openfiles/actions?query=workflow%3AGo)
[![Release](https://github.com/moul/openfiles/workflows/Release/badge.svg)](https://github.com/moul/openfiles/actions?query=workflow%3ARelease)
[![GolangCI](https://golangci.com/badges/github.com/moul/openfiles.svg)](https://golangci.com/r/github.com/moul/openfiles)
[![codecov](https://codecov.io/gh/moul/openfiles/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/openfiles)
[![Go Report Card](https://goreportcard.com/badge/moul.io/openfiles)](https://goreportcard.com/report/moul.io/openfiles)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/openfiles/badge)](https://www.codefactor.io/repository/github/moul/openfiles)


## Usage

```go
import "moul.io/openfiles"

count, err := openfiles.Count()
fmt.Println(count) // 42

emfile := openfiles.IsTooManyError(err)
fmt.Println(emfile) // true
```

## Install

### Using go

```console
$ go get -u moul.io/openfiles
```

### Releases

See https://github.com/moul/openfiles/releases

## Contributing

<!-- FIXME: contributing.gif -->

I really welcome contributions. Your input is the most precious material. I'm well aware of that and I thank you in advance. Everyone is encouraged to look at what they can do on their own scale; no effort is too small.

Everything on contribution is sum up here: [CONTRIBUTING.md](./CONTRIBUTING.md)

### Contributors ✨

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg)](#contributors)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://manfred.life"><img src="https://avatars1.githubusercontent.com/u/94029?v=4" width="100px;" alt=""/><br /><sub><b>Manfred Touron</b></sub></a><br /><a href="#maintenance-moul" title="Maintenance">🚧</a> <a href="https://github.com/moul/openfiles/commits?author=moul" title="Documentation">📖</a> <a href="https://github.com/moul/openfiles/commits?author=moul" title="Tests">⚠️</a> <a href="https://github.com/moul/openfiles/commits?author=moul" title="Code">💻</a></td>
    <td align="center"><a href="https://manfred.life/moul-bot"><img src="https://avatars1.githubusercontent.com/u/41326314?v=4" width="100px;" alt=""/><br /><sub><b>moul-bot</b></sub></a><br /><a href="#maintenance-moul-bot" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

### Stargazers over time

[![Stargazers over time](https://starchart.cc/moul/openfiles.svg)](https://starchart.cc/moul/openfiles)

## License

© 2020 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
