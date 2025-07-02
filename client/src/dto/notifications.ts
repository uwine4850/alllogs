
export interface MsgNotificationInfo {
    TypNotificationInfoMessage?: unknown;
    PID: string;
    Text: string;
}
export function isMsgNotificationInfo(obj: any): obj is MsgNotificationInfo {
    return typeof obj === 'object' && obj !== null && 'TypNotificationInfoMessage' in obj;
}
