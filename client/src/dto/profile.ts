
export interface ProfileMessage {
    TypProfileMessage?: unknown;
    User: UserMessage | undefined;
    Id: number;
    UserId: number;
    Description: string;
    Avatar: string;
    Token: string;
    Error: string;
}
export function isProfileMessage(obj: any): obj is ProfileMessage {
    return typeof obj === 'object' && obj !== null && 'TypProfileMessage' in obj;
}
export interface UserMessage {
    TypUserMessage?: unknown;
    Id: number;
    Username: string;
}
export function isUserMessage(obj: any): obj is UserMessage {
    return typeof obj === 'object' && obj !== null && 'TypUserMessage' in obj;
}
export interface GenTokenMessage {
    TypGenTokenMessage?: unknown;
    UserId: number;
}
export function isGenTokenMessage(obj: any): obj is GenTokenMessage {
    return typeof obj === 'object' && obj !== null && 'TypGenTokenMessage' in obj;
}
export interface TokenResponse {
    TypTokenResponse?: unknown;
    Token: string;
    Error: string;
}
export function isTokenResponse(obj: any): obj is TokenResponse {
    return typeof obj === 'object' && obj !== null && 'TypTokenResponse' in obj;
}
export interface ProfileUpdateMessage {
    TypProfileUpdateMessage?: unknown;
    PID: number;
    Description: string;
    Avatar: File | null;
    OldAvatarPath: string;
    DelAvatar: boolean;
}
export function isProfileUpdateMessage(obj: any): obj is ProfileUpdateMessage {
    return typeof obj === 'object' && obj !== null && 'TypProfileUpdateMessage' in obj;
}
