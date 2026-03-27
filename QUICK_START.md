# Quick Start - Backend & Frontend Integration

## 🚀 Prerequisites

- PostgreSQL installed and running
- Go 1.21+
- Node.js 18+
- Terminal/Command prompt

## ⚡ Quick Setup (5 minutes)

### Step 1: PostgreSQL Database Setup

```bash
# Open PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE notification_db;
\c notification_db

# Copy all SQL from notifications-service/migrations/001_init.sql and paste it here
# Then exit
\q
```

### Step 2: Start Backend Server

```bash
cd notification-service
go mod download
go run cmd/server/main.go
```

**Expected output:**
```
2026/03/26 12:00:00 Starting Probus Notification System...
2026/03/26 12:00:00 ✓ Connected to PostgreSQL
2026/03/26 12:00:00 Starting HTTP server on :8080
```

Press Ctrl+C to stop.

### Step 3: Start Frontend Server  

**In a new terminal:**

```bash
cd notification-admin-portal
npm install  # Only first time
ng serve --port 4302
```

**Expected output:**
```
✔ Application bundle generation complete. [2.152 seconds]
Watch mode enabled. Watching for file changes...
➜  Local:   http://localhost:4302/
```

### Step 4: Open in Browser

- Frontend: http://localhost:4302
- Backend Health: http://localhost:8080/health

---

## ✅ Verify Integration is Working

### Test Backend API

```bash
# Test health check
curl http://localhost:8080/health

# Get languages from database
curl http://localhost:8080/languages

# Create a language
curl -X POST http://localhost:8080/languages \
  -H "Content-Type: application/json" \
  -d '{"name":"Portuguese","code":"PT","status":"Active"}'
```

### Test Frontend UI

1. Open http://localhost:4302 in browser
2. Navigate to Languages section
3. You should see 3 default languages (English, Hindi, Spanish)
4. Try:
   - **Add** a new language
   - **Edit** an existing language
   - **Delete** a language (will remove it - soft delete)

---

## 🛠️ Troubleshooting

| Issue | Solution |
|-------|----------|
| `connection refused` on :8080 | Backend not running. Check Step 2 |
| `Cannot GET /` on 4302 | Frontend not running. Check Step 3 |
| `CORS error` in browser | Already enabled. Check backend is running |
| Database connection failed | Check PostgreSQL running, credentials in env file |
| No data displaying | Check migration ran successfully in Step 1 |

---

## 📁 Important Files

- **Backend**: `notification-service/cmd/server/main.go` - Entry point
- **Frontend**: `notification-admin-portal/src/app/core/config/api.config.ts` - API configuration  
- **Migrations**: `notification-service/migrations/001_init.sql` - Database schema
- **Environment**: `notification-service/.env.example` - Configuration template

---

## 🔄 Data Flow

```
Browser (http://localhost:4302)
    ↓
Angular Service (src/app/services/language.service.ts)
    ↓
HTTP Call to http://localhost:8080/languages
    ↓
Go Handler (internal/interfaces/http/handlers_basic.go)
    ↓
Database Query (PostgreSQL)
    ↓
Return JSON → Display in UI
```

---

## 📚 Next Steps

1. **Explore the UI**: Try all CRUD operations
2. **Check database**: `SELECT * FROM languages;` in PostgreSQL
3. **Review code**: Understand the full architecture
4. **Add new entity**: Follow patterns in existing entities
5. **Deploy**: See `INTEGRATION_GUIDE.md` for production setup

---

## 💡 Tips

- Both servers run on localhost, perfect for development
- Changes auto-reload in Angular (ng serve watch mode)
- SQL changes require database update and potentially Go restart
- Check browser console (F12) for frontend errors
- Check terminal for backend errors

---

**You're all set! 🎉 Your notification system is ready for development.**
