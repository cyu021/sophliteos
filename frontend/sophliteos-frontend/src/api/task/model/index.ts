import { BasicResponse } from '../../model/baseModel';

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
}

/**
 * @description: Request list return value
 */
export type AbilityListGetResultModel = BasicResponse<Array<AbilityItem>>;
export type UpUrlGetResultModal = BasicResponse<string>;
export type UpUrlAddResultModal = BasicResponse<any>;
