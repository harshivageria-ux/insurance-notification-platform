import { PushNotificationTemplate, PushTemplateVariable } from './template-builder.types';

export const PUSH_TEMPLATE_VARIABLES: PushTemplateVariable[] = [
  { key: 'user_name', token: '{{user_name}}' },
  { key: 'order_id', token: '{{order_id}}' },
  { key: 'app_name', token: '{{app_name}}' }
];

export const MOCK_PUSH_TEMPLATES: PushNotificationTemplate[] = [
  {
    id: 'tpl_order_shipped',
    templateName: 'Order Shipped',
    channel: 'Push Notification',
    title: 'Your order {{order_id}} is shipped',
    message: 'Hi {{user_name}}, your order is on the way. Track it now.',
    targetUrl: 'https://example.com/orders/{{order_id}}',
    iconUrl: 'https://placehold.co/64x64/png?text=Icon',
    bannerUrl: 'https://placehold.co/360x180/png?text=Banner',
    requireInteraction: false,
    priority: 'High',
    ttl: 3600,
    isActive: true,
    updatedAt: new Date(Date.now() - 1000 * 60 * 60 * 6).toISOString()
  },
  {
    id: 'tpl_welcome',
    templateName: 'Welcome Notification',
    channel: 'Push Notification',
    title: 'Welcome to {{app_name}}',
    message: 'Hello {{user_name}}! Thanks for joining {{app_name}}.',
    targetUrl: 'https://example.com/welcome',
    iconUrl: 'https://placehold.co/64x64/png?text=W',
    bannerUrl: '',
    requireInteraction: true,
    priority: 'Normal',
    ttl: 86400,
    isActive: true,
    updatedAt: new Date(Date.now() - 1000 * 60 * 60 * 24 * 2).toISOString()
  },
  {
    id: 'tpl_payment_success',
    templateName: 'Payment Success',
    channel: 'Push Notification',
    title: 'Payment received',
    message: 'Payment successful for order {{order_id}}. Thank you, {{user_name}}.',
    targetUrl: 'https://example.com/payments/{{order_id}}',
    iconUrl: 'https://placehold.co/64x64/png?text=$',
    bannerUrl: 'https://placehold.co/360x180/png?text=Receipt',
    requireInteraction: false,
    priority: 'Normal',
    ttl: 7200,
    isActive: false,
    updatedAt: new Date(Date.now() - 1000 * 60 * 90).toISOString()
  }
];

