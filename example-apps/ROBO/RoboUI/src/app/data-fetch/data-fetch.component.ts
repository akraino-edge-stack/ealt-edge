import { Component, OnInit, TemplateRef, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RoboService } from './../../app/robo.service';

import { cameraData,camerainfo } from './../datainterface'

import { cameraDetails } from './../datainterface';

// import { ToastService } from './toast.service';

import { ToastrService } from 'ngx-toastr';
import {MatTableDataSource} from '@angular/material/table';

import { MatPaginator } from '@angular/material/paginator';

@Component({
  selector: 'app-data-fetch',
  templateUrl: './data-fetch.component.html',
  styleUrls: ['./data-fetch.component.scss']
})
export class DataFetchComponent implements OnInit {

  cameraColumns: string [] = ['cameraID','cameraLocation','cameraNumber','rtspUrl'];
  cameraDataSource = new MatTableDataSource<camerainfo>(CAMERA_INFO_LIST);

  SERVER_URL = "http://localhost:30092/v1/monitor/video";
  videoUploadForm: FormGroup;  
  cameraDetailsForm: FormGroup;

  cameraData = {} as cameraData;
  camerasArray = [];
  location = [];

  selectedCamera: string;
  selectedLocation: string;

  selectedRTSP: string;

  cameradetconcat: string

  url;
  format;


  cameraInfo = {} as cameraDetails;
  cameraArrayList = [];

  selectedCameraId = []
  selectedCameraID: string

  selectedValues = []
  
  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;

  constructor(
    private formBuilder: FormBuilder, 
    private httpClient: HttpClient,
    private roboService: RoboService,
    private toastService: ToastrService
    ) { }

  ngOnInit() {
    this.fetchCameraDetails();
    this.videoUploadForm = this.formBuilder.group({
      video: ['']
    });

    this.cameraDetailsForm = this.formBuilder.group({
      cameraLocation: [''],
      cameraNumber: ['']
    });

    this.camerasArray = 
    [
      {
        "value": "camera01",
        "viewValue": "camera01"
      },
      {
        "value": "camera02",
        "viewValue": "camera02"
      },
      {
        "value": "camera03",
        "viewValue": "camera03"
      }
    ];

    this.location = 
    [
      {
        "value": "Bangalore",
        "viewValue": "Bangalore"
      }
    ];
    
    this.selectedCamera = "Camera"
    this.selectedLocation = "Bangalore"
    // this.fetchCameraDetails();
  }

  onFileSelect(event) {


    if (event.target.files.length > 0) {
      const file = event.target.files[0];
      this.videoUploadForm.get('video').setValue(file);
    }

    const file = event.target.files && event.target.files[0];
    if (file) {
      var reader = new FileReader();
      reader.readAsDataURL(file);
      if (file.type.indexOf('image') > -1) {
        this.format = 'image';
      } else if (file.type.indexOf('video') > -1) {
        this.format = 'video';
      }
      reader.onload = (event) => {
        this.url = (<FileReader>event.target).result;
      }
    }
  }

  onSubmit() {
    const formData = new FormData();
    formData.append('file', this.videoUploadForm.get('video').value);
    debugger;
    this.showFileSuccess();
    this.httpClient.post<any>(this.SERVER_URL, formData).subscribe
    (
      (res) => console.log(res),
      (err) => console.log(err)
    );
  }
  cameraDetailsSubmit() {
    const formData = new FormData();
    
    this.cameraData.cameraNumber = this.selectedCamera;
    this.cameraData.cameraLocation = this.selectedLocation;
    this.cameraData.rtspUrl = this.selectedRTSP;
    
    this.roboService.postCameraDetails(this.cameraData)
        .subscribe(data => {
          debugger;
          if(data.responce == "success"){
            this.showSuccess();
          }
          console.log(data);
        }
      ,error => console.log(error)
      );

  }

  onCameraSelection() {
    console.log("Inside onCameraSelection.....")
  }

  onCameraIDSelection() {
    var index: number
    console.log("Inside onCameraIDSelection.......")
    debugger;
    // this.roboService.postCameraID(this.selectedCameraID)
    index = this.selectedCameraId.indexOf(this.selectedCameraID)
    debugger;
    this.roboService.triggerDetection(this.selectedValues[index])
    .subscribe(data => {
      debugger;
      console.log(data)
     },
     error => console.log(error));

  }


  showSuccess() {
    console.log("Inside showSuccess.... Method")
    this.toastService.success('Uploaded Succesfully!','Camera Data');
  }

  showFileSuccess() {
    console.log("Inside showSuccess.... Method")
    this.toastService.success('Uploaded Succesfully!','Video File');
  }

  fetchCameraDetails() {
    
    debugger;
    this.roboService.getCameraInfo()
      .subscribe(data => {
      debugger;
      console.log(data);
      this.cameraInfo = data;
      
      this.cameraArrayList = data.roboCamera;
      this.cameraDataSource = new MatTableDataSource(this.cameraArrayList);
      this.cameraDataSource.paginator = this.paginator;

      console.log("For loop started.....")
      for (var val of this.cameraArrayList) {
        debugger;
        
        console.log(val);
        this.cameradetconcat = val.cameraNumber + '/'+ val.rtspUrl + '/' +val.cameraLocation
        // this.selectedCameraId.push(val.camera);
        this.selectedCameraId.push(this.cameradetconcat)
        this.selectedValues.push(val.cameraID)

      }
      debugger;
      console.log("SelectedCameraID")
      console.log(this.selectedCameraId)
     },
     error => console.log(error));
   }

   refreshPage() {
     this.fetchCameraDetails();
   }
}

const CAMERA_INFO_LIST: camerainfo[] = [
  { cameraID: '',cameraLocation: '', cameraNumber: '', rtspUrl: '' }
];
