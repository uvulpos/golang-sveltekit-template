import adapter from "@sveltejs/adapter-static";
import { vitePreprocess } from "@sveltejs/kit/vite";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: [vitePreprocess()],
	kit: {
		adapter: adapter({
			pages: "./dist",
			assets: "./dist",
			fallback: "index.html",
			precompress: false,
			strict: false,
		}),
	},
	vitePlugin: {
		inspector: true,
	},
};

export default config;
