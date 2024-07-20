import { getVideosList } from '/@/api/dataSource';
import { getTaskList } from '/@/api/task';
import { defHttp } from '/@/utils/http/axios';

const attrList = () => {
    return defHttp
    .post({ url: '/list', }, { apiUrl: "/attr", isTransformResponse: false, joinParamsToUrl: false, })
    .then((res) => {
        console.log('attr list', res)
        return res.AlgoAttrDict;
    })
}

const videoUrl = (deviceName) => {
    return getVideosList().then((res) => {
        const r = res.filter((i) => i.name === deviceName);

        if (r.length > 0) {
            return r[0].url;
        }

        return '';
    })
}


// const videoUrl = (taskId) => {
//     return getTaskList({ pageNo: -1, pageSize: -1 }).then((list) => {
//         const filterResult = list.items.filter((item) => { return item.taskId === taskId });
//         if (filterResult.length > 0) {
//             return filterResult[0].deviceName;
//         }
//         return '';
//     }).then((deviceName) => {
//         return getVideosList().then((res) => {
//             console.log('video list', res, deviceName);
//             const r = res.filter((i) => i.name === deviceName);
//             console.log('video list', res, deviceName, r);

//             if (r.length > 0) {
//                 return r[0].url;
//             }

//             return '';
//         })
//     })
// }

const defaultExtend = (annotator) => {
    return attrList().then((aRes) => {
        const attr = aRes[annotator];
        if (!attr) {
            throw 'no annotator found in attr list, annoatorname' + annotator;
        }

        const label = Object.keys(attr)[0];
        const kvp = Object.keys(attr[label]['kvp'])[0];
        const rule = {
            Grp: 1,
            Op: '=',
            K: kvp,
            V: attr[label]['kvp'][kvp][0] || 0,
        }

        const value = {
            "FilterName": "规则名称1",
            "FilterPeriod":
            [
                {
                    "DateStart": "20240101",
                    "DateEnd": "20241230",
                    "TimeStart": "0800",
                    "TimeEnd": "2000",
                    "Weekday": "0,1,2,3,4,5,6"
                }
            ],
            "Filter":
            [
                {
                    "Annotator": annotator,
                    "Label": "1420",
                    "Rtsp": "rtsp://192.168.0.158:8554/sample",
                    "Rule": [ rule ],
                }
            ]
        }
        
        return value
    })
}

const ruleTemplate = () => {
    return defHttp
    .get({ url: '/template/get' }, { apiUrl: '/filter', isTransformResponse: false, joinParamsToUrl: false, })
    .then((res) => {
        return res.FilterTemplate;
    })
}

const apis = {
    attrList,
    defaultExtend,
    videoUrl,
    ruleTemplate,
};

export default apis;
