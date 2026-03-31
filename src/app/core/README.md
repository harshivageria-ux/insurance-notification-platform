# Enterprise API Integration Layer - Core Module

**Status:** ✅ Production-Ready | **Version:** 1.0.0 | **Last Updated:** March 26, 2026

This is the core infrastructure layer for the **Notification Admin Portal** - an enterprise-grade Angular application with seamless mock-to-API mode switching.

## 📁 Directory Structure

```
src/app/core/
├── config/
│   ├── api.config.ts         ← API endpoints & base URL configuration
│   └── index.ts              ← Barrel export
├── models/
│   ├── language.model.ts     ← TypeScript interfaces & DTOs
│   └── index.ts              ← Barrel export
├── services/
│   ├── error-handler.service.ts  ← Centralized error handling
│   ├── loader.service.ts         ← Global loading state management
│   └── index.ts                  ← Barrel export
├── ARCHITECTURE.md               ← Comprehensive architecture guide
├── QUICK_REFERENCE.md            ← One-page quick start
├── MODULE_CHECKLIST.md           ← Checklist for adding new modules
├── IMPLEMENTATION_SUMMARY.md     ← Complete implementation overview
├── PRODUCTION_SETUP.md           ← Environment-based configuration
└── README.md                     ← This file
```

## 🎯 Quick Links

- **Getting Started?** → Read [QUICK_REFERENCE.md](QUICK_REFERENCE.md)
- **Understanding Architecture?** → Read [ARCHITECTURE.md](ARCHITECTURE.md)
- **Adding New Modules?** → Follow [MODULE_CHECKLIST.md](MODULE_CHECKLIST.md)
- **Setup for Production?** → Read [PRODUCTION_SETUP.md](PRODUCTION_SETUP.md)
- **Full Overview?** → See [IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md)

## ✨ Key Features

### 🔄 Dual Mode Operation
- **Mock Mode:** Uses in-memory data (default) - perfect for development
- **API Mode:** Uses HttpClient to call real backend APIs

### 🎯 Auto Mode Switching
```typescript
// Automatically switches based on configuration
// If API_CONFIG.BASE_URL is empty → Mock mode
// If API_CONFIG.BASE_URL is set  → API mode
```

### 🛡️ Type Safety
- Full TypeScript support with interfaces
- No 'any' types
- IDE autocomplete and compile-time error detection

### 🚨 Error Handling
- Centralized error handling service
- User-friendly error messages
- HTTP status code classification
- Network error detection
- Logging ready for Sentry integration

### ⏳ Loading State Management
- Global loader service with BehaviorSubject
- Observable stream: `loaderService.isLoading$`
- Counter mechanism for concurrent operations
- Template-friendly reactive pattern

## 🚀 Switching to API Mode (3 Steps)

### Step 1: Update Base URL
```typescript
// File: src/app/core/config/api.config.ts

export const API_CONFIG = {
  BASE_URL: 'https://your-api.com/api/v1',  // ← Your backend API
  // ... rest of config
};
```

### Step 2: Verify Response Format
Your backend must return:
```json
{
  "success": true,
  "data": { /* Language object or array */ }
}
```

### Step 3: Test
The service automatically switches to API mode. Your existing UI continues to work!

## 📦 Core Services

### ErrorHandlerService
Centralized error handling with user-friendly messages.

```typescript
import { ErrorHandlerService } from '@app/core/services';

constructor(private errorHandler: ErrorHandlerService) {}

this.service.getLanguages().subscribe({
  next: (data) => { /* success */ },
  error: (appError) => {
    if (this.errorHandler.isNetworkError(appError)) {
      // Handle network error
    }
  }
});
```

### LoaderService
Global loading state for showing/hiding spinners.

```typescript
import { LoaderService } from '@app/core/services';

constructor(public loaderService: LoaderService) {}

// In component
ngOnInit() {
  this.service.getLanguages().subscribe({...});
}

// In template
<div *ngIf="loaderService.isLoading$ | async" class="spinner">
  Loading...
</div>
```

## 🔌 API Integration Points

All CRUD operations with both mock and API support:

| Operation | Method | Mock | API |
|-----------|--------|------|-----|
| Read All  | `getLanguages()` | ✓ | GET /languages |
| Create    | `addLanguage(request)` | ✓ | POST /languages |
| Update    | `updateLanguage(request)` | ✓ | PUT /languages/:id |
| Delete    | `deleteLanguage(id)` | ✓ | DELETE /languages/:id |

## 📚 Models & Interfaces

### Language
```typescript
export interface Language {
  id: number;
  name: string;
  code: string;
  status: 'Active' | 'Inactive';
  createdAt?: Date;
  updatedAt?: Date;
}
```

