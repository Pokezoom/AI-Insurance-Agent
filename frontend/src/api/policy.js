import request from '@/utils/request'

export const analyzePolicy = (data) => {
  return request.post('/policy/analyze', data)
}

export const getRecords = (params) => {
  return request.get('/policy/records', { params })
}

export const getRecordDetail = (id) => {
  return request.get(`/policy/records/${id}`)
}

export const deleteRecord = (id) => {
  return request.delete(`/policy/records/${id}`)
}
