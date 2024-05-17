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
        },
        enableSearch: {
            type: 'boolean',
            description: 'enables the search input field',
            defaultValue: false
        },
        mobileVersion: {
            type: 'boolean',
            description: 'toggles hamburger menu or desktop navigation',
            defaultValue: false
        },
        environment: {
            type: 'string',
            description: 'displays the current environment',
            defaultValue: ""
        }
    },
    args: {
        roundCorner: true,
        enableSearch: true,
        mobileVersion: false,
    }
} satisfies Meta<Navigation>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Production: Story = {
    args: {
    },
};
export const StagingEnvironment: Story = {
    args: {
        environment: "Staging"
    },
};

export const Mobile: Story = {
    args: {
        mobileVersion: true
    },
};