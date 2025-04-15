import type { AxiosRequestConfig, AxiosResponse } from "axios";
import axios from "axios";

export class AsyncRequest{
    private url: string;
    private config: AxiosRequestConfig;
    private onResponseFn: ((response: AxiosResponse) => void) | undefined;
    private onErrorFn: ((error: unknown) => void) | undefined;
    constructor(url: string, config: AxiosRequestConfig){
        this.url = url;
        this.config = config;
    }

    public onResponse(fn: (response: AxiosResponse) => void){
        this.onResponseFn = fn;
    }

    public onError(fn: (error: unknown) => void){
        this.onErrorFn = fn;
    }

    public async get(){
        try{
            const response = await axios.get(this.url, this.config);
            if (this.onResponseFn){
                this.onResponseFn(response);
            }
        } catch (error){
            if (this.onErrorFn){
                this.onErrorFn(error);
            }
        }
    }

    public async post<D = any>(data: D){
        try{
            const response = await axios.post(this.url, data, this.config);
            if (this.onResponseFn){
                this.onResponseFn(response);
            }
        } catch (error){
            if (this.onErrorFn){
                this.onErrorFn(error);
            }
        }
    }
}