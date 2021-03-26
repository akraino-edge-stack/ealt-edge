export interface nodeDetails {
    eliotNodes: nodinfo[];
}

export interface nodinfo {
  nodeName: string;
  nodeStatus: string;
  nodeRole: string;
  age: string;
  version: string;
  internalIp: string;
  externalIp: string;
  osImage: string;
  kernel: string;
  containerRuntime: string;
}

export interface uploadImageData {
  uploadFile: any;
}

export interface pcbdetectObject {
  
}

