import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), "");
  const apiBase = env.VITE_API_BASE_URL || "/api";
  const proxyTarget = env.VITE_PROXY_TARGET || "http://127.0.0.1:8094";

  return {
    plugins: [react()],
    server: apiBase.startsWith("http")
      ? undefined
      : {
          proxy: {
            [apiBase]: {
              target: proxyTarget,
              changeOrigin: true,
              rewrite: (path) => path.replace(new RegExp(`^${apiBase}`), ""),
            },
          },
        },
  };
});
