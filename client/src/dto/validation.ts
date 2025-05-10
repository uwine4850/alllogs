import type { LoginResponseMessage } from './auth'

export function IsLoginResponseMessage(data: any): data is LoginResponseMessage {
  const keys = Object.keys(data)
  const expectedKeys = ['JWT', 'UID', 'Error']

  return (
    typeof data.JWT === 'string' &&
    typeof data.UID === 'number' &&
    typeof data.Error === 'string' &&
    keys.length === expectedKeys.length
  )
}
