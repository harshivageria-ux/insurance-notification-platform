import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

import { MOCK_PUSH_TEMPLATES } from './template-mock-data';
import { PushNotificationTemplate } from './template-builder.types';

@Injectable({ providedIn: 'root' })
export class TemplateStateService {
  private readonly templatesSubject = new BehaviorSubject<PushNotificationTemplate[]>(
    MOCK_PUSH_TEMPLATES.map((t) => ({ ...t }))
  );

  templates$(): Observable<PushNotificationTemplate[]> {
    return this.templatesSubject.asObservable();
  }

  snapshot(): PushNotificationTemplate[] {
    return this.templatesSubject.value;
  }

  create(template: Omit<PushNotificationTemplate, 'id' | 'updatedAt'>): void {
    const id = `tpl_${Math.random().toString(16).slice(2)}_${Date.now().toString(16)}`;
    const updatedAt = new Date().toISOString();
    const next: PushNotificationTemplate = { ...template, id, updatedAt };
    this.templatesSubject.next([next, ...this.templatesSubject.value]);
  }

  update(id: string, patch: Partial<Omit<PushNotificationTemplate, 'id'>>): void {
    const updatedAt = new Date().toISOString();
    this.templatesSubject.next(
      this.templatesSubject.value.map((t) => (t.id === id ? { ...t, ...patch, updatedAt } : t))
    );
  }

  toggleActive(id: string): void {
    const t = this.templatesSubject.value.find((x) => x.id === id);
    if (!t) return;
    this.update(id, { isActive: !t.isActive });
  }
}

