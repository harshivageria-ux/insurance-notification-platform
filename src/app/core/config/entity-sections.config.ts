import { EntityRecord, EntitySectionConfig } from '../models/entity-crud.model';

const booleanOptions = [
  { label: 'Yes', value: true },
  { label: 'No', value: false }
] as const;

const codeFieldRules = {
  minLength: 2,
  maxLength: 30,
  uppercase: true,
  pattern: '^[A-Z0-9_]{2,30}$',
  patternMessage: 'Use 2 to 30 uppercase letters, numbers, or underscores only.'
} as const;

const nameFieldRules = {
  minLength: 2,
  maxLength: 100
} as const;

const descriptionFieldRules = {
  maxLength: 500
} as const;

const positiveIntegerRules = {
  min: 1,
  integer: true
} as const;

function text(value: unknown): string | null {
  if (value === null || value === undefined) {
    return null;
  }

  const normalized = String(value).trim();
  return normalized ? normalized : null;
}

function uuid(value: unknown): string | null {
  return text(value);
}

function intValue(value: unknown): number {
  if (value === null || value === undefined) {
    return 0;
  }

  const normalized = String(value).trim();
  if (!normalized) {
    return 0;
  }

  const n = Number(normalized);
  return Number.isFinite(n) ? n : 0;
}

function numberValue(value: unknown): number {
  return Number(value ?? 0);
}

function booleanValue(value: unknown): boolean {
  return value === true || value === 'true';
}

