import '@testing-library/jest-dom/vitest';

// Mock SvelteKit modules
vi.mock('$app/environment', () => ({
	browser: false,
	dev: true,
	building: false,
	version: 'test',
}));

vi.mock('$app/navigation', () => ({
	goto: vi.fn(),
	replaceState: vi.fn(),
	pushState: vi.fn(),
	back: vi.fn(),
	forward: vi.fn(),
}));

vi.mock('$app/stores', () => {
	const getStores = () => {
		const navigating = writable(null);
		const page = writable({
			url: new URL('http://localhost'),
			params: {},
			route: {
				id: null,
			},
			status: 200,
			error: null,
			data: {},
			form: undefined,
		});
		const updated = { subscribe: readable(false).subscribe, check: async () => false };

		return { navigating, page, updated };
	};

	const page = {
		subscribe(fn: any) {
			return getStores().page.subscribe(fn);
		},
	};
	const navigating = {
		subscribe(fn: any) {
			return getStores().navigating.subscribe(fn);
		},
	};
	const updated = {
		subscribe(fn: any) {
			return getStores().updated.subscribe(fn);
		},
		check: async () => false,
	};

	return {
		getStores,
		page,
		navigating,
		updated,
	};
});

import { readable, writable } from 'svelte/store';

// Add any additional global test setup here
