import { defineStore } from "pinia";
import { ref } from 'vue';

export const useErrorStore = defineStore("error", () => {
    const errorEl = ref<HTMLElement | null>(null);

    function setText(txt: string){
        if (errorEl.value) {
            const textPlace = errorEl.value.querySelector("#text") as HTMLElement;
            if(textPlace){
                textPlace.innerText = txt;
                errorEl.value.style.display = "flex";
            }
        }
    }

    function setErrorElement(el: HTMLElement){
        errorEl.value = el;
    }
    return {setText, setErrorElement}
})