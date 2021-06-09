import request from '@/utils/request.js'

export function PluginList () {
  return request({
    url: '/pluginList',
    method: 'get'
  })
}

export function PluginCheck (data) {
  return request({
    url: '/check',
    method: 'post',
    data: data
  })
}
