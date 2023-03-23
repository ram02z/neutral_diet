import * as path from "path";
import { build, defineConfig } from "vite";
import react from "@vitejs/plugin-react";

const firebaseSwPlugin: import('vite').Plugin = {
  name: 'firebaseSwPlugin',
  apply: 'build',
  enforce: 'post',
  buildEnd: async () => {
    const conf = {
      build: {
        lib: {
          entry: path.resolve(__dirname, './src/firebase-messaging-sw.js'),
          formats: ['es'],
        },
        rollupOptions: {
          output: {
            entryFileNames: 'firebase-messaging-sw.js',
          },
        },
        outDir: path.resolve(__dirname, './dist'),
        emptyOutDir: false,
      },
      configFile: false,
    };

    await build(conf);
  }

}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    firebaseSwPlugin,
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});
