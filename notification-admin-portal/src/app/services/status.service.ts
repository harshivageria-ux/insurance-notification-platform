import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { getApiUrl } from '../core/config/api.config';

export interface Status {
  id: number;
  name: string;
  description: string;
  status: 'Active' | 'Inactive';
}

@Injectable({ providedIn: 'root' })
export class StatusService {
  private endpoint = 'statuses';

  constructor(private http: HttpClient) {}

  /**
   * Map backend status response to frontend Status model
   */
  private mapBackendToFrontend(status: any): Status {
    return {
      id: status.status_id,
      name: status.status_name,
      description: '', // Backend doesn't have description
      status: status.is_active ? 'Active' : 'Inactive'
    };
  }

  getAll(): Observable<Status[]> {
    return this.http.get<any[]>(getApiUrl(this.endpoint)).pipe(
      map(statuses => statuses.map(stat => this.mapBackendToFrontend(stat)))
    );
  }

  getById(id: number): Observable<Status> {
    return this.http.get<any>(getApiUrl(`${this.endpoint}/${id}`)).pipe(
      map(status => this.mapBackendToFrontend(status))
    );
  }

  create(data: Omit<Status, 'id'>): Observable<Status> {
    const payload = {
      status_name: data.name,
      status: data.status
    };
    return this.http.post<any>(getApiUrl(this.endpoint), payload).pipe(
      map(status => this.mapBackendToFrontend(status))
    );
  }

  update(id: number, data: Partial<Status>): Observable<Status> {
    const payload = {
      status_name: data.name,
      status: data.status
    };
    return this.http.put<any>(getApiUrl(`${this.endpoint}/${id}`), payload).pipe(
      map(status => this.mapBackendToFrontend(status))
    );
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(getApiUrl(`${this.endpoint}/${id}`));
  }
}
