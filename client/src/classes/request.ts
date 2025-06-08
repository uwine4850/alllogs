import { isLoginResponseMessage, type LoginResponseMessage } from '@/dto/auth'
import router from '@/router'
import type { AxiosRequestConfig, AxiosResponse } from 'axios'
import axios from 'axios'

export class AsyncRequest<D = any> {
  protected url: string
  protected config: AxiosRequestConfig
  protected response: AxiosResponse | undefined
  protected data: D | undefined
  protected currentRequest: (() => Promise<void>) | undefined
  protected onResponseFn: ((response: AxiosResponse) => void) | undefined
  protected onErrorFn: ((error: unknown) => void) | undefined
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

  public onError(fn: (error: unknown) => void) {
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
    } catch (error) {
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
    } catch (error) {
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
    } catch (error) {
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
    } catch (error) {
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
          if (loginResponse.Error != '') {
            sessionStorage.removeItem('authJWT')
            sessionStorage.removeItem('profile')
            router.push('/login')
            if (this.onErrorFn) {
              this.onErrorFn(loginResponse.Error)
            }
            return
          } else if (loginResponse.JWT != '') {
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
}
