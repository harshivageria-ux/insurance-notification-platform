/**
 * Language Model
 * TypeScript interface for Language domain entity
 */

export interface Language {
  id: number;
  name: string;
  code: string;
  status: 'Active' | 'Inactive';
  createdAt?: Date;
  updatedAt?: Date;
}

/**
 * Language request payload for API
 */
export interface CreateLanguageRequest {
  name: string;
  code: string;
  status: 'Active' | 'Inactive';
}

/**
 * Language update payload for API
 */
export interface UpdateLanguageRequest extends CreateLanguageRequest {
  id: number;
}

/**
 * API Response wrapper
 */
export interface ApiResponse<T> {
  success: boolean;
  data?: T;
  message?: string;
  errors?: string[];
}

/**
 * Paginated response wrapper
 */
export interface PaginatedResponse<T> {
  success: boolean;
  data: T[];
  total: number;
  page: number;
  pageSize: number;
  totalPages: number;
}
