
export interface Register {
    Username: string;
    Password: string;
    RepeatPassword: string;
}
export interface Login {
    Username: string;
    Password: string;
}
export interface LoginResponse {
    JWT: string;
    Error: string;
}
