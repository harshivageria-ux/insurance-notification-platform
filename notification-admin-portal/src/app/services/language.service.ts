import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of, throwError } from 'rxjs';
import { catchError, tap, finalize, map } from 'rxjs/operators';

import { Language, CreateLanguageRequest, UpdateLanguageRequest, ApiResponse } from '../core/models/language.model';
import { API_CONFIG, buildEndpoint, getApiUrl } from '../core/config/api.config';
import { ErrorHandlerService } from '../core/services/error-handler.service';
import { LoaderService } from '../core/services/loader.service';

@Injectable({
  providedIn: 'root'
})
export class LanguageService {

  private useMock: boolean;

  private mockLanguages: Language[] = [
    { id: 1, name: 'English', code: 'EN', status: 'Active' },
    { id: 2, name: 'Hindi', code: 'HI', status: 'Active' },
    { id: 3, name: 'Spanish', code: 'ES', status: 'Inactive' }
  ];

  constructor(
    private http: HttpClient,
    private errorHandler: ErrorHandlerService,
    private loaderService: LoaderService
  ) {
    this.useMock = !API_CONFIG.BASE_URL;
    console.log('🔧 LanguageService initialized:');
    console.log('   API_CONFIG.BASE_URL:', API_CONFIG.BASE_URL);
    console.log('   isMockMode:', this.useMock);
    console.log('   Full API_CONFIG:', API_CONFIG);
  }

  // ================= MAPPER =================
  private mapBackendToFrontend(language: any): Language {
    return {
      id: Number(language.id), // ✅ always number
      name: language.name,
      code: language.code,
      status: language.is_active ? 'Active' : 'Inactive',
      createdAt: language.created_at ? new Date(language.created_at) : undefined,
      updatedAt: language.updated_at ? new Date(language.updated_at) : undefined
    };
  }

  // ================= PUBLIC METHODS =================

  getLanguages(): Observable<Language[]> {
    this.loaderService.show();
    console.log(`📡 getLanguages() called - Mode: ${this.useMock ? 'MOCK' : 'API'}`);

    return (this.useMock
      ? this.getMockLanguages()
      : this.getLanguagesFromApi()
    ).pipe(
      catchError(error => {
        const appError = this.errorHandler.handleHttpError(error);
        console.error('❌ Failed to fetch languages:', appError);
        return this.getMockLanguages();
      }),
      finalize(() => this.loaderService.hide())
    );
  }

  addLanguage(request: CreateLanguageRequest): Observable<Language> {
    this.loaderService.show();
    console.log(`📡 addLanguage() called - Mode: ${this.useMock ? 'MOCK' : 'API'}`, request);

    return (this.useMock
      ? this.addMockLanguage(request)
      : this.addLanguageViaApi(request)
    ).pipe(
      tap(resp => console.log('✅ Language added:', resp)),
      catchError(error => {
        const appError = this.errorHandler.handleHttpError(error);
        console.error('❌ Failed to create language:', appError);
        return throwError(() => appError);
      }),
      finalize(() => this.loaderService.hide())
    );
  }

  updateLanguage(request: UpdateLanguageRequest): Observable<Language> {
    this.loaderService.show();
    console.log(`📡 updateLanguage() called - Mode: ${this.useMock ? 'MOCK' : 'API'}`, request);

    return (this.useMock
      ? this.updateMockLanguage(request)
      : this.updateLanguageViaApi(request)
    ).pipe(
      tap(resp => console.log('✅ Language updated:', resp)),
      catchError(error => {
        const appError = this.errorHandler.handleHttpError(error);
        console.error('❌ Failed to update language:', appError);
        return throwError(() => appError);
      }),
      finalize(() => this.loaderService.hide())
    );
  }

  deleteLanguage(id: number): Observable<boolean> {
    this.loaderService.show();

    return (this.useMock
      ? this.deleteMockLanguage(id) // ✅ FIXED
      : this.deleteLanguageViaApi(id)
    ).pipe(
      catchError(error => {
        const appError = this.errorHandler.handleHttpError(error);
        console.error('Failed to delete language:', appError);
        return throwError(() => appError);
      }),
      finalize(() => this.loaderService.hide())
    );
  }

