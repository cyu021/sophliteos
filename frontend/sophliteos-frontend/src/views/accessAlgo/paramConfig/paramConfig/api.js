import { defHttp } from '/@/utils/http/axios';

const attrList = () => {
    return defHttp
    .post({ url: '/list', }, { apiUrl: "/attr", isTransformResponse: false, joinParamsToUrl: false, })
    .then((res) => {
        console.log('attr list', res)
        return res.AlgoAttrDict;
    })
}

const apis = {
    attrList,
};

export default apis;
