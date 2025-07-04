export class WrappedObserver {
  private className: string
  public intersectionObserver: IntersectionObserver | null = null
  public mutationObserver: MutationObserver | null = null
  private onTrigerFn: ((el: HTMLElement) => void) | undefined
  constructor(className: string) {
    this.className = className
  }

  public onTrigger(onTrigerFn: (el: HTMLElement) => void) {
    this.onTrigerFn = onTrigerFn
  }

  public observeNow(el: HTMLElement) {
    this.intersectionObserver?.disconnect()
    this.intersectionObserver = new IntersectionObserver(
      (entries) => {
        const entry = entries[0]
        if (entry.isIntersecting) {
          if (this.onTrigerFn) {
            this.onTrigerFn(el)
          }
        }
      },
      {
        threshold: 0.01,
      },
    )

    this.intersectionObserver.observe(el)
  }

  public watch() {
    this.mutationObserver = new MutationObserver(() => {
      const el = document.querySelector<HTMLElement>(`.${this.className}`)
      if (el) {
        this.observeNow(el)
      }
    })

    this.mutationObserver.observe(document.body, {
      childList: true,
      subtree: true,
      attributes: true,
      attributeFilter: ['class'],
    })
  }
}
