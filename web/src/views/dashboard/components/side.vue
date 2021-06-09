<template>
    <el-menu
      :router="true"
      default-active="1"
      class="el-menu-vertical-demo"
      background-color="#E9EEF3">
      <el-menu-item index="host-status">
        <i class="el-icon-menu"></i>
        <span slot="title">
            服务器状态
        </span>
      </el-menu-item>
      <el-menu-item index="plugin-check">
        <i class="el-icon-setting"></i>
        <span slot="title">安全检查</span>
      </el-menu-item>
      <el-menu-item index="/plugin-list">
        <i class="el-icon-setting"></i>
        <span slot="title">
          功能列表
        </span>
      </el-menu-item>
      <el-menu-item index="about">
        <i class="el-icon-setting"></i>
        <span slot="title"> 关于</span>
      </el-menu-item>
    </el-menu>
</template>
