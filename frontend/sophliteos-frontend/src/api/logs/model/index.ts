export interface AlarmRecordParams {
  event: number;
  startTime: string;
  endTime: string;
  pageNo: number;
  pageSize: number;
}

export interface LogFunctionInfo {
  userName: string[];
  operationIp: string[];
  operationFunc: string[];
}
