import {ChangeDetectionStrategy, Component} from '@angular/core'

@Component({
  standalone: true,
  selector: 'client-home',
  styleUrl: './home.component.css',
  templateUrl: './home.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HomeComponent {}
