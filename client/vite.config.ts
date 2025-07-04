import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'
import path from 'path'

// https://vite.dev/config/
export default defineConfig({
  server: {
    watch: {
      usePolling: true,
    },
    host: '0.0.0.0',
    port: 3000,
  },
  plugins: [
    vue(),
    vueJsx(),
    vueDevTools(),
    createSvgIconsPlugin({
      iconDirs: [path.resolve(__dirname, 'src/assets/svg')],
      symbolId: 'icon-[name]',
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
