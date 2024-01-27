import {HttpClient, HttpErrorResponse} from '@angular/common/http'
import {Injectable, inject} from '@angular/core'

import {catchError, delay, map, throwError} from 'rxjs'

import {Item, ItemApiResponse, ItemsApiResponse} from './items.types'

@Injectable({providedIn: 'root'})
export class ItemsService {
  #httpClient = inject(HttpClient)

  fetchItem(id: string) {
    return this.#httpClient
      .get<ItemApiResponse>(`https://your-api.com/api/v1/items/${id}`)
      .pipe(
        delay(100),
        map(response => {
          return response.data.item as Item
        }),
        catchError(this.handleError)
      )
  }

  fetchItems() {
    return this.#httpClient
      .get<ItemsApiResponse>(`https://your-api.com/api/v1/items`)
      .pipe(
        delay(100),
        map(response => {
          return response.data.items as Item[]
        }),
        catchError(this.handleError)
      )
  }

  private handleError(error: HttpErrorResponse) {
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error)
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${error.status}, body was: `,
        error.error
      )
    }
    // Return an observable with a user-facing error message.
    return throwError(
      () => new Error('Something bad happened; please try again later.')
    )
  }
}
