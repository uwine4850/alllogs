
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
