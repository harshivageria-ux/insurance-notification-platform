import { CommonModule } from '@angular/common';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';

import { ToastMessage, ToastService } from './toast.service';

@Component({
  selector: 'app-toast-host',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="toast-wrap" *ngIf="toast">
      <div class="toast" [class.toast-success]="toast.kind === 'success'" [class.toast-error]="toast.kind === 'error'">
        <div class="toast-title">{{ toast.title }}</div>
        <div class="toast-message" *ngIf="toast.message">{{ toast.message }}</div>
      </div>
    </div>
  `,
  styleUrls: ['./toast-host.component.scss']
})
export class ToastHostComponent implements OnInit, OnDestroy {
  toast: ToastMessage | null = null;
  private sub?: Subscription;

  constructor(private readonly toastService: ToastService) {}

  ngOnInit(): void {
    this.sub = this.toastService.toast$().subscribe((t) => (this.toast = t));
  }

  ngOnDestroy(): void {
    this.sub?.unsubscribe();
  }
}

