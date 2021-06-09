<template>
    <div>
        <div>
            <div>
              <div>cpu使用率</div>
              <el-progress type="circle" :percentage="info.Cpu[0].使用率"></el-progress>
              <div>
                <div>型号: {{info.Cpu[0].型号}}</div>
                <div>数量: {{info.Cpu[0].数量}}</div>
              </div>
            </div>
        </div>
        <el-divider></el-divider>
        <div>
            <div>内存使用率</div>
            <el-progress type="circle" :percentage="info.Mem.使用率"></el-progress>
            <div>可用: {{info.Mem.可用}}</div>
            <div>已使用: {{info.Mem.已使用}}</div>
            <div>总量: {{info.Mem.总量}}</div>
            <div>空闲: {{info.Mem.空闲}}</div>
        </div>
        <el-divider></el-divider>
        <div>
            <div>主机名称:{{info.Host.主机名称}}</div>
            <div>内核:{{info.Host.内核}}</div>
            <div>平台:{{info.Host.平台}}</div>
            <div>系统:{{info.Host.系统}}</div>
        </div>
    </div>
</template>

<script>
import { GetStatus } from '@/api/status.js'
export default {
  data () {
    return {
      info: {
        Cpu: [],
        Disk: null,
        Host: {},
        Mem: {}
      }
    }
  },
  computed: {

  },
  methods: {
    getInfo: function () {
      GetStatus().then(Response => {
        this.info = Response
      })
    }
  },
  created () {
    setInterval(this.getInfo, 15000)
  }
}
</script>
