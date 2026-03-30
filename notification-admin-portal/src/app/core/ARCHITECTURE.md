/**
 * ENTERPRISE API INTEGRATION ARCHITECTURE
 * =======================================
 * 
 * This document outlines the enterprise-level API integration setup for the
 * Notification Admin Portal. The architecture is designed to support seamless
 * transition from mock data to real backend APIs.
 * 
 * 📁 FOLDER STRUCTURE
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * src/app/core/
 * ├── config/
 * │   ├── api.config.ts          ← API endpoints & configuration
 * │   └── index.ts               ← Barrel export
 * ├── models/
 * │   ├── language.model.ts      ← TypeScript interfaces (Language, Request/Response)
 * │   └── index.ts               ← Barrel export
 * └── services/
 *     ├── error-handler.service.ts   ← Centralized error handling
 *     ├── loader.service.ts          ← Global loading state management
 *     └── index.ts                   ← Barrel export
 * 
 * src/app/services/
 * └── language.service.ts        ← Business logic (mock + API modes)
 * 
 * src/app/modules/languages/
 * └── languages.component.ts     ← UI logic (properly typed)
 * 
 * 
 * 🎯 KEY PRINCIPLES
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * 1. SEPARATION OF CONCERNS
 *    - Core services: Error handling, loading state, communication
 *    - Business services: Language service with mock + API logic
 *    - Components: UI logic, user interactions
 *    - Models: TypeScript interfaces for type safety
 *    - Config: Centralized API configuration
 * 
 * 2. TYPE SAFETY
 *    - All data uses TypeScript interfaces from core/models/
 *    - No 'any' types in production code
 *    - Better IDE autocomplete and refactoring support
 * 
 * 3. MOCK MODE VS API MODE
 *    - Mock mode: useMock = true (uses in-memory data)
 *    - API mode: useMock = false (uses HttpClient)
 *    - Automatic switching based on API_CONFIG.BASE_URL
 *    - Fallback to mock data on API errors
 * 
 * 4. ERROR HANDLING
 *    - Centralized error handler with user-friendly messages
 *    - Proper HTTP status code handling
 *    - Logging for debugging in development
 *    - Error propagation with typed AppError
 * 
 * 5. LOADING STATE
 *    - Global loader service with counter mechanism
 *    - Handles multiple concurrent async operations
 *    - Observable stream for UI subscription
 *    - Prevents UI glitches with proper show/hide timing
 * 
 * 
 * 🚀 QUICK START: SWITCHING TO API MODE
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * Step 1: Configure API Base URL
 * ────────────────────────────────────────────────────────────────────────────
 * 
 * File: src/app/core/config/api.config.ts
 * 
 * BEFORE:
 *   export const API_CONFIG = {
 *     BASE_URL: '',  // ← Empty for mock mode
 *     ...
 *   }
 * 
 * AFTER:
 *   export const API_CONFIG = {
 *     BASE_URL: 'https://api.notification.com/api/v1',  // ← Your backend URL
 *     ...
 *   }
 * 
 * The LanguageService will automatically switch to API mode!
 * 
 * 
 * Step 2: Backend API Contract
 * ────────────────────────────────────────────────────────────────────────────
 * 
 * Your backend should implement these endpoints:
 * 
 * 1. GET /languages
 *    Response: { success: true, data: Language[] }
 *    
 * 2. POST /languages
 *    Request: { name, code, status }
 *    Response: { success: true, data: Language }
 *    Maps to: sp_languages_insert
 *    
 * 3. PUT /languages/:id
 *    Request: { name, code, status }
 *    Response: { success: true, data: Language }
 *    Maps to: sp_languages_update
 *    
 * 4. DELETE /languages/:id
 *    Response: { success: true, data: boolean }
 *    Maps to: sp_languages_deactivate (soft delete)
 * 
 * 
 * Step 3: Response Format
 * ────────────────────────────────────────────────────────────────────────────
 * 
 * The backend must return responses in this format:
 * 
 * Success:
 *   {
 *     "success": true,
 *     "data": { /* Language object or array */ }
 *   }
 * 
 * Error:
 *   {
 *     "success": false,
 *     "message": "Error description",
 *     "errors": ["Detailed error 1", "Detailed error 2"]
 *   }
 * 
 * 
 * 📌 USAGE IN COMPONENTS
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * Example: Using LanguageService in a component
 * ────────────────────────────────────────────────────────────────────────────
 * 
 * import { Component } from '@angular/core';
 * import { LanguageService } from '../../services/language.service';
 * import { Language } from '../../core/models/language.model';
 * import { LoaderService } from '../../core/services/loader.service';
 * 
 * @Component({
 *   selector: 'app-language-list',
 *   template: `
 *     <div *ngIf="loaderService.isLoading$ | async" class="loader">
 *       Loading...
 *     </div>
 *     <ul>
 *       <li *ngFor="let lang of languages">{{ lang.name }}</li>
 *     </ul>
 *   `
 * })
 * export class LanguageListComponent {
 *   languages: Language[] = [];
 * 
 *   constructor(
 *     private languageService: LanguageService,
 *     public loaderService: LoaderService
 *   ) {}
 * 
 *   ngOnInit() {
 *     this.languageService.getLanguages().subscribe({
 *       next: (languages: Language[]) => {
 *         this.languages = languages;
 *       },
 *       error: (error) => {
 *         console.error('Failed to load languages:', error);
 *       }
 *     });
 *   }
 * }
 * 
 * 
 * 🔧 ADVANCED TOPICS
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * 1. Manual Mode Switching (for testing/debugging)
 * ────────────────────────────────────────────────────────────────────────────
 * 
 * // In any component or service
 * this.languageService.setMockMode(true);   // Use mock data
 * this.languageService.setMockMode(false);  // Use API
 * 
 * // Check current mode
 * if (this.languageService.isMockMode()) {
 *   console.log('Running in mock mode');
 * }
 * 
 * 
 * 2. Error Handling Pattern
 * ────────────────────────────────────────────────────────────────────────────
 * 
 * import { ErrorHandlerService, AppError } from '../../core/services';
 * 
 * this.languageService.getLanguages().subscribe({
 *   next: (languages) => { /* success */ },
 *   error: (appError: AppError) => {
 *     if (this.errorHandler.isNetworkError(appError)) {
 *       this.showNetworkErrorMsg = true;
 *     } else if (this.errorHandler.isErrorCode(appError, 'ERROR_403')) {
 *       this.showPermissionErrorMsg = true;
 *     }
 *   }
 * });
 * 
 * 
 * 3. Loading State Pattern
 * ────────────────────────────────────────────────────────────────────────────
 * 
 * // In template
 * <div *ngIf="loaderService.isLoading$ | async" class="spinner">
 *   <p>Processing...</p>
 * </div>
 * 
 * // The loader service automatically manages the loading state
 * // via show() and hide() calls in the language service
 * 
 * 
 * 💾 ADDING NEW API MODULES (e.g., Users, Tenants)
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * Follow the same pattern for any new API module:
 * 
 * 1. Create model in src/app/core/models/user.model.ts
 * 2. Add endpoints to src/app/core/config/api.config.ts
 * 3. Create service in src/app/services/user.service.ts
 * 4. Import types and use in components
 * 
 * 
 * 🧪 TESTING WITH MOCK DATA
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * Mock data is preserved and can be used for:
 * - Local development without backend
 * - Unit testing components
 * - E2E testing with deterministic data
 * - Fallback when API is unavailable
 * 
 * To use mock data in tests:
 * 
 * beforeEach(() => {
 *   languageService.setMockMode(true);
 * });
 * 
 * 
 * 🔐 SECURITY CONSIDERATIONS
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * 1. CORS: Configure CORS on your backend for the frontend URL
 * 2. Authentication: Add JWT/Bearer token interceptor if needed
 * 3. HTTPS: Always use HTTPS in production
 * 4. Sensitive Data: Don't log passwords or tokens to console
 * 5. API Keys: Store in environment files, never in code
 * 
 * 
 * 📝 ENVIRONMENT-BASED CONFIGURATION
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * For production-ready apps, use Angular environment files:
 * 
 * src/environments/environment.ts (development)
 * src/environments/environment.prod.ts (production)
 * 
 * Then in api.config.ts:
 * 
 * import { environment } from '../../environments/environment';
 * 
 * export const API_CONFIG = {
 *   BASE_URL: environment.apiBaseUrl,
 *   ...
 * };
 * 
 * 
 * 📚 SERVICE FILE MAPPING
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * File                                     Purpose
 * ────────────────────────────────────────────────────────────────────────────
 * src/app/core/config/api.config.ts        API endpoints & base URL
 * src/app/core/models/language.model.ts    TypeScript interfaces
 * src/app/core/services/error-handler.ts   Error handling & logging
 * src/app/core/services/loader.service.ts  Loading state management
 * src/app/services/language.service.ts     Business logic (mock + API)
 * src/app/modules/languages/...            UI components (standalone)
 * 
 * 
 * ⚠️ COMMON PITFALLS
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * 1. Forgetting to update api.config.ts with backend URL
 *    → Service will stay in mock mode
 * 
 * 2. API returns different response format than expected
 *    → Update language.model.ts interfaces
 * 
 * 3. Mixing mock and API data
 *    → Never call mock methods when in API mode
 *    → LanguageService automatically switches via useMock flag
 * 
 * 4. Not handling loading state in UI
 *    → Use loaderService.isLoading$ | async in templates
 * 
 * 5. Not catching errors from API calls
 *    → Always implement error handler in subscribe()
 * 
 * 
 * 🎓 NEXT STEPS
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * 1. Review the generated files to understand the structure
 * 2. Implement your backend API endpoints following the contract
 * 3. Update api.config.ts with BASE_URL when backend is ready
 * 4. Test with mock mode first (existing functionality)
 * 5. Switch to API mode gradually (module by module)
 * 6. Add error handling UI in templates
 * 7. Implement loading spinners using loaderService
 * 8. Add user feedback notifications for success/error
 * 
 */
