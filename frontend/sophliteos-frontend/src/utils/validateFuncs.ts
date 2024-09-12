import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();

// IP check  粗略校验 {0~255}.{0~255}.{0~255}.{0~255}
export const IpCheck = (_rule, value) => {
  if (value.trim() === '') {
    return Promise.reject(t('maintenance.newworkSettings.inputIp'));
  } else {
    const ip = /^((\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
    if (!ip.test(value)) {
      return Promise.reject(t('maintenance.newworkSettings.validateWrong'));
    } else {
      return Promise.resolve();
    }
  }
};

// subnetMask check  粗略校验 {0~255}.{0~255}.{0~255}.{0~255}
export const subnetMaskCheck = (_rule, value) => {
  if (value.trim() === '') {
    return Promise.reject(t('maintenance.newworkSettings.inpuSubnetMask'));
  } else {
    const subnetMask = /^((\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
    if (!subnetMask.test(value)) {
      return Promise.reject(t('maintenance.newworkSettings.validateWrong'));
    } else {
      return Promise.resolve();
    }
  }
};

// gateway check  粗略校验 {0~255}.{0~255}.{0~255}.{0~255}
export const gatewayCheck = (_rule, value) => {
  if (value.trim() === '') {
    return Promise.reject(t('maintenance.newworkSettings.inputGateway'));
  } else {
    const gateway = /^((\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
    if (!gateway.test(value)) {
      return Promise.reject(t('maintenance.newworkSettings.validateWrong'));
    } else {
      return Promise.resolve();
    }
  }
};

// dns check, 粗略校验 {0~255}.{0~255}.{0~255}.{0~255}
export const dnsCheck = (_rule, value) => {
  if (value.trim() === '') {
    return Promise.reject(t('maintenance.newworkSettings.inputDns'));
  } else {
    const dns = /^((\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
    if (!dns.test(value)) {
      return Promise.reject(t('maintenance.newworkSettings.validateWrong'));
    } else {
      return Promise.resolve();
    }
  }
};

export const protocolCheck = (_rule, value) => {
  console.log("protocolCheck on "+value);
  const protocolRegex = /^http[s]?$/;
  if(!protocolRegex.test(value)) {
    return Promise.reject(t('maintenance.newworkSettings.inputProtocol'));
  } else {
    return Promise.resolve();
  }
};

export const rotateRetentionCheck = (_rule, value) => {
  console.log("rotateRetentionCheck on "+value);
  if(value < 1 || value > 180) {
    return Promise.reject(t('maintenance.rotateConfig.inputRotateRetention'));
  } else {
    return Promise.resolve();
  }
};

export const rotateRetryRowsCheck = (_rule, value) => {
  console.log("rotateRetryRowsCheck on "+value);
  if(value < 1000 || value > 5000) {
    return Promise.reject(t('maintenance.rotateConfig.inputRotateRetryRows'));
  } else {
    return Promise.resolve();
  }
};

export const portCheck = (_rule, value) => {
  console.log("portCheck on "+value);
  const portRegex = /[1-6]?[0-9]*[0-5]/;
  if(!portRegex.test(value)) {
    return Promise.reject(t('maintenance.newworkSettings.inputPort'));
  } else {
    var port2Check = Number(value)
    if (port2Check < 1 || port2Check > 65535) {
      return Promise.reject(t('maintenance.newworkSettings.validateWrong'));
    } else {
      return Promise.resolve();
    }
  }
};

// multi IP check，
export const IpDoubleCheck = (_rule, value) => {
  if (value.trim() === '') {
    return Promise.reject(t('maintenance.newworkSettings.inputIp'));
  } else {
    const doubleIp =
      /^((\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d{1,2}|1\d\d|2[0-4]\d|25[0-5]),((\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
    if (!doubleIp.test(value)) {
      return Promise.reject(t('maintenance.newworkSettings.validateWrong'));
    } else {
      return Promise.resolve();
    }
  }
};

// 密码校验：大写字母，小写字母、数字，特征字符至少3种，长度8-16位
export const passwordCheck = (_rule, value) => {
  if (value === '') {
    return Promise.reject(t('layout.header.inputOldPassword'));
  } else {
    const password =
      /^(?![\da-z]+$)(?![\dA-Z]+$)(?![\d!@#$%^&*]+$)(?![a-zA-Z]+$)(?![a-z!@#$%^&*]+$)(?![A-Z!@#$%^&*]+$)[\da-zA-z!@#$%^&*]{8,16}$/;
    if (!password.test(value)) {
      return Promise.reject(t('layout.header.passwordFormatError'));
    } else {
      return Promise.resolve();
    }
  }
};
