
export interface ProjectMessage {
    Author: ProjectAuthor;
    UserId: string;
    Name: string;
    Description: string;
    Error: string;
}
export interface ProjectAuthor {
    PID: string;
    Username: string;
    Avatar: string;
}
