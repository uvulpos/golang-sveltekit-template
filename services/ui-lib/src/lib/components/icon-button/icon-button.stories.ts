import type { Meta, StoryObj } from '@storybook/svelte';
import { IconButton } from '.';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
const meta = {
    title: 'base-components/Icon-Button',
    component: IconButton,
    tags: ['autodocs'],
    argTypes: {
        label: {
            type: 'string',
            description: 'Content of the button',
            defaultValue: '',
        },
        icon: {
            type: 'string',
            description: 'Filename of the required icon. Use IconType enum for safer use',
            defaultValue: 'unknown-icon',
        },
        animateButton: {
            type: 'boolean',
            description: 'If the hover effekt should be used or not',
            defaultValue: false,
        },
        href: {
            type: 'string',
            description: 'URL, if you want to change site on press',
            defaultValue: '',
        },
        targetBlank: {
            type: 'boolean',
            description: 'Open url in new window',
            defaultValue: false,
        }
    },
    args: {
        icon: 'home',
        label: 'My-Button',
        animateButton: true,
        href: 'https://www.google.com',
        targetBlank: true
    }
} satisfies Meta<IconButton>;

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Primary: Story = {
    args: {
    },
};

// export const Secondary: Story = {
//     args: {
//         icon: 'user',
//         label: 'Button',
//     },
// };
