import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-notification-preview',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="preview-shell">
      <div class="chrome-label">Live Preview</div>

      <div class="push-card">
        <div class="push-row">
          <div class="icon">
            <img
              *ngIf="iconUrl; else fallbackIcon"
              [src]="iconUrl"
              alt="Icon"
              (error)="iconFailed = true"
              (load)="iconFailed = false"
            />
            <ng-template #fallbackIcon>
              <div class="icon-fallback">PN</div>
            </ng-template>
          </div>

          <div class="content">
            <div class="title">{{ title || 'Notification title' }}</div>
            <div class="message">{{ message || 'Notification message preview will appear here.' }}</div>
          </div>
        </div>

        <div class="banner" *ngIf="bannerUrl">
          <img [src]="bannerUrl" alt="Banner" (error)="bannerFailed = true" (load)="bannerFailed = false" />
          <div class="banner-fallback" *ngIf="bannerFailed">Banner image failed to load</div>
        </div>

        <div class="meta">
          <span class="meta-pill">{{ priority || 'Normal' }} Priority</span>
          <span class="meta-pill">TTL {{ ttl || 0 }}s</span>
          <span class="meta-pill" *ngIf="requireInteraction">Requires interaction</span>
        </div>
      </div>
    </div>
  `,
  styleUrls: ['./notification-preview.component.scss']
})
export class NotificationPreviewComponent {
  @Input() title = '';
  @Input() message = '';
  @Input() iconUrl = '';
  @Input() bannerUrl = '';
  @Input() priority: 'High' | 'Normal' = 'Normal';
  @Input() ttl = 0;
  @Input() requireInteraction = false;

  iconFailed = false;
  bannerFailed = false;
}

