import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';

import { NgxChartsModule } from '@swimlane/ngx-charts';

import { RingChartComponent } from './ring-chart/ring-chart.component';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { StackedChartComponent } from './stacked-chart/stacked-chart.component';

@NgModule({
  declarations: [
    AppComponent,
    RingChartComponent,
    StackedChartComponent
  ],
  imports: [
    BrowserModule, 
    FormsModule,
    NgxChartsModule,
    BrowserAnimationsModule,    
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule { }
