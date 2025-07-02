
export interface MsgProfile {
    TypProfileMessage?: unknown;
    User: MsgUser | undefined;
    UserId: number;
    Description: string;
    Avatar: string;
    Token: string;
    Error: string;
}
export function isMsgProfile(obj: any): obj is MsgProfile {
    return typeof obj === 'object' && obj !== null && 'TypProfileMessage' in obj;
}
export interface MsgUser {
    TypUserMessage?: unknown;
    Id: number;
    Username: string;
}
export function isMsgUser(obj: any): obj is MsgUser {
    return typeof obj === 'object' && obj !== null && 'TypUserMessage' in obj;
}
export interface MsgGenToken {
    TypGenTokenMessage?: unknown;
    UserId: number;
}
export function isMsgGenToken(obj: any): obj is MsgGenToken {
    return typeof obj === 'object' && obj !== null && 'TypGenTokenMessage' in obj;
}
export interface MsgTokenResponse {
    TypTokenResponse?: unknown;
    Token: string;
    Error: string;
}
export function isMsgTokenResponse(obj: any): obj is MsgTokenResponse {
    return typeof obj === 'object' && obj !== null && 'TypTokenResponse' in obj;
}
export interface MsgProfileUpdate {
    TypProfileUpdateMessage?: unknown;
    UID: number;
    Description: string;
    Avatar: File | null;
    OldAvatarPath: string;
    DelAvatar: boolean;
}
export function isMsgProfileUpdate(obj: any): obj is MsgProfileUpdate {
    return typeof obj === 'object' && obj !== null && 'TypProfileUpdateMessage' in obj;
}
