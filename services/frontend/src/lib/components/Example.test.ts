import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import Example from './Example.svelte';

describe('Example Component', () => {
	it('renders with default props', () => {
		render(Example);
		expect(screen.getByText('Hello World')).toBeInTheDocument();
	});

	it('renders with custom message', () => {
		render(Example, { props: { message: 'Custom Message' } });
		expect(screen.getByText('Custom Message')).toBeInTheDocument();
	});

	it('handles click events', async () => {
		const { component } = render(Example);
		let clicked = false;

		component.$on('click', () => {
			clicked = true;
		});

		const button = screen.getByRole('button');
		await button.click();

		expect(clicked).toBe(true);
	});
});
