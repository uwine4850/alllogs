
export interface MsgRegister {
    TypRegisterMessage?: unknown;
    Username: string;
    Password: string;
    RepeatPassword: string;
}
export function isMsgRegister(obj: any): obj is MsgRegister {
    return typeof obj === 'object' && obj !== null && 'TypRegisterMessage' in obj;
}
export interface MsgLogin {
    TypLoginMessage?: unknown;
    Username: string;
    Password: string;
}
export function isMsgLogin(obj: any): obj is MsgLogin {
    return typeof obj === 'object' && obj !== null && 'TypLoginMessage' in obj;
}
export interface MsgLoginResponse {
    TypLoginResponseMessage?: unknown;
    JWT: string;
    UID: number;
    Error: string;
}
export function isMsgLoginResponse(obj: any): obj is MsgLoginResponse {
    return typeof obj === 'object' && obj !== null && 'TypLoginResponseMessage' in obj;
}
export interface MsgLogout {
    TypLogoutMessage?: unknown;
    UID: number;
}
export function isMsgLogout(obj: any): obj is MsgLogout {
    return typeof obj === 'object' && obj !== null && 'TypLogoutMessage' in obj;
}
