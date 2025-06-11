import { isLoginResponseMessage, type LoginResponseMessage } from '@/dto/auth'
import { isClientErrorMessage, type ClientErrorMessage, type ServerErrorMessage } from '@/dto/common'
import router from '@/router'
import type { useErrorStore } from '@/stores/error'
import type { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'
import axios from 'axios'

export class AsyncRequest<D = any> {
  protected url: string
  protected config: AxiosRequestConfig
  protected response: AxiosResponse | undefined
  protected data: D | undefined
  protected currentRequest: (() => Promise<void>) | undefined
  protected onResponseFn: ((response: AxiosResponse) => void) | undefined
  protected onErrorFn: ((error: AxiosError) => void) | undefined
  constructor(url: string, config: AxiosRequestConfig) {
    this.url = url
    this.config = config
    if (!this.config.headers) {
      this.config.headers = {}
    }
    this.config.headers.Authorization = sessionStorage.getItem('authJWT')
  }

  protected updateAuthToken() {
    if (this.config.headers) {
      this.config.headers.Authorization = sessionStorage.getItem('authJWT')
    }
  }

  public setData(data: D) {
    this.data = data
  }

  public onResponse(fn: (response: AxiosResponse) => void) {
    this.onResponseFn = fn
  }

  public onError(fn: (error: AxiosError) => void) {
    this.onErrorFn = fn
  }

  public async get() {
    this.currentRequest = this.get.bind(this)
    try {
      const response = await axios.get(this.url, this.config)
      this.response = response
      if (this.onResponseFn) {
        this.onResponseFn(response)
      }
    } catch (error: any) {
      if (this.onErrorFn) {
        this.onErrorFn(error)
      }
    }
  }

  public async post() {
    this.currentRequest = this.post.bind(this)
    try {
      const response = await axios.post(this.url, this.data, this.config)
      this.response = response
      if (this.onResponseFn) {
        this.onResponseFn(response)
      }
    } catch (error: any) {
      if (this.onErrorFn) {
        this.onErrorFn(error)
      }
    }
  }

  public async delete() {
    this.currentRequest = this.delete.bind(this)
    try {
      const cnf = this.config
      const response = await axios.delete(this.url, cnf)

      this.response = response
      if (this.onResponseFn) {
        this.onResponseFn(response)
      }
    } catch (error: any) {
      if (this.onErrorFn) {
        this.onErrorFn(error)
      }
    }
  }

  public async put() {
    this.currentRequest = this.put.bind(this)
    try {
      const response = await axios.put(this.url, this.data, this.config)
      this.response = response
      if (this.onResponseFn) {
        this.onResponseFn(response)
      }
    } catch (error: any) {
      if (this.onErrorFn) {
        this.onErrorFn(error)
      }
    }
  }
  
  public async patch() {
    this.currentRequest = this.patch.bind(this)
    try {
      const response = await axios.patch(this.url, this.data, this.config)
      this.response = response
      if (this.onResponseFn) {
        this.onResponseFn(response)
      }
    } catch (error: any) {
      if (this.onErrorFn) {
        this.onErrorFn(error)
      }
    }
  }
}

export class AsyncRequestWithAuthorization extends AsyncRequest {
  public onResponse(fn: (response: AxiosResponse) => void) {
    this.onResponseFn = async () => {
      if (this.response) {
        if (isLoginResponseMessage(this.response.data)) {
          const loginResponse = this.response.data as LoginResponseMessage
          if (loginResponse.JWT != '') {
            sessionStorage.setItem('authJWT', loginResponse.JWT)
            this.updateAuthToken()
            if (this.currentRequest) {
              await this.currentRequest()
            }
            return
          }
        }
        fn(this.response)
      }
    }
  }
  public override onError(fn: (error: AxiosError) => void, errorStore?: ReturnType<typeof useErrorStore>) {
    this.onErrorFn = async (error: AxiosError) => {
      if(error) {
        if (error.response?.data && isClientErrorMessage(error.response.data)) {
          const clientErrorMessage = error.response.data as ClientErrorMessage
          catchClientError(clientErrorMessage, errorStore)
          return
        }
        fn(error)
      }
    }
  }
}

export function catchClientError(data: ClientErrorMessage, errorStore?: ReturnType<typeof useErrorStore>){
  switch (data.Code){
    case 400:
      errorStore?.setText(data.Text)
      break;
    case 401:
      sessionStorage.removeItem('authJWT')
      sessionStorage.removeItem('profile')
      router.push('/login')
      break;
    case 403:
      router.push(`/error?code=${"403 Forbidden"}&text=${data.Text}`)
      break;
    case 409:
      errorStore?.setText(data.Text)
      break;
  }
}

export function catchServerError(data: ServerErrorMessage, errorStore?: ReturnType<typeof useErrorStore>){
  errorStore?.setText(data.Text)
}