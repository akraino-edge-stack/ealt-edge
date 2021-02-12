import { Component, OnInit, ViewChild } from '@angular/core';

import {MatTableDataSource} from '@angular/material/table';

import { MatPaginator } from '@angular/material/paginator';

import { cameraData,monitorinfo, monitorDetails } from './../datainterface'
import { RoboService } from './../../app/robo.service';

import { DomSanitizer, SafeUrl } from "@angular/platform-browser";


@Component({
  selector: 'app-data-monitor',
  templateUrl: './data-monitor.component.html',
  styleUrls: ['./data-monitor.component.scss']
})
export class DataMonitorComponent implements OnInit {

  monitorColumns: string [] = ['shelfName','objType','currentCount','totalCount','time'];
  monitorDataSource = new MatTableDataSource<monitorinfo>(MONITOR_INFO_LIST);

  monitorArrayList = [];

  imageBlobUrl : any;
  image: any;
  thumbnail: any;

  monitorInfo = {} as monitorDetails;

  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;

  constructor(
    private roboService: RoboService,
    private sanitizer: DomSanitizer
    ) { }

  ngOnInit(): void {
    this.monitorDetails();
    this.inventoryDetails();
  }

  monitorDetails() {
    debugger;
    this.roboService.getMonitorImage()
        .subscribe( (data:any) => {
      debugger;
      console.log(data);
      
      let objectURL = 'data:image/jpeg;base64,' + data.image;
      this.thumbnail = this.sanitizer.bypassSecurityTrustUrl(objectURL);

     },
     error => console.log(error));
  }

  inventoryDetails() {
    this.roboService.getMonitorInfo()
      .subscribe(data => {
      console.log(data);
      this.monitorInfo = data;
      this.monitorArrayList = data.InventryData;
      this.monitorDataSource = new MatTableDataSource(this.monitorArrayList);
      this.monitorDataSource.paginator = this.paginator;
     },
     error => console.log(error));
   }

}

const MONITOR_INFO_LIST: monitorinfo[] = [
  { shelfName: '', ObjType: '', currentCount: '', totalCount: '', time: '' }
];
