import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output, ViewChild } from '@angular/core';
import { PushNotificationTemplate } from '../template-builder.types';
import { PushTemplateFormComponent } from './push-template-form.component';
import { NotificationPreviewComponent } from './notification-preview.component';

@Component({
  selector: 'app-template-drawer',
  standalone: true,
  imports: [CommonModule, PushTemplateFormComponent, NotificationPreviewComponent],
  template: `
    <div class="backdrop" *ngIf="open" (click)="requestClose.emit()"></div>

    <aside class="drawer" [class.drawer-open]="open" (click)="$event.stopPropagation()">
      <header class="drawer-header">
        <div>
          <div class="kicker">Push Notification Builder</div>
          <div class="title">{{ modeTitle }}</div>
        </div>

        <button type="button" class="icon-btn" (click)="requestClose.emit()" aria-label="Close drawer">×</button>
      </header>

      <div class="drawer-body">
        <div class="split">
          <section class="panel">
            <div class="panel-title">Template</div>
            <app-push-template-form #formCmp></app-push-template-form>
          </section>

          <section class="panel">
            <app-notification-preview
              [iconUrl]="formCmp.form.controls.iconUrl.value"
              [bannerUrl]="formCmp.form.controls.bannerUrl.value"
              [title]="formCmp.form.controls.title.value"
              [message]="formCmp.form.controls.message.value"
              [priority]="formCmp.form.controls.priority.value"
              [ttl]="formCmp.form.controls.ttl.value"
              [requireInteraction]="formCmp.form.controls.requireInteraction.value"
            ></app-notification-preview>
          </section>
        </div>
      </div>

      <footer class="drawer-footer">
        <button type="button" class="secondary-btn" (click)="requestClose.emit()">Cancel</button>
        <button type="button" class="primary-btn" [disabled]="mode === 'view' || formCmp.form.invalid" (click)="saveClicked()">
          {{ mode === 'edit' ? 'Update Template' : 'Save Template' }}
        </button>
      </footer>
    </aside>
  `,
  styleUrls: ['./template-drawer.component.scss']
})
export class TemplateDrawerComponent {
  @Input() open = false;
  @Input() mode: 'create' | 'edit' | 'view' = 'create';
  @Input() template: PushNotificationTemplate | null = null;
  @Output() requestClose = new EventEmitter<void>();
  @Output() save = new EventEmitter<{ mode: 'create' | 'edit'; value: any }>();

  @ViewChild('formCmp') formCmp?: PushTemplateFormComponent;

  get modeTitle(): string {
    if (this.mode === 'edit') return 'Edit Template';
    if (this.mode === 'view') return 'View Template';
    return 'Create Template';
  }

  ngOnChanges(): void {
    // ensure form sync when inputs change
    queueMicrotask(() => {
      if (!this.formCmp) return;
      this.formCmp.patchFromTemplate(this.template);
      this.formCmp.setDisabled(this.mode === 'view');
    });
  }

  saveClicked(): void {
    const form = this.formCmp?.form;
    if (!form) return;
    form.markAllAsTouched();
    if (form.invalid) return;
    this.save.emit({ mode: this.mode === 'edit' ? 'edit' : 'create', value: form.getRawValue() });
  }
}

