import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LanguagesComponent } from './languages.component';
import { LanguagesRoutingModule } from './languages-routing.module';

@NgModule({
  imports: [
    CommonModule,
    LanguagesRoutingModule,
    LanguagesComponent   // ✅ IMPORTANT: standalone component goes here
  ]
})
export class LanguagesModule {}