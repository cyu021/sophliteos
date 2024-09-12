import { BasicResponse, BasicResponseCap } from '../../model/baseModel';

export interface MediaServerParams {
  pageNo: number;
  pageSize: number;
}

export interface AbilityItem {
  type: number;
  name: string;
}

export interface upUrlParams {
  ip: string;
  port: string;
  protocol: string;
  endpoint: string;
}

export interface rotateCfgParams {
  record: number;
  serviceLog: number;
  retryRows: number; //rows, 1000-5000
  enableRetry: boolean;
}

/**
 * @description: Request list return value
 */
export type AbilityListGetResultModel = BasicResponse<Array<AbilityItem>>;
export type UpUrlGetResultModal = BasicResponseCap<string>;
export type UpUrlAddResultModal = BasicResponseCap<any>;
