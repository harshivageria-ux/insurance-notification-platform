import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-image-preview-input',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  template: `
    <div class="image-field">
      <label class="label">{{ label }}</label>
      <input class="input" type="url" [placeholder]="placeholder" [formControl]="control" />
      <div class="preview" *ngIf="control.value">
        <img [src]="control.value" [alt]="label" (error)="hasError = true" (load)="hasError = false" />
        <div class="preview-error" *ngIf="hasError">Unable to load image</div>
      </div>
    </div>
  `,
  styleUrls: ['./image-preview-input.component.scss']
})
export class ImagePreviewInputComponent {
  @Input({ required: true }) label = '';
  @Input() placeholder = 'https://...';
  @Input({ required: true }) control!: FormControl<string>;
  hasError = false;
}

