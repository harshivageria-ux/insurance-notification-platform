import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { getApiUrl } from '../core/config/api.config';

export interface Channel {
  id: number;
  name: string;
  type: string;
  is_active: boolean;
}

@Injectable({ providedIn: 'root' })
export class ChannelService {
  private endpoint = 'channels';

  constructor(private http: HttpClient) {}

  getAll(): Observable<Channel[]> {
    return this.http.get<Channel[]>(getApiUrl(this.endpoint));
  }

  getById(id: number): Observable<Channel> {
    return this.http.get<Channel>(getApiUrl(`${this.endpoint}/${id}`));
  }

  create(data: Omit<Channel, 'id'>): Observable<Channel> {
    return this.http.post<Channel>(getApiUrl(this.endpoint), data);
  }

  update(id: number, data: Partial<Channel>): Observable<Channel> {
    return this.http.put<Channel>(getApiUrl(`${this.endpoint}/${id}`), data);
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(getApiUrl(`${this.endpoint}/${id}`));
  }
}
