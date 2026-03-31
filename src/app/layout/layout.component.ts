import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet, RouterLink, RouterLinkActive } from '@angular/router';

import { ENTITY_SECTIONS } from '../core/config/entity-sections.config';

@Component({
  selector: 'app-layout',
  standalone: true,
  imports: [CommonModule, RouterOutlet, RouterLink, RouterLinkActive],
  template: `
    <div class="app-container">
      <aside class="sidebar">
        <div class="sidebar-header">
          <p class="sidebar-kicker">Notification Admin</p>
          <h1>Notification System</h1>
        </div>

        <nav class="nav-menu">
          <a routerLink="/dashboard" routerLinkActive="active" [routerLinkActiveOptions]="{ exact: true }" class="nav-item">
            <span class="icon">DB</span>
            <span>Dashboard</span>
          </a>

          <button type="button" class="nav-master" (click)="toggleSettings()" aria-expanded="{{ settingsExpanded }}">
            <span class="icon">ST</span>
            <span>Settings</span>
          </button>

          <div *ngIf="settingsExpanded" class="nav-submenu">
            <div *ngFor="let group of settingsGroups" class="nav-group">
              <div class="nav-group-title">{{ group.title }}</div>

              <a
                *ngFor="let route of group.routes"
                [routerLink]="'/' + route"
                routerLinkActive="active"
                class="nav-item"
              >
                <span class="icon">{{ getInitials(sectionByRoute(route)?.title || route) }}</span>
                <span>{{ sectionByRoute(route)?.title || route }}</span>
              </a>
            </div>
          </div>
        </nav>
      </aside>

      <main class="main-content">
        <router-outlet></router-outlet>
      </main>
    </div>
  `,
  styles: [`
    .app-container {
      display: flex;
      min-height: 100vh;
      background: #f3f6fb;
    }

    .sidebar {
      width: 280px;
      background: linear-gradient(180deg, #1b2537 0%, #253753 100%);
      color: #ffffff;
      box-shadow: 16px 0 40px rgba(18, 38, 63, 0.12);
      overflow-y: auto;
      padding-bottom: 24px;
    }

    .sidebar-header {
      padding: 28px 24px 20px;
      border-bottom: 1px solid rgba(255, 255, 255, 0.08);
    }

    .sidebar-kicker {
      margin: 0 0 8px;
      color: #9fb6d6;
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.14em;
      font-weight: 700;
    }

    .sidebar-header h1 {
      margin: 0;
      font-size: 24px;
      line-height: 1.25;
      font-weight: 700;
    }

    .nav-menu {
      display: flex;
      flex-direction: column;
      gap: 6px;
      padding: 18px 12px 0;
    }

    .nav-master {
      display: flex;
      align-items: center;
      gap: 12px;
      padding: 12px 14px;
      color: #dfe7f3;
      text-decoration: none;
      border: 1px solid transparent;
      border-radius: 14px;
      background: transparent;
      cursor: pointer;
      font: inherit;
      font-weight: 700;
      text-align: left;
      transition: all 0.2s ease;
    }

    .nav-master:hover {
      background: rgba(255, 255, 255, 0.08);
      border-color: rgba(255, 255, 255, 0.08);
    }

    .nav-submenu {
      padding-left: 2px;
      display: flex;
      flex-direction: column;
      gap: 12px;
    }

    .nav-group-title {
      padding: 0 10px;
      color: #9fb6d6;
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 0.14em;
      font-weight: 700;
    }

    .nav-item {
      display: flex;
      align-items: center;
      gap: 12px;
      padding: 12px 14px;
      color: #dfe7f3;
      text-decoration: none;
      border: 1px solid transparent;
      border-radius: 14px;
      transition: all 0.2s ease;
    }

    .nav-item:hover {
      background: rgba(255, 255, 255, 0.08);
      border-color: rgba(255, 255, 255, 0.08);
    }

    .nav-item.active {
      background: #ffffff;
      border-color: #ffffff;
      color: #162033;
      box-shadow: 0 10px 24px rgba(9, 30, 66, 0.18);
    }

    .nav-item.active .icon {
      background: #162033;
      color: #ffffff;
    }

    .icon {
      width: 32px;
      height: 32px;
      border-radius: 10px;
      display: inline-flex;
      align-items: center;
      justify-content: center;
      background: rgba(255, 255, 255, 0.1);
      font-size: 11px;
      font-weight: 700;
      letter-spacing: 0.08em;
      flex: 0 0 auto;
    }

    .main-content {
      flex: 1;
      overflow-y: auto;
      background: #f3f6fb;
    }

    @media (max-width: 900px) {
      .app-container {
        flex-direction: column;
      }

      .sidebar {
        width: 100%;
      }
    }
  `]
})
export class LayoutComponent {
  protected settingsExpanded = true;

  protected readonly settingsGroups = [
    {
      title: 'Master Tables',
      routes: ['languages', 'priorities', 'statuses', 'schedule-types', 'categories'],
    },
    {
      title: 'Channels',
      routes: ['channels', 'channel-providers'],
    },
    {
      title: 'Templates & Routing',
      routes: ['template-groups', 'templates', 'routing-rules', 'mappings'],
    },
  ];

  protected sectionByRoute(route: string) {
    return ENTITY_SECTIONS.find((s) => s.route === route);
  }

  protected toggleSettings(): void {
    this.settingsExpanded = !this.settingsExpanded;
  }

  protected getInitials(title: string): string {
    return title
      .split(' ')
      .map((part) => part[0])
      .join('')
      .slice(0, 2)
      .toUpperCase();
  }
}
