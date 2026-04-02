import { CommonModule } from '@angular/common';
import { Component, OnInit, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, RouterLink } from '@angular/router';

import { CrudFieldConfig, CrudFieldOption, EntityIdentifier, EntityRecord, EntitySectionConfig } from '../../core/models/entity-crud.model';
import { EntityCrudService } from '../../services/entity-crud.service';

@Component({
  selector: 'app-entity-management',
  standalone: true,
  imports: [CommonModule, FormsModule, RouterLink],
  template: `
    <section class="page-shell" *ngIf="section">
      <header class="page-header">
        <div>
          <p class="eyebrow">Admin Portal</p>
          <h1>{{ section.title }}</h1>
          <p class="subtitle">{{ section.subtitle }}</p>
        </div>
      </header>

      <div class="stats-grid">
        <article class="stat-card">
          <span class="stat-label">Total Records</span>
          <strong>{{ filteredItems.length }}</strong>
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
        <input
          type="text"
          class="search-input"
          [placeholder]="section.searchPlaceholder || 'Search...'"
          [(ngModel)]="searchText"
          (ngModelChange)="currentPage = 1"
        />

        <div class="toolbar-actions" *ngIf="section.listParamKey">
          <ng-container *ngIf="listParamField?.type === 'select'; else textFilter">
            <select class="filter-input" [(ngModel)]="listParamValue">
              <option [ngValue]="''">{{ section.listParamPlaceholder || 'Select' }}</option>
              <option *ngFor="let option of getFieldOptions(listParamField!)" [ngValue]="option.value">{{ option.label }}</option>
            </select>
          </ng-container>
          <ng-template #textFilter>
            <input
              type="text"
              class="filter-input"
              [placeholder]="section.listParamPlaceholder || ''"
              [(ngModel)]="listParamValue"
            />
          </ng-template>
          <button class="secondary-btn" (click)="loadItems()">Load {{ section.listParamLabel || 'Data' }}</button>
        </div>

        <button class="primary-btn" (click)="openCreateModal()">Add {{ singularTitle }}</button>
      </div>

      <div class="table-card">
        <div class="empty-state" *ngIf="errorMessage">{{ errorMessage }}</div>

        <div class="table-wrap" *ngIf="!errorMessage">
          <table>
            <thead>
              <tr>
                <th *ngFor="let field of tableFields" [style.width]="field.width || null">{{ field.label }}</th>
                <th *ngIf="section.showActiveBadgeKey">Active</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr *ngFor="let item of paginatedItems; let rowIndex = index">
                <td *ngFor="let field of tableFields">
                  <span [class.clamped]="field.type === 'textarea'">{{ displayValue(item, field) }}</span>
                </td>
                <td *ngIf="section.showActiveBadgeKey">
                  <span class="badge" [class.badge-active]="!!item[section.showActiveBadgeKey!]" [class.badge-inactive]="!item[section.showActiveBadgeKey!]">
                    {{ item[section.showActiveBadgeKey!] ? 'Yes' : 'No' }}
                  </span>
                </td>
                <td class="actions-cell">
                  <button class="ghost-btn" (click)="openEditModal(item)">Edit</button>
                  <a
                    *ngIf="section.rowDetailRoute && section.rowDetailLabel"
                    class="ghost-btn"
                    [routerLink]="buildRowDetailRoute(item)"
                  >
                    {{ section.rowDetailLabel }}
                  </a>
                  <button
                    *ngIf="section.supportsToggle && section.toggleEndpoint && section.showActiveBadgeKey"
                    class="ghost-btn"
                    (click)="toggleItem(item)"
                  >
                    {{ item[section.showActiveBadgeKey] ? 'Disable' : 'Enable' }}
                  </button>
                  <button class="danger-btn" (click)="deleteItem(item)">Delete</button>
                </td>
              </tr>
              <tr *ngIf="!paginatedItems.length">
                <td class="empty-row" [attr.colspan]="tableColumnCount">No records found.</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div class="pagination" *ngIf="filteredItems.length">
        <button class="secondary-btn" [disabled]="currentPage === 1" (click)="currentPage = currentPage - 1">Prev</button>
        <span>Page {{ currentPage }} of {{ totalPages }}</span>
        <button class="secondary-btn" [disabled]="currentPage === totalPages" (click)="currentPage = currentPage + 1">Next</button>
      </div>
    </section>

    <div class="modal-backdrop" *ngIf="showModal" (click)="closeModal()">
      <div class="modal-card" (click)="$event.stopPropagation()">
        <div class="modal-header">
          <h2>{{ isEditMode ? 'Edit' : 'Add' }} {{ singularTitle }}</h2>
          <button class="icon-btn" (click)="closeModal()">x</button>
        </div>

        <div class="form-grid">
          <label class="field" *ngFor="let field of formFields">
            <span>{{ field.label }}</span>

            <input
              *ngIf="field.type === 'text' || field.type === 'number'"
              [type]="field.type"
              [placeholder]="field.placeholder || ''"
              [(ngModel)]="formModel[field.key]"
              (ngModelChange)="onFieldValueChange(field)"
              [class.input-error]="fieldErrors[field.key]"
            />

            <textarea
              *ngIf="field.type === 'textarea'"
              rows="7"
              [placeholder]="field.placeholder || ''"
              [(ngModel)]="formModel[field.key]"
              (ngModelChange)="onFieldValueChange(field)"
              [class.input-error]="fieldErrors[field.key]"
            ></textarea>

            <select *ngIf="field.type === 'select'" [(ngModel)]="formModel[field.key]" (ngModelChange)="onFieldValueChange(field)" [class.input-error]="fieldErrors[field.key]">
              <option [ngValue]="''">Select {{ field.label }}</option>
              <option *ngFor="let option of getFieldOptions(field)" [ngValue]="option.value">{{ option.label }}</option>
            </select>

            <small class="field-error" *ngIf="fieldErrors[field.key]">{{ fieldErrors[field.key] }}</small>
          </label>
        </div>

        <p class="modal-error" *ngIf="modalErrorMessage">{{ modalErrorMessage }}</p>

        <div class="modal-actions">
          <button class="secondary-btn" (click)="closeModal()">Cancel</button>
          <button class="primary-btn" (click)="saveItem()">{{ isEditMode ? 'Update' : 'Save' }}</button>
        </div>
      </div>
    </div>

    <div class="modal-backdrop" *ngIf="showDeleteConfirm" (click)="closeDeleteConfirm()">
      <div class="modal-card confirm-card" (click)="$event.stopPropagation()">
        <div class="modal-header">
          <h2>Delete {{ singularTitle }}</h2>
          <button class="icon-btn" (click)="closeDeleteConfirm()">x</button>
        </div>

        <p class="confirm-message">
          This will deactivate the record (soft delete). You can re-enable it later from the toggle action.
        </p>

        <div class="confirm-details" *ngIf="deleteCandidateLabel">
          <span class="confirm-pill">{{ deleteCandidateLabel }}</span>
        </div>

        <div class="modal-actions">
          <button class="secondary-btn" (click)="closeDeleteConfirm()">Cancel</button>
          <button class="danger-btn" (click)="confirmDelete()">Delete</button>
        </div>
      </div>
    </div>
  `,
  styles: [`
    :host {
      display: block;
    }

    .page-shell {
      padding: 32px;
      background: #f3f6fb;
      min-height: 100vh;
      color: #162033;
    }

    .page-header h1 {
      margin: 0;
      font-size: 2rem;
      font-weight: 700;
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

    .stats-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
      gap: 18px;
      margin: 24px 0;
    }

    .stat-card,
    .table-card,
    .modal-card {
      background: #ffffff;
      border: 1px solid #dbe4f0;
      border-radius: 20px;
      box-shadow: 0 16px 40px rgba(18, 38, 63, 0.08);
    }

    .stat-card {
      padding: 22px 24px;
    }

    .stat-label {
      display: block;
      color: #60708d;
      font-size: 0.92rem;
      margin-bottom: 8px;
    }

    .stat-card strong {
      font-size: 1.9rem;
    }

    .toolbar {
      display: flex;
      gap: 12px;
      align-items: center;
      flex-wrap: wrap;
      margin-bottom: 20px;
    }

    .toolbar-actions {
      display: flex;
      gap: 12px;
      align-items: center;
      flex-wrap: wrap;
    }

    .search-input,
    .filter-input,
    .field input,
    .field textarea,
    .field select {
      border: 1px solid #cfd9e6;
      border-radius: 14px;
      padding: 12px 14px;
      font: inherit;
      background: #ffffff;
      color: #162033;
    }

    .search-input {
      min-width: 260px;
      flex: 1 1 280px;
    }

    .filter-input {
      width: 180px;
    }

    .primary-btn,
    .secondary-btn,
    .ghost-btn,
    .danger-btn,
    .icon-btn {
      border: none;
      border-radius: 12px;
      padding: 11px 16px;
      font: inherit;
      font-weight: 600;
      cursor: pointer;
      transition: transform 0.15s ease, box-shadow 0.15s ease, background 0.15s ease;
    }

    .primary-btn {
      background: #1ea96b;
      color: #ffffff;
      box-shadow: 0 10px 24px rgba(30, 169, 107, 0.22);
    }

    .secondary-btn,
    .ghost-btn {
      background: #e9eef6;
      color: #20304f;
    }

    .danger-btn {
      background: #eb5757;
      color: #ffffff;
    }

    .icon-btn {
      background: transparent;
      color: #50607f;
      padding: 8px 10px;
    }

    .table-card {
      overflow: hidden;
    }

    .table-wrap {
      overflow-x: auto;
    }

    table {
      width: 100%;
      border-collapse: collapse;
    }

    thead {
      background: #162033;
      color: #ffffff;
    }

    th,
    td {
      padding: 18px 16px;
      text-align: left;
      border-bottom: 1px solid #e8eef5;
      vertical-align: top;
    }

    tbody tr:hover {
      background: #f8fbff;
    }

    .clamped {
      display: inline-block;
      max-width: 280px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .actions-cell {
      display: flex;
      gap: 8px;
      flex-wrap: wrap;
    }

    .badge {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      min-width: 76px;
      padding: 6px 10px;
      border-radius: 999px;
      font-size: 0.85rem;
      font-weight: 600;
    }

    .badge-active {
      background: #def7ea;
      color: #137a4a;
    }

    .badge-inactive {
      background: #fde7e7;
      color: #b13a3a;
    }

    .empty-row,
    .empty-state {
      text-align: center;
      color: #60708d;
      padding: 28px;
    }

    .pagination {
      display: flex;
      align-items: center;
      justify-content: flex-end;
      gap: 12px;
      margin-top: 18px;
    }

    .modal-backdrop {
      position: fixed;
      inset: 0;
      background: rgba(11, 18, 32, 0.56);
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 20px;
      z-index: 1000;
    }

    .modal-card {
      width: min(760px, 100%);
      padding: 24px;
    }

    .modal-header,
    .modal-actions {
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 12px;
    }

    .modal-header {
      margin-bottom: 18px;
    }

    .modal-header h2 {
      margin: 0;
    }

    .form-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
      gap: 16px;
      margin-bottom: 18px;
    }

    .field {
      display: flex;
      flex-direction: column;
      gap: 8px;
      color: #31415f;
      font-weight: 600;
    }

    .input-error {
      border-color: #d64545;
      background: #fff6f6;
    }

    .field-error,
    .modal-error {
      color: #c62828;
      font-size: 0.85rem;
      font-weight: 600;
      margin: 0;
    }

    .field textarea {
      resize: vertical;
      min-height: 180px;
    }

    .confirm-card {
      width: min(560px, 100%);
    }

    .confirm-message {
      margin: 8px 0 0;
      color: #60708d;
      line-height: 1.5;
      font-weight: 600;
    }

    .confirm-details {
      margin-top: 14px;
    }

    .confirm-pill {
      display: inline-flex;
      align-items: center;
      padding: 8px 12px;
      border-radius: 999px;
      background: #f1f6ff;
      border: 1px solid #dbe7ff;
      color: #20304f;
      font-weight: 800;
      max-width: 100%;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    @media (max-width: 900px) {
      .page-shell {
        padding: 20px;
      }

      .pagination {
        justify-content: flex-start;
        flex-wrap: wrap;
      }
    }
  `]
})
export class EntityManagementComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly entityCrudService = inject(EntityCrudService);
  private readonly hiddenItemIds = new Set<string>();

  section!: EntitySectionConfig;
  items: EntityRecord[] = [];
  formModel: EntityRecord = {};
  searchText = '';
  currentPage = 1;
  pageSize = 8;
  showModal = false;
  showDeleteConfirm = false;
  private deleteCandidateId: EntityIdentifier | null = null;
  deleteCandidateLabel = '';
  isEditMode = false;
  editingId: EntityIdentifier | null = null;
  errorMessage = '';
  modalErrorMessage = '';
  listParamValue: string | number | null = null;
  fieldOptions: Record<string, CrudFieldOption[]> = {};
  fieldErrors: Record<string, string> = {};

  ngOnInit(): void {
    this.section = this.route.snapshot.data['section'] as EntitySectionConfig;
    this.listParamValue = this.section.listParamDefault ?? null;
    this.resetForm();
    this.loadFieldOptions();
    this.loadItems();
  }

  get singularTitle(): string {
    return this.section.title.endsWith('s') ? this.section.title.slice(0, -1) : this.section.title;
  }

  get tableFields(): CrudFieldConfig[] {
    return this.section.fields.filter((field) => field.table !== false);
  }

  get listParamField(): CrudFieldConfig | undefined {
    return this.section.listParamKey
      ? this.section.fields.find((field) => field.key === this.section.listParamKey)
      : undefined;
  }

  get filteredItems(): EntityRecord[] {
    const search = this.searchText.trim().toLowerCase();
    if (!search) {
      return this.items;
    }

    return this.items.filter((item) =>
      this.section.fields.some((field) => {
        if (field.searchable === false) {
          return false;
        }

        const value = item[field.key];
        return value !== null && value !== undefined && String(value).toLowerCase().includes(search);
      })
    );
  }

  get paginatedItems(): EntityRecord[] {
    const startIndex = (this.currentPage - 1) * this.pageSize;
    return this.filteredItems.slice(startIndex, startIndex + this.pageSize);
  }

  get totalPages(): number {
    return Math.max(1, Math.ceil(this.filteredItems.length / this.pageSize));
  }

  get activeCount(): number {
    return this.items.filter((item) => this.isActiveItem(item)).length;
  }

  get inactiveCount(): number {
    return this.items.filter((item) => !this.isActiveItem(item)).length;
  }

  get tableColumnCount(): number {
    // columns: (table fields) + (optional Active badge) + (Actions)
    return this.tableFields.length + (this.section.showActiveBadgeKey ? 1 : 0) + 1;
  }

  get formFields(): CrudFieldConfig[] {
    return this.section.fields.filter((field) => field.showInForm !== false);
  }

  loadItems(): void {
    if (
      this.section.listParamKey &&
      (this.listParamValue === null || this.listParamValue === undefined || String(this.listParamValue).trim() === '')
    ) {
      this.items = [];
      this.errorMessage = `Enter a ${this.section.listParamLabel || 'filter value'} to load records.`;
      return;
    }

    this.errorMessage = '';
    this.entityCrudService.list(this.section.endpoint, this.section.listParamKey, this.listParamValue).subscribe({
      next: (items) => {
        this.items = items.filter((item) => !this.hiddenItemIds.has(this.toHiddenItemKey(this.getItemId(item))));
        this.currentPage = 1;
        if (!items.length) {
          this.errorMessage = '';
        }
      },
      error: () => {
        this.items = [];
        this.errorMessage = `Unable to load ${this.section.title.toLowerCase()}.`;
      }
    });
  }

  openCreateModal(): void {
    this.isEditMode = false;
    this.editingId = null;
    this.resetForm();
    this.fieldErrors = {};
    this.modalErrorMessage = '';
    this.showModal = true;
  }

  openEditModal(item: EntityRecord): void {
    this.isEditMode = true;
    this.editingId = this.getItemId(item);
    this.formModel = this.createFormModel(item);
    this.fieldErrors = {};
    this.modalErrorMessage = '';
    this.showModal = true;
  }

  closeModal(): void {
    this.showModal = false;
    this.fieldErrors = {};
    this.modalErrorMessage = '';
    this.resetForm();
  }

  saveItem(): void {
    this.fieldErrors = this.validateForm();
    if (Object.keys(this.fieldErrors).length > 0) {
      this.modalErrorMessage = 'Please fix the highlighted fields before saving.';
      return;
    }

    let payload: EntityRecord;

    try {
      payload = this.normalizePayload({ ...this.formModel });
    } catch {
      this.modalErrorMessage = 'Please enter valid JSON for the JSON fields before saving.';
      return;
    }

    this.modalErrorMessage = '';

    if (this.isEditMode && this.editingId !== null) {
      this.entityCrudService.update(this.section.endpoint, this.editingId, payload).subscribe({
        next: () => {
          this.closeModal();
          this.loadItems();
        },
        error: () => {
          this.errorMessage = `Unable to update ${this.singularTitle.toLowerCase()}.`;
        }
      });
      return;
    }

    this.entityCrudService.create(this.section.endpoint, payload).subscribe({
      next: () => {
        this.closeModal();
        this.loadItems();
      },
      error: () => {
        this.errorMessage = `Unable to create ${this.singularTitle.toLowerCase()}.`;
      }
    });
  }

  deleteItem(item: EntityRecord): void {
    const id = this.getItemId(item);
    if (id === null || id === undefined || id === '') {
      return;
    }
    this.deleteCandidateId = id;
    this.deleteCandidateLabel = this.buildDeleteLabel(item);
    this.showDeleteConfirm = true;
  }

  closeDeleteConfirm(): void {
    this.showDeleteConfirm = false;
    this.deleteCandidateId = null;
    this.deleteCandidateLabel = '';
  }

  confirmDelete(): void {
    if (this.deleteCandidateId === null || this.deleteCandidateId === undefined || this.deleteCandidateId === '') {
      this.closeDeleteConfirm();
      return;
    }

    const id = this.deleteCandidateId;
    this.closeDeleteConfirm();

    // Soft-delete via backend DELETE if available; otherwise hide locally.
    this.entityCrudService.delete(this.section.endpoint, id).subscribe({
      next: () => {
        this.hiddenItemIds.add(this.toHiddenItemKey(id));
        this.items = this.items.filter((existingItem) => this.getItemId(existingItem) !== id);
        this.currentPage = Math.min(this.currentPage, this.totalPages);
      },
      error: () => {
        // Fallback: hide locally so UI continues smoothly.
        this.hiddenItemIds.add(this.toHiddenItemKey(id));
        this.items = this.items.filter((existingItem) => this.getItemId(existingItem) !== id);
        this.currentPage = Math.min(this.currentPage, this.totalPages);
      }
    });
  }

  private buildDeleteLabel(item: EntityRecord): string {
    const parts: string[] = [];
    const primaryKeys = ['name', 'code', 'priority_code', 'status_code', 'schedule_code', 'setting_key', 'title_template'];
    for (const key of primaryKeys) {
      const value = item[key];
      if (value !== null && value !== undefined && String(value).trim() !== '') {
        parts.push(String(value).trim());
      }
      if (parts.length >= 2) break;
    }
    if (parts.length) return parts.join(' · ');
    return `${this.singularTitle} #${String(this.getItemId(item))}`;
  }

  toggleItem(item: EntityRecord): void {
    if (!this.section.toggleEndpoint || !this.section.showActiveBadgeKey) {
      return;
    }

    const id = this.getItemId(item);
    const nextState = !Boolean(item[this.section.showActiveBadgeKey]);

    this.entityCrudService.toggle(this.section.toggleEndpoint, id, nextState).subscribe({
      next: () => this.loadItems(),
      error: () => {
        this.errorMessage = `Unable to toggle ${this.singularTitle.toLowerCase()}.`;
      }
    });
  }

  displayValue(item: EntityRecord, field: CrudFieldConfig): string {
    const value = item[field.key];
    if (value === null || value === undefined || value === '') {
      return '-';
    }

    if (this.isJsonField(field.key)) {
      return typeof value === 'string' ? value : JSON.stringify(value);
    }

    if (typeof value === 'boolean') {
      return value ? 'Yes' : 'No';
    }

    const matchedOption = this.getFieldOptions(field).find((option) => option.value === value);
    if (matchedOption) {
      return matchedOption.label;
    }

    return String(value);
  }

  getFieldOptions(field: CrudFieldConfig): CrudFieldOption[] {
    return this.fieldOptions[field.key] || field.options || [];
  }

  private resetForm(): void {
    this.formModel = { ...this.section.defaultItem };
    if (this.section.listParamKey && this.listParamValue !== null && this.listParamValue !== '') {
      this.formModel[this.section.listParamKey] = this.listParamValue;
    }
  }

  onFieldValueChange(field: CrudFieldConfig): void {
    if (field.uppercase && typeof this.formModel[field.key] === 'string') {
      this.formModel[field.key] = String(this.formModel[field.key]).toUpperCase();
    }

    if (this.fieldErrors[field.key]) {
      delete this.fieldErrors[field.key];
      this.fieldErrors = { ...this.fieldErrors };
    }

    if (this.modalErrorMessage) {
      this.modalErrorMessage = '';
    }
  }

  private normalizePayload(payload: EntityRecord): EntityRecord {
    for (const field of this.formFields) {
      if ((field.type === 'text' || field.type === 'textarea') && typeof payload[field.key] === 'string') {
        const normalized = String(payload[field.key]).trim();
        payload[field.key] = field.uppercase ? normalized.toUpperCase() : normalized;
      }

      if (field.type === 'number') {
        payload[field.key] = Number(payload[field.key] ?? 0);
      }

      if (this.isJsonField(field.key) && typeof payload[field.key] === 'string') {
        const rawValue = String(payload[field.key] || '').trim();
        payload[field.key] = rawValue ? JSON.parse(rawValue) : {};
      }
    }

    delete payload['id'];
    delete payload['is_active'];
    delete payload['created_at'];
    delete payload['updated_at'];
    delete payload['deleted_at'];
    delete payload['version'];

    if (this.section.idKey && this.section.idKey !== 'id') {
      delete payload[this.section.idKey];
    }

    if (this.isEditMode && this.section.updatePayload) {
      return this.section.updatePayload(payload, this.formModel);
    }

    if (!this.isEditMode && this.section.createPayload) {
      return this.section.createPayload(payload, this.formModel);
    }

    return payload;
  }

  private validateForm(): Record<string, string> {
    const errors: Record<string, string> = {};

    for (const field of this.formFields) {
      const value = this.formModel[field.key];
      const isEmpty = value === null || value === undefined || value === '';

      if (field.required && isEmpty) {
        errors[field.key] = `${field.label} is required.`;
        continue;
      }

      if (isEmpty) {
        continue;
      }

      if (field.type === 'text' || field.type === 'textarea') {
        const normalized = String(value).trim();
        const candidate = field.uppercase ? normalized.toUpperCase() : normalized;

        if (field.minLength !== undefined && normalized.length < field.minLength) {
          errors[field.key] = `${field.label} must be at least ${field.minLength} characters.`;
          continue;
        }

        if (field.maxLength !== undefined && normalized.length > field.maxLength) {
          errors[field.key] = `${field.label} must be ${field.maxLength} characters or fewer.`;
          continue;
        }

        if (field.pattern && !new RegExp(field.pattern).test(candidate)) {
          errors[field.key] = field.patternMessage || `${field.label} format is invalid.`;
          continue;
        }
      }

      if (field.type === 'number') {
        const numericValue = Number(value);

        if (!Number.isFinite(numericValue)) {
          errors[field.key] = `${field.label} must be a valid number.`;
          continue;
        }

        if (field.integer && !Number.isInteger(numericValue)) {
          errors[field.key] = `${field.label} must be a whole number.`;
          continue;
        }

        if (field.min !== undefined && numericValue < field.min) {
          errors[field.key] = `${field.label} must be at least ${field.min}.`;
          continue;
        }

        if (field.max !== undefined && numericValue > field.max) {
          errors[field.key] = `${field.label} must be ${field.max} or less.`;
          continue;
        }
      }

      if (this.isJsonField(field.key)) {
        const rawValue = String(value).trim();
        if (field.required && !rawValue) {
          errors[field.key] = `${field.label} is required.`;
          continue;
        }

        if (rawValue) {
          try {
            JSON.parse(rawValue);
          } catch {
            errors[field.key] = `${field.label} must be valid JSON.`;
          }
        }
      }
    }

    return errors;
  }

  private isActiveItem(item: EntityRecord): boolean {
    if (this.section.showActiveBadgeKey) {
      return Boolean(item[this.section.showActiveBadgeKey]);
    }

    return item['status'] === 'Active';
  }

  private createFormModel(item: EntityRecord): EntityRecord {
    const formModel = { ...item };

    for (const field of this.section.fields) {
      if (this.isJsonField(field.key) && formModel[field.key] !== undefined && formModel[field.key] !== null) {
        formModel[field.key] = typeof formModel[field.key] === 'string'
          ? formModel[field.key]
          : JSON.stringify(formModel[field.key], null, 2);
      }
    }

    return formModel;
  }

  private isJsonField(key: string): boolean {
    return key === 'variables' || key === 'condition';
  }

  getItemId(item: EntityRecord): EntityIdentifier {
    const idKey = this.section.idKey || 'id';
    return item[idKey] as EntityIdentifier;
  }

  getDisplayId(item: EntityRecord, rowIndex: number): EntityIdentifier {
    if ((this.section.idKey || 'id') === 'id') {
      return (this.currentPage - 1) * this.pageSize + rowIndex + 1;
    }

    return this.getItemId(item);
  }

  buildRowDetailRoute(item: EntityRecord): string {
    if (!this.section.rowDetailRoute) return '';
    const id = this.getItemId(item);
    return this.section.rowDetailRoute.replace(':id', String(id));
  }

  private loadFieldOptions(): void {
    const relationFields = this.section.fields.filter((field) => field.dataEndpoint);

    for (const field of relationFields) {
      this.entityCrudService.list(field.dataEndpoint!).subscribe({
        next: (items) => {
          const filteredItems = field.optionFilterActive
            ? items.filter((item) => item['is_active'] !== false)
            : items;

          this.fieldOptions[field.key] = filteredItems.map((item) => ({
            value: item[field.optionValueKey || 'id'] as string | number | boolean,
            label: this.buildOptionLabel(field, item)
          }));

          if (this.section.listParamKey === field.key && !this.listParamValue && this.fieldOptions[field.key].length > 0) {
            this.listParamValue = this.fieldOptions[field.key][0].value as string | number;
            this.formModel[field.key] = this.listParamValue;
            this.loadItems();
          }
        },
        error: () => {
          this.fieldOptions[field.key] = field.options || [];
        }
      });
    }
  }

  private buildOptionLabel(field: CrudFieldConfig, item: EntityRecord): string {
    const labelKeys = field.optionLabelKeys || [field.optionValueKey || 'id'];
    const parts = labelKeys
      .map((key) => item[key])
      .filter((value) => value !== null && value !== undefined && String(value).trim() !== '')
      .map((value) => String(value));

    if (parts.length === 0) {
      return String(item[field.optionValueKey || 'id'] || '');
    }

    if (parts.length === 1) {
      return parts[0];
    }

    return `${parts[0]} (${parts.slice(1).join(' / ')})`;
  }

  private toHiddenItemKey(id: EntityIdentifier): string {
    return String(id);
  }
}
