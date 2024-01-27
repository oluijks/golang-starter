export interface ItemsApiResponse {
  status: number
  message: string
  data: Items
}

export interface ItemApiResponse {
  status: number
  message: string
  data: {
    item: Item
  }
}

export interface Items {
  items: Item[]
}

export interface Item {
  _id: string
  title: string
  created_at: string
  updated_at: string
}

export type ItemsState = {
  items: Item[]
  isLoading: boolean
  error: string | null
}

export const initialItemsState: ItemsState = {
  items: [] as Item[],
  isLoading: false,
  error: null,
}

export type ItemState = {
  item: Item
  isLoading: boolean
  error: string | null
}

export const initialItemState: ItemState = {
  item: {} as Item,
  isLoading: false,
  error: null,
}
