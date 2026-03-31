/**
 * ADDING NEW API MODULES - CHECKLIST
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * When adding a new feature (e.g., Users, Tenants, Notifications), follow
 * this checklist to maintain consistency with the enterprise architecture.
 * 
 * 
 * ✅ STEP 1: Create Models
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * File: src/app/core/models/{feature}.model.ts
 * 
 * Contents:
 *   - Main entity interface (e.g., User, Tenant)
 *   - Create request interface
 *   - Update request interface
 *   - Any other related DTOs
 * 
 * Example structure:
 *   export interface User { id, name, email, status, createdAt?, updatedAt? }
 *   export interface CreateUserRequest { name, email, role }
 *   export interface UpdateUserRequest extends CreateUserRequest { id }
 *   export interface ApiResponse<T> { success, data?, message?, errors? }
 * 
 * Then update src/app/core/models/index.ts to export the new model.
 * 
 * 
 * ✅ STEP 2: Add API Endpoints
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * File: src/app/core/config/api.config.ts
 * 
 * Add to API_CONFIG.ENDPOINTS:
 * 
 *   USERS: {
 *     GET_ALL: '/users',
 *     CREATE: '/users',
 *     UPDATE: '/users/:id',
 *     DELETE: '/users/:id',
 *     SEARCH: '/users/search'  // Optional
 *   }
 * 
 * Use consistent HTTP verbs:
 *   - GET for retrieval
 *   - POST for creation (maps to sp_insert_*)
 *   - PUT for updates (maps to sp_update_*)
 *   - DELETE for deletion/soft delete (maps to sp_deactivate_*)
 * 
 * 
 * ✅ STEP 3: Create Business Service
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * File: src/app/services/{feature}.service.ts
 * 
 * Structure:
 *   @Injectable({ providedIn: 'root' })
 *   export class UserService {
 *     private useMock = true;  // ← Important!
 *     private mockUsers: User[] = [ /* initial data */ ];
 * 
 *     constructor(
 *       private http: HttpClient,
 *       private errorHandler: ErrorHandlerService,
 *       private loaderService: LoaderService
 *     ) {
 *       this.useMock = !API_CONFIG.BASE_URL;  // ← Auto-switch!
 *     }
 * 
 *     // Public methods (both mock and API support)
 *     getUsers(): Observable<User[]> { ... }
 *     addUser(request: CreateUserRequest): Observable<User> { ... }
 *     updateUser(request: UpdateUserRequest): Observable<User> { ... }
 *     deleteUser(id: number): Observable<boolean> { ... }
 *     setMockMode(useMock: boolean): void { ... }
 *     isMockMode(): boolean { ... }
 * 
 *     // Private mock methods
 *     private getMockUsers(): Observable<User[]> { ... }
 *     private addMockUser(request): Observable<User> { ... }
 *     private updateMockUser(request): Observable<User> { ... }
 *     private deleteMockUser(id): Observable<boolean> { ... }
 * 
 *     // Private API methods
 *     private getUsersFromApi(): Observable<User[]> { ... }
 *     private addUserViaApi(request): Observable<User> { ... }
 *     private updateUserViaApi(request): Observable<User> { ... }
 *     private deleteUserViaApi(id): Observable<boolean> { ... }
 *   }
 * 
 * Then update src/app/services/index.ts to export the new service.
 * 
 * 
 * ✅ STEP 4: Update Components
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * File: src/app/modules/{feature}/{feature}.component.ts
 * 
 * Key points:
 *   - Use proper typing (User[], not any[])
 *   - Import service and LoaderService
 *   - Subscribe to isLoading$ in template for loader
 *   - Implement proper error handling
 *   - Follow the existing patterns in LanguagesComponent
 * 
 * Example constructor:
 *   constructor(
 *     private userService: UserService,
 *     public loaderService: LoaderService
 *   ) {}
 * 
 * 
 * ✅ STEP 5: Update Templates
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * Show loader during API calls:
 *   <div *ngIf=\"loaderService.isLoading$ | async\" class=\"spinner\">
 *     Loading...
 *   </div>
 * 
 * Show content when loaded:
 *   <div *ngIf=\"!(loaderService.isLoading$ | async)\">
 *     <!-- Your content here -->
 *   </div>
 * 
 * 
 * ✅ STEP 6: Testing & Verification
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * 1. Test with mock mode (default)
 *    - Verify business logic works with mock data
 *    - Check UI renders correctly
 * 
 * 2. Test with API mode (when backend ready)
 *    service.setMockMode(false);
 *    - Verify API calls are made
 *    - Check error handling
 *    - Verify loading state
 * 
 * 3. Manual testing
 *    - Open browser DevTools
 *    - Network tab to see API calls
 *    - Console for any errors
 * 
 * 
 * 🎯 BEST PRACTICES CHECKLIST
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * ✓ Use TypeScript interfaces, NOT 'any' types
 * ✓ Always use ErrorHandlerService for error handling
 * ✓ Always use LoaderService for loading state
 * ✓ Keep mock data for testing/fallback
 * ✓ Separate mock and API methods with private keywords
 * ✓ Use RxJS operators (tap, catchError, finalize)
 * ✓ Implement proper unsubscribe patterns (OnDestroy)
 * ✓ Add comments explaining complex logic
 * ✓ Follow the Language module pattern
 * ✓ Use barrel exports (index.ts) for cleaner imports\n * ✓ Test with both mock and API modes
 * \n * ✓ Update this checklist when adding new patterns
 * 
 * \n * 🔍 COMMON MISTAKES TO AVOID
 * ═══════════════════════════════════════════════════════════════════════════
 * \n * ❌ Using 'any' type instead of interfaces\n *    Reason: Loses type safety and IDE support\n *    Fix: Create interface in core/models/\n * \n * ❌ Forgetting loaderService.show()/hide()\n *    Reason: UI doesn't show loading state\n *    Fix: Always call show() at start, hide() in finalize\n * \n * ❌ Not handling errors in subscribe()\n *    Reason: Failed API calls silently fail\n *    Fix: Always implement error handler\n * \n * ❌ Mixing mock and API calls\n *    Reason: Inconsistent behavior\n *    Fix: Let useMock flag control which methods are called\n * \n * ❌ Direct HTTP calls in components\n *    Reason: Breaks separation of concerns\n *    Fix: Always go through service layer\n * \n * ❌ Not updating api.config.ts\n *    Reason: Service stays in mock mode\n *    Fix: Set BASE_URL in api.config.ts when backend ready\n * \n * \n * 📋 QUICK COPY-PASTE TEMPLATES\n * ═══════════════════════════════════════════════════════════════════════════\n * \n * [See bottom of file for template code]\n * \n */\n \n// ============================================================================\n// TEMPLATE CODE - Copy & customize for new modules\n// ============================================================================\n \n/*\n \n// 1. MODEL TEMPLATE (src/app/core/models/example.model.ts)\n export interface Example {\n   id: number;\n   name: string;\n   status: 'Active' | 'Inactive';\n   createdAt?: Date;\n   updatedAt?: Date;\n }\n \n export interface CreateExampleRequest {\n   name: string;\n   status: 'Active' | 'Inactive';\n }\n \n export interface UpdateExampleRequest extends CreateExampleRequest {\n   id: number;\n }\n \n export interface ApiResponse<T> {\n   success: boolean;\n   data?: T;\n   message?: string;\n   errors?: string[];\n }\n \n \n// 2. SERVICE TEMPLATE (src/app/services/example.service.ts)\n import { Injectable } from '@angular/core';\n import { HttpClient } from '@angular/common/http';\n import { Observable, of, throwError } from 'rxjs';\n import { catchError, tap, finalize } from 'rxjs/operators';\n import { Example, CreateExampleRequest, UpdateExampleRequest } from '../core/models/example.model';\n import { API_CONFIG, buildEndpoint, getApiUrl } from '../core/config/api.config';\n import { ErrorHandlerService } from '../core/services/error-handler.service';\n import { LoaderService } from '../core/services/loader.service';\n \n @Injectable({ providedIn: 'root' })\n export class ExampleService {\n   private useMock = true;\n   private mockExamples: Example[] = [\n     { id: 1, name: 'Example 1', status: 'Active' }\n   ];\n \n   constructor(\n     private http: HttpClient,\n     private errorHandler: ErrorHandlerService,\n     private loaderService: LoaderService\n   ) {\n     this.useMock = !API_CONFIG.BASE_URL;\n   }\n \n   getExamples(): Observable<Example[]> {\n     this.loaderService.show();\n     return (this.useMock\n       ? this.getMockExamples()\n       : this.getExamplesFromApi()\n     ).pipe(\n       catchError(error => {\n         this.errorHandler.handleHttpError(error);\n         return this.getMockExamples();\n       }),\n       finalize(() => this.loaderService.hide())\n     );\n   }\n \n   private getMockExamples(): Observable<Example[]> {\n     return of([...this.mockExamples]);\n   }\n \n   private getExamplesFromApi(): Observable<Example[]> {\n     const url = getApiUrl(API_CONFIG.ENDPOINTS.EXAMPLES.GET_ALL);\n     return this.http.get<ApiResponse<Example[]>>(url);\n   }\n }\n \n*/\n