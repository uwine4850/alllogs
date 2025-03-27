import { h, render, type Component, type Ref } from 'vue'

export function addComponent(placeID: string, component: Component) {
  const container = document.getElementById(placeID) as HTMLElement
  if (!container) {
    console.error('Container not found')
    return
  }
  clearHTMLElement(container)
  const vnode = h(component)
  render(vnode, container)
}

export function clearHTMLElement(element: HTMLElement) {
  render(null, element);
}
