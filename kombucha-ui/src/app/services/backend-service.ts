import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Stats } from '../model/stats.model';

@Injectable({
  providedIn: 'root'
})
export class BackendService {
  private backendUrl = "http://localhost:8080"
  private httpClient = inject(HttpClient)

  constructor() { }

  /**
   * @deprecated will be removed in favour of {@link getCurrentStats}
   * @returns 
   */
  getCurrentCount(): Observable<number> {
    console.log("requesting count from backend")
    return this.httpClient.get<number>(this.backendUrl)
  }

  getCurrentStats(): Observable<Stats> {
    console.log("requesting stats from backend")
    return this.httpClient.get<Stats>(this.backendUrl + "/stats")
  }

}
