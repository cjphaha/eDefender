<template>
    <div>
        <el-table
        :data="list">
            <el-table-column
            label="id"
            >
            <template slot-scope="scope">
                {{scope.row.references.kpid}}
            </template>
            </el-table-column>
            <el-table-column
            label="名称"
            prop="name">
            </el-table-column>
            <el-table-column
            label="描述"
            prop="remarks">
            </el-table-column>
            <el-table-column
            label="检测种类"
            prop="type">
            </el-table-column>
            <el-table-column
            label="目标插件"
            prop="target">
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import { PluginList } from '@/api/api'
export default {
  data () {
    return {
      list: []
    }
  },
  methods: {
    getInfo: function () {
      PluginList().then(Response => {
        this.list = Response
      })
    }
  },
  created () {
    this.getInfo()
  }
}
</script>
