
export interface ProfileMessage {
    User: UserMessage;
    Id: string;
    UserId: string;
    Description: string;
    Avatar: string;
    Token: string;
    Error: string;
}
export interface UserMessage {
    Id: string;
    Username: string;
}
export interface GenTokenMessage {
    UserId: string;
}
export interface TokenResponse {
    Token: string;
    Error: string;
}
export interface ProfileUpdateMessage {
    PID: string;
    Description: string;
    Avatar: File | null;
    OldAvatarPath: string;
    DelAvatar: boolean;
}
