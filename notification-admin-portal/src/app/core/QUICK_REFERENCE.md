/**
 * QUICK REFERENCE: API INTEGRATION
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * 🔧 SWITCHING FROM MOCK TO API
 * ===================================
 * 
 * 1️⃣  Update Base URL
 *     File: src/app/core/config/api.config.ts
 *     Change: BASE_URL: '' → BASE_URL: 'https://api.notification.com/api/v1'
 * 
 * 2️⃣  Verify Response Format
 *     Backend should return: { success: true, data: { ... } }
 *     See ARCHITECTURE.md for details
 * 
 * 3️⃣  Test in Browser Console
 *     Let service = inject(LanguageService);
 *     service.setMockMode(false);
 *     service.getLanguages().subscribe(console.log);
 * 
 * 
 * 📦 FILE LOCATIONS
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * Configuration:
 *   📄 src/app/core/config/api.config.ts
 * 
 * Models (TypeScript Interfaces):
 *   📄 src/app/core/models/language.model.ts
 * 
 * Services:
 *   📄 src/app/services/language.service.ts          (Business logic)
 *   📄 src/app/core/services/error-handler.service.ts  (Error handling)
 *   📄 src/app/core/services/loader.service.ts        (Loading state)
 * 
 * Components:
 *   📄 src/app/modules/languages/languages.component.ts (UI)
 * 
 * 
 * 🎯 ENDPOINTS IMPLEMENTED
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * GET /languages               → getLanguages()
 * POST /languages              → addLanguage(request)
 * PUT /languages/:id           → updateLanguage(request)
 * DELETE /languages/:id        → deleteLanguage(id)
 * 
 * 
 * 💻 COMPONENT USAGE EXAMPLE
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * import { LanguageService } from '../../services/language.service';
 * import { Language } from '../../core/models/language.model';
 * import { LoaderService } from '../../core/services/loader.service';
 * 
 * export class MyComponent {
 *   languages: Language[] = [];
 * \n *   constructor(
 *     private languageService: LanguageService,
 *     public loaderService: LoaderService
 *   ) {}
 * \n *   ngOnInit() {
 *     this.languageService.getLanguages().subscribe({
 *       next: (data) => this.languages = data,
 *       error: (err) => console.error(err)
 *     });
 *   }
 * }
 * 
 * 
 * 🎨 TEMPLATE USAGE EXAMPLE
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * <!-- Show loader during API calls -->
 * <div *ngIf="loaderService.isLoading$ | async" class=\"loader\">
 *   Loading...
 * </div>
 * 
 * <!-- Display data -->
 * <div *ngIf=\"!(loaderService.isLoading$ | async)\" class=\"content\">
 *   <div *ngFor=\"let lang of languages\">
 *     <h3>{{ lang.name }}</h3>
 *     <p>Code: {{ lang.code }} - Status: {{ lang.status }}</p>
 *   </div>
 * </div>
 * 
 * 
 * 🔄 REQUEST/RESPONSE FORMAT
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * CREATE Request:
 *   POST /languages
 *   {
 *     \"name\": \"French\",
 *     \"code\": \"FR\",
 *     \"status\": \"Active\"
 *   }
 * 
 * CREATE Response:
 *   {
 *     \"success\": true,
 *     \"data\": {
 *       \"id\": 4,
 *       \"name\": \"French\",
 *       \"code\": \"FR\",
 *       \"status\": \"Active\"
 *     }
 *   }
 * 
 * UPDATE Request:
 *   PUT /languages/4
 *   {
 *     \"name\": \"French (Updated)\",
 *     \"code\": \"FR\",\n *     \"status\": \"Inactive\"
 *   }
 * 
 * DELETE Request:
 *   DELETE /languages/4
 * 
 * DELETE Response:
 *   {
 *     \"success\": true,
 *     \"data\": true
 *   }
 * 
 * 
 * ⚠️  ERROR HANDLING
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * try {
 *   this.languageService.addLanguage(request).subscribe({
 *     next: (language) => {
 *       // Success
 *     },
 *     error: (appError) => {
 *       // Automatically handled by service
 *       // Using centralized ErrorHandlerService
 *       console.error(appError.message);
 *     }
 *   });
 * } catch (err) {
 *   // Unexpected error
 * }
 * 
 * 
 * 🧪 TESTING WITH MOCK MODE
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * // Force mock mode in component/test
 * beforeEach(() => {
 *   const service = TestBed.inject(LanguageService);
 *   service.setMockMode(true);
 * });
 * 
 * // Check current mode
 * if (this.languageService.isMockMode()) {
 *   console.log('Running tests with mock data');
 * }
 * 
 * 
 * 🔗 RELATED FILES
 * ═══════════════════════════════════════════════════════════════════════════
 * 
 * ARCHITECTURE.md  → Comprehensive architecture guide\n * README.md        → Project documentation
 * 
 */
