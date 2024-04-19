import type { Meta, StoryObj } from '@storybook/svelte';
import { Navigation } from '.';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
const meta = {
    title: 'base-components/Navigation',
    component: Navigation,
    tags: ['autodocs'],
    argTypes: {
        roundCorner: {
            type: 'boolean',
            description: 'storybook cosmetics'
        }
    },
    args: {
        roundCorner: true,
    }
} satisfies Meta<Navigation>;

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Primary: Story = {
    args: {
    },
};