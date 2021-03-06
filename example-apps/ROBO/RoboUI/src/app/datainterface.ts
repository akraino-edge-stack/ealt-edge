export interface cameraData {
    cameraLocation: string;
    cameraNumber: string;
    rtspUrl: string;
    // videoName: string;
}

export interface cameraDetails {
  roboCamera: camerainfo[];
}

export interface camerainfo {
  cameraID: any;
  cameraLocation: string;
	cameraNumber: string;
	rtspUrl: string;
}

export interface monitorDetails {
  InventryData: monitorinfo[];
}

export interface monitorImage {
  image: Blob;
}


export interface monitorinfo {
  shelfName: string;
  ObjType: string;
  currentCount: string;
  maxCount: string;
  status: string;
  time: string;
}


export interface cameraID {
  cameraID: string;
}


export interface appsinfo {
  namespace: string;
  name: string;
  status: string;
  ip: string;
  node: string;
}

export interface appsPvcs {
  appsData: appsinfo[];
  pvcData: pvpvsinfo[];
}

export interface pvpvsinfo {
  namespace: string;
  name: string;
  status: string;
  volume: string;
  storageclass: string;
  volumemode: string;
}

export interface backupRestore {
  backupsData: backupsinfo[];
  restoresData: restoresinfo[];
}

export interface backupsinfo {
  name: string;
  status: string;
  errors: string;
  warnings: string;
  created: string;
}

// export interface backups {
//   backupsData: backupsinfo[];
// }

export interface restoresinfo {
  name: string;
  backup: string;
  status: string;
}

export interface backupData {
  backupName: string;
  namespace: string;
}

export interface restoreData {
  restoreName: string;
  backupName: string;
}

// export interface restores {
//   restoresData: restoresinfo[];
// }