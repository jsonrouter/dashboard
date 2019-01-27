import { NgModule } from '@angular/core'
import { FlexLayoutModule } from '@angular/flex-layout'
import { CommonModule } from '@angular/common'
import { EndpointsComponent } from '../../comps/pages/endpoints/endpoints.component'
import { EndpointsTreeComponent } from '../../comps/trees/endpoints-tree/endpoints-tree.component'
import { MatTreeModule, MatIconModule, MatButtonModule, MatInputModule, MatSelectModule, MatRadioModule, MatCardModule } from '@angular/material'
import { TestTreeComponent } from '../../comps/trees/test-tree/test-tree.component'
import { AddressFormComponent } from '../../comps/forms/address-form/address-form.component'
import { ReactiveFormsModule } from '@angular/forms'
import { HomeComponent } from '../../comps/pages/home/home.component'
import { MaterialModule } from '../../modules/material/material.module'

@NgModule({
  declarations: [EndpointsComponent, EndpointsTreeComponent, TestTreeComponent, AddressFormComponent, HomeComponent],
  imports: [
    CommonModule,
    MaterialModule,
    FlexLayoutModule,
    MatTreeModule,
    MatIconModule,
    MatButtonModule,
    MatInputModule,
    MatSelectModule,
    MatRadioModule,
    MatCardModule,
    ReactiveFormsModule
  ],
  // exports: [
  //   MaterialModule,
  //   CommonModule,
  //   MaterialModule,
  //   FlexLayoutModule,
  //   MatTreeModule,
  //   MatIconModule,
  //   MatButtonModule,
  //   MatInputModule,
  //   MatSelectModule,
  //   MatRadioModule,
  //   MatCardModule,
  //   ReactiveFormsModule
  // ]
})
export class SharedModule { }
