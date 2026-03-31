import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, map } from 'rxjs';

import { buildEndpoint, getApiUrl } from '../core/config/api.config';
import { ApiEnvelope, EntityIdentifier } from '../core/models/entity-crud.model';

@Injectable({ providedIn: 'root' })
export class EntityCrudService {
  constructor(private readonly http: HttpClient) {}

  list(endpoint: string, listParamKey?: string, listParamValue?: string | number | null): Observable<Record<string, unknown>[]> {
    const hasListParamValue = listParamValue !== null && listParamValue !== undefined && String(listParamValue).trim() !== '';
    const finalEndpoint = listParamKey && hasListParamValue
      ? buildEndpoint(`${endpoint}/:${listParamKey}`, { [listParamKey]: listParamValue })
      : endpoint;

    return this.http.get<Record<string, unknown>[] | ApiEnvelope<Record<string, unknown>[]>>(getApiUrl(finalEndpoint)).pipe(
      map((response) => this.unwrapArray(response))
    );
  }

  create(endpoint: string, payload: Record<string, unknown>): Observable<Record<string, unknown>> {
    return this.http.post<Record<string, unknown> | ApiEnvelope<Record<string, unknown>>>(getApiUrl(endpoint), payload).pipe(
      map((response) => this.unwrapObject(response))
    );
  }

  update(endpoint: string, id: EntityIdentifier, payload: Record<string, unknown>): Observable<Record<string, unknown>> {
    return this.http.put<Record<string, unknown> | ApiEnvelope<Record<string, unknown>>>(
      getApiUrl(buildEndpoint(`${endpoint}/:id`, { id })),
      payload
    ).pipe(map((response) => this.unwrapObject(response)));
  }

  delete(endpoint: string, id: EntityIdentifier): Observable<unknown> {
    return this.http.delete(getApiUrl(buildEndpoint(`${endpoint}/:id`, { id })));
  }

  toggle(toggleEndpoint: string, id: EntityIdentifier, isActive: boolean): Observable<Record<string, unknown>> {
    return this.http.patch<Record<string, unknown> | ApiEnvelope<Record<string, unknown>>>(
      getApiUrl(buildEndpoint(toggleEndpoint, { id })),
      { is_active: isActive }
    ).pipe(map((response) => this.unwrapObject(response)));
  }

  private unwrapArray(
    response: Record<string, unknown>[] | ApiEnvelope<Record<string, unknown>[]>
  ): Record<string, unknown>[] {
    if (Array.isArray(response)) {
      return response;
    }

    return Array.isArray(response.data) ? response.data : [];
  }

  private unwrapObject(
    response: Record<string, unknown> | ApiEnvelope<Record<string, unknown>>
  ): Record<string, unknown> {
    if ('data' in response && response.data && typeof response.data === 'object') {
      return response.data as Record<string, unknown>;
    }

    return response as Record<string, unknown>;
  }
}
