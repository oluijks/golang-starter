import {computed, inject} from '@angular/core'
import {tapResponse} from '@ngrx/operators'
import {
  patchState,
  signalStore,
  withComputed,
  withHooks,
  withMethods,
  withState,
} from '@ngrx/signals'
import {rxMethod} from '@ngrx/signals/rxjs-interop'
import {pipe, switchMap, tap} from 'rxjs'

import {throttleInput} from '../shared/operators/throttle-input'

import {ItemsService} from './items.service'
import {
  ItemState,
  ItemsState,
  initialItemState,
  initialItemsState,
} from './items.types'

export const ItemsStore = signalStore(
  {providedIn: 'root'},
  withState<ItemState>(initialItemState),
  withMethods((store, itemsService = inject(ItemsService)) => ({
    loadItem: rxMethod<string>(
      pipe(
        tap(() => patchState(store, {isLoading: true})),
        switchMap((id: string) =>
          itemsService.fetchItem(id).pipe(
            tapResponse({
              next: item => patchState(store, {item: item}),
              error: console.error,
              finalize: () => patchState(store, {isLoading: false}),
            })
          )
        )
      )
    ),
  })),
  withState<ItemsState>(initialItemsState),
  withComputed(store => ({
    itemsCount: computed(() => store.items().length),
  })),
  withMethods((store, itemsService = inject(ItemsService)) => ({
    loadItems: rxMethod<void>(
      pipe(
        throttleInput(),
        tap(() => patchState(store, {isLoading: true})),
        switchMap(() =>
          itemsService.fetchItems().pipe(
            tapResponse({
              next: items => patchState(store, {items: items}),
              error: console.error,
              finalize: () => patchState(store, {isLoading: false}),
            })
          )
        )
      )
    ),
  })),
  withHooks({
    onInit({loadItems}) {
      loadItems()
    },
  })
)
