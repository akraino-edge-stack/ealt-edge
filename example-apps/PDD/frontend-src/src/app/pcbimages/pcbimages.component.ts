import { Component, OnInit } from '@angular/core';

import { EaltserviceService } from '../ealtservice.service'

import { DomSanitizer, SafeUrl } from "@angular/platform-browser";

@Component({
  selector: 'app-pcbimages',
  templateUrl: './pcbimages.component.html',
  styleUrls: ['./pcbimages.component.scss']
})
export class PcbimagesComponent implements OnInit {

  input_images = ["./../../assets/images/input_images/01_spurious_copper_03.jpg","./../../assets/images/input_images/04_missing_hole_04.jpg","./../../assets/images/input_images/07_spurious_copper_05.jpg","./../../assets/images/input_images/09_missing_hole_04.jpg","./../../assets/images/input_images/09_spurious_copper_09.jpg","./../../assets/images/input_images/11_open_circuit_03.jpg"];
  inputimg = true;

  thumbnail01 : any;
  thumbnail02 : any;
  thumbnail03 : any;
  thumbnail04 : any;
  thumbnail05 : any;
  thumbnail : any;
  data: any;

  cam: any;
  edge: any;

  constructor(
    private serviceobj: EaltserviceService,
    private sanitizer: DomSanitizer
  ) { }

  ngOnInit(): void {
    // this.inputimg == false;
    // this.getImage(this.data);
    this.cam = JSON.parse(localStorage.getItem('camera'));
    this.edge = JSON.parse(localStorage.getItem('edgesite'));
    this.data = this.edge + '/' + this.cam;

    this.getImage(this.data);
  }

  getImage(data) {

    //let myCompOneObj = new CompOneComponent();
    debugger;
    this.serviceobj.getInputImage(data)
        .subscribe( (data:any) => {
          
      debugger;

      let objectURL = 'data:image/jpeg;base64,' + data.image01;
      this.thumbnail = this.sanitizer.bypassSecurityTrustUrl(objectURL);
      
      // let objectURL01 = 'data:image/jpeg;base64,' + data.image01;
      let objectURL02 = 'data:image/jpeg;base64,' + data.image02;
      let objectURL03 = 'data:image/jpeg;base64,' + data.image03;
      let objectURL04 = 'data:image/jpeg;base64,' + data.image04;
      let objectURL05 = 'data:image/jpeg;base64,' + data.image05;


      // this.thumbnail01 = this.sanitizer.bypassSecurityTrustUrl(objectURL01);
      this.thumbnail02 = this.sanitizer.bypassSecurityTrustUrl(objectURL02);
      this.thumbnail03 = this.sanitizer.bypassSecurityTrustUrl(objectURL03);
      this.thumbnail04 = this.sanitizer.bypassSecurityTrustUrl(objectURL04);
      this.thumbnail05 = this.sanitizer.bypassSecurityTrustUrl(objectURL05);

     },
     error => console.log(error));    

  }

}
