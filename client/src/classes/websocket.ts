import { ref, onBeforeUnmount } from 'vue'

const websockets: { [key: string]: MyWebsocket } = {}

export function getSocket(key: string): MyWebsocket {
  return websockets[key]
}

export interface SockedMessage {
  Type: number
  AID: number
  Payload: any
}

export class MyWebsocket<msg = any> {
  public socket: WebSocket
  private onOpen: () => void = () => {}
  private onClose: () => void = () => {}
  private onMessage: (event: MessageEvent) => void = (event: MessageEvent) => {}
  public isConnected: boolean = false

  constructor(key: string, url: string) {
    this.socket = new WebSocket(url)
    websockets[key] = this
  }

  public OnOpen(fn: () => void) {
    this.onOpen = fn
  }

  public OnMessage(fn: (event: MessageEvent) => void) {
    this.onMessage = fn
  }

  public OnClose(fn: () => void) {
    this.onClose = fn
  }

  public Close() {
    if (this.socket.readyState === WebSocket.OPEN) {
      this.socket.close()
    }
  }

  public Send(data: msg) {
    if (this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify(data))
    }
  }

  public Watch() {
    this.socket.onopen = () => {
      this.isConnected = true
      this.onOpen()
    }
    this.socket.onclose = () => {
      this.isConnected = false
      this.onClose()
    }
    this.socket.onmessage = this.onMessage
  }
}
