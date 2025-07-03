import { Component, inject } from '@angular/core';
import { BackendService } from '../services/backend-service';
import { output } from '@angular/core';

@Component({
  selector: 'app-counter',
  imports: [],
  templateUrl: './counter.html',
  styleUrl: './counter.css'
})
export class Counter {
  private backendService = inject(BackendService)
  counterClicked = output<void>();

  onCounterClick() {
    this.backendService.incCount()
    this.counterClicked.emit()
  }

}
