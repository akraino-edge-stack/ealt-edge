import { Component, OnInit } from '@angular/core';

import { EaltserviceService } from '../ealtservice.service'

import { DomSanitizer, SafeUrl } from "@angular/platform-browser";

@Component({
  selector: 'app-pcboutputimages',
  templateUrl: './pcboutputimages.component.html',
  styleUrls: ['./pcboutputimages.component.scss']
})
export class PcboutputimagesComponent implements OnInit {

  output_images = ["./../../assets/images/output_images/01_spurious_copper_03.jpg","./../../assets/images/output_images/04_missing_hole_04.jpg","./../../assets/images/output_images/07_spurious_copper_05.jpg","./../../assets/images/output_images/09_missing_hole_04.jpg","./../../assets/images/output_images/09_spurious_copper_09.jpg","./../../assets/images/output_images/11_open_circuit_03.jpg"];
  outputimg = true;

  thumbnail01 : any;
  thumbnail02 : any;
  thumbnail03 : any;
  thumbnail04 : any;
  thumbnail05 : any;

  constructor(
    private serviceobj: EaltserviceService,
    private sanitizer: DomSanitizer
  ) { }

  ngOnInit(): void {
    this.getImage();
  }

  getImage() {
    debugger;
    localStorage.removeItem('camera');
    localStorage.removeItem('edgesite');
    this.serviceobj.getOutputImage()
        .subscribe( (data:any) => {
          
      debugger;
      
      let objectURL01 = 'data:image/jpeg;base64,' + data.image01;
      let objectURL02 = 'data:image/jpeg;base64,' + data.image02;
      let objectURL03 = 'data:image/jpeg;base64,' + data.image03;
      let objectURL04 = 'data:image/jpeg;base64,' + data.image04;
      let objectURL05 = 'data:image/jpeg;base64,' + data.image05;

      this.thumbnail01 = this.sanitizer.bypassSecurityTrustUrl(objectURL01);
      this.thumbnail02 = this.sanitizer.bypassSecurityTrustUrl(objectURL02);
      this.thumbnail03 = this.sanitizer.bypassSecurityTrustUrl(objectURL03);
      this.thumbnail04 = this.sanitizer.bypassSecurityTrustUrl(objectURL04);
      this.thumbnail05 = this.sanitizer.bypassSecurityTrustUrl(objectURL05);

     },
     error => console.log(error));    

  }

}
