import { Component, inject } from '@angular/core';
import { BackendService } from '../services/backend-service';

@Component({
  selector: 'app-dashboard',
  imports: [],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css'
})
export class Dashboard {
  private backendService = inject(BackendService)

  drinkedBottles: number = 0

  // loading stats on startup 
  ngOnInit() {
    console.log("running onInit for Dashboard")
    this.backendService.getCurrentStats().subscribe({
      next: (response) => {
        console.log("number of bottles: ", response)
        this.drinkedBottles = response
      }
    })
  }

}
