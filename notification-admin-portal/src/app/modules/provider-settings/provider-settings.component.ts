import { CommonModule } from '@angular/common';
import { Component, OnInit, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, RouterLink } from '@angular/router';

import { EntityCrudService } from '../../services/entity-crud.service';

type ProtocolKey = 'http' | 'smtp' | 'pop';

type HttpSettings = {
  baseUrl: string;
  username: string;
  password: string;
};

type SmtpSettings = {
  host: string;
  port: number;
  username: string;
  password: string;
  useTls: boolean;
};

type PopSettings = {
  host: string;
  port: number;
  username: string;
  password: string;
  useSsl: boolean;
};

@Component({
  selector: 'app-provider-settings',
  standalone: true,
  imports: [CommonModule, FormsModule, RouterLink],
  template: `
    <section class="page-shell" *ngIf="providerId > 0">
      <header class="page-header">
        <div>
          <p class="eyebrow">Admin Portal</p>
          <h1>Provider Settings</h1>
          <p class="subtitle">Configure HTTP / SMTP / POP credentials for provider ID {{ providerId }}</p>
        </div>
        <div class="header-actions">
          <a routerLink="/channel-providers" class="secondary-link">Back</a>
        </div>
      </header>

      <div class="tabs">
        <button class="tab-btn" [class.active]="activeTab === 'http'" (click)="activeTab = 'http'">HTTP</button>
        <button class="tab-btn" [class.active]="activeTab === 'smtp'" (click)="activeTab = 'smtp'">SMTP</button>
        <button class="tab-btn" [class.active]="activeTab === 'pop'" (click)="activeTab = 'pop'">POP</button>
      </div>

      <div class="card">
        <!-- HTTP -->
        <div *ngIf="activeTab === 'http'">
          <div class="form-grid">
            <label class="field">
              <span>Base URL</span>
              <input type="text" [(ngModel)]="http.baseUrl" placeholder="https://example.com/api" />
            </label>

            <label class="field">
              <span>Username</span>
              <input type="text" [(ngModel)]="http.username" placeholder="username" />
            </label>

            <label class="field">
              <span>Password</span>
              <input type="password" [(ngModel)]="http.password" placeholder="password" />
            </label>
          </div>
        </div>

        <!-- SMTP -->
        <div *ngIf="activeTab === 'smtp'">
          <div class="form-grid">
            <label class="field">
              <span>Host/IP</span>
              <input type="text" [(ngModel)]="smtp.host" placeholder="smtp.example.com" />
            </label>

            <label class="field">
              <span>Port</span>
              <input type="number" [(ngModel)]="smtp.port" placeholder="587" />
            </label>

            <label class="field">
              <span>Username</span>
              <input type="text" [(ngModel)]="smtp.username" placeholder="username" />
            </label>

            <label class="field">
              <span>Password</span>
              <input type="password" [(ngModel)]="smtp.password" placeholder="password" />
            </label>

            <label class="field checkbox">
              <span>Use TLS</span>
              <input type="checkbox" [(ngModel)]="smtp.useTls" />
            </label>
          </div>
        </div>

        <!-- POP -->
        <div *ngIf="activeTab === 'pop'">
          <div class="form-grid">
            <label class="field">
              <span>Host/IP</span>
              <input type="text" [(ngModel)]="pop.host" placeholder="pop.example.com" />
            </label>

            <label class="field">
              <span>Port</span>
              <input type="number" [(ngModel)]="pop.port" placeholder="110" />
            </label>

            <label class="field">
              <span>Username</span>
              <input type="text" [(ngModel)]="pop.username" placeholder="username" />
            </label>

            <label class="field">
              <span>Password</span>
              <input type="password" [(ngModel)]="pop.password" placeholder="password" />
            </label>

            <label class="field checkbox">
              <span>Use SSL</span>
              <input type="checkbox" [(ngModel)]="pop.useSsl" />
            </label>
          </div>
        </div>
      </div>

      <div class="actions-row">
        <div class="status">{{ saveMessage }}</div>
        <div class="actions">
          <button class="primary-btn" (click)="saveActiveTab()">Save {{ activeTab.toUpperCase() }}</button>
        </div>
      </div>
    </section>
  `,
  styles: [
    `
      .page-shell {
        padding: 32px;
        min-height: 100vh;
        background: #f3f6fb;
        color: #162033;
      }
      .page-header {
        display: flex;
        gap: 16px;
        justify-content: space-between;
        align-items: flex-start;
        margin-bottom: 18px;
      }
      .eyebrow {
        margin: 0 0 8px;
        color: #4f6b95;
        font-size: 0.82rem;
        text-transform: uppercase;
        letter-spacing: 0.14em;
        font-weight: 700;
      }
      .subtitle {
        margin: 10px 0 0;
        color: #60708d;
        max-width: 720px;
      }
      .tabs {
        display: flex;
        gap: 10px;
        margin: 12px 0;
      }
      .tab-btn {
        border: 1px solid #dbe4f0;
        background: #ffffff;
        color: #162033;
        border-radius: 14px;
        padding: 10px 14px;
        cursor: pointer;
        font-weight: 700;
      }
      .tab-btn.active {
        background: #162033;
        border-color: #162033;
        color: #ffffff;
      }
      .card {
        background: #ffffff;
        border: 1px solid #dbe4f0;
        border-radius: 20px;
        box-shadow: 0 16px 40px rgba(18, 38, 63, 0.08);
        padding: 20px;
      }
      .form-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
        gap: 16px;
      }
      .field {
        display: flex;
        flex-direction: column;
        gap: 8px;
        color: #31415f;
        font-weight: 600;
      }
      input[type='text'],
      input[type='number'],
      input[type='password'] {
        border: 1px solid #cfd9e6;
        border-radius: 14px;
        padding: 12px 14px;
        font: inherit;
        background: #ffffff;
        color: #162033;
      }
      .checkbox {
        flex-direction: row;
        align-items: center;
      }
      .actions-row {
        display: flex;
        gap: 16px;
        justify-content: space-between;
        align-items: center;
        margin-top: 16px;
      }
      .status {
        color: #60708d;
        font-weight: 700;
      }
      .primary-btn {
        background: #1ea96b;
        color: #ffffff;
        border: none;
        border-radius: 12px;
        padding: 11px 16px;
        font: inherit;
        font-weight: 700;
        cursor: pointer;
      }
      .secondary-link {
        color: #1f66d1;
        text-decoration: none;
        font-weight: 700;
      }
    `,
  ],
})
export class ProviderSettingsComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly entityCrudService = inject(EntityCrudService);

  providerId = 0;
  activeTab: ProtocolKey = 'http';
  saveMessage = '';

  // Local UI model (stored in provider_settings.setting_value as JSON)
  http: HttpSettings = { baseUrl: '', username: '', password: '' };
  smtp: SmtpSettings = { host: '', port: 587, username: '', password: '', useTls: true };
  pop: PopSettings = { host: '', port: 110, username: '', password: '', useSsl: false };

  // Keep record IDs for updates
  private settingIdByProtocol: Partial<Record<ProtocolKey, number>> = {};

  ngOnInit(): void {
    const raw = this.route.snapshot.params['providerId'];
    this.providerId = Number(raw);
    if (!Number.isFinite(this.providerId) || this.providerId <= 0) {
      return;
    }

    this.loadSettings();
  }

  private loadSettings(): void {
    this.saveMessage = '';
    this.entityCrudService.list('/provider-settings', 'provider_id', this.providerId).subscribe({
      next: (rows) => {
        for (const row of rows) {
          const settingKey = String(row['setting_key'] ?? '');
          const settingValue = String(row['setting_value'] ?? '');
          const id = Number(row['id']);

          if (!settingKey) continue;

          if (settingKey === 'http' || settingKey === 'smtp' || settingKey === 'pop') {
            const protocol = settingKey as ProtocolKey;
            if (Number.isFinite(id) && id > 0) {
              this.settingIdByProtocol[protocol] = id;
            }

            // setting_value is expected to be JSON
            if (settingValue.trim()) {
              try {
                const parsed = JSON.parse(settingValue) as unknown;
                if (protocol === 'http') this.http = { ...this.http, ...(parsed as Partial<HttpSettings>) };
                if (protocol === 'smtp') this.smtp = { ...this.smtp, ...(parsed as Partial<SmtpSettings>) };
                if (protocol === 'pop') this.pop = { ...this.pop, ...(parsed as Partial<PopSettings>) };
              } catch {
                // If old data is not JSON, ignore silently (won't break saving)
              }
            }
          }
        }
      },
      error: () => {
        this.saveMessage = 'Unable to load provider settings.';
      },
    });
  }

  private getActivePayload(): { settingKey: ProtocolKey; settingValue: string; isSensitive: boolean } {
    if (this.activeTab === 'http') {
      const isSensitive = Boolean(this.http.password && this.http.password.trim());
      return { settingKey: 'http', settingValue: JSON.stringify(this.http), isSensitive };
    }
    if (this.activeTab === 'smtp') {
      const isSensitive = Boolean(this.smtp.password && this.smtp.password.trim());
      return { settingKey: 'smtp', settingValue: JSON.stringify(this.smtp), isSensitive };
    }

    const isSensitive = Boolean(this.pop.password && this.pop.password.trim());
    return { settingKey: 'pop', settingValue: JSON.stringify(this.pop), isSensitive };
  }

  saveActiveTab(): void {
    if (this.providerId <= 0) return;

    const payload = this.getActivePayload();
    const existingId = this.settingIdByProtocol[payload.settingKey];

    this.saveMessage = 'Saving...';

    if (existingId && existingId > 0) {
      this.entityCrudService.update('/provider-settings', existingId, {
        setting_key: payload.settingKey,
        setting_value: payload.settingValue,
        is_sensitive: payload.isSensitive,
      }).subscribe({
        next: () => (this.saveMessage = `${payload.settingKey.toUpperCase()} saved successfully.`),
        error: () => (this.saveMessage = 'Save failed.'),
      });
      return;
    }

    this.entityCrudService.create('/provider-settings', {
      provider_id: this.providerId,
      setting_key: payload.settingKey,
      setting_value: payload.settingValue,
      is_sensitive: payload.isSensitive,
      created_by: 'admin',
    }).subscribe({
      next: () => (this.saveMessage = `${payload.settingKey.toUpperCase()} created successfully.`),
      error: () => (this.saveMessage = 'Save failed.'),
    });
  }
}

