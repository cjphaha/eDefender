<template>
    <div>
        <div>
            <div>cpu使用率</div>
            <el-progress type="circle" :percentage="25"></el-progress>
            <div>型号</div>
            <div>数量</div>
        </div>
        <el-divider></el-divider>
        <div>
            <div>内存使用率</div>
            <el-progress type="circle" :percentage="25"></el-progress>
            <div>型号</div>
            <div>数量</div>
        </div>
        <div>
            <div>主机名称</div>
            <div>系统</div>
        </div>
    </div>
</template>

<script>
import { GetStatus } from '@/api/status.js'
export default {
  data () {
    return {
      info: {}
    }
  },
  methods: {
    getInfo: function () {
      GetStatus().then(Response => {
        this.info = Response
        console.log(this.info)
      })
    }
  },
  created () {
    setInterval(this.getInfo, 10000)
  }
}
</script>
