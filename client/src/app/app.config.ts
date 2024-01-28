import {ApplicationConfig} from '@angular/core'
import {provideHttpClient} from '@angular/common/http'
import {provideRouter} from '@angular/router'

import {DOCUMENT, WINDOW} from './app.tokens'

import {appRoutes} from './app.routes'

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(appRoutes),
    provideHttpClient(),
    {
      provide: WINDOW,
      useFactory: () => window,
    },
    {
      provide: DOCUMENT,
      useFactory: () => document,
    },
  ],
}
