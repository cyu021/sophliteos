export const options = [
  { label: '吸烟检测', value: 'Smoking' },
  { label: '人员入侵', value: 'HumanBreakIn' },
  { label: '未戴口罩', value: 'PeopleWithoutMask' },
  { label: '非机动车乱停', value: 'NoneMotorVehicleParking' },
  { label: '突发性事件', value: 'UnexpectedEvents' },
  { label: '占道经营', value: 'RoadOccupatedInOperation' },
  { label: '电动车检测', value: 'ElectricCarEntersElevator' },
  { label: '占用消防通道', value: 'OccupationOfFireAccess' },
  { label: '垃圾暴露', value: 'GarbageExposed' },
  { label: '垃圾满溢', value: 'GarbageOverflow' },
  { label: '烟雾检测', value: 'SmokeDetected' },
  { label: '火焰监测', value: 'FireDetected' },
  { label: '裸土识别', value: 'LoessExposed' },
  { label: '积水识别', value: 'RoadPonding' },
  { label: '机动车违停', value: 'MotorVehicleParking' },
  { label: '出店经营', value: 'OffStoreOperation' },
  { label: '户外广告牌识别', value: 'IllegalOutdoorAdvertising' },
  { label: '拉横幅识别', value: 'BannerOrSlogansHungging' },
  { label: '沿街晾晒', value: 'HangDownTheStreet' },
  { label: '井盖识别', value: 'ManholeCoverDamaged' },
  { label: '渣土车抓拍', value: 'ConstructionTruck' },
  { label: '机动车识别（不带车牌）', value: 'MotorVehicleBreakIn' },
  { label: '安全帽识别（未带）', value: 'WithoutHelmetOnSite' },
  { label: '工程车辆检测', value: 'EngineeringVehicle' },
  { label: '施工占道', value: 'RoadOccupatedInConstruction' },
  { label: '道路破损', value: 'RoadDamaged' },
  { label: '乱堆物料', value: 'HeapOfMaterial' },
  { label: '占用无障碍通道', value: 'OccupationOfBarrierFreeAccess' },
  { label: '玩手机打电话检测', value: 'PlayPhone' },
  { label: '非法垂钓', value: 'Fishing' },
  { label: '离岗检测', value: 'OffDutyAlarm' },
  { label: '安全头盔识别（未带）', value: 'WithoutSafeHelmet' },
  { label: '非机动车识别', value: 'NoneMotorVehicleBreakIn' },
  { label: '扬尘监测', value: 'GroundDust' },
  { label: '人员攀爬', value: 'Climbing' },
  { label: '人流密度', value: 'PersonNumber' },
  { label: '人员尾随入户', value: 'FollowIntoHousePerson' },
  { label: '厨师帽识别', value: 'WithoutChefHat' },
  { label: '白色厨师服识别', value: 'WearDetection' },
  { label: '人员摔倒', value: 'PersonnelFalls' },
  { label: '越界识别', value: 'HumanCrossTheBorder' },
  { label: '客流统计', value: 'PassengerFlow' },
  { label: '车辆逆行识别', value: 'Retrograde' },
  { label: '工程车冒黑烟', value: 'SmokyEngineeringVehicle' },
  { label: '徘徊识别', value: 'HumanHover' },
];

export const option = {
  Smoking: '吸烟检测',
  HumanBreakIn: '人员入侵',
  PeopleWithoutMask: '未戴口罩',
  NoneMotorVehicleParking: '非机动车乱停',
  UnexpectedEvents: '突发性事件',
  RoadOccupatedInOperation: '占道经营',
  ElectricCarEntersElevator: '电动车检测',
  OccupationOfFireAccess: '占用消防通道',
  GarbageExposed: '垃圾暴露',
  GarbageOverflow: '垃圾满溢',
  SmokeDetected: '烟雾检测',
  FireDetected: '火焰监测',
  LoessExposed: '裸土识别',
  RoadPonding: '积水识别',
  MotorVehicleParking: '机动车违停',
  OffStoreOperation: '出店经营',
  IllegalOutdoorAdvertising: '户外广告牌识别',
  BannerOrSlogansHungging: '拉横幅识别',
  HangDownTheStreet: '沿街晾晒',
  ManholeCoverDamaged: '井盖识别',
  ConstructionTruck: '渣土车抓拍',
  MotorVehicleBreakIn: '机动车识别（不带车牌）',
  WithoutHelmetOnSite: '安全帽识别（未带）',
  EngineeringVehicle: '工程车辆检测',
  RoadOccupatedInConstruction: '施工占道',
  RoadDamaged: '道路破损',
  HeapOfMaterial: '乱堆物料',
  OccupationOfBarrierFreeAccess: '占用无障碍通道',
  PlayPhone: '玩手机打电话检测',
  Fishing: '非法垂钓',
  OffDutyAlarm: '离岗检测',
  WithoutSafeHelmet: '安全头盔识别（未带）',
  NoneMotorVehicleBreakIn: '非机动车识别',
  GroundDust: '扬尘监测',
  Climbing: '人员攀爬',
  PersonNumber: '人流密度',
  FollowIntoHousePerson: '人员尾随入户',
  WithoutChefHat: '厨师帽识别',
  WearDetection: '白色厨师服识别',
  PersonnelFalls: '人员摔倒',
  HumanCrossTheBorder: '越界识别',
  PassengerFlow: '客流统计',
  Retrograde: '车辆逆行识别',
  SmokyEngineeringVehicle: '工程车冒黑烟',
  HumanHover: '徘徊识别',
};

export const revOption = {};

for (const key in option) {
  if (option.hasOwnProperty(key)) {
    revOption[option[key]] = key;
  }
}
