import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BackendService {
  private httpClient = inject(HttpClient)

  constructor() { }

  getCurrentStats(): Observable<number> {
    console.log("requesting stats from backend")
    return this.httpClient.get<number>("http://localhost:8080")
  }

}
