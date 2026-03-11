# tim

Timezone conversion CLI tool.

## Install

```
cargo install --git https://github.com/yumazak/tim
```

Or download a binary from [Releases](https://github.com/yumazak/tim/releases).

## Usage

### `h` — Convert hour (0-23)

```bash
# JST 9:00 → UTC (default: Asia/Tokyo → UTC)
tim h 9
# 0

# UTC → JST
tim h -f UTC -t Asia/Tokyo 15
# 0

# From stdin
printf "9\n12\n15\n" | tim h
# 0
# 3
# 6
```

### `dt` — Convert datetime

```bash
# Naive datetime (interpreted as --from timezone)
tim dt 2024-01-15T09:00:00
# 2024-01-15T00:00:00+00:00

# With explicit timezones
tim dt -f UTC -t Asia/Tokyo 2024-01-15T00:00:00
# 2024-01-15T09:00:00+09:00

# RFC3339 input (offset in input takes precedence over --from)
tim dt 2024-01-15T09:00:00+09:00
# 2024-01-15T00:00:00+00:00

# Seconds can be omitted
tim dt 2024-01-15T09:00

# From stdin
printf "2024-01-15T09:00:00\n2024-06-15T12:30:00\n" | tim dt
```

### `tz` — List all IANA time zones

```bash
tim tz
# Africa/Abidjan
# Africa/Accra
# ...
```

### Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--from` | `-f` | `Asia/Tokyo` | Source timezone (IANA name) |
| `--to` | `-t` | `UTC` | Target timezone (IANA name) |

### Supported datetime formats

| Format | Example |
|--------|---------|
| RFC3339 | `2024-01-15T09:00:00+09:00` |
| ISO 8601 (seconds) | `2024-01-15T09:00:00` |
| ISO 8601 (no seconds) | `2024-01-15T09:00` |
| Space separator | `2024-01-15 09:00:00` |
| Space separator (no seconds) | `2024-01-15 09:00` |

## License

MIT
