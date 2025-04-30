import type { GenTokenMessage, ProfileMessage, TokenResponse } from '@/dto/profile'
import { AsyncRequestWithAuthorization } from '@/classes/request'
import type { BaseResponseMessage } from '@/dto/common'
import { type AxiosResponse } from 'axios'
import { ref, type Ref } from 'vue'
import { useErrorStore } from '@/stores/error'

export const getProfileData = async (
  _profileDataRef: Ref<ProfileMessage | null>,
  _tokenRef: Ref<string | null> | null,
  id: string,
  errorStore: ReturnType<typeof useErrorStore>,
) => {
  const req = new AsyncRequestWithAuthorization(`http://localhost:8000/profile/${id}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    withCredentials: true,
  })
  req.onResponse((response: AxiosResponse) => {
    const profileResponse = response.data as ProfileMessage
    if (profileResponse.Error !== '') {
      errorStore.setText(profileResponse.Error)
    } else {
      _profileDataRef.value = profileResponse
      if (_tokenRef) {
        _tokenRef.value = _profileDataRef.value.Token
      }
    }
  })
  req.onError((error: unknown) => {
    errorStore.setText(String(error))
  })
  await req.get()
}

export const generateTokenForm = async (
  _tokenRef: Ref<string | null>,
  profileData: ProfileMessage | null,
  errorStore: ReturnType<typeof useErrorStore>,
) => {
  if (!profileData) {
    return
  }
  const formData = ref<GenTokenMessage>({
    UserId: profileData.Id,
  })

  const req = new AsyncRequestWithAuthorization('http://localhost:8000/gen-token', {
    headers: {
      'Content-Type': 'application/json',
    },
    withCredentials: true,
  })
  req.onResponse((response: AxiosResponse) => {
    const tokenResponse = response.data as TokenResponse
    if (tokenResponse.Error !== '') {
      errorStore.setText(tokenResponse.Error)
    } else {
      _tokenRef.value = tokenResponse.Token
    }
  })
  req.onError((error: unknown) => {
    errorStore.setText(String(error))
  })
  req.setData(formData.value)
  req.post()
}

export const deleteToken = async (
  _tokenRef: Ref<string | null>,
  profileData: ProfileMessage | null,
  errorStore: ReturnType<typeof useErrorStore>,
) => {
  if (!profileData) {
    return
  }
  const req = new AsyncRequestWithAuthorization('http://localhost:8000/del-token', {
    headers: {
      'Content-Type': 'application/json',
    },
    withCredentials: true,
  })
  req.onResponse((response: AxiosResponse) => {
    const baseResponse = response.data as BaseResponseMessage
    if (baseResponse.Error != '') {
      errorStore.setText(baseResponse.Error)
    } else {
      _tokenRef.value = ''
    }
  })
  req.onError((error: unknown) => {
    errorStore.setText(String(error))
  })
  req.delete()
}
