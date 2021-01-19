import { Injectable } from '@angular/core';

import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from '@angular/common/http';
import { Observable,throwError } from 'rxjs'
import { timer, Subscription, pipe } from 'rxjs';

import { cameraData, camerainfo, cameraDetails, monitorDetails, monitorinfo, cameraID } from './datainterface'


@Injectable({
  providedIn: 'root'
})
export class RoboService {

  private baseUrl = 'http://localhost:9996';

  private postCameraDetailsUrl = this.baseUrl + '/v1/monitor/cameras'
  private cameraDetailsUrl = this.baseUrl + '/v1/monitor/cameras'
  private cameraDetails_url = './../assets/data/camera.json'

  private monitorDetails_url = './../assets/data/inventory.json'
  private monitorDetailsUrl = this.baseUrl + '/v1/inventry/table'

  private monitorImageUrl = this.baseUrl + '/v1/monitor/image'

  private triggerObjUrl = this.baseUrl + '/v1/monitor/cameras/'

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

  getMonitorImage(): Observable<Blob> {
    debugger;
    return this.http.get<Blob>(this.monitorImageUrl);
  }


  triggerDetection(data): Observable<any> {
    console.log(data);
    debugger;
    this.triggerObjUrl = this.triggerObjUrl + data;
    return this.http.get<any>(this.triggerObjUrl)
  }
}
