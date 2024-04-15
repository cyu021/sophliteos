
const serverUrl = 'http://localhost:38081/dynamic';

const list = () => {
  const url = `${serverUrl}/getDynamicModel`;
  return fetch(url, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  }).then((res) => res.json());
};

const upload = (file) => {
  const url = `${serverUrl}/uploadDynamicModel`;
  return fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'multipart/form-data',
    },
    body: file,
  }).then((res) => res.json());
};

const start = (name) => {
  const url = `${serverUrl}/onoffDynamicModel?annotatorName=${name}&onoff=true`;
  return fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  }).then((res) => res.json());
};

const stop = (name) => {
  const url = `${serverUrl}/onoffDynamicModel?annotatorName=${name}&onoff=false`;
  return fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  }).then((res) => res.json());
};

const deleteFile = (name) => {
  const url = `${serverUrl}/deleteDynamicModel?annotatorName=${name}`;
  return fetch(url, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
  }).then((res) => res.json());
};

const apis = {
  list,
  upload,
  start,
  stop,
  deleteFile,
};

export default apis;
