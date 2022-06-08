import axios from 'axios'

import config from '../../config'

const API = axios.create({ baseURL: config.apiPath })

/**
 * Devolte os Headers HTTP padrão e customizados
 * @param {*} headers objeto ex: {'Authorization': 'Basic qwerqwerqwer'}
 */
export function getHeaders (headers = {}) {
  let lheaders = {}

  if (window.localStorage.autenticated) {
    lheaders.Authorization = 'Bearer ' + JSON.parse(window.localStorage.autenticated).token
  }

  if (headers) {
    lheaders = { ...lheaders, ...headers }
  }

  return lheaders
}

/**
 * Metodos HTTP para interações Ajax
 */

export const getData = async (url, headers) => {
  return API.get(url, { headers: getHeaders(headers) })
}

export const postData = async (url, data, headers) => {
  return API.post(url, data, { headers: getHeaders(headers) })
}

export const putData = async (url, data, headers) => {
  return API.put(url, data, { headers: getHeaders(headers) })
}

export const deleteData = async (url, headers) => {
  return API.delete(url, { headers: getHeaders(headers) })
}

export default {
  getHeaders,
  getData,
  postData,
  putData,
  deleteData
}
