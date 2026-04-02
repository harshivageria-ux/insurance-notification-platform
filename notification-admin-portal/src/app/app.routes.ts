import { Routes } from '@angular/router';
import { LayoutComponent } from './layout/layout.component';
import { DashboardComponent } from './modules/dashboard/dashboard.component';
import { EntityManagementComponent } from './modules/entity-management/entity-management.component';
import { ENTITY_SECTION_MAP } from './core/config/entity-sections.config';
import { ProviderSettingsComponent } from './modules/provider-settings/provider-settings.component';
import { TemplatePageComponent } from './modules/templates/template-page.component';

export const routes: Routes = [
  {
    path: '',
    component: LayoutComponent,
    children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: 'dashboard', component: DashboardComponent },
      { path: 'languages', component: EntityManagementComponent, data: { section: ENTITY_SECTION_MAP['languages'] } },
      { path: 'priorities', component: EntityManagementComponent, data: { section: ENTITY_SECTION_MAP['priorities'] } },
      { path: 'statuses', component: EntityManagementComponent, data: { section: ENTITY_SECTION_MAP['statuses'] } },
      { path: 'schedule-types', component: EntityManagementComponent, data: { section: ENTITY_SECTION_MAP['schedule-types'] } },
      { path: 'categories', component: EntityManagementComponent, data: { section: ENTITY_SECTION_MAP['categories'] } },
      { path: 'channels', component: EntityManagementComponent, data: { section: ENTITY_SECTION_MAP['channels'] } },
      { path: 'channel-providers', component: EntityManagementComponent, data: { section: ENTITY_SECTION_MAP['channel-providers'] } },
      { path: 'channel-providers/:providerId/provider-settings', component: ProviderSettingsComponent },
      { path: 'template-groups', component: EntityManagementComponent, data: { section: ENTITY_SECTION_MAP['template-groups'] } },
      { path: 'templates', component: TemplatePageComponent },
      { path: 'routing-rules', component: EntityManagementComponent, data: { section: ENTITY_SECTION_MAP['routing-rules'] } }
    ]
  }
];
