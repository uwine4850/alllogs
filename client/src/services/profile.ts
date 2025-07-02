import type { MsgGenToken, MsgProfile, MsgTokenResponse } from '@/dto/profile'
import { AsyncRequestWithAuthorization } from '@/classes/request'
import type { MsgBaseResponse } from '@/dto/common'
import { AxiosError, type AxiosResponse } from 'axios'
import { ref, type Ref } from 'vue'
import { useErrorStore } from '@/stores/error'

export const getProfileData = async (
  _profileDataRef: Ref<MsgProfile | null>,
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
    const profileResponse = response.data as MsgProfile
    if (profileResponse.Error !== '') {
      errorStore.setText(profileResponse.Error)
    } else {
      _profileDataRef.value = profileResponse
      if (_tokenRef) {
        _tokenRef.value = _profileDataRef.value.Token
      }
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore)
  await req.get()
}

export const generateTokenForm = async (
  _tokenRef: Ref<string | null>,
  profileData: MsgProfile | null,
  errorStore: ReturnType<typeof useErrorStore>,
) => {
  if (!profileData) {
    return
  }
  const formData = ref<MsgGenToken>({
    UserId: profileData.UserId,
  })

  const req = new AsyncRequestWithAuthorization('http://localhost:8000/gen-token', {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    withCredentials: true,
  })
  req.onResponse((response: AxiosResponse) => {
    const tokenResponse = response.data as MsgTokenResponse
    if (tokenResponse.Error !== '') {
      errorStore.setText(tokenResponse.Error)
    } else {
      _tokenRef.value = tokenResponse.Token
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore)
  req.setData(formData.value)
  req.post()
}

export const deleteToken = async (
  uid: any,
  _tokenRef: Ref<string | null>,
  profileData: MsgProfile | null,
  errorStore: ReturnType<typeof useErrorStore>,
) => {
  if (!profileData) {
    return
  }
  const req = new AsyncRequestWithAuthorization(`http://localhost:8000/del-token/user/${uid}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    withCredentials: true,
  })
  req.onResponse((response: AxiosResponse) => {
    const baseResponse = response.data as MsgBaseResponse
    if (baseResponse.Error != '') {
      errorStore.setText(baseResponse.Error)
    } else {
      _tokenRef.value = ''
    }
  })
  req.onError((error: AxiosError) => {
    errorStore.setText("unexpected error: " + error.message)
  }, errorStore)
  req.delete()
}
