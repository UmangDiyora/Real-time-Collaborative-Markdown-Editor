# Real-time Collaborative Markdown Editor - Complete Implementation Blueprint

## Project Overview
Build a production-ready real-time collaborative markdown editor similar to HackMD/CodiMD but in Go. Multiple users can edit the same document simultaneously with real-time synchronization, conflict resolution, and live preview.

## Core Requirements
- Real-time multi-user editing
- Operational Transformation (OT) or CRDT for conflict resolution
- WebSocket-based synchronization
- User authentication and authorization
- Document persistence and versioning
- Live markdown preview
- Cursor position sharing
- User presence indicators
- Document sharing and permissions
- Export to multiple formats
- Syntax highlighting
- Auto-save functionality

## Architecture Components

### 1. Core Components
- **WebSocket Server**: Handles real-time connections
- **Sync Engine**: Manages document synchronization
- **OT/CRDT Engine**: Conflict resolution
- **Auth Service**: User authentication/authorization
- **Document Store**: Document persistence
- **Session Manager**: User session handling
- **Presence Service**: Track active users
- **Export Service**: Convert to various formats
- **API Server**: RESTful endpoints

### 2. Client Components
- **Editor Interface**: CodeMirror/Monaco integration
- **Sync Client**: WebSocket client for sync
- **Markdown Renderer**: Real-time preview
- **Collaboration UI**: User cursors and selections

## Detailed Implementation Steps

### Phase 1: Project Setup and Architecture

#### Step 1.1: Project Structure
```
markdown-collab/
├── cmd/
│   ├── server/          # Main server entry point
│   └── migrate/         # Database migration tool
├── internal/
│   ├── auth/            # Authentication logic
│   ├── document/        # Document management
│   ├── session/         # Session handling
│   ├── sync/            # Synchronization engine
│   ├── transform/       # OT/CRDT implementation
│   ├── websocket/       # WebSocket handling
│   ├── presence/        # User presence tracking
│   └── export/          # Document export
├── pkg/
│   ├── models/          # Data models
│   ├── protocol/        # Wire protocol definitions
│   ├── store/           # Storage interfaces
│   └── utils/           # Utility functions
├── web/
│   ├── static/          # Frontend assets
│   ├── templates/       # HTML templates
│   └── src/             # Frontend source (if separate)
├── migrations/          # Database migrations
├── configs/             # Configuration files
├── docker/              # Docker configurations
└── tests/              # Test suites
```

#### Step 1.2: Data Models

**User Model:**
- ID (UUID)
- Username (string)
- Email (string)
- PasswordHash (string)
- Avatar (string)
- CreatedAt (timestamp)
- LastLoginAt (timestamp)
- Settings (JSON)

**Document Model:**
- ID (UUID)
- Title (string)
- Content (text)
- OwnerID (User ID)
- Version (int)
- CreatedAt (timestamp)
- UpdatedAt (timestamp)
- LastEditedBy (User ID)
- Tags (array)
- IsPublic (boolean)
- ShareToken (string)

**Document Version:**
- ID (UUID)
- DocumentID (Document ID)
- Version (int)
- Content (text)
- AuthorID (User ID)
- Operations (JSON) - for OT
- CreatedAt (timestamp)

**Session Model:**
- ID (string)
- UserID (User ID)
- DocumentID (Document ID)
- CursorPosition (int)
- SelectionStart (int)
- SelectionEnd (int)
- Color (string)
- LastActivity (timestamp)
- ConnectionID (string)

**Permission Model:**
- ID (UUID)
- DocumentID (Document ID)
- UserID (User ID)
- Permission (enum: VIEW, EDIT, ADMIN)
- CreatedAt (timestamp)

### Phase 2: Authentication System

#### Step 2.1: User Registration/Login
- Email/password registration
- Password hashing with bcrypt
- Email verification
- Password reset functionality
- OAuth integration (Google, GitHub)

#### Step 2.2: JWT Implementation
- Access token (short-lived, 15 min)
- Refresh token (long-lived, 7 days)
- Token rotation on refresh
- Secure storage in httpOnly cookies
- CSRF protection

#### Step 2.3: Session Management
- Redis-based session store
- Session expiration handling
- Multiple device support
- Session invalidation on logout

### Phase 3: WebSocket Implementation

#### Step 3.1: Connection Management
- Upgrade HTTP to WebSocket
- Connection authentication
- Heartbeat/ping-pong
- Automatic reconnection
- Connection pooling

