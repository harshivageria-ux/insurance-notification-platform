import { environment } from '../../../environments/environment';

/**
 * API Configuration
 * Centralized API endpoints and base URL configuration
 * Supports both mock and real API modes
 */

export const API_CONFIG = {
  // Base URL - set to your backend server URL
  // Current backend target: Go API host
  BASE_URL: environment.apiBaseUrl,

  // API Endpoints
  ENDPOINTS: {
    LANGUAGES: {
      GET_ALL: '/languages',
      CREATE: '/languages',
      UPDATE: '/languages/:id',
      DELETE: '/languages/:id'
    },
    PRIORITIES: {
      GET_ALL: '/priorities',
      CREATE: '/priorities',
      UPDATE: '/priorities/:id',
      DELETE: '/priorities/:id'
    },
    STATUSES: {
      GET_ALL: '/statuses',
      CREATE: '/statuses',
      UPDATE: '/statuses/:id',
      DELETE: '/statuses/:id'
    },
    SCHEDULE_TYPES: {
      GET_ALL: '/schedule-types',
      CREATE: '/schedule-types',
      UPDATE: '/schedule-types/:id',
      DELETE: '/schedule-types/:id'
    },
    CATEGORIES: {
      GET_ALL: '/categories',
      CREATE: '/categories',
      UPDATE: '/categories/:id',
      DELETE: '/categories/:id'
    },
    CHANNELS: {
      GET_ALL: '/channels',
      CREATE: '/channels',
      UPDATE: '/channels/:id',
      TOGGLE: '/channels/:id/toggle',
      DELETE: '/channels/:id'
    },
    CHANNEL_PROVIDERS: {
      GET_ALL: '/channel-providers',
      CREATE: '/channel-providers',
      UPDATE: '/channel-providers/:id',
      TOGGLE: '/channel-providers/:id/toggle',
      DELETE: '/channel-providers/:id'
    },
    PROVIDER_SETTINGS: {
      GET_ALL: '/provider-settings/:provider_id',
      CREATE: '/provider-settings',
      UPDATE: '/provider-settings/:id',
      DELETE: '/provider-settings/:id'
    },
    TEMPLATE_GROUPS: {
      GET_ALL: '/template-groups',
      CREATE: '/template-groups',
      UPDATE: '/template-groups/:id',
      DELETE: '/template-groups/:id'
    },
    TEMPLATES: {
      GET_ALL: '/templates',
      CREATE: '/templates',
      UPDATE: '/templates/:id',
      PREVIEW: '/templates/:id/preview',
      DELETE: '/templates/:id'
    },
    ROUTING_RULES: {
      GET_ALL: '/routing-rules',
      CREATE: '/routing-rules',
      UPDATE: '/routing-rules/:id',
      TOGGLE: '/routing-rules/:id/toggle',
      DELETE: '/routing-rules/:id'
    }
  },

  // Stored Procedure Mapping (for backend reference)
  // Backend will use:
  // - sp_languages_insert (POST)
  // - sp_languages_update (PUT)
  // - sp_languages_deactivate (DELETE - soft delete)

  // HTTP Configuration
  HTTP_CONFIG: {
    TIMEOUT: 30000, // 30 seconds
    RETRY_ATTEMPTS: 3,
    RETRY_DELAY: 1000 // milliseconds
  }
};

/**
 * Helper function to build full API URL
 * @param endpoint - relative endpoint path
 * @returns full API URL
 */
export function getApiUrl(endpoint: string): string {
  if (!API_CONFIG.BASE_URL) {
    return endpoint; // For mock mode, return as is
  }
  return `${API_CONFIG.BASE_URL}${endpoint}`;
}

/**
 * Helper function to replace path parameters
 * @param endpoint - endpoint with :param syntax
 * @param params - object with parameter values
 * @returns endpoint with replaced parameters
 */
export function buildEndpoint(endpoint: string, params?: Record<string, any>): string {
  if (!params) {
    return endpoint;
  }

  let url = endpoint;
  Object.keys(params).forEach(key => {
    url = url.replace(`:${key}`, params[key]);
  });
  return url;
}
