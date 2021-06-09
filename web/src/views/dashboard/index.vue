<template>
    <el-container>
        <el-header>
            e-defender
        </el-header>
        <el-container>
            <el-aside width="200px">
            <side></side>
        </el-aside>
        <el-main>
            <router-view></router-view>
        </el-main>
        </el-container>

        <el-footer></el-footer>
    </el-container>
</template>

<script>
import side from '@/views/dashboard/components/side'
export default {
  components: {
    side
  }
}
</script>

<style>
  .el-header, .el-footer {
    color: #333;
    text-align: center;
    line-height: 60px;
  }

  .el-aside {
    background-color: #E9EEF3;
    color: #333;
    text-align: center;
    line-height: 200px;
  }

</style>
