import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

export interface ToastMessage {
  id: string;
  kind: 'success' | 'info' | 'error';
  title: string;
  message?: string;
}

@Injectable({ providedIn: 'root' })
export class ToastService {
  private readonly toastSubject = new BehaviorSubject<ToastMessage | null>(null);

  toast$(): Observable<ToastMessage | null> {
    return this.toastSubject.asObservable();
  }

  show(toast: Omit<ToastMessage, 'id'>, ttlMs = 2500): void {
    const id = `${Date.now().toString(16)}_${Math.random().toString(16).slice(2)}`;
    this.toastSubject.next({ id, ...toast });
    setTimeout(() => {
      if (this.toastSubject.value?.id === id) {
        this.toastSubject.next(null);
      }
    }, ttlMs);
  }
}

