import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { API_CONFIG, getApiUrl } from '../core/config/api.config';

export interface Priority {
  id: number;
  name: string;
  level: number;
  status: 'Active' | 'Inactive';
}

@Injectable({ providedIn: 'root' })
export class PriorityService {
  private endpoint = 'priorities';

  constructor(private http: HttpClient) {}

  getAll(): Observable<Priority[]> {
    return this.http.get<Priority[]>(getApiUrl(this.endpoint));
  }

  getById(id: number): Observable<Priority> {
    return this.http.get<Priority>(getApiUrl(`${this.endpoint}/${id}`));
  }

  create(data: Omit<Priority, 'id'>): Observable<Priority> {
    return this.http.post<Priority>(getApiUrl(this.endpoint), data);
  }

  update(id: number, data: Partial<Priority>): Observable<Priority> {
    return this.http.put<Priority>(getApiUrl(`${this.endpoint}/${id}`), data);
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(getApiUrl(`${this.endpoint}/${id}`));
  }
}
