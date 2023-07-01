import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/kit/vite';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),

	kit: {
		// See https://kit.svelte.dev/docs/adapters for more information about adapters.
		adapter: adapter({

            pages: './dist',
            assets: './dist',
            fallback: "index.html",
            precompress: false,
            strict: false,
        })
	}
};

export default config;
