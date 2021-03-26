import { Component, OnInit, ViewChild } from '@angular/core';

import {MatTableDataSource} from '@angular/material/table';

import { MatPaginator } from '@angular/material/paginator';

import { cameraData,monitorinfo, monitorDetails } from './../datainterface'
import { RoboService } from './../../app/robo.service';

import { DomSanitizer, SafeUrl } from "@angular/platform-browser";

import { timer } from "rxjs"

@Component({
  selector: 'app-data-monitor',
  templateUrl: './data-monitor.component.html',
  styleUrls: ['./data-monitor.component.scss']
})
export class DataMonitorComponent implements OnInit {

  monitorColumns: string [] = ['shelfName','objType','currentCount','maxCount','status','time'];
  monitorDataSource = new MatTableDataSource<monitorinfo>(MONITOR_INFO_LIST);

  monitorArrayList = [];

  imageBlobUrl : any;
  image: any;
  thumbnail: any;
  myvar: any

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

  timerFunc() {
    debugger;
    setInterval(this.refreshPage, 2000);
    // setTimeout(this.refreshPage , 2000);
  }

  refreshPage() {
    console.log("Inside refreshPage....")
    this.inventoryDetails();
    this.monitorDetails();
  }

  monitorDetails() {
    debugger;
    // setInterval(function(){ 
    //   console.log("Oooo Yeaaa!");
    //   this.roboService.getMonitorImage()
    //     .subscribe( (data:any) => {
          
    //   debugger;
    //   console.log(data);
    //   // referred https://stackoverflow.com/questions/55591871/view-blob-response-as-image-in-angular
      
    //   let objectURL = 'data:image/jpeg;base64,' + data.image;
    //   this.thumbnail = this.sanitizer.bypassSecurityTrustUrl(objectURL);

    //  },
    //  error => console.log(error));


    // }, 2000);
    this.roboService.getMonitorImage()
        .subscribe( (data:any) => {
          
      debugger;
      
      let objectURL = 'data:image/jpeg;base64,' + data.image;
      this.thumbnail = this.sanitizer.bypassSecurityTrustUrl(objectURL);
      debugger;

     },
     error => console.log(error));    

  }

  inventoryDetails() {
    console.log("inventoryDetails started....")

    // setInterval(function(){ 
    //   console.log("Oooo!");
    //   this.roboService.getMonitorInfo()
    //   .subscribe(data => {
    //   console.log(data);
    //   this.monitorInfo = data;
      
    //   this.monitorArrayList = data.InventryData;
    //   this.monitorDataSource = new MatTableDataSource(this.monitorArrayList);
    //   this.monitorDataSource.paginator = this.paginator;
    //  },
    //  error => console.log(error));
    // }, 2000);
    this.roboService.getMonitorInfo()
      .subscribe(data => {
      console.log(data);
      this.monitorInfo = data;
      
      this.monitorArrayList = data.InventryData;
      this.monitorDataSource = new MatTableDataSource(this.monitorArrayList);
      this.monitorDataSource.paginator = this.paginator;
     },
     error => console.log(error));
     console.log("inventoryDetails finished...")
   }

   getColor(clr) {
      if(clr == "Mostly Filled"){
        return "green"
      }
      if(clr == "Partially Filled"){
        return "orange"
      }
      if(clr == "Needs Filling"){
        return "red"
      }
   }

}

const MONITOR_INFO_LIST: monitorinfo[] = [
  { shelfName: '', ObjType: '', currentCount: '', maxCount: '', status: '', time: '' }
];
