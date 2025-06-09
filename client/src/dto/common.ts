
export interface BaseResponseMessage {
    TypBaseResponseMessage?: unknown;
    Ok: boolean;
    Error: string;
}
export function isBaseResponseMessage(obj: any): obj is BaseResponseMessage {
    return typeof obj === 'object' && obj !== null && 'TypBaseResponseMessage' in obj;
}
export interface ClientErrorMessage {
    TypClientErrorMessage?: unknown;
    Code: number;
    Text: string;
}
export function isClientErrorMessage(obj: any): obj is ClientErrorMessage {
    return typeof obj === 'object' && obj !== null && 'TypClientErrorMessage' in obj;
}
