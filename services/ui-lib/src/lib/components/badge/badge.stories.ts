import type { Meta, StoryObj } from '@storybook/svelte';
import { Badge } from '.';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
const meta = {
    title: 'base-components/Badge',
    component: Badge,
    tags: ['autodocs'],
    argTypes: {
        label: {
            type: "string",
            description: "Value of the Badge"
        }
    },
    args: {
        label: "I like 👍🏻"
    }
} satisfies Meta<Badge>;

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Primary: Story = {
    args: {
    },
};