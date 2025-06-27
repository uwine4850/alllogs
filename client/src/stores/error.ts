import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useErrorStore = defineStore('error', () => {
  const errors = ref<Record<string, HTMLElement | null>>({})
  const customErrorTextId = ref("text")

  function setCustomErrorTextId(idValue: string){
    customErrorTextId.value = idValue
  }

  function setText(txt: string, id: string = "errorEl") {
    if (errors.value[id]){
      const errorElement = errors.value[id]
      const textPlace = errorElement.querySelector('span') as HTMLElement
      if (textPlace) {
        textPlace.innerText = txt
        errorElement.style.display = 'flex'
      }
    }
  }

  function setErrorElement(el: HTMLElement, id: string = "errorEl") {
    errors.value[id] = el
  }
  return { setText, setErrorElement, setCustomErrorTextId }
})
