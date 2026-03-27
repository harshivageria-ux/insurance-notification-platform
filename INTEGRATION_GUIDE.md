# Backend-Frontend Integration Guide

## System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│           Angular Admin Portal (Frontend)                   │
│  - Port: 4302                                               │
│  - Framework: Angular 21.2.0 (Standalone Components)        │
│  - State Management: RxJS + TypeScript                      │
└──────────────────────────┬──────────────────────────────────┘
                           │ HTTP REST API
                           │ (CORS Enabled)
                           ↓
┌─────────────────────────────────────────────────────────────┐
│         Go Notification Service (Backend)                   │
│  - Port: 8080                                               │
│  - Framework: Chi v5 Router                                 │
│  - Database: PostgreSQL (Soft Delete Pattern)               │
│  - Architecture: Clean Architecture (Domain → Handlers)     │
└─────────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────────┐
│              PostgreSQL Database                            │
│  - 11 Tables (Languages, Priorities, Channels, etc.)        │
│  - Soft Delete Pattern (deleted_at timestamp)               │
│  - Timestamps: created_at, updated_at, deleted_at           │
└─────────────────────────────────────────────────────────────┘
```

## Prerequisites

### Backend Requirements
- Go 1.21+
- PostgreSQL 13+
- Git

### Frontend Requirements
- Node.js 18+
- npm or yarn
- Angular CLI 17+

## Backend Setup

### Step 1: Initialize Go Module

```bash
cd notification-service
go mod download
go mod tidy
```

### Step 2: Create PostgreSQL Database

```sql
-- Create database
CREATE DATABASE notification_db;

-- Connect to database
\c notification_db

-- Run migrations (copy content from migrations/001_init.sql)
-- Or use a migration tool like golang-migrate
```

### Step 3: Create Environment File

```bash
# Copy example env file
cp .env.example .env

# Edit .env with your database credentials
```

### Step 4: Run Backend Server

```bash
cd cmd/server
go run main.go

# Or build and run binary
go build -o notification-service main.go
./notification-service
```

Expected output:
```
2026/03/26 12:00:00 Starting Probus Notification System...
2026/03/26 12:00:00 ✓ Connected to PostgreSQL
2026/03/26 12:00:00 Starting HTTP server on :8080
```

## Frontend Setup

### Step 1: Install Dependencies

```bash
cd notification-admin-portal
npm install
```

### Step 2: Verify Backend Configuration

The frontend is already configured to use the backend API. Check [src/app/core/config/api.config.ts](src/app/core/config/api.config.ts):

```typescript
BASE_URL: 'http://localhost:8080'  // Points to your Go backend
```

### Step 3: Start Frontend Development Server

```bash
ng serve --port 4302

# Or use npm
npm start
```

Expected output:
```
✔ Application bundle generation complete. [2.152 seconds]
Watch mode enabled. Watching for file changes...
➜  Local:   http://localhost:4302/
```

## API Endpoints

### Languages
```
GET    /languages           - List all languages
POST   /languages           - Create language
PUT    /languages/{id}      - Update language
DELETE /languages/{id}      - Deactivate language (soft delete)
```

### Priorities
```
GET    /priorities          - List all priorities
POST   /priorities          - Create priority
PUT    /priorities/{id}     - Update priority
DELETE /priorities/{id}     - Deactivate priority
```

### Statuses
```
GET    /statuses            - List all statuses
POST   /statuses            - Create status
PUT    /statuses/{id}       - Update status
DELETE /statuses/{id}       - Deactivate status
```

### Schedule Types
```
GET    /schedule-types      - List all schedule types
POST   /schedule-types      - Create schedule type
PUT    /schedule-types/{id} - Update schedule type
DELETE /schedule-types/{id} - Deactivate schedule type
```

### Categories
```
GET    /categories          - List all categories
POST   /categories          - Create category
PUT    /categories/{id}     - Update category
DELETE /categories/{id}     - Deactivate category
```

### Channels
```
GET    /channels            - List all channels
POST   /channels            - Create channel
PUT    /channels/{id}       - Update channel
PATCH  /channels/{id}/toggle - Toggle channel active status
DELETE /channels/{id}       - Deactivate channel
```

### Channel Providers
```
GET    /channel-providers             - List all providers
POST   /channel-providers             - Create provider
PUT    /channel-providers/{id}        - Update provider
PATCH  /channel-providers/{id}/toggle - Toggle provider active status
DELETE /channel-providers/{id}        - Deactivate provider
```

### Provider Settings
```
GET    /provider-settings/{provider_id} - List settings for provider
POST   /provider-settings              - Create setting
PUT    /provider-settings/{id}         - Update setting
DELETE /provider-settings/{id}         - Deactivate setting
```

### Template Groups
```
GET    /template-groups     - List all template groups
POST   /template-groups     - Create template group
PUT    /template-groups/{id} - Update template group
DELETE /template-groups/{id} - Deactivate template group
```

### Templates
```
GET    /templates           - List all templates
POST   /templates           - Create template
PUT    /templates/{id}      - Update template
POST   /templates/{id}/preview - Preview template rendering
DELETE /templates/{id}      - Deactivate template
```

### Routing Rules
```
GET    /routing-rules             - List all routing rules
POST   /routing-rules             - Create routing rule
PUT    /routing-rules/{id}        - Update routing rule
PATCH  /routing-rules/{id}/toggle - Toggle rule active status
DELETE /routing-rules/{id}        - Deactivate routing rule
```

## Testing the Integration

### Manual Testing with cURL

```bash
# Health check
curl http://localhost:8080/health