export const ENTITY_SECTIONS: EntitySectionConfig[] = [
  {
    key: 'languages',
    route: 'languages',
    title: 'Languages',
    subtitle: 'Manage supported languages in the PostgreSQL master table.',
    endpoint: '/languages',
    idKey: 'id',
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'name', label: 'Name', type: 'text', placeholder: 'English', required: true, table: true, searchable: true, ...nameFieldRules },
      { key: 'code', label: 'Code', type: 'text', placeholder: 'EN', required: true, table: true, searchable: true, width: '120px', ...codeFieldRules },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: { name: '', code: '', created_by: 'admin' },
    createPayload: (payload) => ({
      code: text(payload['code']),
      name: text(payload['name']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      code: text(payload['code']),
      name: text(payload['name']),
      updated_by: text(payload['created_by'])
    }),
    searchPlaceholder: 'Search languages...'
  },
  {
    key: 'priorities',
    route: 'priorities',
    title: 'Priorities',
    subtitle: 'Manage notification priority values and codes.',
    endpoint: '/priorities',
    idKey: 'priority_id',
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'priority_id', label: 'Priority ID', type: 'number', placeholder: '1', required: false, table: false, showInForm: false, width: '140px', ...positiveIntegerRules },
      { key: 'priority_code', label: 'Priority Code', type: 'text', placeholder: 'HIGH', table: true, searchable: true, ...codeFieldRules },
      { key: 'description', label: 'Description', type: 'textarea', placeholder: 'High priority', table: true, searchable: true, ...descriptionFieldRules },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: { priority_code: '', description: '', created_by: 'admin' },
    createPayload: (payload) => ({
      priority_code: text(payload['priority_code']),
      description: text(payload['description']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      priority_code: text(payload['priority_code']),
      description: text(payload['description']),
      created_by: text(payload['created_by'])
    }),
    searchPlaceholder: 'Search priorities...'
  },
  {
    key: 'statuses',
    route: 'statuses',
    title: 'Statuses',
    subtitle: 'Manage notification lifecycle statuses.',
    endpoint: '/statuses',
    idKey: 'status_id',
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'status_id', label: 'Status ID', type: 'number', placeholder: '1', required: false, table: false, showInForm: false, width: '120px', ...positiveIntegerRules },
      { key: 'status_code', label: 'Status Code', type: 'text', placeholder: 'QUEUED', table: true, searchable: true, ...codeFieldRules },
      { key: 'name', label: 'Name', type: 'text', placeholder: 'Queued', required: true, table: true, searchable: true, ...nameFieldRules },
      { key: 'description', label: 'Description', type: 'textarea', placeholder: 'Current status description', table: true, searchable: true, ...descriptionFieldRules },
      { key: 'is_final', label: 'Final Status', type: 'select', required: true, table: true, options: [...booleanOptions], width: '140px' },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: { status_code: '', name: '', description: '', is_final: false, created_by: 'admin' },
    createPayload: (payload) => ({
      status_code: text(payload['status_code']),
      name: text(payload['name']),
      description: text(payload['description']),
      is_final: booleanValue(payload['is_final']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      status_code: text(payload['status_code']),
      name: text(payload['name']),
      description: text(payload['description']),
      is_final: booleanValue(payload['is_final'])
    }),
    searchPlaceholder: 'Search statuses...'
  },
  {
    key: 'schedule-types',
    route: 'schedule-types',
    title: 'Schedule Types',
    subtitle: 'Manage available notification schedule types.',
    endpoint: '/schedule-types',
    idKey: 'schedule_type_id',
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'schedule_type_id', label: 'Schedule Type ID', type: 'number', placeholder: '1', required: false, table: false, showInForm: false, width: '160px', ...positiveIntegerRules },
      { key: 'schedule_code', label: 'Schedule Code', type: 'text', placeholder: 'IMMEDIATE', table: true, searchable: true, ...codeFieldRules },
      { key: 'description', label: 'Description', type: 'textarea', placeholder: 'Immediate delivery', table: true, searchable: true, ...descriptionFieldRules },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: { schedule_code: '', description: '', created_by: 'admin' },
    createPayload: (payload) => ({
      schedule_code: text(payload['schedule_code']),
      description: text(payload['description']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      schedule_code: text(payload['schedule_code']),
      description: text(payload['description'])
    }),
    searchPlaceholder: 'Search schedule types...'
  },
  {
    key: 'categories',
    route: 'categories',
    title: 'Categories',
    subtitle: 'Manage notification categories.',
    endpoint: '/categories',
    idKey: 'id',
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'code', label: 'Code', type: 'text', placeholder: 'PAYMENT', required: true, table: true, searchable: true, width: '160px', ...codeFieldRules },
      { key: 'name', label: 'Name', type: 'text', placeholder: 'Payment', required: true, table: true, searchable: true, ...nameFieldRules },
      { key: 'description', label: 'Description', type: 'textarea', placeholder: 'Category description', table: true, searchable: true, ...descriptionFieldRules },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: { code: '', name: '', description: '', created_by: 'admin' },
    createPayload: (payload) => ({
      code: text(payload['code']),
      name: text(payload['name']),
      description: text(payload['description']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      code: text(payload['code']),
      name: text(payload['name']),
      description: text(payload['description'])
    }),
    searchPlaceholder: 'Search categories...'
  },
  {
    key: 'channels',
    route: 'channels',
    title: 'Channels',
    subtitle: 'Manage delivery channels such as email, SMS, and push.',
    endpoint: '/channels',
    idKey: 'id',
    toggleEndpoint: '/channels/:id/toggle',
    supportsToggle: true,
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'code', label: 'Code', type: 'text', placeholder: 'EMAIL', required: true, table: true, searchable: true, width: '140px', ...codeFieldRules },
      { key: 'name', label: 'Name', type: 'text', placeholder: 'Email', required: true, table: true, searchable: true, ...nameFieldRules },
      { key: 'description', label: 'Description', type: 'textarea', placeholder: 'Channel description', table: true, searchable: true, ...descriptionFieldRules },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: { code: '', name: '', description: '', created_by: 'admin' },
    createPayload: (payload) => ({
      code: text(payload['code']),
      name: text(payload['name']),
      description: text(payload['description']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      code: text(payload['code']),
      name: text(payload['name']),
      description: text(payload['description'])
    }),
    searchPlaceholder: 'Search channels...'
  },
  {
    key: 'channel-providers',
    route: 'channel-providers',
    title: 'Channel Providers',
    subtitle: 'Manage providers assigned to a delivery channel.',
    endpoint: '/channel-providers',
    rowDetailRoute: '/channel-providers/:id/provider-settings',
    rowDetailLabel: 'Provider Settings',
    idKey: 'id',
    toggleEndpoint: '/channel-providers/:id/toggle',
    supportsToggle: true,
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'channel_id', label: 'Channel', type: 'select', required: true, table: false, searchable: true, dataEndpoint: '/channels', optionValueKey: 'id', optionLabelKeys: ['name', 'code'], optionFilterActive: true },
      { key: 'name', label: 'Name', type: 'text', placeholder: 'Twilio', required: true, table: true, searchable: true, ...nameFieldRules },
      { key: 'code', label: 'Code', type: 'text', placeholder: 'TWILIO', required: true, table: true, searchable: true, ...codeFieldRules },
      { key: 'priority', label: 'Priority', type: 'number', placeholder: '1', table: true, width: '120px', ...positiveIntegerRules },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: { channel_id: '', name: '', code: '', priority: 1, created_by: 'admin' },
    createPayload: (payload) => ({
      channel_id: intValue(payload['channel_id']),
      name: text(payload['name']),
      code: text(payload['code']),
      priority: numberValue(payload['priority']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      name: text(payload['name']),
      priority: numberValue(payload['priority'])
    }),
    searchPlaceholder: 'Search channel providers...'
  },
  {
    key: 'provider-settings',
    route: 'provider-settings',
    title: 'Provider Settings',
    subtitle: 'Manage provider-level settings, including encrypted values.',
    endpoint: '/provider-settings',
    idKey: 'id',
    showActiveBadgeKey: 'is_active',
    listParamKey: 'provider_id',
    listParamLabel: 'Provider',
    listParamPlaceholder: 'Select provider',
    listParamDefault: '',
    fields: [
      { key: 'provider_id', label: 'Provider', type: 'select', required: true, table: true, searchable: true, dataEndpoint: '/channel-providers', optionValueKey: 'id', optionLabelKeys: ['name', 'code'], optionFilterActive: true },
      { key: 'setting_key', label: 'Setting Key', type: 'text', placeholder: 'api_key', required: true, table: true, searchable: true, minLength: 2, maxLength: 100, pattern: '^[A-Za-z0-9_.-]{2,100}$', patternMessage: 'Use 2 to 100 letters, numbers, dots, hyphens, or underscores.' },
      { key: 'setting_value', label: 'Setting Value', type: 'textarea', placeholder: 'Setting value', required: true, table: true, maxLength: 2000 },
      { key: 'is_sensitive', label: 'Sensitive', type: 'select', required: true, table: true, options: [...booleanOptions], width: '120px' },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: { provider_id: '', setting_key: '', setting_value: '', is_sensitive: false, created_by: 'admin' },
    createPayload: (payload) => ({
      provider_id: intValue(payload['provider_id']),
      setting_key: text(payload['setting_key']),
      setting_value: text(payload['setting_value']),
      is_sensitive: booleanValue(payload['is_sensitive']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      setting_key: text(payload['setting_key']),
      setting_value: text(payload['setting_value']),
      is_sensitive: booleanValue(payload['is_sensitive'])
    }),
    searchPlaceholder: 'Search provider settings...'
  },
  {
    key: 'template-groups',
    route: 'template-groups',
    title: 'Template Groups',
    subtitle: 'Manage template groups and their linked categories.',
    endpoint: '/template-groups',
    idKey: 'id',
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'name', label: 'Name', type: 'text', placeholder: 'Payment Alerts', required: true, table: true, searchable: true, ...nameFieldRules },
      { key: 'category_id', label: 'Category', type: 'select', required: true, table: false, searchable: true, dataEndpoint: '/categories', optionValueKey: 'id', optionLabelKeys: ['name', 'code'], optionFilterActive: true },
      { key: 'description', label: 'Description', type: 'textarea', placeholder: 'Group description', table: true, searchable: true, ...descriptionFieldRules },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: { name: '', category_id: '', description: '', created_by: 'admin' },
    createPayload: (payload) => ({
      name: text(payload['name']),
      category_id: intValue(payload['category_id']),
      description: text(payload['description']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      name: text(payload['name']),
      category_id: intValue(payload['category_id']),
      description: text(payload['description'])
    }),
    searchPlaceholder: 'Search template groups...'
  },
  {
    key: 'templates',
    route: 'templates',
    title: 'Templates',
    subtitle: 'Manage notification templates and placeholder support.',
    endpoint: '/templates',
    idKey: 'id',
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'template_group_id', label: 'Template Group', type: 'select', required: true, table: false, searchable: true, dataEndpoint: '/template-groups', optionValueKey: 'id', optionLabelKeys: ['name'], optionFilterActive: true },
      { key: 'channel_id', label: 'Channel', type: 'select', required: true, table: false, searchable: true, dataEndpoint: '/channels', optionValueKey: 'id', optionLabelKeys: ['name', 'code'], optionFilterActive: true },
      { key: 'language_id', label: 'Language', type: 'select', required: true, table: false, searchable: true, dataEndpoint: '/languages', optionValueKey: 'id', optionLabelKeys: ['name', 'code'], optionFilterActive: true },
      { key: 'title_template', label: 'Title Template', type: 'textarea', placeholder: 'Policy Update', table: true, maxLength: 250 },
      { key: 'message_template', label: 'Message Template', type: 'textarea', placeholder: 'Hello {{name}}', required: true, table: true, maxLength: 5000 },
      { key: 'has_variables', label: 'Has Variables', type: 'select', required: true, table: true, options: [...booleanOptions], width: '140px' },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: {
      template_group_id: '',
      channel_id: '',
      language_id: '',
      title_template: '',
      message_template: '',
      has_variables: false,
      created_by: 'admin'
    },
    createPayload: (payload) => ({
      template_group_id: intValue(payload['template_group_id']),
      channel_id: intValue(payload['channel_id']),
      language_id: intValue(payload['language_id']),
      title_template: text(payload['title_template']),
      message_template: text(payload['message_template']),
      has_variables: booleanValue(payload['has_variables']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      template_group_id: intValue(payload['template_group_id']),
      channel_id: intValue(payload['channel_id']),
      language_id: intValue(payload['language_id']),
      title_template: text(payload['title_template']),
      message_template: text(payload['message_template']),
      has_variables: booleanValue(payload['has_variables'])
    }),
    searchPlaceholder: 'Search templates...'
  },
  {
    key: 'routing-rules',
    route: 'routing-rules',
    title: 'Routing Rules',
    subtitle: 'Manage preferred and fallback providers per template group.',
    endpoint: '/routing-rules',
    idKey: 'id',
    toggleEndpoint: '/routing-rules/:id/toggle',
    supportsToggle: true,
    showActiveBadgeKey: 'is_active',
    fields: [
      { key: 'template_group_id', label: 'Template Group', type: 'select', required: true, table: true, searchable: true, dataEndpoint: '/template-groups', optionValueKey: 'id', optionLabelKeys: ['name'], optionFilterActive: true },
      { key: 'channel_id', label: 'Channel', type: 'select', required: true, table: true, searchable: true, dataEndpoint: '/channels', optionValueKey: 'id', optionLabelKeys: ['name', 'code'], optionFilterActive: true },
      { key: 'preferred_provider_id', label: 'Preferred Provider', type: 'select', required: true, table: true, searchable: true, dataEndpoint: '/channel-providers', optionValueKey: 'id', optionLabelKeys: ['name', 'code'], optionFilterActive: true },
      { key: 'fallback_provider_id', label: 'Fallback Provider', type: 'select', table: true, searchable: true, dataEndpoint: '/channel-providers', optionValueKey: 'id', optionLabelKeys: ['name', 'code'], optionFilterActive: true },
      { key: 'created_by', label: 'Created By', type: 'text', placeholder: 'admin', table: false }
    ],
    defaultItem: {
      template_group_id: '',
      channel_id: '',
      preferred_provider_id: '',
      fallback_provider_id: '',
      created_by: 'admin'
    },
    createPayload: (payload) => ({
      template_group_id: intValue(payload['template_group_id']),
      channel_id: intValue(payload['channel_id']),
      preferred_provider_id: intValue(payload['preferred_provider_id']),
      fallback_provider_id: intValue(payload['fallback_provider_id']),
      created_by: text(payload['created_by'])
    }),
    updatePayload: (payload) => ({
      preferred_provider_id: intValue(payload['preferred_provider_id']),
      fallback_provider_id: intValue(payload['fallback_provider_id'])
    }),
    searchPlaceholder: 'Search routing rules...'
  }
];

export const ENTITY_SECTION_MAP = ENTITY_SECTIONS.reduce<Record<string, EntitySectionConfig>>((acc, section) => {
  acc[section.route] = section;
  return acc;
}, {});
