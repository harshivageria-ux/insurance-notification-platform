import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';

import { PushTemplateVariable } from '../template-builder.types';

@Component({
  selector: 'app-variable-chips',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="chips">
      <button
        type="button"
        class="chip"
        *ngFor="let v of variables"
        (click)="insert.emit(v.token)"
        [attr.aria-label]="'Insert variable ' + v.token"
      >
        {{ v.token }}
      </button>
    </div>
  `,
  styleUrls: ['./variable-chips.component.scss']
})
export class VariableChipsComponent {
  @Input() variables: PushTemplateVariable[] = [];
  @Output() insert = new EventEmitter<string>();
}

