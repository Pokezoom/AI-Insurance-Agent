import request from '@/utils/request'

export const getUserList = (params) => {
  return request.get('/admin/users', { params })
}

export const updateUserStatus = (userId, data) => {
  return request.put(`/admin/users/${userId}/status`, data)
}
