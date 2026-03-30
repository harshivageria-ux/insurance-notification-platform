export type EntityStatus = 'Active' | 'Inactive';

export interface ApiEnvelope<T> {
  success?: boolean;
  data?: T;
  message?: string;
  error?: string;
  errors?: string[];
}

export interface CrudFieldOption {
  label: string;
  value: string | number | boolean;
}

export interface CrudFieldConfig {
  key: string;
  label: string;
  type: 'text' | 'textarea' | 'number' | 'select';
  placeholder?: string;
  required?: boolean;
  /**
   * When false, the field is not rendered in the modal form.
   * (Independent of `table`, which only controls table columns.)
   */
  showInForm?: boolean;
  minLength?: number;
  maxLength?: number;
  min?: number;
  max?: number;
  integer?: boolean;
  uppercase?: boolean;
  pattern?: string;
  patternMessage?: string;
  options?: CrudFieldOption[];
  table?: boolean;
  searchable?: boolean;
  width?: string;
  dataEndpoint?: string;
  optionValueKey?: string;
  optionLabelKeys?: string[];
  optionFilterActive?: boolean;
}

export type EntityRecord = Record<string, unknown>;
export type EntityIdentifier = string | number;
export type EntityPayloadTransformer = (payload: EntityRecord, item?: EntityRecord) => EntityRecord;

export interface EntitySectionConfig {
  key: string;
  route: string;
  title: string;
  subtitle: string;
  endpoint: string;
  idKey?: string;
  deleteEndpoint?: string;
  updateEndpoint?: string;
  toggleEndpoint?: string;
  /**
   * Optional "details" action shown in the table row actions cell.
   * Use `:id` placeholder to be replaced with the row id.
   */
  rowDetailRoute?: string;
  rowDetailLabel?: string;
  fields: CrudFieldConfig[];
  defaultItem: Record<string, unknown>;
  createPayload?: EntityPayloadTransformer;
  updatePayload?: EntityPayloadTransformer;
  searchPlaceholder?: string;
  listParamKey?: string;
  listParamLabel?: string;
  listParamPlaceholder?: string;
  listParamDefault?: string | number;
  supportsToggle?: boolean;
  showStatusBadgeKey?: string;
  showActiveBadgeKey?: string;
}
