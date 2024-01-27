import {NgIf} from '@angular/common'
import {Component, Renderer2, effect, inject, signal} from '@angular/core'

import {DOCUMENT, WINDOW} from '../../../../app.tokens'

enum DarkMode {
  Dark = 'dark',
  Light = 'light',
}

@Component({
  standalone: true,
  selector: 'client-theme-switcher',
  styleUrl: './theme-switcher.component.css',
  templateUrl: './theme-switcher.component.html',
  imports: [NgIf],
})
export class ThemeSwitcherComponent {
  #key = 'darkMode'
  #window: Window = inject(WINDOW)
  #document: Document = inject(DOCUMENT)
  #renderer: Renderer2 = inject(Renderer2)
  #htmlEl = this.#document.getElementsByTagName('html')[0]

  darkMode = signal<boolean | null>(
    JSON.parse(this.#window.localStorage.getItem(this.#key)!)
  )

  constructor() {
    effect(() => {
      this.#window.localStorage.setItem(
        this.#key,
        JSON.stringify(this.darkMode())
      )

      this.darkMode()
        ? this.#renderer.addClass(this.#htmlEl, DarkMode.Dark)
        : this.#renderer.removeClass(this.#htmlEl, DarkMode.Dark)
    })
  }
}
