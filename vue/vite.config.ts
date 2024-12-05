import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons';
export default defineConfig({
  base:'/static/',
  plugins: [
    vue()
    ,createSvgIconsPlugin({
      // 指定需要缓存的图标文件夹
      iconDirs: [path.resolve(process.cwd(), "src/assets/svgs")],
      // 指定symbolId格式
      symbolId: "[name]",
      customDomId: "turing-planet-svgs", // 避免多项目互相影响
    })

  ],
  resolve:{
    // 导入以下文件可以不用后缀名
    extensions:['.vue','.ts'],
    // 配置路径别名
    alias:{
      "@":path.resolve(__dirname,"src")
    }
  },
  server:{
    // 配置代理
    proxy:{
      "/api":{
        target:'http://127.0.0.1:8008',//设置代理目标
        changeOrigin:true,//是否改变请求源地址
        rewrite: (path) => {
          return path.replace(/\/api/, '')
      }
      }
    }
  },
})