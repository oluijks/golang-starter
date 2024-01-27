import {debounceTime, distinctUntilChanged, pipe} from 'rxjs'

export const throttleInput = ({time} = {time: 100}) => {
  return pipe(debounceTime(time), distinctUntilChanged())
}
