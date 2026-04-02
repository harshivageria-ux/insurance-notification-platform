export type PushTemplatePriority = 'High' | 'Normal';

export interface PushNotificationTemplate {
  id: string;
  templateName: string;
  channel: 'Push Notification';
  title: string;
  message: string;
  targetUrl: string;
  iconUrl: string;
  bannerUrl: string;
  requireInteraction: boolean;
  priority: PushTemplatePriority;
  ttl: number;
  isActive: boolean;
  updatedAt: string; // ISO
}

export interface PushTemplateVariable {
  key: string; // e.g. user_name
  token: string; // e.g. {{user_name}}
}

