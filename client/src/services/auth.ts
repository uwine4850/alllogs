import { AsyncRequestWithAuthorization } from "@/classes/request"
import type { MsgLogout } from "@/dto/auth"
import type { MsgBaseResponse } from "@/dto/common"
import router from "@/router"
import type { AxiosResponse } from "axios"


export async function logout(UID: number){
  const formData: MsgLogout = {
    UID: UID,
  }
    const req = new AsyncRequestWithAuthorization(`http://localhost:8000/logout`, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        withCredentials: true,
      })
      req.onResponse((response: AxiosResponse) => {
        const profileResponse = response.data as MsgBaseResponse
        if (profileResponse.Error !== '') {
          throw profileResponse.Error
        } else {
          sessionStorage.removeItem('authJWT')
          sessionStorage.removeItem('profile')
          router.push("/login")
        }
      })
      req.onError((error: unknown) => {
        throw error
      })
    req.setData(formData)
    await req.post()
}