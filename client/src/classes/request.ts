import { isLoginResponseMessage, type LoginResponseMessage } from '@/dto/auth'
import { isClientErrorMessage, isServerErrorMessage, type BaseResponseMessage, type ClientErrorMessage, type ServerErrorMessage } from '@/dto/common'
import router from '@/router'
import type { useErrorStore } from '@/stores/error'
import { AxiosError, type AxiosRequestConfig, type AxiosResponse } from 'axios'
import axios from 'axios'

export class AsyncRequest<D = any> {
  protected url: string
  protected config: AxiosRequestConfig
  protected response: AxiosResponse | undefined
  protected data: D | undefined
  protected withCSRF: boolean
  protected currentRequest: (() => Promise<void>) | undefined
  protected onResponseFn: ((response: AxiosResponse) => void) | undefined
  protected onErrorFn: ((error: AxiosError) => void) | undefined
  constructor(url: string, config: AxiosRequestConfig, withCSRF: boolean = true) {
    this.url = url
    this.config = config
    this.withCSRF = withCSRF
    if (!this.config.headers) {
      this.config.headers = {}
    }
    this.config.headers.Authorization = sessionStorage.getItem('authJWT')
  }

  protected async csrfToken(){
    if (!this.withCSRF){
      return
    }
    const csrfToken = getCookie("CSRF_TOKEN")
    if(!csrfToken){
      try {
        await setToken()
      } catch (err: any) {
        const axiosError = axios.isAxiosError(err)
        ? err
        : new AxiosError(err || String(err))
        if(this.onErrorFn){
          this.onErrorFn(axiosError)
        }
      } finally{
        const newToken = getCookie("CSRF_TOKEN")
        this.config.headers = {
          ...this.config.headers,
          'X-CSRF-TOKEN': newToken,
        }
      }
    } else {
      this.config.headers = {
        ...this.config.headers,
        'X-CSRF-TOKEN': csrfToken,
      }
    }
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
    await this.csrfToken()
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
    await this.csrfToken()
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
    await this.csrfToken()
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
    await this.csrfToken()
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
    await this.csrfToken()
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
        if (error.response?.data && isServerErrorMessage(error.response.data)) {
          const serverErrorMessage = error.response.data as ServerErrorMessage
          catchServerError(serverErrorMessage, errorStore)
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

export function getCookie(name: string): string | null {
  const value = `; ${document.cookie}`
  const parts = value.split(`; ${name}=`)
  if (parts.length === 2) {
    return parts.pop()!.split(';').shift() || null
  }
  return null
}

async function setToken(): Promise<void> {
  const req = new AsyncRequest(`http://localhost:8000/set-csrf`, {
      headers: {
        'Content-Type': 'application/json',
      },
      withCredentials: true,
  }, false)
  return new Promise((resolve, reject) => {
    req.onResponse((response: AxiosResponse) => {
      const baseResponse = response.data as BaseResponseMessage
      if (baseResponse.Error && baseResponse.Error == '') {
        reject(baseResponse.Error)
      } else {
        resolve()
      }
    })
    
    req.onError((error: AxiosError) => {
      reject(error)
    })

    req.get()
  })
}