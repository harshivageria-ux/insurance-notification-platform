import { CommonModule } from '@angular/common';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Subscription } from 'rxjs';

import { TemplateDrawerComponent } from './builder/template-drawer.component';
import { ToggleSwitchComponent } from './shared/toggle-switch.component';
import { ToastHostComponent } from './shared/toast-host.component';
import { ToastService } from './shared/toast.service';
import { PushNotificationTemplate } from './template-builder.types';
import { TemplateStateService } from './template-state.service';

@Component({
  selector: 'app-template-page',
  standalone: true,
  imports: [CommonModule, FormsModule, TemplateDrawerComponent, ToggleSwitchComponent, ToastHostComponent],
  template: `
    <section class="page-shell">
      <header class="page-header">
        <div>
          <p class="eyebrow">Templates</p>
          <h1>Templates</h1>
          <p class="subtitle">Manage push notification templates with live preview and variable support.</p>
        </div>
      </header>

      <div class="stats-grid">
        <article class="stat-card">
          <span class="stat-label">Total</span>
          <strong>{{ templates.length }}</strong>
        </article>
        <article class="stat-card">
          <span class="stat-label">Active</span>
          <strong>{{ activeCount }}</strong>
        </article>
        <article class="stat-card">
          <span class="stat-label">Inactive</span>
          <strong>{{ inactiveCount }}</strong>
        </article>
      </div>

      <div class="toolbar">
        <input class="search-input" type="text" placeholder="Search templates..." [(ngModel)]="searchText" />
        <button class="primary-btn" (click)="openCreate()">Add Template</button>
      </div>

      <div class="table-card">
        <ng-container *ngIf="loading; else content">
          <div class="skeleton">
            <div class="skeleton-row" *ngFor="let _ of [1,2,3,4,5]"></div>
          </div>
        </ng-container>

        <ng-template #content>
          <div class="empty-state" *ngIf="!filteredTemplates.length">
            <div class="empty-title">No templates found</div>
            <div class="empty-subtitle">Create a push notification template to get started.</div>
            <button class="primary-btn" (click)="openCreate()">Add Template</button>
          </div>

          <div class="table-wrap" *ngIf="filteredTemplates.length">
            <table>
              <thead>
                <tr>
                  <th>Template Name</th>
                  <th>Channel</th>
                  <th>Title</th>
                  <th>Status</th>
                  <th>Last Updated</th>
                  <th class="actions">Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr *ngFor="let t of filteredTemplates">
                  <td class="strong">{{ t.templateName }}</td>
                  <td><span class="badge">{{ t.channel }}</span></td>
                  <td><span class="clamp">{{ t.title }}</span></td>
                  <td>
                    <div class="status-cell">
                      <app-toggle-switch [checked]="t.isActive" (checkedChange)="toggleStatus(t)"></app-toggle-switch>
                      <span class="status-label" [class.active]="t.isActive">{{ t.isActive ? 'Active' : 'Inactive' }}</span>
                    </div>
                  </td>
                  <td>{{ formatDate(t.updatedAt) }}</td>
                  <td class="actions-cell">
                    <button class="ghost-btn" (click)="openEdit(t)">Edit</button>
                    <button class="ghost-btn" (click)="openView(t)">View</button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </ng-template>
      </div>
    </section>

    <app-template-drawer
      [open]="drawerOpen"
      [mode]="drawerMode"
      [template]="drawerTemplate"
      (requestClose)="closeDrawer()"
      (save)="handleSave($event)"
    ></app-template-drawer>

    <app-toast-host></app-toast-host>
  `,
  styleUrls: ['./template-page.component.scss']
})
export class TemplatePageComponent implements OnInit, OnDestroy {
  templates: PushNotificationTemplate[] = [];
  searchText = '';
  loading = true;

  drawerOpen = false;
  drawerMode: 'create' | 'edit' | 'view' = 'create';
  drawerTemplate: PushNotificationTemplate | null = null;

  private sub?: Subscription;

  constructor(private readonly state: TemplateStateService, private readonly toast: ToastService) {}

  ngOnInit(): void {
    // Simulated loading state
    setTimeout(() => (this.loading = false), 450);
    this.sub = this.state.templates$().subscribe((t) => (this.templates = t));
  }

  ngOnDestroy(): void {
    this.sub?.unsubscribe();
  }

  get filteredTemplates(): PushNotificationTemplate[] {
    const q = this.searchText.trim().toLowerCase();
    if (!q) return this.templates;
    return this.templates.filter((t) => {
      return (
        t.templateName.toLowerCase().includes(q) ||
        t.title.toLowerCase().includes(q) ||
        t.message.toLowerCase().includes(q)
      );
    });
  }

  get activeCount(): number {
    return this.templates.filter((t) => t.isActive).length;
  }

  get inactiveCount(): number {
    return this.templates.filter((t) => !t.isActive).length;
  }

  formatDate(iso: string): string {
    const d = new Date(iso);
    return d.toLocaleString();
  }

  openCreate(): void {
    this.drawerMode = 'create';
    this.drawerTemplate = null;
    this.drawerOpen = true;
  }

  openEdit(t: PushNotificationTemplate): void {
    this.drawerMode = 'edit';
    this.drawerTemplate = t;
    this.drawerOpen = true;
  }

  openView(t: PushNotificationTemplate): void {
    this.drawerMode = 'view';
    this.drawerTemplate = t;
    this.drawerOpen = true;
  }

  closeDrawer(): void {
    this.drawerOpen = false;
  }

  toggleStatus(t: PushNotificationTemplate): void {
    this.state.toggleActive(t.id);
    this.toast.show({ kind: 'info', title: 'Status updated', message: `${t.templateName} is now ${!t.isActive ? 'Active' : 'Inactive'}.` });
  }

  handleSave(evt: { mode: 'create' | 'edit'; value: any }): void {
    const value = evt.value;
    const payload = {
      templateName: value.templateName,
      channel: 'Push Notification' as const,
      title: value.title,
      message: value.message,
      targetUrl: value.targetUrl,
      iconUrl: value.iconUrl,
      bannerUrl: value.bannerUrl,
      requireInteraction: value.requireInteraction,
      priority: value.priority,
      ttl: Number(value.ttl || 0),
      isActive: true
    };

    if (evt.mode === 'edit' && this.drawerTemplate) {
      this.state.update(this.drawerTemplate.id, payload);
      this.toast.show({ kind: 'success', title: 'Template updated', message: this.drawerTemplate.templateName });
    } else {
      this.state.create(payload);
      this.toast.show({ kind: 'success', title: 'Template saved', message: payload.templateName });
    }
    this.closeDrawer();
  }
}