#### Step 3.2: Message Protocol
Define message types:
```
Operation Types:
- INSERT: Insert text at position
- DELETE: Delete text at position
- FORMAT: Apply formatting
- CURSOR: Update cursor position
- SELECT: Update selection
- PRESENCE: User presence update
- SYNC: Full document sync
- ACK: Operation acknowledgment
- ERROR: Error notification

Message Structure:
{
  "type": "operation_type",
  "documentId": "doc_uuid",
  "userId": "user_uuid",
  "timestamp": 1234567890,
  "version": 42,
  "payload": {
    // Operation-specific data
  }
}
```

#### Step 3.3: Room Management
- Document-based rooms
- User join/leave handling
- Broadcast to room members
- Room state management
- Presence tracking per room

### Phase 4: Operational Transformation Implementation

#### Step 4.1: OT Fundamentals
Implement core OT operations:
- Insert(position, text)
- Delete(position, length)
- Transform(op1, op2) - transform concurrent ops
- Compose(op1, op2) - combine sequential ops

#### Step 4.2: OT Algorithm
```
Transform Rules:
1. Insert vs Insert:
   - If same position, order by user ID
   - Adjust positions accordingly

2. Insert vs Delete:
   - Adjust delete position if after insert

3. Delete vs Delete:
   - Handle overlapping deletes
   - Adjust ranges appropriately

4. Maintain intention preservation
5. Ensure convergence
```

#### Step 4.3: Version Vector
- Track document version per client
- Server authoritative version
- Operation buffering for out-of-order ops
- Conflict resolution strategy

### Phase 5: CRDT Alternative Implementation

#### Step 5.1: CRDT Selection (Alternative to OT)
Implement Yjs-like CRDT:
- Tree-based CRDT structure
- Unique character IDs
- Causal ordering
- Tombstone for deletions
- Eventual consistency

#### Step 5.2: CRDT Operations
- Local operations immediately applied
- Propagate to other clients
- Merge concurrent operations
- Garbage collection for tombstones

### Phase 6: Document Management

#### Step 6.1: Document CRUD
- Create new document
- Load existing document
- Auto-save every 5 seconds
- Manual save trigger
- Delete with soft-delete option

#### Step 6.2: Versioning System
- Save snapshots periodically
- Store operations between snapshots
- Version history browsing
- Diff view between versions
- Restore to previous version

#### Step 6.3: Document Sharing
- Generate shareable links
- Public/private documents
- Read-only sharing option
- Collaboration invites
- Permission management

### Phase 7: Real-time Synchronization

#### Step 7.1: Client Sync
- Initial document load
- Apply local changes optimistically
- Send operations to server
- Receive and transform remote ops
- Handle conflicts gracefully

#### Step 7.2: Server Sync
- Receive client operations
- Transform against concurrent ops
- Broadcast to other clients
- Persist to database
- Handle disconnections

#### Step 7.3: Sync Optimization
- Batch operations
- Throttle/debounce updates
- Delta compression
- Binary protocol for efficiency

### Phase 8: Presence and Awareness

#### Step 8.1: User Presence
- Show online users
- User avatar/color
- "User is typing" indicator
- Last seen timestamp
- Connection status

#### Step 8.2: Cursor Tracking
- Broadcast cursor positions
- Show remote cursors
- Smooth cursor animations
- Handle cursor in deleted text

#### Step 8.3: Selection Sharing
- Share text selections
- Highlight other users' selections
- Different colors per user
- Selection labels

### Phase 9: Frontend Implementation

#### Step 9.1: Editor Integration
- CodeMirror 6 or Monaco Editor
- Custom key bindings
- Syntax highlighting
- Auto-completion
- Find and replace

#### Step 9.2: Markdown Renderer
- Real-time preview pane
- Syntax highlighting in code blocks
- Math formula rendering (KaTeX)
- Mermaid diagram support
- Table of contents generation

#### Step 9.3: UI Components
- Split-pane editor/preview
- Toolbar with formatting buttons
- File tree for multiple documents
- Search functionality
- Theme switching (dark/light)

### Phase 10: Storage Layer

#### Step 10.1: PostgreSQL Schema
```sql
Tables:
- users
- documents
- document_versions
- operations
- permissions
- sessions
- share_links

Indexes:
- documents(owner_id, updated_at)
- operations(document_id, version)
- permissions(user_id, document_id)
```

#### Step 10.2: Redis Usage
- Session storage
- Presence data
- Operation buffer
- Document cache
- Rate limiting

#### Step 10.3: File Storage
- Store document attachments
- Image uploads
- Local filesystem or S3
- CDN integration

