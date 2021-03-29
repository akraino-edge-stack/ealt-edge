import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { EaltserviceService } from './ealtservice.service';

import {MatExpansionModule} from '@angular/material/expansion';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {MatCardModule} from '@angular/material/card';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import {MatMenuModule} from '@angular/material/menu';

import { HttpClientModule } from '@angular/common/http';

// import { MatIconModule, MatSidenavModule, MatListModule, MatButtonModule } from '@angular/material';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import {MatPaginatorModule} from '@angular/material/paginator';
import {MatTableModule} from '@angular/material/table';
import {MatSelectModule} from '@angular/material/select';
import {MatTooltipModule} from '@angular/material/tooltip';

import {MatCheckboxModule} from '@angular/material/checkbox';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MarkdownModule } from 'ngx-markdown';
import { NgxMdModule } from 'ngx-md';

import {MatDialogModule} from '@angular/material/dialog';
import {MatRadioModule} from '@angular/material/radio';

import { BotDetectCaptchaModule } from 'angular-captcha'; 
import { NgTerminalModule } from 'ng-terminal';

import { NgxPermissionsModule } from 'ngx-permissions';

import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import { jqxChartModule } from 'jqwidgets-ng/jqxchart';
import { MDBBootstrapModule } from 'angular-bootstrap-md';
import { ChartsModule, WavesModule } from 'angular-bootstrap-md';
import * as cors from "cors";


import 'hammerjs';
import 'particles.js';
import 'chart.js';
import { HomeComponent } from './home/home.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { ToastrModule } from 'ngx-toastr';

import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { CarouselModule } from 'ngx-bootstrap/carousel';

import {MatTabsModule} from '@angular/material/tabs';
import {MatStepperModule} from '@angular/material/stepper';
import { MenuItemComponent } from './menu-item/menu-item.component';
import { PcbComponent } from './pcb/pcb.component';
import { PcbimagesComponent } from './pcbimages/pcbimages.component';
import { PcboutputimagesComponent } from './pcboutputimages/pcboutputimages.component';


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    ToolbarComponent,
    MenuItemComponent,
    PcbComponent,
    PcbimagesComponent,
    PcboutputimagesComponent
  ],

  imports: [
    BrowserModule,
    AppRoutingModule,
    MatExpansionModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatCardModule,
    MatToolbarModule,
    MatIconModule,
    MatSidenavModule,
    MatListModule,
    MatButtonModule,
    HttpClientModule,
    MatSnackBarModule,
    MatMenuModule,
    MatTableModule,
    MatPaginatorModule,
    MatSelectModule,
    MatTooltipModule,
    ReactiveFormsModule,
    ToastrModule,
    BotDetectCaptchaModule,
    NgbModule,
    CarouselModule.forRoot(),
    MatTabsModule,
    MatStepperModule,
    MatCheckboxModule,
    NgTerminalModule,
    MarkdownModule,
    NgxMdModule,
    // NgxFloatButtonModule,
    MatDialogModule,
    NgxPermissionsModule.forRoot(),
    MatRadioModule,
    MatSlideToggleModule,
    jqxChartModule,
    ChartsModule,
    WavesModule,
    // MDBBootstrapModule
    MDBBootstrapModule.forRoot()
    // ParticlesModule
  ],

  exports: [],

  providers: [ 
    EaltserviceService,
   ],

  entryComponents: [
  ],
  
  bootstrap: [AppComponent]

})
export class AppModule {
}
