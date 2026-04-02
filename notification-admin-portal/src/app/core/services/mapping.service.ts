import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, throwError } from 'rxjs';

export interface MappingCategoryChannel {
  id: number;
  category_id: number;
  channel_id: number;
  is_active: boolean;
  created_at: string;
}

export interface MappingChannelProvider {
  id: number;
  channel_id: number;
  provider_id: number;
  priority?: number;
  is_active: boolean;
}

export interface MappingTemplateChannelLanguage {
  id: number;
  template_group_id?: number;
  template_id: number;
  channel_id: number;
  language_id: number;
  is_active: boolean;
  created_at: string;
}

@Injectable({ providedIn: 'root' })
export class MappingService {
  private mappingBase = 'http://localhost:9100/api';
  private coreBase = 'http://localhost:9000';

  constructor(private http: HttpClient) {}

  private handleError(err: any) {
    const message = err?.error?.error || err?.message || 'An error occurred';
    return throwError(() => ({ message }));
  }

  getCategories(): Observable<any[]> {
    return this.http.get<any[]>(`${this.coreBase}/categories`).pipe(catchError(this.handleError));
  }

  getChannels(): Observable<any[]> {
    return this.http.get<any[]>(`${this.coreBase}/channels`).pipe(catchError(this.handleError));
  }

  getProviders(): Observable<any[]> {
    return this.http.get<any[]>(`${this.coreBase}/channel-providers`).pipe(catchError(this.handleError));
  }

  getLanguages(): Observable<any[]> {
    return this.http.get<any[]>(`${this.coreBase}/languages`).pipe(catchError(this.handleError));
  }

  getTemplateGroups(): Observable<any[]> {
    return this.http.get<any[]>(`${this.coreBase}/template-groups`).pipe(catchError(this.handleError));
  }

  getTemplates(): Observable<any[]> {
    return this.http.get<any[]>(`${this.coreBase}/templates`).pipe(catchError(this.handleError));
  }

  getCategoryChannels(limit = 25, offset = 0): Observable<{ items: MappingCategoryChannel[] }> {
    const params = new HttpParams().set('limit', String(limit)).set('offset', String(offset));
    return this.http
      .get<{ items: MappingCategoryChannel[] }>(`${this.mappingBase}/notification-category-channel`, { params })
      .pipe(catchError(this.handleError));
  }

  createCategoryChannel(payload: { category_id: number; channel_id: number }): Observable<MappingCategoryChannel> {
    return this.http
      .post<MappingCategoryChannel>(`${this.mappingBase}/notification-category-channel`, payload)
      .pipe(catchError(this.handleError));
  }

  deleteCategoryChannel(id: number): Observable<void> {
    return this.http.delete<void>(`${this.mappingBase}/notification-category-channel/${id}`).pipe(catchError(this.handleError));
  }

  getChannelProviders(limit = 25, offset = 0): Observable<{ items: MappingChannelProvider[] }> {
    const params = new HttpParams().set('limit', String(limit)).set('offset', String(offset));
    return this.http
      .get<{ items: MappingChannelProvider[] }>(`${this.mappingBase}/channel-provider-map`, { params })
      .pipe(catchError(this.handleError));
  }

  createChannelProvider(payload: { channel_id: number; provider_id: number; priority: number }): Observable<MappingChannelProvider> {
    return this.http
      .post<MappingChannelProvider>(`${this.mappingBase}/channel-provider-map`, payload)
      .pipe(catchError(this.handleError));
  }

  deleteChannelProvider(id: number): Observable<void> {
    return this.http.delete<void>(`${this.mappingBase}/channel-provider-map/${id}`).pipe(catchError(this.handleError));
  }

  getTemplateChannelLanguages(limit = 25, offset = 0): Observable<{ items: MappingTemplateChannelLanguage[] }> {
    const params = new HttpParams().set('limit', String(limit)).set('offset', String(offset));
    return this.http
      .get<{ items: MappingTemplateChannelLanguage[] }>(`${this.mappingBase}/template-channel-language-map`, { params })
      .pipe(catchError(this.handleError));
  }

  createTemplateChannelLanguage(
    payload: { template_group_id: number; template_id: number; channel_id: number; language_id: number }
  ): Observable<MappingTemplateChannelLanguage> {
    return this.http
      .post<MappingTemplateChannelLanguage>(`${this.mappingBase}/template-channel-language-map`, payload)
      .pipe(catchError(this.handleError));
  }

  deleteTemplateChannelLanguage(id: number): Observable<void> {
    return this.http.delete<void>(`${this.mappingBase}/template-channel-language-map/${id}`).pipe(catchError(this.handleError));
  }
}