  setMockMode(useMock: boolean): void {
    this.useMock = useMock;
    console.log(`Language Service: ${useMock ? 'MOCK' : 'API'} mode activated`);
  }

  isMockMode(): boolean {
    return this.useMock;
  }

  // ================= MOCK METHODS =================

  private getMockLanguages(): Observable<Language[]> {
    return of([...this.mockLanguages]);
  }

  private addMockLanguage(request: CreateLanguageRequest): Observable<Language> {
    const nextId = Math.max(...this.mockLanguages.map(l => l.id), 0) + 1;

    const newLanguage: Language = {
      id: nextId,
      ...request
    };

    this.mockLanguages.push(newLanguage);
    return of(newLanguage);
  }

  private updateMockLanguage(request: UpdateLanguageRequest): Observable<Language> {
    const index = this.mockLanguages.findIndex(l => l.id === request.id);

    if (index !== -1) {
      this.mockLanguages[index] = { ...this.mockLanguages[index], ...request };
      return of(this.mockLanguages[index]);
    }

    return throwError(() => new Error(`Language with id ${request.id} not found`));
  }

  private deleteMockLanguage(id: number): Observable<boolean> { // ✅ FIXED TYPE
    const index = this.mockLanguages.findIndex(l => l.id === id);

    if (index !== -1) {
      this.mockLanguages.splice(index, 1);
      return of(true);
    }

    return throwError(() => new Error(`Language with id ${id} not found`));
  }

  // ================= API METHODS =================

  private getLanguagesFromApi(): Observable<Language[]> {
    const url = getApiUrl(API_CONFIG.ENDPOINTS.LANGUAGES.GET_ALL);
    console.log('🌐 GET request to:', url);

    return this.http.get<ApiResponse<Language[]> | Language[]>(url).pipe(
      tap(response => console.log('✅ API response received:', response)),
      map(response => {
        const languages = Array.isArray(response) ? response : response.data || [];
        return languages.map(lang => this.mapBackendToFrontend(lang));
      })
    );
  }

  private addLanguageViaApi(request: CreateLanguageRequest): Observable<Language> {
    const url = getApiUrl(API_CONFIG.ENDPOINTS.LANGUAGES.CREATE);
    console.log('🌐 POST request to:', url, 'Body:', request);

    return this.http.post<ApiResponse<Language> | Language>(url, request).pipe(
      tap(response => console.log('✅ API response received:', response)),
      map(response => this.unwrapLanguageResponse(response))
    );
  }

  private updateLanguageViaApi(request: UpdateLanguageRequest): Observable<Language> {
    const endpoint = buildEndpoint(API_CONFIG.ENDPOINTS.LANGUAGES.UPDATE, { id: request.id });
    const url = getApiUrl(endpoint);
    console.log('🌐 PUT request to:', url, 'Body:', request);

    return this.http.put<ApiResponse<Language> | Language>(url, request).pipe(
      tap(response => console.log('✅ API response received:', response)),
      map(response => this.unwrapLanguageResponse(response))
    );
  }

  private deleteLanguageViaApi(id: number): Observable<boolean> {
    const endpoint = buildEndpoint(API_CONFIG.ENDPOINTS.LANGUAGES.DELETE, { id });
    const url = getApiUrl(endpoint);
    console.log('🌐 DELETE request to:', url);

    return this.http.delete<ApiResponse<boolean> | { message?: string }>(url).pipe(
      tap(response => console.log('✅ API response received:', response)),
      map(response => {
        if (typeof response === 'object' && response !== null && 'data' in response) {
          return response.data || true;
        }
        return true;
      })
    );
  }

  private unwrapLanguageResponse(response: ApiResponse<Language> | Language): Language {
    let language: any;

    if (response && typeof response === 'object' && 'data' in response && response.data) {
      language = response.data;
    } else {
      language = response;
    }

    return this.mapBackendToFrontend(language);
  }
}