import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, throwError } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class MappingService {
  private base = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  private handleError(err: any) {
    const message = err?.error?.error || err?.message || 'An error occurred';
    return throwError(() => ({ message }));
  }

  // Category <-> Channel
  getCategoryChannels(limit = 25, offset = 0): Observable<{ items: MappingCategoryChannel[] }> {
    let params = new HttpParams().set('limit', String(limit)).set('offset', String(offset));
    return this.http.get<{ items: MappingCategoryChannel[] }>(`${this.base}/notification-category-channel`, { params }).pipe(catchError(this.handleError));
  }

  createCategoryChannel(payload: { category_id: number; channel_id: number }): Observable<MappingCategoryChannel> {
    return this.http.post<MappingCategoryChannel>(`${this.base}/notification-category-channel`, payload).pipe(catchError(this.handleError));
  }

  deleteCategoryChannel(id: number): Observable<void> {
    return this.http.delete<void>(`${this.base}/notification-category-channel/${id}`).pipe(catchError(this.handleError));
  }

  // Channel <-> Provider
  getChannelProviders(limit = 25, offset = 0): Observable<{ items: MappingChannelProvider[] }> {
    let params = new HttpParams().set('limit', String(limit)).set('offset', String(offset));
    return this.http.get<{ items: MappingChannelProvider[] }>(`${this.base}/channel-provider-map`, { params }).pipe(catchError(this.handleError));
  }

  createChannelProvider(payload: { channel_id: number; provider_id: number; priority: number }): Observable<MappingChannelProvider> {
    return this.http.post<MappingChannelProvider>(`${this.base}/channel-provider-map`, payload).pipe(catchError(this.handleError));
  }

  deleteChannelProvider(id: number): Observable<void> {
    return this.http.delete<void>(`${this.base}/channel-provider-map/${id}`).pipe(catchError(this.handleError));
  }

  // Template <-> Channel <-> Language
  getTemplateChannelLanguages(limit = 25, offset = 0): Observable<{ items: MappingTemplateChannelLanguage[] }> {
    let params = new HttpParams().set('limit', String(limit)).set('offset', String(offset));
    return this.http.get<{ items: MappingTemplateChannelLanguage[] }>(`${this.base}/template-channel-language-map`, { params }).pipe(catchError(this.handleError));
  }

  createTemplateChannelLanguage(payload: { template_id: number; channel_id: number; language_id: number }): Observable<MappingTemplateChannelLanguage> {
    return this.http.post<MappingTemplateChannelLanguage>(`${this.base}/template-channel-language-map`, payload).pipe(catchError(this.handleError));
  }

  deleteTemplateChannelLanguage(id: number): Observable<void> {
    return this.http.delete<void>(`${this.base}/template-channel-language-map/${id}`).pipe(catchError(this.handleError));
  }

  // Category-Channel mapping
  getCategoryChannels(limit = 25, offset = 0): Observable<any> {
    const params = new HttpParams().set('limit', limit.toString()).set('offset', offset.toString());
    return this.http.get<any>(`${this.base}/category-channel`, { params }).pipe(catchError(this.handleError));
  }

  createCategoryChannel(payload: { category_id: number; channel_id: number }): Observable<any> {
    return this.http.post<any>(`${this.base}/category-channel`, payload).pipe(catchError(this.handleError));
  }

  deleteCategoryChannel(id: number): Observable<any> {
    return this.http.delete<any>(`${this.base}/category-channel/${id}`).pipe(catchError(this.handleError));
  }

  // Channel-Provider mapping
  getChannelProviders(limit = 25, offset = 0): Observable<any> {
    const params = new HttpParams().set('limit', limit.toString()).set('offset', offset.toString());
    return this.http.get<any>(`${this.base}/channel-provider`, { params }).pipe(catchError(this.handleError));
  }

  createChannelProvider(payload: { channel_id: number; provider_id: number; priority: number }): Observable<any> {
    return this.http.post<any>(`${this.base}/channel-provider`, payload).pipe(catchError(this.handleError));
  }

  deleteChannelProvider(id: number): Observable<any> {
    return this.http.delete<any>(`${this.base}/channel-provider/${id}`).pipe(catchError(this.handleError));
  }

  // Template mapping
  getTemplateMappings(limit = 25, offset = 0): Observable<any> {
    const params = new HttpParams().set('limit', limit.toString()).set('offset', offset.toString());
    return this.http.get<any>(`${this.base}/template-map`, { params }).pipe(catchError(this.handleError));
  }

  createTemplateMapping(payload: { template_id: number; channel_id: number; language_id: number }): Observable<any> {
    return this.http.post<any>(`${this.base}/template-map`, payload).pipe(catchError(this.handleError));
  }

  deleteTemplateMapping(id: number): Observable<any> {
    return this.http.delete<any>(`${this.base}/template-map/${id}`).pipe(catchError(this.handleError));
  }
}

