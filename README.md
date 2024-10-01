# Spider

Spider is a WordPress website (blog) deployment tool. It Automates the process of launching and transfering database assets to a new WordPress blog site. Designed with multisite migration in mind, it is ready to 

![Spider](spider.webp)

## Prerequisites

- Googles' [Go language](https://go.dev) installed to enable building executables from source code.

- An `env.json` file containing enviromental data:

```json
	"production": {
		"url": "example.com",
		"path": "example_com",
		"server": "example.dmz"
	},
```

## Build

From the root folder containing *main.go*, use the command that matches your environment:

### Windows & Mac:

```bash
go build -o [name] .
```

### Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

```bash
[program] [source flag] [destination flag] [website slug]
```

## Flags

Current flages are:

- -s (Staging)

- -p (Production)

- -b (Blog)

- -d (Development)

- -t (Test)

- -e (Engage)

- -f (Forms)

- -w (Working)

- -v (Vanity)

Example deployment:

```bash
spider -s -p antiracism
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/spider/blob/main/LICENSE.md) and is part of the Public Domain.
