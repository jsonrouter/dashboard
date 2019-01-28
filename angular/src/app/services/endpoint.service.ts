import { Injectable } from '@angular/core'
import { Observable } from 'rxjs/internal/Observable'
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http'
import * as env from '../../environments/environment'

@Injectable({
  providedIn: 'root'
})
export class EndpointService {

  constructor(
    private http: HttpClient,
  ) { }

  getEndpoints(): Observable<any> {
    const url = '/openapi.spec.json'
    return this.http.get(url)
  }

  getLocal(): Observable<any> {
    const url = '../../endpoints.json'
    return this.http.get(url)
  }
}
