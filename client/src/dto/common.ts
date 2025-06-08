
export interface BaseResponseMessage {
    TypBaseResponseMessage?: unknown;
    Ok: boolean;
    Error: string;
}
export function isBaseResponseMessage(obj: any): obj is BaseResponseMessage {
    return typeof obj === 'object' && obj !== null && 'TypBaseResponseMessage' in obj;
}
