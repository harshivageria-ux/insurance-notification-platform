import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { LanguageService } from '../../services/language.service';
import { Language, CreateLanguageRequest } from '../../core/models/language.model';
import { LoaderService } from '../../core/services/loader.service';

@Component({
  selector: 'app-languages',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './languages.component.html',
  styleUrls: ['./languages.component.css']
})
export class LanguagesComponent implements OnInit {
  private readonly languageCodePattern = /^[A-Z]{2,10}$/;
  private readonly hiddenLanguageIds = new Set<number>();

  languages: Language[] = [];
  showModal = false;
  loading = false;
  showToastFlag = false;
  toastMessage = '';
  searchText = '';
  currentPage = 1;
  pageSize = 5;
  validationErrors: Record<string, string> = {};

  newLanguage: CreateLanguageRequest = {
    name: '',
    code: '',
    status: 'Active'
  };

  editingLanguageId: number | null = null;
  isEditMode = false;

  constructor(
    private languageService: LanguageService,
    public loaderService: LoaderService
  ) {}

  ngOnInit() {
    this.loadLanguages();
  }

  get paginatedLanguages(): Language[] {
    const filtered = this.languages.filter(lang =>
      lang.name.toLowerCase().includes(this.searchText.toLowerCase()) ||
      lang.code.toLowerCase().includes(this.searchText.toLowerCase())
    );
    
    const startIndex = (this.currentPage - 1) * this.pageSize;
    return filtered.slice(startIndex, startIndex + this.pageSize);
  }

  get totalPages(): number {
    const filtered = this.languages.filter(lang =>
      lang.name.toLowerCase().includes(this.searchText.toLowerCase()) ||
      lang.code.toLowerCase().includes(this.searchText.toLowerCase())
    );
    return Math.ceil(filtered.length / this.pageSize);
  }

  // ✅ GET API CALL
  loadLanguages() {
    this.loading = true;
    this.languageService.getLanguages().subscribe({
      next: (res: Language[]) => {
        this.languages = res.filter((language) => !this.hiddenLanguageIds.has(language.id));
        this.loading = false;
      },
      error: (err) => {
        console.error('Error loading languages', err);
        this.loading = false;
        this.showToast('Error loading languages');
      }
    });
  }

  // ✅ OPEN MODAL
  openModal() {
    this.validationErrors = {};
    this.showModal = true;
  }

  // ✅ ALSO SUPPORT HTML BUTTON NAME
  openAddForm() {
    this.editingLanguageId = null;
    this.isEditMode = false;
    this.resetForm();
    this.validationErrors = {};
    this.showModal = true;
  }

  // ✅ CLOSE MODAL
  closeModal() {
    this.showModal = false;
    this.editingLanguageId = null;
    this.isEditMode = false;
    this.validationErrors = {};
    this.resetForm();
  }

  // ✅ POST API CALL (handles both create and update)
  saveLanguage() {
    this.normalizeFormValues();
    this.validationErrors = this.validateForm();

    if (Object.keys(this.validationErrors).length > 0) {
      this.showToast('Please fix the highlighted errors');
      return;
    }

    if (this.editingLanguageId !== null) {
      // Update mode
      this.languageService.updateLanguage({
        id: this.editingLanguageId,
        ...this.newLanguage
      }).subscribe({
        next: () => {
          this.loadLanguages();
          this.closeModal();
          this.showToast('Language updated successfully');
        },
        error: (err) => {
          console.error('Error updating language', err);
          this.showToast(err?.message || 'Error updating language');
        }
      });
    } else {
      // Create mode
      this.languageService.addLanguage(this.newLanguage).subscribe({
        next: () => {
          this.loadLanguages();
          this.closeModal();
          this.showToast('Language added successfully');
        },
        error: (err) => {
          console.error('Error saving language', err);
          this.showToast(err?.message || 'Error saving language');
        }
      });
    }
  }

  // ✅ EDIT LANGUAGE (opens modal with data)
  editLanguage(lang: Language) {
    this.editingLanguageId = lang.id;
    this.isEditMode = true;
    this.newLanguage = {
      name: lang.name,
      code: lang.code,
      status: lang.status
    };
    this.validationErrors = {};
    this.showModal = true;
  }

  // ✅ DELETE LANGUAGE (API call)
  deleteLanguage(id: number) {
    if (confirm('Are you sure you want to delete?')) {
      this.languageService.deleteLanguage(id).subscribe({
        next: () => {
          this.hiddenLanguageIds.add(id);
          this.languages = this.languages.filter((language) => language.id !== id);
          this.currentPage = Math.min(this.currentPage, Math.max(1, this.totalPages));
          this.showToast('Language removed from the UI');
        },
        error: (err) => {
          console.error('Error deleting language', err);
          this.showToast(err?.message || 'Error deleting language');
        }
      });
    }
  }

  // ✅ RESET FORM
  resetForm() {
    this.newLanguage = {
      name: '',
      code: '',
      status: 'Active'
    };
  }

  onNameInput(): void {
    this.newLanguage.name = this.newLanguage.name.replace(/\s+/g, ' ').trimStart();
    this.clearValidationError('name');
  }

  onCodeInput(): void {
    this.newLanguage.code = this.newLanguage.code.toUpperCase().replace(/[^A-Z]/g, '').slice(0, 10);
    this.clearValidationError('code');
  }

  onStatusChange(): void {
    this.clearValidationError('status');
  }

  private validateForm(): Record<string, string> {
    const errors: Record<string, string> = {};
    const name = this.newLanguage.name.trim();
    const code = this.newLanguage.code.trim().toUpperCase();
    const status = this.newLanguage.status;

    if (!name) {
      errors['name'] = 'Language name is required.';
    } else if (name.length > 100) {
      errors['name'] = 'Language name must be 100 characters or fewer.';
    } else if (this.hasDuplicateName(name)) {
      errors['name'] = 'This language name already exists.';
    }

    if (!code) {
      errors['code'] = 'Language code is required.';
    } else if (!this.languageCodePattern.test(code)) {
      errors['code'] = 'Use 2 to 10 letters only, for example EN or HINDI.';
    } else if (this.hasDuplicateCode(code)) {
      errors['code'] = 'This language code already exists.';
    }

    if (status !== 'Active' && status !== 'Inactive') {
      errors['status'] = 'Please choose a valid status.';
    }

    return errors;
  }

  private normalizeFormValues(): void {
    this.newLanguage = {
      ...this.newLanguage,
      name: this.newLanguage.name.replace(/\s+/g, ' ').trim(),
      code: this.newLanguage.code.trim().toUpperCase(),
      status: this.newLanguage.status
    };
  }

  private hasDuplicateName(name: string): boolean {
    const normalizedName = name.toLowerCase();
    return this.languages.some(lang =>
      lang.id !== this.editingLanguageId &&
      lang.name.trim().toLowerCase() === normalizedName
    );
  }

  private hasDuplicateCode(code: string): boolean {
    const normalizedCode = code.toUpperCase();
    return this.languages.some(lang =>
      lang.id !== this.editingLanguageId &&
      lang.code.trim().toUpperCase() === normalizedCode
    );
  }

  private clearValidationError(field: string): void {
    if (this.validationErrors[field]) {
      delete this.validationErrors[field];
      this.validationErrors = { ...this.validationErrors };
    }
  }

  // ✅ SHOW TOAST MESSAGE
  showToast(message: string) {
    this.toastMessage = message;
    this.showToastFlag = true;
    setTimeout(() => {
      this.showToastFlag = false;
    }, 3000);
  }
}
