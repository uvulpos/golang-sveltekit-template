import type { Meta, StoryObj } from '@storybook/svelte';
import { Author } from '.';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories
const meta = {
    title: 'base-components/Author',
    component: Author,
    tags: ['autodocs'],
    argTypes: {
        roundCorner: {
            type: 'boolean',
            description: 'storybook cosmetics'
        },
    },
    args: {
        roundCorner: true,
    }
} satisfies Meta<Author>;

export default meta;
type Story = StoryObj<typeof meta>;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Primary: Story = {
    args: {
        authorData: [
            {
                header: "Author",
                value: "Tim Riedl",
                imagePath: "/assets/author/example.jpg",
            },
            {
                header: "Published on",
                value: "13.04.2021",
            },
            {
                header: "Execution status",
                value: "success",
                iconName: "check-green",
            },
            {
                "header": "Publish status",
                "value": "Unpublished",
                "iconName": "cross-red"
            }
        ]
    },
};