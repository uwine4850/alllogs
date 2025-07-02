
export interface MsgBaseResponse {
    TypBaseResponseMessage?: unknown;
    Ok: boolean;
    Error: string;
}
export function isMsgBaseResponse(obj: any): obj is MsgBaseResponse {
    return typeof obj === 'object' && obj !== null && 'TypBaseResponseMessage' in obj;
}
export interface MsgClientError {
    TypClientErrorMessage?: unknown;
    Code: number;
    Text: string;
}
export function isMsgClientError(obj: any): obj is MsgClientError {
    return typeof obj === 'object' && obj !== null && 'TypClientErrorMessage' in obj;
}
export interface MsgServerError {
    TypServerErrorMessage?: unknown;
    Code: number;
    Text: string;
}
export function isMsgServerError(obj: any): obj is MsgServerError {
    return typeof obj === 'object' && obj !== null && 'TypServerErrorMessage' in obj;
}
