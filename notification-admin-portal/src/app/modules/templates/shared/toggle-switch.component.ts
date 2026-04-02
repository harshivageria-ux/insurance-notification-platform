import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-toggle-switch',
  standalone: true,
  imports: [CommonModule],
  template: `
    <button
      type="button"
      class="toggle"
      [class.toggle-on]="checked"
      [class.toggle-disabled]="disabled"
      [attr.aria-pressed]="checked"
      [disabled]="disabled"
      (click)="onToggle()"
    >
      <span class="thumb"></span>
    </button>
  `,
  styleUrls: ['./toggle-switch.component.scss']
})
export class ToggleSwitchComponent {
  @Input() checked = false;
  @Input() disabled = false;
  @Output() checkedChange = new EventEmitter<boolean>();

  onToggle(): void {
    if (this.disabled) return;
    this.checkedChange.emit(!this.checked);
  }
}

