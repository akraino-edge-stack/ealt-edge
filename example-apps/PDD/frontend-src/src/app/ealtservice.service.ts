import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from '@angular/common/http';
import { Observable,throwError } from 'rxjs'
import { timer, Subscription, pipe } from 'rxjs';
import { switchMap } from 'rxjs/operators';

import { delay } from 'rxjs/operators';

import { nodeDetails } from './datainterface';


import { retry,catchError } from 'rxjs/operators';
import { v4 as uuid } from 'uuid';



@Injectable({
  providedIn: 'root'
})
export class EaltserviceService {


  private baseUrl = 'http://localhost:30281/';

  private imageUploadUrl = this.baseUrl+'uploadimageinput';

  private nodes_url = './../assets/data/nodes.json';
  private packageUploadUrl = this.baseUrl+'uploadimageinput';
  
  private historyIdUrl = this.baseUrl+'history/files';

  private outputImageUrl = this.baseUrl

  private inputImageUrl = this.baseUrl

  private pcbDetectUrl = this.baseUrl

  constructor(private http:HttpClient) {
  }

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type':'application/json'
    })
  }

  getNodesInfoo(): Observable<nodeDetails> {
    return this.http.get<nodeDetails>(this.nodes_url);
  }

  postHistoryId(data): Observable<Blob> {
    return this.http.get<Blob>(this.historyIdUrl, {params: data} );
  }

  postDeploymentPackage(data): Observable<any> {
    return this.http.post<any>(this.packageUploadUrl, data);
  }

  postInputImages(data): Observable<any> {
    return this.http.post<any>(this.imageUploadUrl, data)
  }

  getOutputImage(): Observable<any> {
    debugger;
    this.outputImageUrl = this.baseUrl + 'v1/pcb/resultimage';
    return this.http.get<any>(this.outputImageUrl);
  }

  getInputImage(data): Observable<any> {
    debugger;
    this.inputImageUrl = this.baseUrl + 'v1/pcb/preview/' + data;
    // return this.http.get<any>(this.monitorImageUrl);
    return this.http.get<any>(this.inputImageUrl);
  }

  pcbDetect(data): Observable<any> {
    debugger;
    this.pcbDetectUrl = this.baseUrl + 'v1/pcb/detection/' + data;
    return this.http.get<any>(this.pcbDetectUrl);
  }

}
