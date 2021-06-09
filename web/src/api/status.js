import request from '@/utils/request.js'

export function GetStatus () {
  return request({
    url: '/info',
    method: 'get'
  })
}
