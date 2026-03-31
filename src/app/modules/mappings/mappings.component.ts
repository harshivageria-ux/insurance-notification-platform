import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MappingService, MappingCategoryChannel, MappingChannelProvider, MappingTemplateChannelLanguage } from '../../core/services/mapping.service';

@Component({
  selector: 'app-mappings',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './mappings.component.html',
  styleUrls: ['./mappings.component.css']
})
export class MappingsComponent implements OnInit {
  categoryChannelForm: FormGroup;
  channelProviderForm: FormGroup;
  templateChannelLanguageForm: FormGroup;

  categoryChannels: MappingCategoryChannel[] = [];
  channelProviders: MappingChannelProvider[] = [];
  templateChannelLanguages: MappingTemplateChannelLanguage[] = [];

  categories: any[] = [];
  channels: any[] = [];
  providers: any[] = [];
  languages: any[] = [];
  templateGroups: any[] = [];
  templates: any[] = [];

  successMessage = '';
  errorMessage = '';

  constructor(private fb: FormBuilder, private mappingService: MappingService) {
    this.categoryChannelForm = this.fb.group({ category_id: [null, Validators.required], channel_id: [null, Validators.required] });
    this.channelProviderForm = this.fb.group({ channel_id: [null, Validators.required], provider_id: [null, Validators.required], priority: [1, [Validators.required, Validators.min(1)]] });
    this.templateChannelLanguageForm = this.fb.group({ template_group_id: [null, Validators.required], channel_id: [null, Validators.required], language_id: [null, Validators.required], template_id: [null, Validators.required] });
  }

  ngOnInit() {
    this.loadMasters();
    this.refreshMappings();
  }

  private loadMasters() {
    this.mappingService.getCategories().subscribe(res => this.categories = res, err => this._setError(err));
    this.mappingService.getChannels().subscribe(res => this.channels = res, err => this._setError(err));
    this.mappingService.getProviders().subscribe(res => this.providers = res, err => this._setError(err));
    this.mappingService.getLanguages().subscribe(res => this.languages = res, err => this._setError(err));
    this.mappingService.getTemplateGroups().subscribe(res => this.templateGroups = res, err => this._setError(err));
    this.mappingService.getTemplates().subscribe(res => this.templates = res, err => this._setError(err));
  }

  private refreshMappings() {
    this.mappingService.getCategoryChannels().subscribe(res => this.categoryChannels = res.items, err => this._setError(err));
    this.mappingService.getChannelProviders().subscribe(res => this.channelProviders = res.items, err => this._setError(err));
    this.mappingService.getTemplateChannelLanguages().subscribe(res => this.templateChannelLanguages = res.items, err => this._setError(err));
  }

  createCategoryChannel() {
    if (this.categoryChannelForm.invalid) return;
    const payload = this.categoryChannelForm.value;

    this.mappingService.createCategoryChannel(payload).subscribe(() => {
      this._setSuccess('Category-channel mapping created successfully');
      this.categoryChannelForm.reset();
      this.refreshMappings();
    }, err => this._setError(err));
  }

  createChannelProvider() {
    if (this.channelProviderForm.invalid) return;
    this.mappingService.createChannelProvider(this.channelProviderForm.value).subscribe(() => {
      this._setSuccess('Channel-provider mapping created successfully');
      this.channelProviderForm.reset({ priority: 1 });
      this.refreshMappings();
    }, err => this._setError(err));
  }

  createTemplateChannelLanguage() {
    if (this.templateChannelLanguageForm.invalid) return;

    const payload = {
      template_id: this.templateChannelLanguageForm.value.template_id,
      channel_id: this.templateChannelLanguageForm.value.channel_id,
      language_id: this.templateChannelLanguageForm.value.language_id,
    };

    this.mappingService.createTemplateChannelLanguage(payload).subscribe(() => {
      this._setSuccess('Template-channel-language mapping created successfully');
      this.templateChannelLanguageForm.reset();
      this.refreshMappings();
    }, err => this._setError(err));
  }

  deleteCategoryChannel(id: number) {
    this.mappingService.deleteCategoryChannel(id).subscribe(() => {
      this._setSuccess('Category-channel mapping deleted');
      this.refreshMappings();
    }, err => this._setError(err));
  }

  deleteChannelProvider(id: number) {
    this.mappingService.deleteChannelProvider(id).subscribe(() => {
      this._setSuccess('Channel-provider mapping deleted');
      this.refreshMappings();
    }, err => this._setError(err));
  }

  deleteTemplateChannelLanguage(id: number) {
    this.mappingService.deleteTemplateChannelLanguage(id).subscribe(() => {
      this._setSuccess('Template-channel-language mapping deleted');
      this.refreshMappings();
    }, err => this._setError(err));
  }

  private _setSuccess(message: string) {
    this.errorMessage = '';
    this.successMessage = message;
    setTimeout(() => this.successMessage = '', 4000);
  }

  private _setError(err: any) {
    this.successMessage = '';
    this.errorMessage = err?.message || 'An error occurred';
    setTimeout(() => this.errorMessage = '', 5000);
  }
}
