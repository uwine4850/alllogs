
export interface RegisterMessage {
    TypRegisterMessage?: unknown;
    Username: string;
    Password: string;
    RepeatPassword: string;
}
export function isRegisterMessage(obj: any): obj is RegisterMessage {
    return typeof obj === 'object' && obj !== null && 'TypRegisterMessage' in obj;
}
export interface LoginMessage {
    TypLoginMessage?: unknown;
    Username: string;
    Password: string;
}
export function isLoginMessage(obj: any): obj is LoginMessage {
    return typeof obj === 'object' && obj !== null && 'TypLoginMessage' in obj;
}
export interface LoginResponseMessage {
    TypLoginResponseMessage?: unknown;
    JWT: string;
    UID: number;
    Error: string;
}
export function isLoginResponseMessage(obj: any): obj is LoginResponseMessage {
    return typeof obj === 'object' && obj !== null && 'TypLoginResponseMessage' in obj;
}
export interface LogoutMessage {
    TypLogoutMessage?: unknown;
    AID: number;
}
export function isLogoutMessage(obj: any): obj is LogoutMessage {
    return typeof obj === 'object' && obj !== null && 'TypLogoutMessage' in obj;
}
