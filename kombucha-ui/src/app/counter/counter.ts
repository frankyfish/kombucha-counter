import { Component, inject } from '@angular/core';
import { BackendService } from '../services/backend-service';

@Component({
  selector: 'app-counter',
  imports: [],
  templateUrl: './counter.html',
  styleUrl: './counter.css'
})
export class Counter {
  private backendService = inject(BackendService)

  onCounterClick() {
    this.backendService.incCount()
  }


}
