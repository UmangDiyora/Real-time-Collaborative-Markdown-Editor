# Real-time Collaborative Markdown Editor

A production-ready real-time collaborative markdown editor built with Go, similar to HackMD/CodiMD. Multiple users can edit the same document simultaneously with real-time synchronization and conflict resolution.

## Features

- **Real-time Collaboration**: Multiple users can edit documents simultaneously
- **Conflict Resolution**: Operational Transformation (OT) or CRDT for handling concurrent edits
- **WebSocket Sync**: Low-latency real-time synchronization
- **User Authentication**: JWT-based authentication with OAuth support
- **Document Management**: Create, edit, share, and version documents
- **Live Preview**: Real-time markdown rendering with syntax highlighting
- **Presence Awareness**: See who's editing, cursor positions, and selections
- **Export/Import**: Support for MD, HTML, PDF, DOCX, LaTeX formats
- **Security**: TLS, WSS, rate limiting, and OWASP compliance

## Architecture

```
markdown-collab/
├── cmd/                 # Entry points
│   ├── server/         # Main server
│   └── migrate/        # Database migrations
├── internal/           # Private application code
│   ├── auth/          # Authentication logic
│   ├── document/      # Document management
│   ├── session/       # Session handling
│   ├── sync/          # Synchronization engine
│   ├── transform/     # OT/CRDT implementation
│   ├── websocket/     # WebSocket handling
│   ├── presence/      # User presence tracking
│   └── export/        # Document export
├── pkg/               # Public libraries
│   ├── models/        # Data models
│   ├── protocol/      # Wire protocol
│   ├── store/         # Storage interfaces
│   └── utils/         # Utilities
├── web/               # Frontend
├── migrations/        # Database migrations
└── configs/           # Configuration files
```

## Tech Stack

- **Backend**: Go 1.21+
- **Database**: PostgreSQL 14+
- **Cache**: Redis 7+
- **WebSocket**: gorilla/websocket
- **Frontend**: CodeMirror 6 / Monaco Editor
- **Deployment**: Docker + Kubernetes

## Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 14+
- Redis 7+
- Node.js 18+ (for frontend)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/UmangDiyora/markdown-collab.git
cd markdown-collab
```

2. Install dependencies:
```bash
go mod download
```

3. Configure the application:
```bash
cp configs/config.example.yaml configs/config.yaml
# Edit configs/config.yaml with your settings
```

4. Run database migrations:
```bash
go run cmd/migrate/main.go up
```

5. Start the server:
```bash
go run cmd/server/main.go
```

## Configuration

See `configs/config.example.yaml` for all configuration options.

## Development

### Running Tests
```bash
go test ./...
```

### Running with Docker
```bash
docker-compose up
```

## Performance Targets

- Support 100+ concurrent users per document
- Sub-100ms sync latency
- 99.9% uptime
- <2s document load time
- Handle 1000+ operations per second

## Security

- OWASP Top 10 compliance
- GDPR compliance
- TLS/WSS encryption
- Rate limiting
- Input sanitization

## Contributing

Contributions are welcome! Please read CONTRIBUTING.md for details.

## License

MIT License - see LICENSE file for details

## Roadmap

- [x] Project setup
- [ ] Authentication system
- [ ] WebSocket infrastructure
- [ ] OT/CRDT implementation
- [ ] Document management
- [ ] Real-time sync
- [ ] Frontend interface
- [ ] Export/Import
- [ ] Deployment setup
- [ ] Monitoring & analytics

## Support

For issues and questions, please open a GitHub issue.