### Request DTOs
```typescript
export interface CreateLanguageRequest {
  name: string;
  code: string;
  status: 'Active' | 'Inactive';
}

export interface UpdateLanguageRequest extends CreateLanguageRequest {
  id: number;
}
```

### API Response Wrapper
```typescript
export interface ApiResponse<T> {
  success: boolean;
  data?: T;
  message?: string;
  errors?: string[];
}
```

## 🎨 Component Integration

```typescript
import { LanguageService } from '@app/services/language.service';
import { Language, CreateLanguageRequest } from '@app/core/models';
import { LoaderService } from '@app/core/services';

@Component({...})
export class LanguageComponent {
  languages: Language[] = [];

  constructor(
    private languageService: LanguageService,
    public loaderService: LoaderService
  ) {}

  ngOnInit() {
    this.languageService.getLanguages().subscribe({
      next: (data: Language[]) => {
        this.languages = data;
      },
      error: (error) => console.error(error)
    });
  }

  addLanguage(request: CreateLanguageRequest) {
    this.languageService.addLanguage(request).subscribe({
      next: () => this.loadLanguages(),
      error: (error) => console.error(error)
    });
  }
}
```

## 🔐 Security Considerations

- ✓ CORS: Configure on backend for frontend domain
- ✓ HTTPS: Always use in production
- ✓ API Keys: Use environment variables or secrets management
- ✓ Sensitive Data: Never log passwords or tokens
- ✓ Authentication: Add JWT/Bearer interceptor when needed

## 🧪 Testing

### Mock Mode (Default)
```typescript
// Mock mode is enabled by default
// No network calls, instant responses
// Perfect for development and testing
```

### API Mode
```typescript
// For testing with API
service.setMockMode(false);
service.getLanguages().subscribe(console.log);

// Check current mode
if (service.isMockMode()) {
  console.log('Running in mock mode');
}
```

## 📊 Architecture Flow

```
Component
   ↓
LanguageService
   ├─ Check: useMock?
   ├─ True  → Mock Methods (of(), immediate return)
   └─ False → API Methods (HttpClient, network call)
       ↓
   ErrorHandlerService (on error)
       ↓
   LoaderService (show/hide)
       ↓
   Observable<Language[]>
       ↓
Component (renders)
```

## ✅ Verification Checklist

- ✓ TypeScript compiles without errors
- ✓ All imports resolve correctly
- ✓ Mock data works (preserved from original)
- ✓ Component renders successfully
- ✓ Loading spinner shows/hides correctly
- ✓ Error handling catches exceptions
- ✓ No console warnings
- ✓ Application runs in development

## 🚀 Production Deployment

See [PRODUCTION_SETUP.md](PRODUCTION_SETUP.md) for:
- Environment-based configuration
- Build commands for different environments
- CI/CD pipeline examples
- Security best practices
- Deployment checklist

## 📖 Documentation Map

| Document | Purpose | Read When |
|----------|---------|-----------|
| [QUICK_REFERENCE.md](QUICK_REFERENCE.md) | 1-page quick start | Starting integration |
| [ARCHITECTURE.md](ARCHITECTURE.md) | Deep dive architecture | Understanding system |
| [MODULE_CHECKLIST.md](MODULE_CHECKLIST.md) | Guide for new modules | Adding features |
| [IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md) | Complete overview | Getting oriented |
| [PRODUCTION_SETUP.md](PRODUCTION_SETUP.md) | Production config | Deploying to production |

## 🎓 Next Steps

1. **Read** [QUICK_REFERENCE.md](QUICK_REFERENCE.md) (5 min)
2. **Review** [ARCHITECTURE.md](ARCHITECTURE.md) (15 min)
3. **Test** mock mode by running the app
4. **Configure** backend API base URL when ready
5. **Test** API mode against your backend
6. **Deploy** to production following [PRODUCTION_SETUP.md](PRODUCTION_SETUP.md)

## 🤝 Contributing

When adding new features:
1. Follow [MODULE_CHECKLIST.md](MODULE_CHECKLIST.md)
2. Maintain consistency with existing patterns
3. Add TypeScript interfaces in `core/models/`
4. Use centralized services
5. Document your changes

## 📞 Support & Questions

For questions about:
- **Architecture** → See [ARCHITECTURE.md](ARCHITECTURE.md)
- **Quick Start** → See [QUICK_REFERENCE.md](QUICK_REFERENCE.md)
- **Adding Features** → See [MODULE_CHECKLIST.md](MODULE_CHECKLIST.md)
- **Production** → See [PRODUCTION_SETUP.md](PRODUCTION_SETUP.md)

## 📋 Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | Mar 26, 2026 | Initial implementation |

## 📄 License

Part of the Insurance Notification Platform enterprise application.

---

**Ready to integrate APIs?** Start with [QUICK_REFERENCE.md](QUICK_REFERENCE.md) 🚀
