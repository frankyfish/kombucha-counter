import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Dashboard } from "./dashboard/dashboard";
import { BackendService } from './services/backend-service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Dashboard],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected title = 'Kombucha Counter 🍺';
  private backendService = inject(BackendService) // NOTICE: no provider configured
}
