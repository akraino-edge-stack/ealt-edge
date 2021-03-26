import { NgModule } from '@angular/core';
import { Routes, RouterModule, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { HomeComponent } from './home/home.component';

import { PcbComponent } from './pcb/pcb.component';
import { PcbimagesComponent } from './pcbimages/pcbimages.component';
import { PcboutputimagesComponent } from './pcboutputimages/pcboutputimages.component';




const routes: Routes = [

  {
    path: '',
    component: HomeComponent
    // canActivate: [AuthGuard]
  },
  {
    path: 'akrainowiki',
    component: HomeComponent,
    resolve: {
      url: 'externalUrlRedirectResolver'
    },
    data: {
      externalUrl: 'https://wiki.akraino.org/display/AK/ELIOT%3A+Edge+Lightweight+and+IoT+Blueprint+Family?src=contextnavpagetreemode'
    }
  },
  {
    path: 'pcb',
    component: PcbComponent,
  },
  {
    path: 'displayimage',
    component: PcbimagesComponent,
  },
  {
    path: 'results',
    component: PcboutputimagesComponent,
  },
  {
    path: '**',
    redirectTo: ''
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {
    scrollPositionRestoration: 'enabled'
  })],
  exports: [RouterModule],
  providers: [
    {
        provide: 'externalUrlRedirectResolver',
        useValue: (route: ActivatedRouteSnapshot, state: RouterStateSnapshot) =>
        {
            // window.location.href = (route.data as any).externalUrl;
            window.open((route.data as any).externalUrl);
            debugger;
            console.log(route.url);
            route.url[0].path="";

        }
    }
]
})
export class AppRoutingModule { }
