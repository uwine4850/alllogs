
export interface ProfileMessage {
    User: UserMessage | undefined;
    Id: number;
    UserId: number;
    Description: string;
    Avatar: string;
    Token: string;
    Error: string;
}
export interface UserMessage {
    Id: number;
    Username: string;
}
export interface GenTokenMessage {
    UserId: number;
}
export interface TokenResponse {
    Token: string;
    Error: string;
}
export interface ProfileUpdateMessage {
    PID: number;
    Description: string;
    Avatar: File | null;
    OldAvatarPath: string;
    DelAvatar: boolean;
}
