import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { getApiUrl } from '../core/config/api.config';

export interface Template {
  id: number;
  name: string;
  content: string;
  template_group_id: number;
  variables: string;
}

@Injectable({ providedIn: 'root' })
export class TemplateService {
  private endpoint = 'templates';

  constructor(private http: HttpClient) {}

  getAll(): Observable<Template[]> {
    return this.http.get<Template[]>(getApiUrl(this.endpoint));
  }

  getById(id: number): Observable<Template> {
    return this.http.get<Template>(getApiUrl(`${this.endpoint}/${id}`));
  }

  create(data: Omit<Template, 'id'>): Observable<Template> {
    return this.http.post<Template>(getApiUrl(this.endpoint), data);
  }

  update(id: number, data: Partial<Template>): Observable<Template> {
    return this.http.put<Template>(getApiUrl(`${this.endpoint}/${id}`), data);
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(getApiUrl(`${this.endpoint}/${id}`));
  }
}
