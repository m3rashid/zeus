/// <reference types="vitest" />
/// <reference types="vite/client" />

import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';

const config = defineConfig({
  plugins: [solidPlugin()],
  server: {
    port: 3000,
  },
  build: {
    target: 'esnext',
  },
  test: {
    environment: 'jsdom',
    globals: true,
    testTransformMode: { web: ['/.[jt]sx?$/'] },
    coverage: {
      provider: 'v8',
      reporter: ['text', 'json', 'html'],
      include: ['src/**/*'],
    },
  },
});

export default config;
