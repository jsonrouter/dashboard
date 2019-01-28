import { Component, OnInit } from '@angular/core'
import { EndpointService } from '../../../services/endpoint.service'

@Component({
  selector: 'app-endpoints',
  templateUrl: './endpoints.component.html',
  styleUrls: ['./endpoints.component.sass']
})
export class EndpointsComponent implements OnInit {
  endpoints: any

  constructor(
    private end: EndpointService,
  ) {
    this.end.getEndpoints()
      .subscribe((data) => {
        console.log(data)
        this.endpoints = data
      }, (error) => {
        console.error('could not get the endpoints', error.error)
      })
  }

  keys(obj): Array<string> {
    return Object.keys(obj)
  }

  ngOnInit() {
  }

}
