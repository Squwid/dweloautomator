import { CUSTOM_ELEMENTS_SCHEMA, NgModule } from '@angular/core';
import { MatSelectModule } from '@angular/material';
import { MatButtonModule } from '@angular/material/button';
// stuff for material
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgxMaterialTimepickerModule } from 'ngx-material-timepicker';
import { AppComponent } from './app.component';
import { TimepageComponent } from './timepage/timepage.component';

@NgModule({
  declarations: [
    AppComponent,
    TimepageComponent
  ],
  imports: [
    BrowserModule,
    MatCardModule,
    MatFormFieldModule,
    MatSelectModule,
    BrowserAnimationsModule,
    NgxMaterialTimepickerModule,
    MatSnackBarModule,
    MatButtonModule
  ],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
