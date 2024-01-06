import multiAdapter from '@macfja/svelte-multi-adapter'
import staticAdapter from "@sveltejs/adapter-static";
import autoAdatper from "@sveltejs/adapter-auto";
import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: [vitePreprocess()],
	kit: {
		adapter: multiAdapter([
			staticAdapter({
				pages: "./dist",
				assets: "./dist",
				fallback: "index.html",
				precompress: false,
				strict: false,
			}),
			autoAdatper(),
		]),
	},
	vitePlugin: {
		inspector: true,
	},
};

export default config;
