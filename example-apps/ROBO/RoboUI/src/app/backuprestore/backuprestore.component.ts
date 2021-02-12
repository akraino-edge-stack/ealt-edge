import { Component, OnInit, ViewChild } from '@angular/core';

import {MatTableDataSource} from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { RoboService } from './../../app/robo.service';

import { appsinfo,pvpvsinfo,backupsinfo,restoresinfo,backupData,restoreData } from './../../app/datainterface'

import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-backuprestore',
  templateUrl: './backuprestore.component.html',
  styleUrls: ['./backuprestore.component.scss']
})
export class BackuprestoreComponent implements OnInit {

  appsColumns: string [] = ['namespace','name','ready','status','restarts','age','ip','node','nominatednode','readinessgates']
  appsDataSource = new MatTableDataSource<appsinfo>(APPS_INFO_LIST);

  appsArrayList = [];

  pvsColumns: string [] = ['namespace','name','status','volume','capacity','accessmodes','storageclass','age','volumemode']
  pvsDataSource = new MatTableDataSource<pvpvsinfo>(PVS_INFO_LIST);

  pvsArrayList = [];

  backupsColumns: string [] = ['name','status','errors','warnings','created']
  backupsDataSource = new MatTableDataSource<backupsinfo>(BACKUPS_INFO_LIST);

  backupsArrayList = [];

  restoresColumns: string [] = ['name','backup','status']
  restoresDataSource = new MatTableDataSource<restoresinfo>(RESTORES_INFO_LIST);

  restoresArrayList = [];

  selectedNamespace : string
  selectedBackupName : string

  selectedRestoreName: string
  selectedBackupname : string

  appsinfo = {}

  backupData = {} as backupData

  restoreData = {} as restoreData

  @ViewChild(MatPaginator, {static: true}) paginator: MatPaginator;

  constructor(
    private roboService: RoboService,
    private toastService: ToastrService
    ) {}

  ngOnInit(): void {
    this.selectedBackupName = "backup01"
    this.selectedRestoreName = "restore01"
    this.selectedNamespace = "default"
    this.selectedBackupname = "backup01"

    this.getAppsPvcs();

    this.getBackupsRestores();
  }

  getAppsPvcs() {
    this.roboService.getAppsPvcsInfo()
      .subscribe(data => {
      debugger;
      console.log(data);
      this.appsinfo = data;
      
      this.appsArrayList = data.appsData;
      this.appsDataSource = new MatTableDataSource(this.appsArrayList);
      this.appsDataSource.paginator = this.paginator;

      this.pvsArrayList = data.pvcData;
      this.pvsDataSource = new MatTableDataSource(this.pvsArrayList);
      this.pvsDataSource.paginator = this.paginator;
     },
     error => console.log(error)); 
  }

  getBackupsRestores() {
    this.roboService.getBackupRestoreInfo()
      .subscribe(data => {
      debugger;
      console.log(data);
      this.appsinfo = data;
      
      this.backupsArrayList = data.backupsData;
      this.backupsDataSource = new MatTableDataSource(this.backupsArrayList);
      this.backupsDataSource.paginator = this.paginator;

      this.restoresArrayList = data.restoresData;
      this.restoresDataSource = new MatTableDataSource(this.restoresArrayList);
      this.restoresDataSource.paginator = this.paginator;
     },
     error => console.log(error)); 
  }

  refreshPage() {
    debugger;
    this.getBackupsRestores();
    this.getAppsPvcs();
  }

  postBackup() {
    console.log("Inside postBackup.....")
    this.backupData.backupName = this.selectedBackupName;
    this.backupData.namespace = this.selectedNamespace;
    this.showBackupSuccess()
    this.roboService.postBackup(this.backupData)
    .subscribe(data => {
      debugger;
      if(data.responce == "success"){
        this.showBackupSuccess();
      }
      console.log(data);
    }
  ,error => console.log(error)
  );
  }

  restore() {
    console.log("Inside postBackup.....")

    this.restoreData.restoreName = this.selectedRestoreName;
    this.restoreData.backupName = this.selectedBackupname;
    
    this.showRestoreSuccess()
    this.roboService.postRestore(this.restoreData)
    .subscribe(data => {
      debugger;
      if(data.responce == "success"){
        this.showRestoreSuccess();
      }
      console.log(data);
    }
  ,error => console.log(error)
  );
  }

  showBackupSuccess() {
    this.toastService.success('Backup Successful..','Backup Data');
  }

  showRestoreSuccess() {
    this.toastService.success('Restore Successful..','Restore Data');
  }

  simulateDisaster() {
    console.log("Inside simulateDisaster....")
    this.roboService.disturbCluster()
      .subscribe(data => {
      debugger;
      console.log(data);
     },
     error => console.log(error));
  }

}

const APPS_INFO_LIST: appsinfo[] = [
  { namespace: '',name: '', ready: '', status: '', restarts: '', age: '', ip: '', node: '', nominatednode: '', readinessgates: '' }
];

const PVS_INFO_LIST: pvpvsinfo[] = [
  { namespace: '',name: '', status: '', volume: '', capacity: '', accessmodes: '', storageclass: '', age: '', volumemode: '' }
];

const BACKUPS_INFO_LIST: backupsinfo[] = [
  { name: '', status: '', errors: '', warnings: '', created: ''}
];

const RESTORES_INFO_LIST: restoresinfo[] = [
  { name: '', backup: '', status: ''}
];


// "zone.js": "~0.10.2"