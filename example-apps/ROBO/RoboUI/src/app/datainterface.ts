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
  cameraID: string;
  cameraLocation: string;
	cameraNumber: string;
	rtspUrl: string;
}

export interface monitorDetails {
  InventryData: monitorinfo[];
}


export interface monitorinfo {
  shelfName: string;
	ObjType: string;
  currentCount: string;
  totalCount: string;
  time: string;
}


export interface cameraID {
  cameraID: string;
}