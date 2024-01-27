import {Pipe, PipeTransform} from '@angular/core'

import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
dayjs.extend(relativeTime)

@Pipe({
  name: 'dateFromNow',
  standalone: true,
})
export class DateFromNowPipe implements PipeTransform {
  transform(date: string): string {
    return dayjs(date).fromNow()
  }
}
