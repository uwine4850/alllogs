
export interface ProjectMessage {
    TypProjectMessage?: unknown;
    Author: ProjectAuthor | undefined;
    Id: number;
    UserId: number;
    Name: string;
    Description: string;
    Error: string;
}
export function isProjectMessage(obj: any): obj is ProjectMessage {
    return typeof obj === 'object' && obj !== null && 'TypProjectMessage' in obj;
}
export interface ProjectAuthor {
    TypProjectAuthor?: unknown;
    UID: number;
    Username: string;
    Avatar: string;
}
export function isProjectAuthor(obj: any): obj is ProjectAuthor {
    return typeof obj === 'object' && obj !== null && 'TypProjectAuthor' in obj;
}
export interface ProjectLogGroupMessage {
    TypProjectLogGroupMessage?: unknown;
    Id: number;
    ProjectId: number;
    Name: string;
    Description: string;
    Error: string;
}
export function isProjectLogGroupMessage(obj: any): obj is ProjectLogGroupMessage {
    return typeof obj === 'object' && obj !== null && 'TypProjectLogGroupMessage' in obj;
}
