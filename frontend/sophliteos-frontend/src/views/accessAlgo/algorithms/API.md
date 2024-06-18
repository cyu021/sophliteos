//显示算法包+启用状态+算能侧认知 => curl -H 'Content-Type: application/x-www-form-urlencoded' -X GET http://localhost:38081/dynamic/getDynamicModelcurl -XPOST http://localhost:5566/abilities/list

//上传算法包 => curl -H 'Content-Type: multipart/form-data' -F "modelFile=@/home/admin/full_structure.zip;type=application/x-zip-compressed" -X POST http://localhost:38081/dynamic/uploadDynamicModel

//启用算法包 => curl -XPOST http://localhost:5566/abilities/token/get curl -XPOST http://localhost:5566/abilities/list curl -X POST "http://localhost:38081/dynamic/onoffDynamicModel?annotatorName=full_structure&onoff=true" -H "accept: application/json«  curl -XPOST -H 'Content-Type: application/json' -d '{ "token": "7d0ad947-316f-4d62-b56e-e9f70595da9b", « full_structure": "16"}' http://localhost:5566/abilities/set curl -XPOST -H 'Content-Type: application/json' -d '{ "token": "7d0ad947-316f-4d62-b56e-e9f70595da9b"}' http://localhost:5566/abilities/token/release

//停用算法包 curl -XPOST http://localhost:5566/abilities/token/get curl -XPOST http://localhost:5566/abilities/list curl -X POST "http://localhost:38081/dynamic/onoffDynamicModel?annotatorName=full_structure&onoff=false" -H "accept: application/json«  curl -XPOST -H 'Content-Type: application/json' -d '{ "token": "7d0ad947-316f-4d62-b56e-e9f70595da9b", « full_structure": "16"}' http://localhost:5566/abilities/unset curl -XPOST -H 'Content-Type: application/json' -d '{ "token": "7d0ad947-316f-4d62-b56e-e9f70595da9b"}' http://localhost:5566/abilities/token/release

//删除算法包 curl -X DELETE "http://localhost:38081/dynamic/deleteDynamicModel?annotatorName=full_structure" -H "accept: application/json"

=========================

//查询算法包+启用状态 curl -H 'Content-Type: application/x-www-form-urlencoded' -X GET http://localhost:38081/dynamic/getDynamicModel | jq -c '.data[] | {annotator_name, sts}' => {"annotator_name":"full_structure","sts":1} {"annotator_name":"full_structure","sts":1}

//上传算法包 curl -H 'Content-Type: multipart/form-data' -F "modelFile=@/home/admin/full_structure.zip;type=application/x-zip-compressed" -X POST http://localhost:38081/dynamic/uploadDynamicModel | jq . => { "code": 200, "data": null, "msg": "success. full_structure.zip", "trace_id": "e4d9d281-eca2-47cd-a8a1-41153a8382e6" }

//启用算法包 curl -X POST "http://localhost:38081/dynamic/onoffDynamicModel?annotatorName=full_structure&onoff=true" -H "accept: application/json" => {"code":200,"data":"","msg":"success","trace_id":"56e93be5-dfd1-46eb-ba10-2f65c51b9f35"}

//停用算法包 curl -X POST "http://localhost:38081/dynamic/onoffDynamicModel?annotatorName=full_structure&onoff=false" -H "accept: application/json" => {"code":200,"data":"","msg":"success","trace_id":"3f8d447c-e9b8-414a-a95c-288af00c774e"}

//删除算法包 curl -X DELETE "http://localhost:38081/dynamic/deleteDynamicModel?annotatorName=full_structure" -H "accept: application/json" | jq . => { "code": 200, "data": "full_structure", "msg": "delete success", "trace_id": "3a6a19c0-2306-424c-99ad-255410a29c0d" }
