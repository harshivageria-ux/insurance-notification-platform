# Frontend Debug Guide - Save Data to Backend

## 🔍 Step 1: Verify Angular App Compiled with New Logging

1. The `ng serve` terminal should show:
   ```
   ✔ Application bundle generation complete. [X.XXX seconds]
   ```
   - If you see **recompilation errors**, stop `ng serve` and run `npm install` again

2. **Browser**: Open Developer Tools (`F12`)
   - Go to **Console** tab
   - Look for initialization logs starting with 🔧:
     ```
     🔧 LanguageService initialized:
        API_CONFIG.BASE_URL: http://localhost:9000
        isMockMode: false
        Full API_CONFIG: {...}
     ```

### ⚠️ If you see:
- `API_CONFIG.BASE_URL: ''` or `undefined` → **frontend is in MOCK mode** (not calling backend)
- `isMockMode: true` → **confirm you need to fix environment configuration**

---

## 🧪 Step 2: Test Adding a Language via UI

### Actions:
1. In browser (`http://localhost:4302`):
   - Click **Add Language** button
   - Enter:
     - Name: `TestDebugLang`
     - Code: `TDL`
   - Click **Save**

2. **Open Network Tab** (`F12` → **Network**)
   - Look for a request to: `POST http://localhost:9000/languages`
   - Status should be: **201** (created)
   - Response body should contain the new language with `id` field

3. **Check Console** for:
   ```
   📡 addLanguage() called - Mode: API
   🌐 POST request to: http://localhost:9000/languages Body: {...}
   ✅ Language added: {id: 189, name: "TestDebugLang", ...}
   ✅ Language updated: {...}
   ```

### ✅ Expected Behavior:
- Language appears in the list immediately
- Console shows `📡 addLanguage() called - Mode: API`
- Network tab shows POST to backend with 201 response

### ❌ If NOT Working:
- Check Console for error messages
- Look in Network tab for the POST request
- If no POST request appears → frontend is in MOCK mode

---

## 🔴 Error Diagnosis

### Error 1: "No POST request in Network Tab"
**Problem**: Frontend is still using mock mode
**Solution**:
```bash
cd notification-admin-portal
# Kill ng serve (Ctrl+C)
npm install
ng serve --port 4302
```
- Check console for `isMockMode: false`

### Error 2: "POST fails with CORS error"
**Console shows**: `Access to XMLHttpRequest blocked by CORS policy`
**Solution**: Backend CORS is not set correctly
- Backend already has CORS enabled, but verify in  [server.go](notification-service/internal/interfaces/http/server.go#L190)

### Error 3: "POST fails with 500 error"
**Console shows**: `❌ Failed to create language: Error: ...`
**Solution**:
1. Check backend terminal for error message
2. Backend logs should show the exact database error
3. Common issues:
   - Duplicate name/code already in DB
   - Missing `created_by` field in request
   - Database connection issue

### Error 4: "POST succeeds but data doesn't persist"
**Console shows**: `✅ Language added: {...}` but data is gone after refresh
**Solution**:
1. Verify data in database:
   ```sql
   SELECT * FROM languages_master WHERE code = 'TDL' ORDER BY id DESC LIMIT 1;
   ```
2. If empty → backend didn't actually save (check backend logs)
3. If data exists → frontend is showing stale mock data (clear browser cache)

---

## 🛠️ How to Collect Debug Info for Support

If nothing works, collect this info:

### 1. Browser Console Output
```bash
# Right-click on Console output → Select All → Copy
# Paste into a text file
```

### 2. Network Request Details
```bash
# Right-click on the POST request in Network tab → Copy as cURL
# Share the cURL command
```

### 3. Backend Log Output
```bash
# From the terminal where you ran "go run cmd/server/main.go"
# Copy and share the last 50 lines of terminal output
```

### 4. Database State
```bash
# Connect to PostgreSQL and run:
SELECT COUNT(*) FROM languages_master;
SELECT * FROM languages_master ORDER BY id DESC LIMIT 5;
```

---

## ✅ Full End-to-End Test

Run this checklist:

- [ ] Backend running: `go run cmd/server/main.go` (port 9000)
- [ ] Frontend running: `ng serve --port 4302`
- [ ] Browser open: `http://localhost:4302`
- [ ] Console shows: ✅ `isMockMode: false`
- [ ] Try add language
- [ ] Network tab shows: ✅ `POST 201` to backend
- [ ] Console shows: ✅ `📡 addLanguage() called - Mode: API`
- [ ] Language appears in list
- [ ] Refresh page → language still there ✅ (confirms DB save)

---

## 📊 Expected Console Output (Working State)

```
🔧 LanguageService initialized:
   API_CONFIG.BASE_URL: http://localhost:9000
   isMockMode: false
   Full API_CONFIG: {...}

📡 getLanguages() called - Mode: API
🌐 GET request to: http://localhost:9000/languages
✅ API response received: [...]

📡 addLanguage() called - Mode: API
   {name: 'TestDebugLang', code: 'TDL', status: 'Active'}
🌐 POST request to: http://localhost:9000/languages Body: {...}
✅ API response received: {id: 189, name: 'TestDebugLang', ...}
✅ Language added: {id: 189, ...}
```

---

## 🚀 Quick Fix Commands

If frontend shows mock mode, force API mode:

```bash
# Terminal 1: Backend
cd notification-service
go run cmd/server/main.go

# Terminal 2: Frontend (fresh install)
cd notification-admin-portal
npm ci
npm audit fix --force
ng serve --port 4302 --poll 2000
```

The `--poll` flag helps with file watching on some systems.
