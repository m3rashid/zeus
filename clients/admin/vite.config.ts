import path from 'path';
import { defineConfig } from 'vite';
import solid from 'vite-plugin-solid';

const config = defineConfig({
  plugins: [solid()],
  server: {
    port: 3001,
  },
  resolve: {
    alias: [{ find: '@admin', replacement: path.resolve(__dirname, 'src') }],
  },
});

export default config;
