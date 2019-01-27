import { Component, ViewChild, ElementRef, OnDestroy } from '@angular/core'
import { Subscription } from 'rxjs'
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout'
import { Observable } from 'rxjs'
import { map } from 'rxjs/operators'
import { MatSidenav } from '@angular/material'

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.sass']
})
export class SidebarComponent implements OnDestroy {
  @ViewChild('sidenav') sidenav: MatSidenav
  isHandset: boolean = true
  breakSub: Subscription

  constructor(private breakpointObserver: BreakpointObserver) {
    this.breakSub = breakpointObserver.observe(Breakpoints.Handset).pipe(map(result => result.matches))
      .subscribe((data) => {
        console.log(data)
        this.isHandset = data
      })
  }

  close(): void {
    if (this.isHandset) {
      this.sidenav.close()
    }
  }

  ngOnDestroy() {
    this.breakSub.unsubscribe()
  }

}
