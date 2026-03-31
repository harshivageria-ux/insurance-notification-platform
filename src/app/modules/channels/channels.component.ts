import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ChannelService, Channel } from '../../services/channel.service';

@Component({
  selector: 'app-channels',
  standalone: true,
  imports: [CommonModule, FormsModule],
  template: `
    <div class="container">
      <div class="header">
        <h1>Channels</h1>
        <button class="btn-add" (click)="openModal()">+ Add Channel</button>
      </div>

      <div class="search-bar">
        <input type="text" placeholder="Search..." [(ngModel)]="searchText">
      </div>

      <table class="table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Type</th>
            <th>Active</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr *ngFor="let item of paginatedItems">
            <td>{{item.id}}</td>
            <td>{{item.name}}</td>
            <td>{{item.type}}</td>
            <td><span class="badge" [ngClass]="item.is_active ? 'active' : 'inactive'">{{item.is_active ? 'Yes' : 'No'}}</span></td>
            <td>
              <button class="btn-edit" (click)="editItem(item)">Edit</button>
              <button class="btn-delete" (click)="deleteItem(item.id)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>

      <div class="modal" *ngIf="showModal" (click)="showModal = false">
        <div class="modal-content" (click)="$event.stopPropagation()">
          <div class="modal-header">
            <h2>{{isEditMode ? 'Edit Channel' : 'Add Channel'}}</h2>
            <button class="close" (click)="showModal = false">×</button>
          </div>
          <div class="modal-body">
            <input type="text" placeholder="Name" [(ngModel)]="newItem.name">
            <input type="text" placeholder="Type (Email, SMS, Push)" [(ngModel)]="newItem.type">
            <label>
              <input type="checkbox" [(ngModel)]="newItem.is_active">
              Active
            </label>
          </div>
          <div class="modal-footer">
            <button class="btn-cancel" (click)="showModal = false">Cancel</button>
            <button class="btn-save" (click)="saveItem()">{{isEditMode ? 'Update' : 'Save'}}</button>
          </div>
        </div>
      </div>
    </div>
  `,
  styles: [`
    .container { padding: 20px; }
    .header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; }
    .btn-add { background: #4CAF50; color: white; padding: 10px 20px; border: none; cursor: pointer; border-radius: 4px; }
    .table { width: 100%; border-collapse: collapse; }
    th, td { padding: 12px; text-align: left; border-bottom: 1px solid #ddd; }
    th { background: #f5f5f5; }
    .badge { padding: 4px 8px; border-radius: 3px; font-size: 12px; }
    .badge.active { background: #4CAF50; color: white; }
    .badge.inactive { background: #f44336; color: white; }
    .btn-edit { background: #2196F3; color: white; padding: 5px 10px; border: none; cursor: pointer; margin-right: 5px; }
    .btn-delete { background: #f44336; color: white; padding: 5px 10px; border: none; cursor: pointer; }
    .modal { display: flex; position: fixed; top: 0; left: 0; right: 0; bottom: 0; background: rgba(0,0,0,0.5); justify-content: center; align-items: center; }
    .modal-content { background: white; padding: 30px; border-radius: 8px; width: 400px; }
    input, select { width: 100%; padding: 10px; margin-bottom: 15px; border: 1px solid #ddd; }
    label { display: block; margin-bottom: 15px; }
    label input[type="checkbox"] { margin-right: 8px; }
    .btn-save { background: #4CAF50; color: white; }
    .btn-cancel { background: #999; color: white; }
    .btn-save, .btn-cancel { padding: 10px 20px; border: none; cursor: pointer; margin-right: 10px; }
  `]
})
export class ChannelsComponent implements OnInit {
  items: Channel[] = [];
  showModal = false;
  searchText = '';
  currentPage = 1;
  pageSize = 5;
  isEditMode = false;
  editingId: number | null = null;

  newItem: any = { name: '', type: '', is_active: true };

  constructor(private service: ChannelService) {}

  ngOnInit() {
    this.loadItems();
  }

  get paginatedItems() {
    const filtered = this.items.filter(item =>
      item.name.toLowerCase().includes(this.searchText.toLowerCase())
    );
    const startIndex = (this.currentPage - 1) * this.pageSize;
    return filtered.slice(startIndex, startIndex + this.pageSize);
  }

  get totalPages() {
    const filtered = this.items.filter(item =>
      item.name.toLowerCase().includes(this.searchText.toLowerCase())
    );
    return Math.ceil(filtered.length / this.pageSize);
  }

  loadItems() {
    this.service.getAll().subscribe({
      next: (data) => this.items = data,
      error: (err) => console.error('Error loading channels', err)
    });
  }

  openModal() {
    this.showModal = true;
    this.isEditMode = false;
    this.newItem = { name: '', type: '', is_active: true };
  }

  editItem(item: Channel) {
    this.editingId = item.id;
    this.newItem = { ...item };
    this.isEditMode = true;
    this.showModal = true;
  }

  saveItem() {
    if (this.isEditMode && this.editingId) {
      this.service.update(this.editingId, this.newItem).subscribe({
        next: () => { this.loadItems(); this.showModal = false; },
        error: (err) => console.error('Error updating channel', err)
      });
    } else {
      this.service.create(this.newItem).subscribe({
        next: () => { this.loadItems(); this.showModal = false; },
        error: (err) => console.error('Error creating channel', err)
      });
    }
  }

  deleteItem(id: number) {
    if (confirm('Are you sure?')) {
      this.service.delete(id).subscribe({
        next: () => this.loadItems(),
        error: (err) => console.error('Error deleting channel', err)
      });
    }
  }
}
