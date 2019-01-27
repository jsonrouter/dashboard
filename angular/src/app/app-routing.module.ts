import { NgModule } from '@angular/core'
import { Routes, RouterModule } from '@angular/router'
import { EndpointsComponent } from './comps/pages/endpoints/endpoints.component'
import { SidebarComponent } from './sidebar/sidebar.component'
import { HomeComponent } from './comps/pages/home/home.component'

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'endpoints', component: EndpointsComponent },

  { path: '**', redirectTo: '', pathMatch: 'full' },
]

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
