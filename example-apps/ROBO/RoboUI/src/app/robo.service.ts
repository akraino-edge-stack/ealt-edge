import { Injectable } from '@angular/core';

import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from '@angular/common/http';
import { Observable,throwError } from 'rxjs'
import { timer, Subscription, pipe } from 'rxjs';

import { cameraData, camerainfo, cameraDetails, monitorDetails, monitorinfo, cameraID, appsPvcs, backupRestore, monitorImage } from './datainterface'


@Injectable({
  providedIn: 'root'
})
export class RoboService {

  private baseUrl = 'http://159.138.33.54:30091';

  private inventoryBaseUrl = 'http://159.138.33.54:30092';

  private postCameraDetailsUrl = this.inventoryBaseUrl + '/v1/monitor/cameras' // 92
  private cameraDetailsUrl = this.inventoryBaseUrl + '/v1/monitor/cameras' //
  //private cameraDetails_url = './../assets/data/camera.json'

  //private monitorDetails_url = './../assets/data/inventory.json'
  private monitorDetailsUrl = this.inventoryBaseUrl + '/v1/inventry/table'

  private monitorImageUrl = this.inventoryBaseUrl + '/v1/inventry/image'

  private triggerObjUrl = this.inventoryBaseUrl + '/v1/monitor/cameras/'

  private appsPvcsDetailsUrl = this.baseUrl + '/v1/robo/apps-pvcs'
  //private appsPvcsDetails_url = './../assets/data/appspvc.json'

  private backupRestoreDetailsUrl = this.baseUrl + '/v1/robo/backup-restore'
  
  //private backupRestoreDetails_url = './../assets/data/backuprestore.json'

  private postBackupDetailsUrl = this.baseUrl + '/v1/robo/backup'

  private postRestoreDetailsUrl = this.baseUrl + '/v1/robo/restore'

  private disasterUrl = this.baseUrl + '/v1/robo/disaster'

  constructor(private http:HttpClient) { }
  
  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type':'application/json'
    })
  }

  httpOptionss = {
    headers: new HttpHeaders({
      'Content-Type':'application/json'
    })
  }

  postCameraDetails(data): Observable<any> {
    console.log(data);
    debugger;
    return this.http.post<any>(this.postCameraDetailsUrl, data)
  }

  getCameraInfo(): Observable<cameraDetails> {
    debugger;
    return this.http.get<cameraDetails>(this.cameraDetailsUrl);
  }

  getMonitorInfo(): Observable<monitorDetails> {
    debugger;
    return this.http.get<monitorDetails>(this.monitorDetailsUrl);
  }

  getMonitorImage(): Observable<any> {
    debugger;
    return this.http.get<any>(this.monitorImageUrl);
  }


  triggerDetection(data): Observable<any> {
    console.log(data);
    debugger;
    this.triggerObjUrl = this.triggerObjUrl + data;
    return this.http.get<any>(this.triggerObjUrl)
  }

  getAppsPvcsInfo(): Observable<appsPvcs> {
    return this.http.get<appsPvcs>(this.appsPvcsDetailsUrl);
  }

  // getBackupRestoreInfo(): Observable<backupRestore> {
    // }
    //   return this.http.get<backupRestore>(this.backupRestoreDetails_url);

  getBackupRestoreInfo(): Observable<any> {
    return this.http.get<any>(this.backupRestoreDetailsUrl);
  }

  disturbCluster(): Observable<any> {
    return this.http.get<any>(this.disasterUrl);
  }

  postBackup(data): Observable<any> {
    console.log(data);
    debugger;
    return this.http.post<any>(this.postBackupDetailsUrl, data)
  }

  postRestore(data): Observable<any> {
    console.log(data);
    debugger;
    return this.http.post<any>(this.postRestoreDetailsUrl, data)
  }
}
