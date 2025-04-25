import type { AuthorizationMessage, LoginResponseMessage } from './auth'

export function IsLoginResponseMessage(data: any): data is LoginResponseMessage {
  const keys = Object.keys(data)
  const expectedKeys = ['JWT', 'UID', 'Error']
  return (
    expectedKeys.every((key) => typeof data[key] === 'string') &&
    keys.length === expectedKeys.length
  )
}
