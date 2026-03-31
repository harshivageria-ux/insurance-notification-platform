import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { getApiUrl } from '../core/config/api.config';

export interface Category {
  id: string;
  name: string;
  description: string;
  status: 'Active' | 'Inactive';
}

@Injectable({ providedIn: 'root' })
export class CategoryService {
  private endpoint = 'categories';

  constructor(private http: HttpClient) {}

  /**
   * Map backend category response to frontend Category model
   */
  private mapBackendToFrontend(category: any): Category {
    return {
      id: category.id,
      name: category.name,
      description: category.description,
      status: category.is_active ? 'Active' : 'Inactive'
    };
  }

  getAll(): Observable<Category[]> {
    return this.http.get<any[]>(getApiUrl(this.endpoint)).pipe(
      map(categories => categories.map(cat => this.mapBackendToFrontend(cat)))
    );
  }

  getById(id: string): Observable<Category> {
    return this.http.get<any>(getApiUrl(`${this.endpoint}/${id}`)).pipe(
      map(category => this.mapBackendToFrontend(category))
    );
  }

  create(data: Omit<Category, 'id'>): Observable<Category> {
    const payload = {
      name: data.name,
      description: data.description,
      status: data.status
    };
    return this.http.post<any>(getApiUrl(this.endpoint), payload).pipe(
      map(category => this.mapBackendToFrontend(category))
    );
  }

  update(id: string, data: Partial<Category>): Observable<Category> {
    const payload = {
      ...data,
      status: data.status
    };
    return this.http.put<any>(getApiUrl(`${this.endpoint}/${id}`), payload).pipe(
      map(category => this.mapBackendToFrontend(category))
    );
  }

  delete(id: string): Observable<void> {
    return this.http.delete<void>(getApiUrl(`${this.endpoint}/${id}`));
  }
}
