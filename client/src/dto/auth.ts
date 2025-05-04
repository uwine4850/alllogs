
export interface RegisterMessage {
    Username: string;
    Password: string;
    RepeatPassword: string;
}
export interface LoginMessage {
    Username: string;
    Password: string;
}
export interface LoginResponseMessage {
    JWT: string;
    UID: number;
    Error: string;
}
