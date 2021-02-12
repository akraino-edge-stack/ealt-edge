import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DataFetchComponent } from './data-fetch/data-fetch.component';
import { DataMonitorComponent } from './data-monitor/data-monitor.component';
import { BackuprestoreComponent } from './backuprestore/backuprestore.component';

const routes: Routes = [
  {
    path: 'dataupload',
    component: DataFetchComponent
  },
  {
    path: 'datamonitor',
    component: DataMonitorComponent
  },

  {
    path: 'backuprestore',
    component: BackuprestoreComponent
  }
  // {
  //   path: '**',
  //   redirectTo: ''
  // }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { 

}
