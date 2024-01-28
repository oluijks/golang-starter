import {Component, OnInit, Renderer2, inject, isDevMode} from '@angular/core'
import {RouterOutlet} from '@angular/router'

import {DOCUMENT, WINDOW} from './app.tokens'

@Component({
  standalone: true,
  selector: 'client-root',
  styleUrl: './app.component.css',
  templateUrl: './app.component.html',
  imports: [RouterOutlet],
})
export class AppComponent implements OnInit {
  title = 'client'

  #className = 'debug-screens'
  #document: Document = inject(DOCUMENT)
  #renderer: Renderer2 = inject(Renderer2)

  ngOnInit(): void {
    if (isDevMode()) {
      this.#renderer.addClass(this.#document.body, this.#className)
    }
  }
}
