import { defineConfig } from 'vite';
import solid from 'vite-plugin-solid';

const config = defineConfig({
  plugins: [solid()],
  server: {
    port: 3001,
  },
});

export default config;
