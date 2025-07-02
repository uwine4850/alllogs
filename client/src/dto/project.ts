
export interface MsgProject {
    TypProjectMessage?: unknown;
    Author: MsgProjectAuthor | undefined;
    Id: number;
    UserId: number;
    Name: string;
    Description: string;
    Error: string;
}
export function isMsgProject(obj: any): obj is MsgProject {
    return typeof obj === 'object' && obj !== null && 'TypProjectMessage' in obj;
}
export interface MsgProjectAuthor {
    TypProjectAuthor?: unknown;
    UID: number;
    Username: string;
    Avatar: string;
}
export function isMsgProjectAuthor(obj: any): obj is MsgProjectAuthor {
    return typeof obj === 'object' && obj !== null && 'TypProjectAuthor' in obj;
}
export interface MsgProjectLogGroup {
    TypProjectLogGroupMessage?: unknown;
    Id: number;
    ProjectId: number;
    Name: string;
    Description: string;
    Error: string;
    AuthorToken: string;
}
export function isMsgProjectLogGroup(obj: any): obj is MsgProjectLogGroup {
    return typeof obj === 'object' && obj !== null && 'TypProjectLogGroupMessage' in obj;
}
export interface MsgLogItem {
    TypLogItemMessage?: unknown;
    Type: number;
    Token: string;
    Error: string;
    Payload: MsgLogItemPayload | undefined;
}
export function isMsgLogItem(obj: any): obj is MsgLogItem {
    return typeof obj === 'object' && obj !== null && 'TypLogItemMessage' in obj;
}
export interface MsgLogItemPayload {
    TypLogItemPayload?: unknown;
    Id: number;
    LogGroupId: number;
    Text: string;
    Type: string;
    Tag: string;
    Datetime: string;
}
export function isMsgLogItemPayload(obj: any): obj is MsgLogItemPayload {
    return typeof obj === 'object' && obj !== null && 'TypLogItemPayload' in obj;
}
export interface MsgLogItemsFilter {
    TypLogItemsFilter?: unknown;
    Text: string;
    Type: string;
    Tag: string;
    DateTime: string;
}
export function isMsgLogItemsFilter(obj: any): obj is MsgLogItemsFilter {
    return typeof obj === 'object' && obj !== null && 'TypLogItemsFilter' in obj;
}
