import { Component, OnInit, TemplateRef, ViewChild, AfterViewInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { ToolbarService } from './../toolbar/toolbar.service';
import { ToastService } from './toast.service';

import { uploadImageData } from './../datainterface'
import { EaltserviceService } from '../ealtservice.service'

import { PcbimagesComponent } from './../pcbimages/pcbimages.component'


import { DomSanitizer, SafeUrl } from "@angular/platform-browser";

export interface Edgesites {
  value: string;
  viewValue: string;
}

export interface Defects {
  value: string;
  viewValue: string;
}

export interface Cameras {
  value: string;
  viewValue: string;
}

@Component({
  selector: 'app-pcb',
  templateUrl: './pcb.component.html',
  styleUrls: ['./pcb.component.scss']
})
export class PcbComponent implements OnInit {

  imageDeployForm: FormGroup;
  modelDeployForm: FormGroup;
  url;
  msg = "";

  modelBool = false;
  imageBool = false;
  detectBool = false;
  detectPendingBool = false;
  selectedEdgeSite: String;
  selectedDefect: String;
  selectedCamera: String;

  parentData: String;
  detectData: String;

  uploadImageData = {} as uploadImageData;
  isTemplate(toast) { return toast.textOrTpl instanceof TemplateRef; }

  pcbInfo = {} ;

  edgesites: Edgesites[] = [
    {value: 'beijinglab', viewValue: 'beijinglab'},
    {value: 'shenzhenlab', viewValue: 'shenzhenlab'},
    {value: 'shanghailab', viewValue: 'shanghailab'}
  ];

  defects: Defects[] = [
    {value: 'all', viewValue: 'All Defects'},
    {value: 'missing_hole', viewValue: 'Missing Hole'},
    {value: 'mouse_bite', viewValue: 'Mouse Bite'},
    {value: 'open_circuit', viewValue: 'Open Circuit'},
    {value: 'short', viewValue: 'Short'},
    {value: 'spur', viewValue: 'Spur'},
    {value: 'spurious_copper', viewValue: 'Spurious Copper'}
  ];

  cameras: Cameras[] = [
    {value: 'camera1', viewValue: 'camera1'},
    {value: 'camera2', viewValue: 'camera2'},
    {value: 'camera3', viewValue: 'camera3'}
  ];

  @ViewChild(PcbimagesComponent) pcbchild;

  constructor(
    private formBuilder: FormBuilder,
    public toolbarService: ToolbarService,
    public toastService: ToastService,
    private serviceobj: EaltserviceService,
    private sanitizer: DomSanitizer
  ) { }

  ngOnInit(): void {

    // let pcbimagesObj = new PcbimagesComponent();

    this.imageDeployForm = this.formBuilder.group(
      {
        deployzipfile: ['']
      }
    );
    this.modelDeployForm = this.formBuilder.group(
      {
        deploymodelfile: ['']
      }
    );

    this.selectedCamera = JSON.parse(localStorage.getItem('camera'));
    this.selectedEdgeSite = JSON.parse(localStorage.getItem('edgesite'));

  }

  ngOnDestroy(): void {
    // localStorage.removeItem('camera')
    // localStorage.removeItem('edgesite')
  }

  fileProgress(event) {
    console.log("Inside fileProgress...")
    debugger;
    if(event.target.files.length > 0 ){
      const deployfile = event.target.files[0];
      let filee = event.target.files[0];
      this.imageDeployForm.get('deployzipfile').setValue(deployfile);
      debugger;
      let filecon: string = '';
      let fileReader: FileReader = new FileReader();
      // console.log(YAML.parse(deployfile));
      debugger;
    }

    var reader = new FileReader();
		reader.readAsDataURL(event.target.files[0]);
		
		reader.onload = (_event) => {
			this.msg = "";
			this.url = reader.result; 
		}
    console.log("Ends fileProgress...")
  }

  modelFileProgress(event) {
    console.log("Inside fileProgress...")
    debugger;
    if(event.target.files.length > 0 ){
      const deployfile = event.target.files[0];
      let filee = event.target.files[0];
      this.modelDeployForm.get('deploymodelfilefile').setValue(deployfile);
      debugger;
      let filecon: string = '';
      let fileReader: FileReader = new FileReader();
      // console.log(YAML.parse(deployfile));
      debugger;
    }
    console.log("Ends modelFileProgress...")
  }

  onSubmit() {
    console.log("Inside onSubmit() ....")
    const formData = new FormData();

    formData.append('deployfile',this.imageDeployForm.get('deployzipfile').value);
    debugger;

    this.uploadImageData.uploadFile = this.imageDeployForm.value.deployzipfile;

    console.log("uploadImagedata...");
    console.log(this.uploadImageData);
    console.log("formData....");
    console.log(formData);
    // formData.append('file', this.fileData);
    // this.deployData.deployFile = formData.get('yamlfile')
    console.log(this.imageDeployForm.value);
    this.imageBool = true;
    this.modelBool = false;
    this.detectBool = false;
    debugger;
    this.showSuccess();

    localStorage.setItem('camera',JSON.stringify(this.selectedCamera))
    localStorage.setItem('edgesite',JSON.stringify(this.selectedEdgeSite))
    
    debugger;

    // this.imageBool = false;
    this.serviceobj.postDeploymentPackage(formData)
        .subscribe(data => {
          console.log(data);
          this.showSuccess();
          // console.log(data);
        }
      ,error => console.log(error)
      );
  }

  onModelSubmit() {
    console.log("Inside onSubmit() ....")
    const formData = new FormData();

    formData.append('deployfile',this.modelDeployForm.get('deploymodelfile').value);

    this.uploadImageData.uploadFile = this.modelDeployForm.value.deploymodelfile;

    console.log(this.uploadImageData);
    // formData.append('file', this.fileData);
    // this.deployData.deployFile = formData.get('yamlfile')
    console.log(this.modelDeployForm.value);
    this.modelBool = true;
    this.imageBool = false;
    this.detectBool = false;
    this.showSuccess();
    // this.modelBool = false;

    this.serviceobj.postDeploymentPackage(formData)
        .subscribe(data => {
          console.log(data);
          this.showSuccess();
          // console.log(data);
        }
      ,error => console.log(error)
      );
  }

  onDetect() {

    this.detectBool = false;
    this.modelBool = false;
    this.imageBool = false;

    this.toastService.show('Defect Detection in PCB images In progress', { classname: 'bg-warning text-dark', delay: 10000 });

    this.detectData = this.selectedEdgeSite + '/' + this.selectedCamera;
    debugger;
    
    this.serviceobj.pcbDetect(this.detectData)
        .subscribe(data => {
          console.log(data);
          if(data.responce == "success"){
            this.detectBool = true;
            this.detectPendingBool = false;
            this.showSuccess();
          }
        }
      ,error => console.log(error)
      );
    
    this.detectData = '';
    debugger;
  }


  // Success message display

  private async showSuccess() {

    //wait for 3 seconds
    await this.delay(3000);
    console.log("Inside showSuccess...")
    
    if(this.imageBool == true){
      this.toastService.show('PCB images from cameras uploaded successfully', { classname: 'bg-success text-light', delay: 9000 });
    }

    if(this.modelBool == true){
      this.toastService.show('PCB Trained model uploaded Successfully', { classname: 'bg-success text-light', delay: 9000 });
    }

    if(this.detectBool == true){
      // this.toastService.show('Defect Detection in PCB images In progress', { classname: 'bg-warning text-dark', delay: 3000 });
      // await this.delay(3000);
      this.toastService.show('Defect Detection in PCB images Completed', { classname: 'bg-success text-light', delay: 9000 });
    }

    if(this.detectPendingBool == true ){
      this.toastService.show('Defect Detection in PCB images In progress', { classname: 'bg-warning text-dark', delay: 9000 });
      // await this.delay(9000);
    }
    
  }

  private delay(ms: number)
  {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  onEdgeSiteSelection() {
    console.log("on Edge Site Selection triggered....");
    console.log(this.selectedEdgeSite);
  }

  onDefectSelection() {
    debugger;
    console.log("on Defect Selection triggered....");
    console.log(this.selectedDefect);
  }

  onCameraSelection() {
    debugger;
    console.log("on Camera Selection triggered....");
    console.log(this.selectedCamera);
  }

  onPreview() {
    console.log("Inside onPreview.....")
    debugger;
    this.parentData = this.selectedEdgeSite + '/' + this.selectedCamera;
    debugger;
    let pcbInputObj = new PcbimagesComponent(this.serviceobj, this.sanitizer);
  }

}
