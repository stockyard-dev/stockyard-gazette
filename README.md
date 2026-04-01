# Stockyard Gazette

**Blogging platform — write in Markdown, publish instantly, RSS out of the box**

Part of the [Stockyard](https://stockyard.dev) family of self-hosted developer tools.

## Quick Start

```bash
docker run -p 9250:9250 -v gazette_data:/data ghcr.io/stockyard-dev/stockyard-gazette
```

Or with docker-compose:

```bash
docker-compose up -d
```

Open `http://localhost:9250` in your browser.

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `9250` | HTTP port |
| `DATA_DIR` | `./data` | SQLite database directory |
| `GAZETTE_LICENSE_KEY` | *(empty)* | Pro license key |

## Free vs Pro

| | Free | Pro |
|-|------|-----|
| Limits | 10 posts | Unlimited posts |
| Price | Free | $2.99/mo |

Get a Pro license at [stockyard.dev/tools/](https://stockyard.dev/tools/).

## Category

Creator & Small Business

## License

Apache 2.0
