import type { Meta, StoryObj } from '@storybook/svelte';
import { Icon, IconType } from './';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
const meta = {
    title: 'base-components/Icon',
    component: Icon,
    tags: ['autodocs'],
    argTypes: {
        icon: {
            description: "define icon-name. fetches Icon from `/assets/icons/${icon}.svg`"
        },
        fallbackicon: {
            type: "string",
            description: "if icon cannoit be loaded, use fallback"
        }
    },
    parameters: {
        backgrounds: {
            default: 'navigation-background',
            values: [
                { name: 'navigation-background', value: '#222425' },
                { name: 'navigation-background2', value: '#f00' },
            ],
        },
    }
} satisfies Meta<Icon>;

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Primary: Story = {
    args: {
        icon: IconType.Home,
        // fallbackicon: IconType.User
    },
};