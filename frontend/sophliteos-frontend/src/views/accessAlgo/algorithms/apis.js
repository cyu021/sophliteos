
const serverUrl = 'http://localhost:38081/dynamic';

const list = () => {
  // generate mock data
  const data = [
    { key: 1, annotator_name: 'full_structure', sts: 1 },
    { key: 2, annotator_name: 'full_structure_2', sts: 0 },
  ];
  // return with a promise
  return new Promise((resolve) => resolve(data));

  // const url = `${serverUrl}/getDynamicModel`;
  // return fetch(url, {
  //   method: 'GET',
  //   headers: {
  //     'Content-Type': 'application/json',
  //   },
  // }).then((res) => res.json());
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
