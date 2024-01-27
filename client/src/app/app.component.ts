import {Component} from '@angular/core'
import {RouterOutlet} from '@angular/router'

@Component({
  standalone: true,
  selector: 'client-root',
  styleUrl: './app.component.css',
  templateUrl: './app.component.html',
  imports: [RouterOutlet],
})
export class AppComponent {
  title = 'client'
}
