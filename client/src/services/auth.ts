import { AsyncRequestWithAuthorization } from "@/classes/request"
import type { LogoutMessage } from "@/dto/auth"
import type { BaseResponseMessage } from "@/dto/common"
import router from "@/router"
import type { AxiosResponse } from "axios"
import { useRouter } from "vue-router"


export async function logout(UID: number){
  const formData: LogoutMessage = {
    UID: UID,
  }
    const req = new AsyncRequestWithAuthorization(`http://localhost:8000/logout`, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
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
          router.push("/login")
        }
      })
      req.onError((error: unknown) => {
        throw error
      })
    req.setData(formData)
    await req.post()
}