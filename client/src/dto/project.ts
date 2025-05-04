
export interface ProjectMessage {
    Author: ProjectAuthor | undefined;
    UserId: number;
    Name: string;
    Description: string;
    Error: string;
}
export interface ProjectAuthor {
    PID: number;
    Username: string;
    Avatar: string;
}
