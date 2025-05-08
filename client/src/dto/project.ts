
export interface ProjectMessage {
    Author: ProjectAuthor | undefined;
    Id: number;
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
export interface ProjectLogGroupMessage {
    Id: number;
    ProjectId: number;
    Name: string;
    Description: string;
    Error: string;
}
