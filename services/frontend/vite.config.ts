import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      src: "/src",
      "@": "/",
    },
  },
  build: {
    assetsDir: "./",
    outDir: "production",
  },
  base: "./",
  preview: {
    host: true,
  },
});
