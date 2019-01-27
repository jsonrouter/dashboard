import { Injectable } from '@angular/core'
import { Observable } from 'rxjs/internal/Observable'
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http'
import * as env from '../../environments/environment'
const environment = env.environment

@Injectable({
  providedIn: 'root'
})
export class EndpointService {

  constructor(
    private http: HttpClient,
  ) { }

  getEndpoints(): Observable<any> {
    const url = environment.apiEndpoint
    return this.http.get(url)
  }

  getLocalEndpoints(): Observable<any> {
    const url = '../../assets/endpoints.json'
    return this.http.get(url)
  }
}