# List languages
curl http://localhost:8080/languages

# Create language
curl -X POST http://localhost:8080/languages \
  -H "Content-Type: application/json" \
  -d '{"name":"French","code":"FR","status":"Active"}'

# Update language
curl -X PUT http://localhost:8080/languages/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"French","code":"FR","status":"Active"}'

# Delete language (soft delete)
curl -X DELETE http://localhost:8080/languages/1
```

### Testing via Angular UI

1. Open browser to `http://localhost:4302`
2. Navigate to Languages module
3. Test CRUD operations:
   - **Create**: Click "+ Add Language" button
   - **Read**: Languages should display from database
   - **Update**: Click "Edit" button on any language
   - **Delete**: Click "Delete" button (soft delete)

## Database Schema

### Languages Table
```sql
CREATE TABLE languages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(10) NOT NULL UNIQUE,
    status VARCHAR(50) NOT NULL DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```

All other tables follow the same pattern with appropriate fields.

## Common Issues & Solutions

### Issue: Connection Refused on localhost:8080
**Solution**: Ensure Go backend is running
```bash
# Check if process is running
lsof -i :8080

# Kill if needed and restart
go run cmd/server/main.go
```

### Issue: CORS Errors in Browser Console
**Solution**: CORS is already enabled in server.go with:
```go
w.Header().Set("Access-Control-Allow-Origin", "*")
```

### Issue: Database Connection Failed
**Solution**: Verify PostgreSQL is running and credentials match .env file
```bash
# Test connection
psql -h localhost -U postgres -d notification_db
```

### Issue: Angular Service Not Finding API
**Solution**: Verify api.config.ts has correct BASE_URL
```typescript
BASE_URL: 'http://localhost:8080'  // Must match backend port
```

## Development Workflow

### Adding a New Entity

1. **Create Domain Model** (e.g., `internal/domain/my_entity/my_entity.go`)
   - Define struct with db tags
   - Define Create/UpdateRequest structs

2. **Create Repository** (e.g., `internal/infrastructure/repository/my_entity_repository.go`)
   - Implement CRUD operations
   - Use pgx for database queries

3. **Create Database Table** (update migrations/001_init.sql)
   - Add table definition
   - Add indexes for performance

4. **Add HTTP Handlers** (in `internal/interfaces/http/handlers_*.go`)
   - List, Create, Update, Delete methods
   - Proper error handling

5. **Register Routes** (in `internal/interfaces/http/server.go`)
   - Add route group to Routes() method
   - Inject repository

6. **Update Angular Models** (in `src/app/core/models/`)
   - Create interface for entity
   - Create service in `src/app/services/`
   - Use service in component

## Environment Variables

```bash
# Database
DATABASE_URL=postgres://user:password@localhost:5432/notification_db

# Server
HTTP_PORT=8080

# Security
ENCRYPTION_KEY=32-character-encryption-key!!!!

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:4302,http://localhost:4300

# Logging
LOG_LEVEL=debug
```

## Production Deployment

### Backend Considerations
- Set `NODE_ENV=production`
- Use strong encryption key
- Configure proper PostgreSQL backups
- Enable database connection pooling
- Use load balancer for multiple instances
- Configure proper CORS origins
- Enable request logging and monitoring

### Frontend Considerations
- Build for production: `ng build --configuration production`
- Set API_CONFIG.BASE_URL to production backend URL
- Enable service workers for offline support
- Configure CDN for static assets
- Enable gzip compression
- Set proper HTTP cache headers

## Monitoring & Logging

The backend includes basic logging. For production, consider:
- ELK Stack (Elasticsearch, Logstash, Kibana)
- Datadog
- New Relic
- Prometheus + Grafana

The frontend includes error logging via ErrorHandlerService. Consider:
- Sentry
- LogRocket
- Rollbar

---

**Last Updated**: March 2026
**Status**: Production Ready ✓
