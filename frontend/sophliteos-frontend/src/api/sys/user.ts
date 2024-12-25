import { defHttp } from '/@/utils/http/axios';
import {
  LoginParams,
  LogoutParams,
  LoginResultModel,
  GetUserInfoModel,
  PasswordParams,
} from './model/userModel';

import { ErrorMessageMode } from '/#/axios';

enum Api {
  Login = '/login',
  Logout = '/logout',
  GetUserInfo = '/getUserInfo',
  GetPermCode = '/getPermCode',
  Password = '/device/password',
  TestRetry = '/testRetry',
  GetAcctPrivSitemap = '/priv/acct/getPrivSitemap',
  GetAcctList = '/priv/acct/getList',
  GetRoleList = '/priv/role/getList',
  PostAcctUpsert = '/priv/acct/upsert',
  PostAcctDelete = '/priv/acct/delete',
}

export function PostAcctDeleteApi(params: {}) {
  return defHttp.post<{}>(
    {
      url: Api.PostAcctDelete,
      params
    },
    { isTransformResponse: false }
  );
}

export function PostAcctUpsertApi(params: {}) {
  return defHttp.post<{}>(
    {
      url: Api.PostAcctUpsert,
      params
    },
    { isTransformResponse: false }
  );
}

export function GetRoleListApi(params: {}) {
  return defHttp.get<{}>(
    {
      url: Api.GetRoleList,
      params
    },
    { isTransformResponse: false }
  );
}

export function GetAcctListApi(params: {}) {
  return defHttp.get<{}>(
    {
      url: Api.GetAcctList,
      params
    },
    { isTransformResponse: false }
  );
}

export function GetWebAcctPrivApi(params: {}) {
  return defHttp.post<{}>(
    {
      url: Api.GetAcctPrivSitemap,
      params
    },
    { isTransformResponse: false }
  );
}

/**
 * @description: user login api
 */
export function loginApi(params: LoginParams, mode: ErrorMessageMode = 'message') {
  return defHttp.post<LoginResultModel>(
    {
      url: Api.Login,
      params,
    },
    {
      errorMessageMode: mode,
      noSuccessMessage: true,
    },
  );
}

/**
 * @description: getUserInfo
 */
export function getUserInfo() {
  return defHttp.get<GetUserInfoModel>({ url: Api.GetUserInfo }, { errorMessageMode: 'none' });
}

export function getPermCode() {
  return defHttp.get<string[]>({ url: Api.GetPermCode });
}

export function doLogout(params: LogoutParams) {
  return defHttp.post<LoginResultModel>({ url: Api.Logout, params }, { noSuccessMessage: true });
}

export function testRetry() {
  return defHttp.get(
    { url: Api.TestRetry },
    {
      retryRequest: {
        isOpenRetry: true,
        count: 5,
        waitTime: 1000,
      },
    },
  );
}

// 修改密码
export function changePassword(params: PasswordParams) {
  return defHttp.post({ url: Api.Password, params });
}
