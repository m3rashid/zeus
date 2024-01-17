export const notificationTypes = [
  'logout',
  'fetch:webui-notifications',
] as const;

export type NotificationType = (typeof notificationTypes)[number];

export type NotificationHandler = Partial<Record<NotificationType, Function>>;
