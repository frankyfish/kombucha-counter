import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Dashboard } from "./dashboard/dashboard";
import { BackendService } from './services/backend-service';
import { Counter } from "./counter/counter";

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Dashboard, Counter],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected title = 'Kombucha Counter 🍺';
  private backendService = inject(BackendService) // NOTICE: no provider configured
}
