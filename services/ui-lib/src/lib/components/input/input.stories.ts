import type { Meta, StoryObj } from '@storybook/svelte';
import { Input } from '.';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
const meta = {
    title: 'base-components/Input',
    component: Input,
    tags: ['autodocs'],
    argTypes: {
        value: {
            type: "string",
            description: "",
        },
        placeholder: {
            type: "string",
            description: ""
        }
    },
    args: {
        placeholder: "empty input field"
    }
} satisfies Meta<Input>;

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Primary: Story = {
    args: {

    },
};