import type { Meta, StoryObj } from '@storybook/svelte';
import { Button } from './';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
const meta = {
    title: 'base-components/Button',
    component: Button,
    tags: ['autodocs'],
    argTypes: {
    },
} satisfies Meta<Button>;

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Primary: Story = {
    args: {
        icon: 'home',
        label: 'Button',
    },
};

// export const Secondary: Story = {
//     args: {
//         icon: 'user',
//         label: 'Button',
//     },
// };
