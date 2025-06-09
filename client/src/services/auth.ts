import { AsyncRequestWithAuthorization } from "@/classes/request"
import type { LogoutMessage } from "@/dto/auth"
import type { BaseResponseMessage } from "@/dto/common"
import type { AxiosResponse } from "axios"

export async function logout(AID: number){
    const formData: LogoutMessage = {
        AID: AID,
    }
    const req = new AsyncRequestWithAuthorization(`http://localhost:8000/logout`, {
        headers: {
          'Content-Type': 'application/json',
        },
        withCredentials: true,
      })
      req.onResponse((response: AxiosResponse) => {
        const profileResponse = response.data as BaseResponseMessage
        if (profileResponse.Error !== '') {
            throw profileResponse.Error
        } else {
            sessionStorage.removeItem('authJWT')
            sessionStorage.removeItem('profile')
        }
      })
      req.onError((error: unknown) => {
        throw error
      })
    req.setData(formData)
    await req.post()
}