import path from 'path';
import { defineConfig } from 'vite';
import solid from 'vite-plugin-solid';

const config = defineConfig({
  plugins: [solid()],
  server: { port: 3000 },
  resolve: {
    alias: [{ find: '@web', replacement: path.resolve(__dirname, 'src') }],
  },
});

export default config;
