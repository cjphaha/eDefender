<template>
    <div>
        <el-form>
            <el-form-item label="type">
                <el-input v-model="formData.type"></el-input>
            </el-form-item>
            <el-form-item label="netloc">
                <el-input v-model="formData.netloc"></el-input>
            </el-form-item>
            <el-form-item label="target">
                <el-input v-model="formData.target"></el-input>
            </el-form-item>
            <el-form-item label="检测结果">
                <el-input type="textarea" v-model="result" placeholder="这里展示测试结果"></el-input>
            </el-form-item>
        </el-form>
        <el-button type="primary" @click="check">检查</el-button>
    </div>
</template>

<script>
import { PluginCheck } from '@/api/api'
export default {
  data () {
    return {
      result: {},
      formData: {
        type: '',
        netloc: '',
        target: ''
      }
    }
  },
  methods: {
    check: function () {
      PluginCheck(this.formData).then(Response => {
        if (Response === null) {
          this.result = 'null'
        } else {
          this.result = Response
        }
      })
    }
  }
}
</script>