### Phase 11: Export and Import

#### Step 11.1: Export Formats
- Markdown (.md)
- HTML with styling
- PDF via headless Chrome
- DOCX via pandoc
- LaTeX
- Plain text

#### Step 11.2: Import Support
- Import markdown files
- Parse and convert DOCX
- Maintain formatting
- Handle images/attachments

### Phase 12: Performance Optimization

#### Step 12.1: Backend Optimization
- Connection pooling
- Query optimization
- Caching strategy
- Load balancing
- Horizontal scaling

#### Step 12.2: Frontend Optimization
- Virtual scrolling for long documents
- Lazy loading
- WebSocket reconnection strategy
- Local storage for offline
- Service worker for caching

#### Step 12.3: Sync Optimization
- Operational transform caching
- Batch operation processing
- Compression (gzip/brotli)
- CDN for static assets

### Phase 13: Security Implementation

#### Step 13.1: Authentication Security
- Rate limiting on login
- Account lockout mechanism
- Two-factor authentication
- Session hijacking prevention
- XSS protection

#### Step 13.2: Authorization
- Document-level permissions
- Role-based access control
- API rate limiting
- Input sanitization
- SQL injection prevention

#### Step 13.3: Data Security
- Encryption at rest
- TLS for all connections
- Secure WebSocket (WSS)
- CORS configuration
- Content Security Policy

### Phase 14: Testing Strategy

#### Step 14.1: Unit Tests
- OT algorithm tests
- CRDT convergence tests
- Auth flow tests
- Permission tests

#### Step 14.2: Integration Tests
- WebSocket connection tests
- Multi-user sync tests
- Database operation tests
- API endpoint tests

#### Step 14.3: End-to-End Tests
- User registration/login flow
- Document creation/editing
- Real-time collaboration
- Export functionality

#### Step 14.4: Performance Tests
- Concurrent user limits
- Large document handling
- Operation throughput
- Memory usage under load

### Phase 15: Deployment

#### Step 15.1: Docker Setup
```dockerfile
# Multi-stage build
# Build stage for Go binary
# Runtime stage with minimal image
# Health checks
# Environment configuration
```

#### Step 15.2: Kubernetes Deployment
- Deployment for main server
- StatefulSet for WebSocket servers
- Redis deployment
- PostgreSQL with persistent volume
- Ingress configuration
- Horizontal pod autoscaler

#### Step 15.3: CI/CD Pipeline
- GitHub Actions workflow
- Automated testing
- Docker image building
- Deployment to staging
- Production deployment

### Phase 16: Monitoring and Analytics

#### Step 16.1: Application Metrics
- Active users count
- Document count
- Operations per second
- Sync latency
- Error rates

#### Step 16.2: Infrastructure Monitoring
- CPU/Memory usage
- WebSocket connections
- Database performance
- Redis metrics
- Network I/O

#### Step 16.3: User Analytics
- User activity tracking
- Document engagement
- Feature usage
- Performance metrics
- Error tracking (Sentry)

## Configuration Schema

```yaml
server:
  host: "0.0.0.0"
  port: 8080
  websocket_port: 8081

database:
  postgres:
    host: "localhost"
    port: 5432
    database: "markdown_collab"
    username: "user"
    password: "pass"
    max_connections: 100

redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0

auth:
  jwt_secret: "secret-key"
  access_token_duration: "15m"
  refresh_token_duration: "168h"
  
sync:
  operation_buffer_size: 1000
  sync_interval: "100ms"
  snapshot_threshold: 100

storage:
  type: "s3" # local, s3
  s3:
    bucket: "markdown-docs"
    region: "us-east-1"
    
limits:
  max_document_size: "10MB"
  max_concurrent_connections: 1000
  rate_limit_requests: 100
  rate_limit_duration: "1m"
```

## Performance Targets
- Support 100+ concurrent users per document
- Sub-100ms sync latency
- 99.9% uptime
- <2s document load time
- Support documents up to 10MB
- Handle 1000+ operations per second

## Security Requirements
- OWASP Top 10 compliance
- GDPR compliance for user data
- Regular security audits
- Penetration testing
- Vulnerability scanning
- Security headers implementation

## Feature Roadmap
1. MVP: Basic collaborative editing
2. Version history and rollback
3. Comments and annotations
4. Plugins and extensions
5. Mobile application
6. Offline mode with sync
7. AI-powered writing assistance
8. Team workspaces

This blueprint provides a complete implementation guide for building a production-ready real-time collaborative markdown editor in Go.
