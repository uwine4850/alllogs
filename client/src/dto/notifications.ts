
export interface NotificationInfoMessage {
    TypNotificationInfoMessage?: unknown;
    PID: string;
    Text: string;
}
export function isNotificationInfoMessage(obj: any): obj is NotificationInfoMessage {
    return typeof obj === 'object' && obj !== null && 'TypNotificationInfoMessage' in obj;
}
