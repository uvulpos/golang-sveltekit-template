import type { Meta, StoryObj } from '@storybook/svelte';
import { Logger } from '.';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
const meta = {
    title: 'base-components/Logger',
    component: Logger,
    tags: ['autodocs'],
    argTypes: {
        showTime: {
            type: "boolean",
            default: true,
            description: "show time column or hide it"
        },
        showEmojis: {
            type: "boolean",
            default: false,
            description: "show emojis column or hide it"
        }
    },
    args: {
    }
} satisfies Meta<Logger>;

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Primary: Story = {
    args: {

    },
};