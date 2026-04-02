import { CommonModule } from '@angular/common';
import { Component, ElementRef, EventEmitter, Input, Output, ViewChild } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

import { PUSH_TEMPLATE_VARIABLES } from '../template-mock-data';
import { PushNotificationTemplate, PushTemplatePriority } from '../template-builder.types';
import { ImagePreviewInputComponent } from './image-preview-input.component';
import { VariableChipsComponent } from './variable-chips.component';
import { ToggleSwitchComponent } from '../shared/toggle-switch.component';

export interface PushTemplateFormValue {
  templateName: string;
  channel: 'Push Notification';
  title: string;
  message: string;
  targetUrl: string;
  iconUrl: string;
  bannerUrl: string;
  requireInteraction: boolean;
  priority: PushTemplatePriority;
  ttl: number;
}

@Component({
  selector: 'app-push-template-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, ImagePreviewInputComponent, VariableChipsComponent, ToggleSwitchComponent],
  template: `
    <form class="form" [formGroup]="form" (ngSubmit)="submit.emit()">
      <div class="grid">
        <label class="field">
          <span class="label">Template Name</span>
          <input class="input" type="text" formControlName="templateName" placeholder="e.g. Order Shipped" />
          <small class="error" *ngIf="form.controls.templateName.touched && form.controls.templateName.invalid">
            Template name is required
          </small>
        </label>

        <label class="field">
          <span class="label">Channel</span>
          <input class="input" type="text" formControlName="channel" [disabled]="true" />
        </label>

        <label class="field">
          <span class="label">Title</span>
          <input
            #titleInput
            class="input"
            type="text"
            formControlName="title"
            placeholder="e.g. Your order is shipped"
            (focus)="activeField = 'title'"
          />
        </label>

        <label class="field">
          <span class="label">Message</span>
          <textarea
            #messageInput
            class="textarea"
            formControlName="message"
            rows="6"
            placeholder="Write a short notification message..."
            (focus)="activeField = 'message'"
          ></textarea>
        </label>

        <div class="variables">
          <div class="variables-head">
            <div>
              <div class="variables-title">Insert Variables</div>
              <div class="variables-subtitle">Click to insert into Title / Message</div>
            </div>
          </div>
          <app-variable-chips [variables]="variables" (insert)="insertVariable($event)"></app-variable-chips>
        </div>

        <label class="field">
          <span class="label">Target URL</span>
          <input class="input" type="url" formControlName="targetUrl" placeholder="https://example.com/path" />
        </label>

        <app-image-preview-input label="Icon Image URL" placeholder="https://..." [control]="form.controls.iconUrl"></app-image-preview-input>
        <app-image-preview-input label="Banner Image URL" placeholder="https://..." [control]="form.controls.bannerUrl"></app-image-preview-input>

        <div class="row">
          <div class="field">
            <span class="label">Require Interaction</span>
            <div class="switch-row">
              <app-toggle-switch [checked]="form.controls.requireInteraction.value" (checkedChange)="form.controls.requireInteraction.setValue($event)"></app-toggle-switch>
              <span class="switch-label">{{ form.controls.requireInteraction.value ? 'On' : 'Off' }}</span>
            </div>
          </div>

          <label class="field">
            <span class="label">Priority</span>
            <select class="input" formControlName="priority">
              <option value="High">High</option>
              <option value="Normal">Normal</option>
            </select>
          </label>

          <label class="field">
            <span class="label">TTL (seconds)</span>
            <input class="input" type="number" min="0" formControlName="ttl" />
          </label>
        </div>
      </div>
    </form>
  `,
  styleUrls: ['./push-template-form.component.scss']
})
export class PushTemplateFormComponent {
  @Input() mode: 'create' | 'edit' | 'view' = 'create';
  @Output() submit = new EventEmitter<void>();

  @ViewChild('titleInput') titleInput?: ElementRef<HTMLInputElement>;
  @ViewChild('messageInput') messageInput?: ElementRef<HTMLTextAreaElement>;

  activeField: 'title' | 'message' = 'title';
  readonly variables = PUSH_TEMPLATE_VARIABLES;

  readonly form = new FormGroup({
    templateName: new FormControl<string>('', { nonNullable: true, validators: [Validators.required] }),
    channel: new FormControl<'Push Notification'>('Push Notification', { nonNullable: true }),
    title: new FormControl<string>('', { nonNullable: true }),
    message: new FormControl<string>('', { nonNullable: true, validators: [Validators.required] }),
    targetUrl: new FormControl<string>('', { nonNullable: true }),
    iconUrl: new FormControl<string>('', { nonNullable: true }),
    bannerUrl: new FormControl<string>('', { nonNullable: true }),
    requireInteraction: new FormControl<boolean>(false, { nonNullable: true }),
    priority: new FormControl<PushTemplatePriority>('Normal', { nonNullable: true }),
    ttl: new FormControl<number>(3600, { nonNullable: true })
  });

  constructor() {
    // Always fixed for push templates.
    this.form.controls.channel.disable({ emitEvent: false });
  }

  setDisabled(disabled: boolean): void {
    if (disabled) {
      this.form.disable({ emitEvent: false });
    } else {
      this.form.enable({ emitEvent: false });
      this.form.controls.channel.disable({ emitEvent: false });
    }
  }

  patchFromTemplate(t?: PushNotificationTemplate | null): void {
    if (!t) {
      this.form.reset(
        {
          templateName: '',
          channel: 'Push Notification',
          title: '',
          message: '',
          targetUrl: '',
          iconUrl: '',
          bannerUrl: '',
          requireInteraction: false,
          priority: 'Normal',
          ttl: 3600
        },
        { emitEvent: false }
      );
      return;
    }
    this.form.patchValue(
      {
        templateName: t.templateName,
        channel: 'Push Notification',
        title: t.title,
        message: t.message,
        targetUrl: t.targetUrl,
        iconUrl: t.iconUrl,
        bannerUrl: t.bannerUrl,
        requireInteraction: t.requireInteraction,
        priority: t.priority,
        ttl: t.ttl
      },
      { emitEvent: false }
    );
  }

  insertVariable(token: string): void {
    const target =
      this.activeField === 'message' ? this.messageInput?.nativeElement : this.titleInput?.nativeElement;
    if (!target) return;

    const start = target.selectionStart ?? target.value.length;
    const end = target.selectionEnd ?? target.value.length;
    const next = target.value.slice(0, start) + token + target.value.slice(end);

    if (this.activeField === 'message') {
      this.form.controls.message.setValue(next);
    } else {
      this.form.controls.title.setValue(next);
    }

    // restore caret after insertion
    queueMicrotask(() => {
      target.focus();
      const pos = start + token.length;
      target.setSelectionRange(pos, pos);
    });
  }
}

