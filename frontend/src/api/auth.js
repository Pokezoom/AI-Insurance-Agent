import request from '@/utils/request'

export const register = (data) => {
  return request.post('/auth/register', data)
}

export const login = (data) => {
  return request.post('/auth/login', data)
}

export const getProfile = () => {
  return request.get('/user/profile')
}

export const changePassword = (data) => {
  return request.put('/user/password', data)
}
