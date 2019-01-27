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
    this.end.getLocalEndpoints()
      .subscribe((data) => {
        console.log(data)
        this.endpoints = data
      })
    // this.end.getEndpoints()
    //   .subscribe((data) => {
    //     console.log(data)
    //     this.endpoints = data
    //   }, (error) => {
    //     console.error('could not get the endpoints', error.error)
    //     this.end.getLocalEndpoints()
    //       .subscribe((data) => {
    //         console.log('got local endpoint', data)
    //         this.endpoints = data
    //       })
    //   })
  }

  keys(obj): Array<string> {
    return Object.keys(obj)
  }

  ngOnInit() {
  }

}
