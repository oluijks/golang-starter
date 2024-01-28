import {ChangeDetectionStrategy, Component} from '@angular/core'

@Component({
  standalone: true,
  selector: 'client-not-found',
  styleUrl: './not-found.component.css',
  templateUrl: './not-found.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class NotFoundComponent {}
