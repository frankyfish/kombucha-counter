import { Component, inject } from '@angular/core';
import { BackendService } from '../services/backend-service';
import { Counter } from '../counter/counter';

@Component({
  selector: 'app-dashboard',
  imports: [Counter],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css'
})
export class Dashboard {
  private backendService = inject(BackendService)

  drinkedBottles: number = 0
  mlDrank: number = 0
  moneySaved: number = 0

  // loading stats on startup 
  ngOnInit() {
    console.log("running onInit for Dashboard")
    this.loadStats()
  }

  private loadStats() {
    this.backendService.getCurrentStats().subscribe({
      next: (response) => {
        console.log("current stats: ", response)
        this.drinkedBottles = response.count
        this.mlDrank = response.ml
        this.moneySaved = response.saved
      }
    })
  }

  onCounterClicked() {
    this.loadStats()
  }

}
